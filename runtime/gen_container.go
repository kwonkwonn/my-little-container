package runtime

import (
	"bufio"
	"fmt"

	"github.com/kwonkwonn/my-little-container/container"
)



func newContainer(reader *bufio.Reader)(container.Container_order, error){
	var containerOrder container.Container_order
	fmt.Println("Enter container name:")
	name,_:= reader.ReadString('\n')
	containerOrder.Name = name[:len(name)-1]
	fmt.Println("Enter executable path:")
	exec_path,_:= reader.ReadString('\n')
	containerOrder.Exec_path = exec_path[:len(exec_path)-1]
	
	fmt.Println("Enter CPU limit (e.g., 1000 for 30%):")
	cpu,_:= reader.ReadString('\n')
	containerOrder.ContainerConf = &container.Container_config{}
	containerOrder.ContainerConf.Cpu = cpu[:len(cpu)-1]
	if AtoI(containerOrder.ContainerConf.Cpu) > 30000{
		return container.Container_order{}, fmt.Errorf("CPU limit exceeds maximum of 30000")
	} 
	// fmt.Println("Enter Memory limit (e.g., 500m for 500MB):")
	// memory,_:= reader.ReadString('\n')
	// containerOrder.ContainerConf.Memory = memory[:len(memory)-1]

	// fmt.Println("Enter volume paths (comma separated):")
	// volumes,_:= reader.ReadString('\n') 
	// not implemented yet , volume, memory


	_, err := container.New_container(containerOrder)
	if err != nil {
		return container.Container_order{}, err
	}

	return containerOrder, nil
}