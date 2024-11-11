package cmd1
 
 type Task struct{
	Name string `yaml: "name"`
	Params map[string]string `yaml: "params"`
	Type string `yaml:type`
 
 }
 
 type Proccess struct{
	Tasks []Task `yaml: tasks`
 }

 