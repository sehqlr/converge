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

package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/asteris-llc/converge/load"
	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "graph the execution of a module",
	Long: `graphing is a convenient way to visualize how your graph and
dependencies are structured.

Pipe the output of this function into something like:

    dot -Tpng -oyourgraph.png`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Need one module filename as argument, got 0")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fname := args[0]

		graph, err := load.Load(fname, nil)
		if err != nil {
			log.Fatalf("[FATAL] %s: could not parse file: %s\n", fname, err)
		}

		out, err := graph.GraphString()
		if err != nil {
			log.Fatalf("[FATAL]: %s: could not render graph: %s\n", fname, err)
		}

		fmt.Println(out)
	},
}

func init() {
	RootCmd.AddCommand(graphCmd)
}