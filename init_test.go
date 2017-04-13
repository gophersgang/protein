// Copyright © 2016 Zenly <hello@zen.ly>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protein

import "go.uber.org/zap"

// -----------------------------------------------------------------------------

// init sets the global logger in dev mode before running tests.
func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.ReplaceGlobals(l)
}