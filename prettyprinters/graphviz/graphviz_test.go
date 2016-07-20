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

package graphviz_test

import (
	"testing"

	"github.com/asteris-llc/converge/prettyprinters/graphviz"
	"github.com/stretchr/testify/assert"
)

var (
	emptyOptsMap = map[string]string{}
)

func Test_NewGraphvizPrinter_WhenMissingOptions_UsesDefaultOptions(t *testing.T) {

	printer, _ := graphviz.NewPrinter(emptyOptsMap, stubPrinter)
	actual := printer.Options()
	expected := graphviz.DefaultOptions()
	assert.Equal(t, actual, expected)
}

func Test_NewGraphvizPrinter_WhenProvidedOptions_UsesProvidedOptions(t *testing.T) {
	opts := map[string]string{"rankdir": "TB"} // not the default
	printer, _ := graphviz.NewPrinter(opts, stubPrinter)
	setOpts := printer.Options()
	rankdir := setOpts["rankdir"]
	assert.Equal(t, "TB", rankdir)
}

func stubPrinter(_ interface{}) (string, error) {
	return "", nil
}