/*
Copyright 2020 Heath Harrelson.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package buildinfo

import "fmt"

var (
	// Version is the current version of suspenders. Set at build time.
	Version string

	// Commit is the short SHA of the Git commit used to make the current build. Set at build time.
	Commit string

	// BuildDate is the date the current build was created. Set at build time.
	BuildDate string
)

// Print returns formatted build information.
func Print() string {
	return fmt.Sprintf("Build info: commit=%s build_date=%s", Commit, BuildDate)
}
