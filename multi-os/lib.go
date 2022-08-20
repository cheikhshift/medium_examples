package lib

import (
	"os"
)

type Exe struct {
	Path string
	OS string
}


func GetExe() ( Exe, error) {

	var result Exe
	path,sys := FindExe()

	if _, err := os.Stat(path); err != nil {
        return result,nil
   	} 

   	result = Exe{ path, sys}
	return result,nil

}