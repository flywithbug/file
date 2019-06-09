package main

import (
	`fmt`
	`github.com/flywithbug/file`
)

func main()  {
	path,_ := file.CurrentDir()
	list ,_ := file.DirList(path)
	fmt.Println(list)
}
