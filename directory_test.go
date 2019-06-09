package file

import (
	`fmt`
	`testing`
)

func TestDirList(t *testing.T) {
	path,_ := CurrentDir()
	list ,_ := DirList(path)
	fmt.Println(list)
}
