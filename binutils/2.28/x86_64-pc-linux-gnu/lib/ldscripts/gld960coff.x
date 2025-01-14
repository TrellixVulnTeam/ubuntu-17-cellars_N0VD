/* Default linker script, for normal executables */
/* Copyright (C) 2014-2017 Free Software Foundation, Inc.
   Copying and distribution of this script, with or without modification,
   are permitted in any medium without royalty provided the copyright
   notice and this notice are preserved.  */
SECTIONS
{
    .text :
    {
	 CREATE_OBJECT_SYMBOLS
	*(.text)
	 _etext = .;

	___CTOR_LIST__ = .;
	LONG((___CTOR_END__ - ___CTOR_LIST__) / 4 - 2)
	*(.ctors)
	LONG(0)
	___CTOR_END__ = .;
	___DTOR_LIST__ = .;
	LONG((___DTOR_END__ - ___DTOR_LIST__) / 4 - 2)
	*(.dtors)
	LONG(0)
	___DTOR_END__ = .;
    }
    .data :
    {
 	*(.data)
	CONSTRUCTORS
	 _edata = .;
    }
    .bss :
    {
	 _bss_start = .;
	*(.bss)
	*(COMMON)
	 _end = .;
    }
}
