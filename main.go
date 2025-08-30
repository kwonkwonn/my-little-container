package main

import (
	"fmt"
	"os"

	"github.com/kwonkwonn/my-little-container/container"
	"github.com/kwonkwonn/my-little-container/runtime"
)


func main(){
	fmt.Println("Hello, World!")
	container.Initial_Setting()
	runtime.Runtime_cli()

	os.Exit(0)
}