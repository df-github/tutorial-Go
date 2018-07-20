package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type CustomWriter struct {
	Content string
}

func NewCustomWriter() *CustomWriter{
	return &CustomWriter{}
}

func (instance *CustomWriter) Write(p []byte) (n int, err error){
	instance.Content = string(p)
	return len(p), nil
}

func main()  {
	response, err := http.Get("http://www.google.de")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	customWriter := NewCustomWriter()
	io.Copy(customWriter, response.Body)

	fmt.Println(customWriter.Content)
}
