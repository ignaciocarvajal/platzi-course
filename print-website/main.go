package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorioWeb struct{}

// implementar interface Write de io.copy
func (escritorioWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	e := escritorioWeb{}
	io.Copy(e, res.Body)

}
