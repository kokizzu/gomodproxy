package main

import (
	"fmt"
	"net/http"

	"github.com/goproxy/goproxy"
)

func main() {
	fmt.Println(`please execute this command:
go env -w GOPROXY=http://localhost:65432,direct`)
	fmt.Println(http.ListenAndServe("localhost:65432", &goproxy.Goproxy{}))
}
