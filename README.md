funcstats: simple function stats for C and Go programs.

This program only supports properly-written C code, i.e. that written to
the BSD KNF style.

## EXAMPLE

```
$ funcstats ./funcstats.go
Function stats for funcstats.go:
	7 functions counted
	scanFuncs is the longest function with 59 lines
	displayFuncStats is the shortest function with 12 lines
	The mean number of lines per function is 27
	The median number of lines per function is 23
Line count stats for funcstats.go:
	186 lines counted
	validLangMap has the longest line with 75 characters
	(global) has shortest line with 5 characters
	The mean lines length in this file is 26
	The median line length in this file is 23
Functions in funcstats.go:
	main: 49 lines
	computeStats: 27 lines
	displayFuncStats: 12 lines
	displayLineStats: 12 lines
	validLangMap: 12 lines
	countLineStats: 23 lines
	scanFuncs: 59 lines
$ funcstats ~/code/c/libschannel/src/*.c
Function stats for /home/kyle/code/c/libschannel/src/schannel.c:
	18 functions counted
	schannel_listen is the longest function with 74 lines
	schannel_send is the shortest function with 1 lines
	The mean number of lines per function is 32
	The median number of lines per function is 31
Line count stats for /home/kyle/code/c/libschannel/src/schannel.c:
	688 lines counted
	schannel_listen has the longest line with 80 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 27
	The median line length in this file is 20
Functions in /home/kyle/code/c/libschannel/src/schannel.c:
	_schannel_send: 40 lines
	schannel_zero: 18 lines
	schannel_listen: 74 lines
	schannel_send: 1 lines
	schannel_init: 17 lines
	validate_keys: 33 lines
	reset_counters: 11 lines
	generate_keypair: 20 lines
	do_kex: 41 lines
	schannel_dial: 72 lines
	unpack_message: 31 lines
	initialise_schannel: 11 lines
	sign_kex: 22 lines
	schannel_recv: 66 lines
	schannel_rekey: 38 lines
	verify_kex: 20 lines
	schannel_recv_kex: 45 lines
	schannel_close: 19 lines
$ funcstats ~/code/c/kam/*.c
Function stats for /home/kyle/code/c/kam/compiler.c:
	8 functions counted
	parse_next is the longest function with 66 lines
	destroy_compiler is the shortest function with 3 lines
	The mean number of lines per function is 27
	The median number of lines per function is 23
Line count stats for /home/kyle/code/c/kam/compiler.c:
	327 lines counted
	backtrace has the longest line with 82 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 25
	The median line length in this file is 19
Functions in /home/kyle/code/c/kam/compiler.c:
	parse: 19 lines
	write_program: 28 lines
	backtrace: 23 lines
	main: 52 lines
	init_compiler: 10 lines
	destroy_compiler: 3 lines
	finalise_number: 18 lines
	parse_next: 66 lines

---
Function stats for /home/kyle/code/c/kam/kamvm.c:
	3 functions counted
	run is the longest function with 40 lines
	interpreter is the shortest function with 2 lines
	The mean number of lines per function is 15
	The median number of lines per function is 4
Line count stats for /home/kyle/code/c/kam/kamvm.c:
	69 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 27
	The median line length in this file is 19
Functions in /home/kyle/code/c/kam/kamvm.c:
	interpreter: 2 lines
	run: 40 lines
	main: 4 lines

---
Function stats for /home/kyle/code/c/kam/qtest.c:
	2 functions counted
	test_queue8 is the longest function with 36 lines
	main is the shortest function with 3 lines
	The mean number of lines per function is 19
	The median number of lines per function is 36
Line count stats for /home/kyle/code/c/kam/qtest.c:
	62 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 28
	The median line length in this file is 19
Functions in /home/kyle/code/c/kam/qtest.c:
	test_queue8: 36 lines
	main: 3 lines

---
Function stats for /home/kyle/code/c/kam/queue.c:
	5 functions counted
	enqueue8 is the longest function with 27 lines
	empty_queue8 is the shortest function with 1 lines
	The mean number of lines per function is 11
	The median number of lines per function is 9
Line count stats for /home/kyle/code/c/kam/queue.c:
	79 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 23
	The median line length in this file is 17
Functions in /home/kyle/code/c/kam/queue.c:
	new_queue8: 9 lines
	enqueue8: 27 lines
	dequeue8: 15 lines
	empty_queue8: 1 lines
	destroy_queue8: 3 lines

---
Function stats for /home/kyle/code/c/kam/stack.c:
	10 functions counted
	pop16 is the longest function with 12 lines
	empty16 is the shortest function with 1 lines
	The mean number of lines per function is 7
	The median number of lines per function is 7
Line count stats for /home/kyle/code/c/kam/stack.c:
	128 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 21
	The median line length in this file is 17
Functions in /home/kyle/code/c/kam/stack.c:
	new_stack8: 7 lines
	pop8: 12 lines
	destroy8: 7 lines
	pop16: 12 lines
	destroy16: 7 lines
	push8: 11 lines
	empty8: 1 lines
	new_stack16: 7 lines
	push16: 11 lines
	empty16: 1 lines

---
Function stats for /home/kyle/code/c/kam/stest.c:
	3 functions counted
	test_stack16 is the longest function with 38 lines
	main is the shortest function with 4 lines
	The mean number of lines per function is 26
	The median number of lines per function is 36
Line count stats for /home/kyle/code/c/kam/stest.c:
	99 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 26
	The median line length in this file is 18
Functions in /home/kyle/code/c/kam/stest.c:
	test_stack8: 36 lines
	test_stack16: 38 lines
	main: 4 lines

---
Function stats for /home/kyle/code/c/kam/vm.c:
	6 functions counted
	vm_step is the longest function with 47 lines
	vm_destroy is the shortest function with 7 lines
	The mean number of lines per function is 18
	The median number of lines per function is 15
Line count stats for /home/kyle/code/c/kam/vm.c:
	163 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 24
	The median line length in this file is 17
Functions in /home/kyle/code/c/kam/vm.c:
	eval: 24 lines
	vm_step: 47 lines
	vm_run: 8 lines
	vm_peek: 12 lines
	vm_new: 15 lines
	vm_destroy: 7 lines

---
Function stats for /home/kyle/code/c/kam/vmtest.c:
	3 functions counted
	test2 is the longest function with 37 lines
	main is the shortest function with 6 lines
	The mean number of lines per function is 25
	The median number of lines per function is 32
Line count stats for /home/kyle/code/c/kam/vmtest.c:
	72 lines counted
	(global) has the longest line with 75 characters
	(global) has shortest line with 2 characters
	The mean lines length in this file is 28
	The median line length in this file is 19
Functions in /home/kyle/code/c/kam/vmtest.c:
	test1: 32 lines
	test2: 37 lines
	main: 6 lines
```

The examples are run over this source code on the first commit,
the [libschannel](https://github.com/kisom/libschannel/) repo,
and the [kam](https://github.com/kisom/kam/) repo.


## LICENSE

`funcstats` is released under the ISC license.

