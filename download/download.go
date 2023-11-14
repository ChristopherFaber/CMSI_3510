package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	f, err := os.Create("data.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.Write(bodyBytes)
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Create("bytes.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	f1.Write(bodyBytes)
	f1.Sync()

	f2, err := os.Create("string.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f2.Close()
	f2.WriteString(bodyString)
	f2.Sync()
	fmt.Println(bodyString)
}

func main() {
	url := os.Args[1]
	download(url)
}
