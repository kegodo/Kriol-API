// Filename: cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "1.0.0"

type config struct {
	port int,
	env string
}

func main(
	config config
	logger *log.logger
)

func main(){
	var cfg config 
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development | staging | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime),
}