package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// struct for flags to compile
type CompileFlags struct {
	compiler  string
	compFlags []string
	ldFlags   []string
}

// func to make a new struct
func newCompileFlags() CompileFlags {
	comp := "g++"
	cflags := []string{"-Wall", "-Wextra", "-Wpedantic", "-pipe", "-O2"}
	ldflags := []string{}
	return CompileFlags{comp, cflags, ldflags}
}

// func for compiling with the given flags
func (flags *CompileFlags) compile() {
	files := getFiles()
	var wg sync.WaitGroup
	oFiles := []string{}

	for _, v := range files {
		wg.Add(1)
		go func() {
			ndFile := strings.Split(v, "/")

			oFile := "bin/" + ndFile[len(ndFile)-1] + ".o"

			oFiles = append(oFiles, oFile)

			fFlags := []string{"-c", v, "-o", oFile}

			fFlags = append(fFlags, flags.compFlags...)

			compComand := exec.Command(flags.compiler, fFlags...)
			compComand.Stdout = os.Stdout
			compComand.Stderr = os.Stderr

			fmt.Println(flags.compiler, fFlags)

			err := compComand.Run()
			if err != nil {
				log.Fatal(err)
			}

			defer wg.Done()
		}()
	}

	wg.Wait()

	cmdf := oFiles
	cmdf = append(cmdf, []string{"-o", "main"}...)
	cmdf = append(cmdf, flags.compFlags...)
	cmdf = append(cmdf, flags.ldFlags...)

	fmt.Println(flags.compiler, cmdf)

	exeCmd := exec.Command(flags.compiler, cmdf...)
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr

	err := exeCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
