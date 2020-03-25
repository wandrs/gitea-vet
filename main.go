// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"gitea.com/jolheiser/gitea-vet/imports"
	"gitea.com/jolheiser/gitea-vet/license"
)

func main() {
	unitchecker.Main(
		license.Analyzer,
		imports.Analyzer,
	)
}
