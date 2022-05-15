package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:3000/get-token?login=root111&password=1111&data=21")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
