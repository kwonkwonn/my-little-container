package container

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)


func New_container(config Container_order)(*Container, error){
	// cgroup v2 방식으로 변경
	newCon:= &Container{
		name: config.Name,
		exec_path: config.Exec_path,
		containerConf: config.ContainerConf,
		path: "/sys/fs/cgroup/my-little-container/"+config.Name,
	}

	cmd:= exec.Command(config.Exec_path)
	
	err:= cmd.Start()
	if err != nil {
			return nil, err
		}	
	newCon.current_pid = cmd.Process.Pid
	fmt.Println("New container PID:", newCon.current_pid)
	err = os.Mkdir(newCon.path, 0755)
	if err != nil {
		return nil, err
	}


	f, err := os.OpenFile(newCon.path+"/cgroup.procs", os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := f.Write([]byte(strconv.Itoa(newCon.current_pid))); err != nil {
		return nil, err
	}

	cpuTime := fmt.Sprintf("%d %d", newCon.containerConf.Cpu, 100000)
	f2, err := os.OpenFile(newCon.path+"/cpu.max", os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f2.Close()

	if _, err := f2.Write([]byte(cpuTime)); err != nil {
		return nil, err
	}

	return newCon,nil
}

