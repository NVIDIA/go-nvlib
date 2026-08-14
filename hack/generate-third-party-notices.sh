#!/usr/bin/env bash
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

OUTPUT="${OUTPUT:-THIRD_PARTY_NOTICES.md}"
MODULES_TXT="${MODULES_TXT:-vendor/modules.txt}"
ASSET_LICENSES_DIR="${ASSET_LICENSES_DIR:-hack/notices/licenses}"
NOTICE_FILE="${NOTICE_FILE:-NOTICE}"

LICENSE_PHRASE_TO_SPDX_ID=(
    "3-clause BSD License|BSD-3-Clause"
    "GNU General Public License version 2|GPL-2.0-only"
    "GNU General Public License version 2 or later|GPL-2.0-or-later"
)

NOTICE_PATH_AND_PHRASE_RE='^The file ([^[:space:]]+) is distributed under the (.+)\.[[:space:]]+Maintained by '
NOTICE_COMPONENT_AND_URL_RE='[[:space:]]from the (.+) at (https?://[^[:space:]]+)[[:space:]]*$'

NOTICE_FORMAT_HELP=(
    ""
    "Describe each bundled file in ${NOTICE_FILE} as one blank-line separated paragraph:"
    ""
    "    The file <repo-relative path> is distributed under the <license>."
    "    Maintained by <credits> from"
    "    the <component> at <url>."
    ""
    "For example:"
    ""
    "    The file pkg/pciids/default_pci.ids is distributed under the 3-clause BSD License."
    "    Maintained by Albert Pool, Martin Mares, and other volunteers from"
    "    the PCI ID Project at https://pci-ids.ucw.cz/."
)

PACKAGE_PATTERNS=("./...")

# Listed explicitly: go-licenses resolves one platform per run and build tags differ.
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

die() {
    printf 'ERROR: %s\n' "$1" >&2
    shift
    if (( $# > 0 )); then
        printf '%s\n' "$@" >&2
    fi
    exit 1
}

log() {
    printf '%s\n' "$*" >&2
}

# A license that is itself Markdown would close a fixed ``` fence early.
fence_for_file() {
    local file="$1" longest_backtick_run fence_width
    # -a: a license holding a NUL byte would print "Binary file ... matches".
    longest_backtick_run=$(LC_ALL=C grep -oaE '`+' "${file}" 2>/dev/null \
        | awk '{ if (length($0) > longest) longest = length($0) } END { print longest+0 }')
    fence_width=$(( longest_backtick_run + 1 ))
    (( fence_width < 3 )) && fence_width=3
    printf '%*s' "${fence_width}" '' | tr ' ' '`'
}

emit_fenced_file() {
    local file="$1" fence
    fence="$(fence_for_file "${file}")"
    printf '%stext\n' "${fence}"
    cat "${file}"
    echo
    printf '%s\n' "${fence}"
    echo
}

check_prerequisites() {
    command -v go >/dev/null 2>&1 || die "go is not installed."

    if ./bin/go-licenses --help >/dev/null 2>&1; then
        GO_LICENSES_BIN="${PWD}/bin/go-licenses"
    else
        die "go-licenses is not installed, or ./bin/go-licenses cannot run on this host." \
            "Delete ./bin/go-licenses if it is present, then run 'make third-party-notices';" \
            "make reinstalls it once it is gone."
    fi

    local required_file
    for required_file in "${MODULES_TXT}" "${NOTICE_FILE}"; do
        [[ -f "${required_file}" ]] \
            || die "${required_file} not found — run 'make third-party-notices' from the repo root."
    done
    [[ -d "${ASSET_LICENSES_DIR}" ]] \
        || die "${ASSET_LICENSES_DIR} not found — run 'make third-party-notices' from the repo root."

    # Inside a Go workspace, go list and go-licenses resolve against every go.work module.
    export GOWORK=off

    LOCAL_MODULE=$(go list -m 2>/dev/null || true)
    [[ -n "${LOCAL_MODULE}" ]] || die "could not determine local module path via 'go list -m'."
    [[ "${LOCAL_MODULE}" != *$'\n'* ]] \
        || die "'go list -m' reported more than one module path:" \
               "" \
               "${LOCAL_MODULE}" \
               "" \
               "This generator describes a single module; run it from the root of one."

    # CGO must stay on: with it off, build constraints exclude every file in
    # go-nvml/pkg/dl and go-licenses cannot load ./... at all.
    export GOFLAGS="-mod=vendor"
    export CGO_ENABLED=1
}

prepare_workspace() {
    # Explicit templates: macOS mktemp ignores TMPDIR without one.
    local template_prefix="${TMPDIR:-/tmp}/go-nvlib-notices"
    PLATFORM_SAVE_ROOT="$(mktemp -d "${template_prefix}.XXXXXX")"
    LICENSES_DIR="$(mktemp -d "${template_prefix}-licenses.XXXXXX")"
    COMBINED_CSV="$(mktemp "${template_prefix}-csv.XXXXXX")"
    GO_INDEX_FILE="$(mktemp "${template_prefix}-go-index.XXXXXX")"
    ASSETS_INDEX_FILE="$(mktemp "${template_prefix}-assets.XXXXXX")"
    UNSORTED_ASSETS_INDEX_FILE="$(mktemp "${template_prefix}-assets-unsorted.XXXXXX")"
    ASSET_LEADING_COMMENT_FILE="$(mktemp "${template_prefix}-asset-comment.XXXXXX")"

    local output_dir
    output_dir="$(dirname "${OUTPUT}")"
    mkdir -p "${output_dir}"
    OUTPUT_TMP="$(mktemp "${output_dir}/.$(basename "${OUTPUT}").XXXXXX")"
    trap 'rm -rf "${PLATFORM_SAVE_ROOT}" "${LICENSES_DIR}"; rm -f "${COMBINED_CSV}" "${GO_INDEX_FILE}" "${ASSETS_INDEX_FILE}" "${UNSORTED_ASSETS_INDEX_FILE}" "${ASSET_LEADING_COMMENT_FILE}" "${OUTPUT_TMP}"' EXIT
}

collect_go_dependencies() {
    local platform goos goarch platform_save_dir

    for platform in "${PLATFORMS[@]}"; do
        goos="${platform%/*}"
        goarch="${platform#*/}"
        log "Collecting licenses for ${goos}/${goarch}..."

        platform_save_dir="${PLATFORM_SAVE_ROOT}/${goos}_${goarch}"

        # No --ignore: it matches raw string prefixes and drops silently, so "go"
        # would take golang.org/x/* and LOCAL_MODULE a real "${LOCAL_MODULE}-extra".
        GOOS="${goos}" GOARCH="${goarch}" "${GO_LICENSES_BIN}" save "${PACKAGE_PATTERNS[@]}" \
            --save_path="${platform_save_dir}" \
            --force

        GOOS="${goos}" GOARCH="${goarch}" "${GO_LICENSES_BIN}" csv "${PACKAGE_PATTERNS[@]}" \
            >> "${COMBINED_CSV}"

        cp -R "${platform_save_dir}/." "${LICENSES_DIR}/"
        chmod -R u+w "${LICENSES_DIR}"
    done
}

# Join rather than pick: go-licenses emits one row per recognized license.
join_licenses_per_package() {
    LC_ALL=C sort -u "$1" | awk -F, '
        {
            package_path = $1
            if (!(package_path in source_url)) {
                source_url[package_path] = $2
                package_order[++package_count] = package_path
            }
            if (!((package_path SUBSEP $3) in seen_package_license)) {
                seen_package_license[package_path SUBSEP $3] = 1
                # Count, do not test "package_path in licenses": mawk instantiates
                # the target before evaluating the RHS, BWK awk does not.
                licenses[package_path] = \
                    (license_count[package_path]++ ? licenses[package_path] " / " : "") $3
            }
        }
        END {
            for (i = 1; i <= package_count; i++) {
                package_path = package_order[i]
                print package_path "," source_url[package_path] "," licenses[package_path]
            }
        }
    '
}

drop_local_module() {
    awk -F, -v local_module="${LOCAL_MODULE}" '
        $1 == local_module || index($1, local_module "/") == 1 { next }
        { print }
    '
}

# Module path, not the URL go-licenses reports: in vendor mode that points into
# this repo at HEAD and stops describing released content once main moves.
append_module_path() {
    awk -v modules_txt="${MODULES_TXT}" '
        BEGIN {
            FS = OFS = ","
            while ((getline line < modules_txt) > 0) {
                if (line !~ /^# /) continue
                split(line, fields, " ")
                if (fields[4] == "=>" || fields[3] == "=>") {
                    replacement_path_field = (fields[4] == "=>") ? 5 : 4
                    if (fields[replacement_path_field + 1] == "") {
                        print "ERROR: " modules_txt " replaces " fields[2] " with a local path;" > "/dev/stderr"
                        print "teach hack/generate-third-party-notices.sh how to attribute it." > "/dev/stderr"
                        exit 1
                    }
                    module_paths[++module_count] = fields[2]
                    attributed_path[fields[2]] = fields[replacement_path_field]
                } else {
                    module_paths[++module_count] = fields[2]
                    attributed_path[fields[2]] = fields[2]
                }
            }
            close(modules_txt)
            if (module_count == 0) {
                print "ERROR: no module lines read from " modules_txt > "/dev/stderr"
                exit 1
            }
        }
        {
            longest_match = ""
            for (i = 1; i <= module_count; i++) {
                module_path = module_paths[i]
                if (($1 == module_path || index($1, module_path "/") == 1) &&
                    length(module_path) > length(longest_match)) longest_match = module_path
            }
            print $0, (longest_match == "" ? "unknown" : attributed_path[longest_match])
        }
    '
}

build_go_index() {
    log "Generating dependency index..."
    join_licenses_per_package "${COMBINED_CSV}" | drop_local_module | append_module_path \
        > "${GO_INDEX_FILE}"

    [[ -s "${GO_INDEX_FILE}" ]] \
        || die "go-licenses produced no entries for ${PACKAGE_PATTERNS[*]} — refusing to write empty notices file."

    # go-licenses reports an unclassifiable license as "Unknown" and exits 0.
    if cut -d, -f3 "${GO_INDEX_FILE}" | LC_ALL=C grep -qE '^$|(^| / )Unknown( / |$)'; then
        die "go-licenses could not classify the license of some packages." \
            "Identify them by hand rather than committing a file that says 'Unknown'."
    fi

    if cut -d, -f4 "${GO_INDEX_FILE}" | LC_ALL=C grep -qE '^$|^unknown$'; then
        die "could not resolve the module path for some packages from ${MODULES_TXT}." \
            "Run 'make vendor' and re-run, rather than committing a file with unattributed entries."
    fi

    local package_path _rest_of_row
    while IFS=, read -r package_path _rest_of_row; do
        [[ -z "${package_path}" ]] && continue
        [[ -n "$(license_files_in "${LICENSES_DIR}/${package_path}")" ]] \
            || die "no license text was saved for ${package_path}." \
                   "Re-run after 'make vendor'; do not commit an entry without its license."
    done < "${GO_INDEX_FILE}"
}

spdx_id_for_phrase() {
    local phrase_mapping
    for phrase_mapping in "${LICENSE_PHRASE_TO_SPDX_ID[@]}"; do
        if [[ "$1" == "${phrase_mapping%%|*}" ]]; then
            printf '%s' "${phrase_mapping##*|}"
            return 0
        fi
    done
    return 1
}

known_phrases_help() {
    local phrase_mapping
    printf '%s\n' "Phrases this generator recognizes:"
    for phrase_mapping in "${LICENSE_PHRASE_TO_SPDX_ID[@]}"; do
        printf '  "%s" -> %s\n' "${phrase_mapping%%|*}" "${phrase_mapping##*|}"
    done
}

note_for_path() {
    case "$1" in
        pkg/pciids/default_pci.ids)
            printf '%s' "Upstream offers this database under either the GNU General Public License (version 2 or later) or the 3-clause BSD License. This project distributes it under the 3-clause BSD License, as recorded in the NOTICE file at the root of this repository. The file is embedded into the pkg/pciids package with go:embed, so it is compiled into every consumer of that package. Refresh it with 'make update-pcidb'."
            ;;
        *)
            return 1
            ;;
    esac
}

emit_notice_record() {
    local paragraph="$1" bundled_file license_phrase component_name source_url spdx_id note

    if [[ ! "${paragraph}" =~ ${NOTICE_PATH_AND_PHRASE_RE} ]]; then
        die "could not derive a bundled file's path and license from ${NOTICE_FILE}:" \
            "" \
            "    ${paragraph}" \
            "${NOTICE_FORMAT_HELP[@]}"
    fi
    bundled_file="${BASH_REMATCH[1]}"
    license_phrase="${BASH_REMATCH[2]}"

    if [[ ! "${paragraph}" =~ ${NOTICE_COMPONENT_AND_URL_RE} ]]; then
        die "could not derive the component name and source URL for '${bundled_file}' from ${NOTICE_FILE}:" \
            "" \
            "    ${paragraph}" \
            "${NOTICE_FORMAT_HELP[@]}"
    fi
    component_name="${BASH_REMATCH[1]}"
    # The URL ends the sentence, so it carries the full stop away with it.
    source_url="${BASH_REMATCH[2]%.}"

    if ! spdx_id="$(spdx_id_for_phrase "${license_phrase}")"; then
        die "${NOTICE_FILE} elects '${license_phrase}' for '${bundled_file}', which this generator does not know." \
            "Phrases are matched exactly, never guessed." \
            "" \
            "$(known_phrases_help)" \
            "" \
            "Add \"<phrase>|<SPDX id>\" to LICENSE_PHRASE_TO_SPDX_ID, and the verbatim license" \
            "text as ${ASSET_LICENSES_DIR}/<SPDX id>.txt."
    fi

    [[ -f "${bundled_file}" ]] \
        || die "${NOTICE_FILE} attributes '${bundled_file}', which does not exist." \
               "Point ${NOTICE_FILE} at the file's new location if it moved; do not drop the attribution."

    [[ -s "${ASSET_LICENSES_DIR}/${spdx_id}.txt" ]] \
        || die "${NOTICE_FILE} elects ${spdx_id} for '${bundled_file}', but ${ASSET_LICENSES_DIR}/${spdx_id}.txt does not exist." \
               "An SPDX identifier on its own is not a notice: add the verbatim upstream text there."

    if ! note="$(note_for_path "${bundled_file}")"; then
        die "${NOTICE_FILE} attributes '${bundled_file}', which note_for_path() has no note for." \
            "The note records what ${NOTICE_FILE} does not: why there is an election to make," \
            "how the file reaches consumers, and how it is refreshed."
    fi

    printf '%s|%s|%s|%s|%s\n' \
        "${bundled_file}" "${component_name}" "${spdx_id}" "${source_url}" "${note}"
}

build_assets_index() {
    log "Reading bundled non-Go file records from ${NOTICE_FILE}..."

    local line paragraph=""
    : > "${UNSORTED_ASSETS_INDEX_FILE}"

    # Read as paragraphs: NOTICE wraps its prose, so rewrapping it is not breaking.
    while IFS= read -r line || [[ -n "${line}" ]]; do
        if [[ -z "${line//[[:space:]]/}" ]]; then
            if [[ -n "${paragraph}" ]]; then
                emit_notice_record "${paragraph% }" >> "${UNSORTED_ASSETS_INDEX_FILE}"
                paragraph=""
            fi
            continue
        fi
        paragraph+="${line} "
    done < "${NOTICE_FILE}"
    if [[ -n "${paragraph}" ]]; then
        emit_notice_record "${paragraph% }" >> "${UNSORTED_ASSETS_INDEX_FILE}"
    fi

    LC_ALL=C sort -u "${UNSORTED_ASSETS_INDEX_FILE}" > "${ASSETS_INDEX_FILE}"

    [[ -s "${ASSETS_INDEX_FILE}" ]] \
        || die "${NOTICE_FILE} describes no bundled non-Go files." \
               "${NOTICE_FORMAT_HELP[@]}"
}

# Read from the file, not copied here: 'make update-pcidb' rewrites it wholesale.
extract_leading_comment() {
    awk '/^#/ { print; next } { exit }' "$1"
}

# Filter by name: for restricted licenses 'go-licenses save' copies whole source.
license_files_in() {
    local package_dir="$1" candidate_file
    [[ -d "${package_dir}" ]] || return 0
    while IFS= read -r -d '' candidate_file; do
        if printf '%s' "$(basename "${candidate_file}")" \
            | LC_ALL=C grep -qiE '^(licen[cs]e|notice|copying|copyright|authors|patents)([-._].*)?$'; then
            printf '%s\n' "${candidate_file}"
        fi
    done < <(find "${package_dir}" -maxdepth 1 -type f -print0 2>/dev/null | LC_ALL=C sort -z)
}

emit_go_index_table() {
    local package_path _source_url license_ids module_path
    printf '| Package | License | Module |\n'
    printf '|---------|---------|--------|\n'
    while IFS=, read -r package_path _source_url license_ids module_path; do
        [[ -z "${package_path}" ]] && continue
        # shellcheck disable=SC2016  # backticks are literal markdown here.
        printf '| `%s` | %s | `%s` |\n' "${package_path}" "${license_ids}" "${module_path}"
    done < "${GO_INDEX_FILE}"
}

emit_go_sections() {
    local package_path _source_url license_ids module_path license_file

    while IFS=, read -r package_path _source_url license_ids module_path; do
        [[ -z "${package_path}" ]] && continue

        printf '### %s\n\n' "${package_path}"
        printf '* License: %s\n' "${license_ids}"
        printf '* Module: %s\n\n' "${module_path}"

        while IFS= read -r license_file; do
            [[ -n "${license_file}" ]] || continue
            printf '#### %s\n\n' "$(basename "${license_file}")"
            emit_fenced_file "${license_file}"
        done < <(license_files_in "${LICENSES_DIR}/${package_path}")
        echo
    done < "${GO_INDEX_FILE}"
}

emit_assets_index_table() {
    local bundled_file component_name spdx_id source_url note
    printf '| File | Component | License | Source |\n'
    printf '|------|-----------|---------|--------|\n'
    while IFS='|' read -r bundled_file component_name spdx_id source_url note; do
        # shellcheck disable=SC2016  # backticks are literal markdown here.
        printf '| `%s` | %s | %s | %s |\n' \
            "${bundled_file}" "${component_name}" "${spdx_id}" "${source_url}"
    done < "${ASSETS_INDEX_FILE}"
}

emit_assets_sections() {
    local bundled_file component_name spdx_id source_url note

    while IFS='|' read -r bundled_file component_name spdx_id source_url note; do
        printf '### %s\n\n' "${component_name}"
        # shellcheck disable=SC2016  # backticks are literal markdown here.
        printf '* File: `%s`\n' "${bundled_file}"
        printf '* License: %s\n' "${spdx_id}"
        printf '* Source: %s\n' "${source_url}"
        printf '* Note: %s\n\n' "${note}"

        extract_leading_comment "${bundled_file}" > "${ASSET_LEADING_COMMENT_FILE}"
        if [[ -s "${ASSET_LEADING_COMMENT_FILE}" ]]; then
            printf '#### Notice as it appears in %s\n\n' "${bundled_file}"
            emit_fenced_file "${ASSET_LEADING_COMMENT_FILE}"
        fi

        printf '#### %s\n\n' "${spdx_id}"
        emit_fenced_file "${ASSET_LICENSES_DIR}/${spdx_id}.txt"
        echo
    done < "${ASSETS_INDEX_FILE}"
}

compose_document() {
    log "Composing ${OUTPUT}..."
    {
        cat <<'EOF'
# Third-Party Notices

NVIDIA go-nvlib

This file lists the third-party dependencies that go-nvlib links into the
packages a consumer imports, along with the verbatim text of each dependency's
license. It covers the **Go modules** in the non-test import closure of the
packages under `pkg/`, and the third-party non-Go files committed into this
repository and compiled into those packages.

Go standard library packages are excluded; they are covered by the license of
the Go distribution itself. NVIDIA's own code is excluded; it is covered by
`LICENSE`. Dependencies reached only from `_test.go` files are excluded; they
are vendored for testing but a consumer does not link them.

The `NOTICE` file at the root of this repository remains part of the
distribution. Its entry for `pkg/pciids/default_pci.ids` is reproduced in full
below, so this file stands on its own.

## Go Dependency Index

EOF
        emit_go_index_table

        cat <<'EOF'

## Bundled Non-Go File Index

EOF
        emit_assets_index_table

        cat <<'EOF'

## Go Dependency License Texts

EOF
        emit_go_sections

        cat <<'EOF'
## Bundled Non-Go File Notices

EOF
        emit_assets_sections
    } > "${OUTPUT_TMP}"
    # mktemp creates 0600, and mv within OUTPUT's directory is an atomic rename.
    chmod 644 "${OUTPUT_TMP}"
    mv -f "${OUTPUT_TMP}" "${OUTPUT}"
}

main() {
    check_prerequisites
    prepare_workspace

    collect_go_dependencies
    build_go_index
    build_assets_index
    compose_document

    local go_package_count bundled_file_count
    go_package_count=$(wc -l < "${GO_INDEX_FILE}" | tr -d ' ')
    bundled_file_count=$(wc -l < "${ASSETS_INDEX_FILE}" | tr -d ' ')
    log "Wrote ${OUTPUT} (${go_package_count} Go packages, ${bundled_file_count} bundled non-Go files)"
}

main "$@"
