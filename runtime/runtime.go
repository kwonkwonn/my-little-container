package runtime

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


func Runtime_cli(){
	logo()
   reader:= bufio.NewReader(os.Stdin)

	for{
		fmt.Println("Type 'help' to see all commands" )
		text,_:= reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch text{
		case "help":
			fmt.Println("list : list all containers")
			fmt.Println("run : run a new container")
			fmt.Println("exit : exit the program")
			fmt.Println("info : show container information")
			fmt.Println(" ")
		case "run":
			data,err := newContainer(reader)
			if err != nil {
				fmt.Println("Error creating container:", err)
				break
			}
			fmt.Println("Running container:", data)
		case "list":
			fmt.Println("list : list all containers")
		case "info":
			
			logo()
			fmt.Println("info : show container information")
		default:
			fmt.Println("Unknown command:", text)
			os.Exit(1)
	}
	time.Sleep(1 * time.Second)
}
}




func logo(){
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[33m" + `      ███████
   ██/     /█|
  █       / █|
 ████████/  █|
 █       █  █|
 █  my-  █  █|
 █ little █  █
 █container█ █
 ██████████ █`)

}