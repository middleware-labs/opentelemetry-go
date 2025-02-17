// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

package resource // import "github.com/middleware-labs/otel/sdk/resource"

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type unameProvider func(buf *unix.Utsname) (err error)

var defaultUnameProvider unameProvider = unix.Uname

var currentUnameProvider = defaultUnameProvider

func setDefaultUnameProvider() {
	setUnameProvider(defaultUnameProvider)
}

func setUnameProvider(unameProvider unameProvider) {
	currentUnameProvider = unameProvider
}

// platformOSDescription returns a human readable OS version information string.
// The final string combines OS release information (where available) and the
// result of the `uname` system call.
func platformOSDescription() (string, error) {
	uname, err := uname()
	if err != nil {
		return "", err
	}

	osRelease := osRelease()
	if osRelease != "" {
		return fmt.Sprintf("%s (%s)", osRelease, uname), nil
	}

	return uname, nil
}

// uname issues a uname(2) system call (or equivalent on systems which doesn't
// have one) and formats the output in a single string, similar to the output
// of the `uname` commandline program. The final string resembles the one
// obtained with a call to `uname -snrvm`.
func uname() (string, error) {
	var utsName unix.Utsname

	err := currentUnameProvider(&utsName)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s %s %s",
		unix.ByteSliceToString(utsName.Sysname[:]),
		unix.ByteSliceToString(utsName.Nodename[:]),
		unix.ByteSliceToString(utsName.Release[:]),
		unix.ByteSliceToString(utsName.Version[:]),
		unix.ByteSliceToString(utsName.Machine[:]),
	), nil
}

// getFirstAvailableFile returns an *os.File of the first available
// file from a list of candidate file paths.
func getFirstAvailableFile(candidates []string) (*os.File, error) {
	for _, c := range candidates {
		file, err := os.Open(c)
		if err == nil {
			return file, nil
		}
	}

	return nil, fmt.Errorf("no candidate file available: %v", candidates)
}
