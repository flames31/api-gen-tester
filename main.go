package main

import (
	"github.com/flames31/api-gen-tester/cmd"
	"github.com/flames31/api-gen-tester/internal/log"
)

func main() {
	cmd.Execute()
	defer log.Sync()
}
