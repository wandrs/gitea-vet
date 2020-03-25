// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package imports

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "imports",
	Doc:  "check for import order.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		level := 0
		for _, im := range file.Imports {
			var lvl int
			val := im.Path.Value
			if strings.HasPrefix(val, "code.gitea.io") {
				lvl = 1
			} else if sliceHasPrefix(val, "xorm.io", "github.com") {
				lvl = 2
			} else {
				lvl = 3
			}

			if lvl < level {
				pass.Reportf(file.Pos(), "Imports are sorted wrong")
				break
			}
			level = lvl
		}
	}
	return nil, nil
}

func sliceHasPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
