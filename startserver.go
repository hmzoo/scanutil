package main

import (
  "./webserver"
	"fmt"
)


func main() {
fmt.Println("STARTING WEBSERVER")
webserver.Serve()
}
