package main

import "flag"

func main() {
	var port int
	var filesBaseDir string = ""

	flag.IntVar(&port, "port", 4221, "Port to listen on")
	flag.StringVar(&filesBaseDir, "directory", "", "Base directory for files")
	flag.Parse()

	s := NewServer(port, filesBaseDir)
	s.Start()
}
