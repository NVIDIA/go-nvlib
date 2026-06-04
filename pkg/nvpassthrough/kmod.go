/*
 * Copyright (c) NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nvpassthrough

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
	"golang.org/x/sys/unix"
)

// modInitFunc supports uncompressed files and gzip and xz compressed files.
//
// This code is taken from:
//
//	https://github.com/pmorjan/kmod/blob/d0ca1d5ed38616f2dc65c69add06b55f7cc091a7/cmd/modprobe/modprobe.go#L132
func modInitFunc(path, params string, flags int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	switch filepath.Ext(path) {
	case ".gz":
		rd, err := gzip.NewReader(f)
		if err != nil {
			return err
		}
		defer rd.Close()
		return initModule(rd, params)
	case ".xz":
		rd, err := xz.NewReader(f)
		if err != nil {
			return err
		}
		return initModule(rd, params)
	case ".zst":
		rd, err := zstd.NewReader(f)
		if err != nil {
			return err
		}
		defer rd.Close()
		return initModule(rd, params)
	}

	// uncompressed file, first try finitModule then initModule
	if err := finitModule(int(f.Fd()), params); err != nil {
		if err == unix.ENOSYS {
			return initModule(f, params)
		}
	}
	return nil
}

// finitModule inserts a module file via syscall finit_module(2).
func finitModule(fd int, params string) error {
	return unix.FinitModule(fd, params, 0)
}

// initModule inserts a module via syscall init_module(2).
func initModule(rd io.Reader, params string) error {
	buf, err := io.ReadAll(rd)
	if err != nil {
		return err
	}
	return unix.InitModule(buf, params)
}
