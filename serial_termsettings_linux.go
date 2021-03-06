//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build linux

package serial

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func (port *unixPort) retrieveTermSettings() (*unix.Termios, error) {
	settings := new(unix.Termios)

	if err := ioctl(port.handle, unix.TCGETS, uintptr(unsafe.Pointer(settings))); err != nil {
		return nil, newOSError(err)
	}

	if settings.Cflag&unix.BOTHER == unix.BOTHER {
		if err := ioctl(port.handle, unix.TCGETS2, uintptr(unsafe.Pointer(settings))); err != nil {
			return nil, newOSError(err)
		}
	}
	
	return settings, nil
}

func (port *unixPort) applyTermSettings(settings *unix.Termios) error {
	req := uint64(unix.TCSETS)

	if settings.Cflag&unix.BOTHER == unix.BOTHER {
		req = unix.TCSETS2
	}

	if err := ioctl(port.handle, req, uintptr(unsafe.Pointer(settings))); err != nil {
		return newOSError(err)
	}
	return nil
}
