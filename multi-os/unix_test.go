//+build unix_test

package lib

import "testing"

func TestFindExe(t * testing.T){
	
	_, os := FindExe()

	unixOs := "Unix"
	if os != unixOs {
		t.Fatalf("Error, expected %v go %v", unixOs, os)
	}

}