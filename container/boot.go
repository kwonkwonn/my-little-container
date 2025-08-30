package container

import (
	"fmt"
	"os"
)





func Initial_Setting(){
	_, err := os.Stat("/sys/fs/cgroup/my-little-container")
	if os.IsNotExist(err) {
		errr:= os.Mkdir("/sys/fs/cgroup/my-little-container", 0755)
		if errr != nil {
			fmt.Println("Error creating cgroup directory:", errr)
		} else {
			fmt.Println("cgroup directory created")
		}

		err = os.WriteFile("/sys/fs/cgroup/cgroup.subtree_control", []byte("+cpu +memory"), 0644)
		if err != nil {
		}

		err = os.WriteFile("/sys/fs/cgroup/my-little-container/cgroup.subtree_control", []byte("+cpu +memory"), 0644)
		if err != nil {
		}
	}

	

	
}