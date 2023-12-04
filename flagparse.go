package main

import "flag"

func processInputArguments() Flags {
	var flags = Flags{}
	flag.BoolVar(&flags.c, "c", false, "a bool")
	flag.BoolVar(&flags.d, "d", false, "a bool")
	flag.BoolVar(&flags.u, "u", false, "a bool")
	flag.BoolVar(&flags.i, "i", false, "a bool")
	flag.IntVar(&flags.f, "f", 0, "an int")
	flag.IntVar(&flags.s, "s", 0, "an int")

	flag.Parse()

	return flags
}
