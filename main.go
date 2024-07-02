package main

import (
	"flag"

	"github.com/ganimtron-10/EquityFusion/api"
)

func main() {
	port := flag.String("port", "8080", "Port at which the Web Server should start listening.")
	flag.Parse()

	api.CreateAndInitServer(*port)
}
