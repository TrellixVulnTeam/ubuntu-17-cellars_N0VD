// Do not edit. Bootstrap copy of /tmp/go-20170517-18751-tjx0zu/go/src/cmd/compile/internal/mips/prog.go

//line /tmp/go-20170517-18751-tjx0zu/go/src/cmd/compile/internal/mips/prog.go:1
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips

import (
	"bootstrap/cmd/compile/internal/gc"
	"bootstrap/cmd/internal/obj"
	"bootstrap/cmd/internal/obj/mips"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [mips.ALAST & obj.AMask]gc.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the MIPS opcode.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// Integer
	mips.AADD & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUB & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AAND & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AOR & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AXOR & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ANOR & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMUL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.AMULU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.ADIV & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.ADIVU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead},
	mips.ASLL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRA & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASRL & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASGT & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASGTU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},

	mips.ACLZ & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightRead},
	mips.ACLO & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightRead},

	// Floating point.
	mips.AADDF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AADDD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ASUBD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMULF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AMULD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ADIVF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.ADIVD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	mips.AABSF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite},
	mips.AABSD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},
	mips.ANEGF & obj.AMask:    {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite},
	mips.ANEGD & obj.AMask:    {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},
	mips.ACMPEQF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPEQD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.ACMPGTF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPGTD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.ACMPGEF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead},
	mips.ACMPGED & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RegRead},
	mips.AMOVFD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVDF & obj.AMask:   {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVFW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVWF & obj.AMask:   {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVDW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.AMOVWD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCFW & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},
	mips.ATRUNCDW & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv},

	mips.ASQRTF & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite},
	mips.ASQRTD & obj.AMask: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},

	// Moves
	mips.AMOVB & obj.AMask:  {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVBU & obj.AMask: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVH & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVHU & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVW & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVF & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	mips.AMOVD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},

	// Conditional moves
	mips.ACMOVN & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | RightRdwr},
	mips.ACMOVZ & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | RightRdwr},
	mips.ACMOVT & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | RightRdwr},
	mips.ACMOVF & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | RightRdwr},

	// Conditional trap
	mips.ATEQ & obj.AMask: {Flags: gc.SizeL | gc.RegRead | gc.RightRead},
	mips.ATNE & obj.AMask: {Flags: gc.SizeL | gc.RegRead | gc.RightRead},

	// Atomic
	mips.ASYNC & obj.AMask: {Flags: gc.OK},
	mips.ALL & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite},
	mips.ASC & obj.AMask:   {Flags: gc.SizeL | LeftRdwr | gc.RightRead},

	// Jumps
	mips.AJMP & obj.AMask:  {Flags: gc.Jump | gc.Break},
	mips.AJAL & obj.AMask:  {Flags: gc.Call},
	mips.ABEQ & obj.AMask:  {Flags: gc.Cjmp},
	mips.ABNE & obj.AMask:  {Flags: gc.Cjmp},
	mips.ABGEZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABLTZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABGTZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABLEZ & obj.AMask: {Flags: gc.Cjmp},
	mips.ABFPF & obj.AMask: {Flags: gc.Cjmp},
	mips.ABFPT & obj.AMask: {Flags: gc.Cjmp},
	mips.ARET & obj.AMask:  {Flags: gc.Break},
	obj.ADUFFZERO:          {Flags: gc.Call},
	obj.ADUFFCOPY:          {Flags: gc.Call},
}

func proginfo(p *obj.Prog) gc.ProgInfo {
	info := progtable[p.As&obj.AMask]

	if info.Flags == 0 {
		gc.Fatalf("proginfo: unknown instruction %v", p)
	}

	if (info.Flags&gc.RegRead != 0) && p.Reg == 0 {
		info.Flags &^= gc.RegRead
		info.Flags |= gc.RightRead
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Sym != nil && (info.Flags&gc.LeftRead != 0) {
		info.Flags &^= gc.LeftRead
		info.Flags |= gc.LeftAddr
	}

	if p.As == mips.AMUL && p.To.Reg != 0 {
		info.Flags |= gc.RightWrite
	}

	return info
}
