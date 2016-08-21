// Copyright © 2016 Asteris, LLC
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

package context

import "github.com/asteris-llc/converge/resource"

const DefaultContext = Context{}

// Context provides various contextual values that can be applied
// to many kinds of resources.
type Context struct {
	WorkDir string
}

// Check doesn't really do anything, since Contexts are final.
func (c *Context) Check() (resource.TaskStatus, error) {
	return &resource.Status{Status: c.String()}, nil
}

// Apply doesn't really do anything, since Contexts are final.
func (*Context) Apply() error {
	return nil
}

// String is the final value of the Param
func (c *Context) String() string {
	return fmt.Sprintf("Work dir: %s", c.WorkDir)
}
