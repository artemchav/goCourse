package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	source := flag.String("s", "", "source file")
	dest := flag.String("o", "", "destination file")

	flag.Parse()
	if *source == "" || *dest == "" {
		log.Fatal("Specify source file and destination file. Use --help to see usage")
		return
	}

	if *source == *dest {
		log.Fatal("=)")
	}

	filestat, err := os.Stat(*source)
	if err != nil {
		log.Fatal(err)
	}

	if !filestat.Mode().IsRegular() {
		log.Fatalf("can't use path '%s' as source file", *source)
	}

	from, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.Create(*dest)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully completed")
}
