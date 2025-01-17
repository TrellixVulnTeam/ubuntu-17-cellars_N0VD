/* Default linker script, for normal executables */
/* Copyright (C) 2014-2017 Free Software Foundation, Inc.
   Copying and distribution of this script, with or without modification,
   are permitted in any medium without royalty provided the copyright
   notice and this notice are preserved.  */
OUTPUT_FORMAT("coff-m68k")
OUTPUT_ARCH(m68k)
SECTIONS
{
.text :
	{
	  *(.text)
	  *(.strings)
   	  _etext = .;
	*(.data)
	_edata = .;
	*(.bss)
	*(COMMON)
	 _end = .;
	 }
}
