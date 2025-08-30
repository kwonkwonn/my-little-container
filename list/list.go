package list

import "github.com/kwonkwonn/my-little-container/container"


func New_list()(*container_list){
	newList:= &container_list{
		containers: make([]container.Container,0),
	}
	return newList
}

type container_list struct {
	containers []container.Container
}

