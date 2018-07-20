package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main()  {
	response, err := http.Get("http://www.google.de")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}


	//buffer := make([]byte, 99999)
	//response.Body.Read(buffer)
	//fmt.Println("Response:",string(buffer))

	io.Copy(os.Stdout, response.Body)
}
