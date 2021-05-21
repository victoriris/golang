package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	lg := logWriter{}

	io.Copy(lg, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("wrote", len(bs), "bytes")
	return len(bs), nil
}
