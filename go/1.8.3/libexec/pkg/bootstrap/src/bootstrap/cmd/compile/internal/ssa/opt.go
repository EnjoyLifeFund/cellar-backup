// Do not edit. Bootstrap copy of /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/ssa/opt.go

//line /private/tmp/go-20170524-61097-1ppddl4/go/src/cmd/compile/internal/ssa/opt.go:1
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// machine-independent optimization
func opt(f *Func) {
	applyRewrite(f, rewriteBlockgeneric, rewriteValuegeneric)
}

func dec(f *Func) {
	applyRewrite(f, rewriteBlockdec, rewriteValuedec)
	if f.Config.IntSize == 4 && f.Config.arch != "amd64p32" {
		applyRewrite(f, rewriteBlockdec64, rewriteValuedec64)
	}
}
