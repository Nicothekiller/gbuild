package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// struct for flags to compile
type CompileFlags struct {
	compiler  string
	compFlags []string
	ldFlags   string
}

// func to make a new struct
func newCompileFlags() CompileFlags {
	comp := "gcc"
	cflags := []string{"-Wall", "-Wextra", "-Wpedantic", "-pipe", "-O2"}
	ldflags := ""
	return CompileFlags{comp, cflags, ldflags}
}

// func for compiling with the given flags
func (flags *CompileFlags) compile() {
	files := getFiles()

	for _, v := range files {
		ndFile := strings.Split(v, "/")

		oFile := "bin/" + ndFile[len(ndFile)-1] + ".o"
		fFlags := []string{"-c", v, "-o", oFile}

		fFlags = append(fFlags, flags.compFlags...)

		compComand := exec.Command(flags.compiler, fFlags...)
		compComand.Stdout = os.Stdout
		compComand.Stderr = os.Stderr

		fmt.Print(flags.compiler, " ")
		fmt.Println(fFlags)

		err := compComand.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
