//+build windows_test

package lib

import "testing"

func TestFindExe(t * testing.T){
	
	_, os := FindExe()

	windows := "Windows"
	if os != windows {
		t.Fatalf("Error, expected %v go %v", windows, os)
	}

}