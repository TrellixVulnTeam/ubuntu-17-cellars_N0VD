// Do not edit. Bootstrap copy of /tmp/go-20170517-18751-tjx0zu/go/src/cmd/compile/internal/syntax/printer_test.go

//line /tmp/go-20170517-18751-tjx0zu/go/src/cmd/compile/internal/syntax/printer_test.go:1
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	ast, err := ParseFile(*src, nil, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	Fprint(os.Stdout, ast, true)
	fmt.Println()
}
