// Copyright 2023 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

//go:build js

package vfs

import (
	"os"

	"github.com/cockroachdb/errors"
)

func wrapOSFileImpl(f *os.File) File {
	panic("Not implemented")
}

func (defaultFS) OpenDir(name string) (File, error) {
	return nil, errors.New("Unimplemented")
}

type jsFile struct {
	*os.File
}

func (*jsFile) Prefetch(offset int64, length int64) error { return nil }
func (*jsFile) Preallocate(offset, length int64) error    { return nil }

func (f *jsFile) SyncData() error {
	return errors.New("Unimplemented")
}

func (f *jsFile) SyncTo(int64) (fullSync bool, err error) {
	return false, errors.New("Unimplemented")
}
