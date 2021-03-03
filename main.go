package main

import (
	"github.com/breathbath/oauth-test/cli"
	"log"
)

func main() {
	log.Fatal(cli.StartServer())
}
