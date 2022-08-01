package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const FileTemplate = "package main\nimport (\n%s\n)\nfunc main(){ \n%s %s\n }"

type Terminal struct {
	iCache, vCache []string
}

func (t *Terminal) Process(cmd string) (response string, err error) {

	tokens := strings.Split(cmd, " ")
	fn := tokens[0]

	switch {
	case fn == "list":
		fmt.Println("Imports : ", t.iCache)
		fmt.Println("Variables : ", t.vCache)
	case fn == "import":
		path := strings.ReplaceAll(tokens[1], "\n", "")
		t.iCache = append(t.iCache, path)
		response = path + " imported!"
	case fn == "clear":
		t.iCache = nil
		t.vCache = nil
		response = "Terminal data cleared!"
	case strings.Contains(cmd, "="):
		t.vCache = append(t.vCache, cmd)
		response = "Variable saved"
	default:
		response, err = t.Build(cmd)
	}

	return
}

func (t *Terminal) Build(cmd string) (response string, err error) {

	imports := strings.Join(t.iCache, "\n")
	vars := strings.Join(t.vCache, "\n")
	fileOutput := "console.go"

	code := fmt.Sprintf(FileTemplate, imports, vars, cmd)

	err = os.WriteFile(fileOutput, []byte(code), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	cm := exec.Command("go", "run", fileOutput)
	stdout,_ := cm.CombinedOutput()

	//remove file right after execution
	os.Remove(fileOutput)
	response = string(stdout)


	return
}
