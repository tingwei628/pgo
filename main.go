package main

import (
	"fmt"
	"pgo/remote_command"
)

func main() {
	rc := remote_command.RC{}
	fmt.Println(rc.Command())
}
