/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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
**/

package nvcaps

import "fmt"

// Option defines a function for passing options to the New() call
type Option func(*nvcapslib)

// New creates a new instance of the 'nvcaps' interface
func New(opts ...Option) (Interface, error) {
	i := &nvcapslib{}
	for _, opt := range opts {
		opt(i)
	}
	if i.root == "" {
		i.root = "/"
	}
	if i.caps == nil {
		caps, err := NewMigCaps()
		if err != nil {
			return nil, fmt.Errorf("failed to construct default MIG capabilities: %v", err)
		}
		i.caps = caps
	}
	return i, nil
}

// WithRoot provides a Option to set the root of the 'nvcaps' interface
func WithRoot(root string) Option {
	return func(i *nvcapslib) {
		i.root = root
	}
}

// WithMigCaps provides an Option to set the mig caps of the 'nvcaps' interface
func WithMigCaps(caps MigCaps) Option {
	return func(i *nvcapslib) {
		i.caps = caps
	}
}
