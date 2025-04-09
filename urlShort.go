package main

import (
	"bytes"
	"path/filepath"
	"fmt"
	"os"
	"net/http"
	"urlshort/urlshort"
	"flag"
	"log"
)

var pathsFile = flag.String("pathsFile", "test/example.yml", "The file containing the shortened paths / URLs")

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file.")
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file")
	}

	return buf.Bytes()
}

// Creates default response for base directory
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

// Function for HandleFunc to call
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w , "Welcome to urlShortener, my first Go project")
}

func main(){
	mux := defaultMux()
	
	flag.Parse()

	ext := filepath.Ext(*pathsFile)

	var handler http.Handler
	var err error
	if ext == ".yml" {
		handler, err = urlshort.YAMLHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(getFileBytes(*pathsFile), mux)
		if err != nil {
			panic(err)
		}
	} else {
		log.Fatal("Paths file needs to be YAML")
	}
	
	fmt.Println("Starting server on port :80")
	err = http.ListenAndServe(":80", handler)
	fmt.Println(err)
}




