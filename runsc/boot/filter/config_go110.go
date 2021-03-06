// Copyright 2018 Google Inc.
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

// +build !go1.11

package filter

import (
	"syscall"

	"gvisor.googlesource.com/gvisor/pkg/seccomp"
)

// TODO: Remove this file and merge config_go111.go back into
// config.go once we no longer build with Go 1.10.

func init() {
	allowedSyscalls[syscall.SYS_PSELECT6] = []seccomp.Rule{}
}
