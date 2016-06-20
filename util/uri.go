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

package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Retrieve takes a URL with some protocol (default is file://) and returns a
// reader of the data at that location.
func Retrieve(url *url.URL) (rdr io.ReadCloser, err error) {
	if url.Scheme == "" {
		url.Scheme = "file"
	}
	switch url.Scheme {
	case "file":
		return os.Open(url.Path)
	case "http", "https":
		resp, err := http.Get(url.String())
		return resp.Body, err
	default:
		return nil, fmt.Errorf("protocol %q is not implemented", url.Scheme)
	}
}
