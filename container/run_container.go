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
	
	// 프로세스가 종료될 때 좀비가 되지 않도록 고루틴에서 대기
	go func() {
		cmd.Wait() // 프로세스 종료를 기다림 (좀비 방지)
		fmt.Printf("Container process %d has terminated\n", newCon.current_pid)
	}()
	
	err = os.Mkdir(newCon.path, 0755)
	if err != nil {
		return nil, err
	}

	err= os.WriteFile(newCon.path+"/cgroup.procs", []byte(strconv.Itoa(newCon.current_pid)), 0644)
	if err != nil {
		fmt.Printf("Error moving process %d to cgroup: %v\n", newCon.current_pid, err)
		return nil, err
	}
	
	// 실제로 cgroup에 들어갔는지 확인
	fmt.Printf("Attempting to move PID %d to cgroup %s\n", newCon.current_pid, newCon.path)
	
	err = os.WriteFile(newCon.path+"/cpu.weight", []byte(newCon.containerConf.Cpu), 0644)
	if err != nil {
		return nil, err
	}
	
	return newCon,nil
}

