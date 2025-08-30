package container

type Container struct {
	current_pid  int
	name         string
	path        string
	exec_path   string
	containerConf *Container_config
}


type Container_config struct {
	Volumes  []string
	Cpu      string // 300000 = total
	Memory   string
}

type Container_order struct {
	Name        string
	Exec_path  string
	ContainerConf *Container_config
}

