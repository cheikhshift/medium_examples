package fuzzing

import (
	"testing"
)

func FuzzParseString(f *testing.F){
	f.Add("1, Name, 23rd Fifth Street")
	f.Add("2, Name2, 1 Main street")

	f.Fuzz(func(t *testing.T, s string){
		err := ParseString(s)

		if err != nil {

			if err.Error() == "Invalid string" {
				return
			}

			t.Errorf("%v", err)
		}
		
	})
}

