// +build windows

package lib

func FindExe() (string,string){
	return "C:\\Windows\\System32\\explorer.exe", "Windows"
}