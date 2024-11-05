package main

func main() {
	strFiles := getFiles()

	println("Found", len(strFiles), "files: ")
	for _, v := range strFiles {
		print(v, " ")
	}
	println()

	comp := newCompileFlags()
	comp.compile()
}
