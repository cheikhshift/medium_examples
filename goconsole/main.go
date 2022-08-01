package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	term := Terminal{}
	console(&term)
}

func console(t *Terminal) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	cmd, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	res, err := t.Process(
		cmd,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	console(
		t,
	)
}
