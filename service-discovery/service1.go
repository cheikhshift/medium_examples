package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func BroadCastMe(port int, service string) {

	cmd := exec.Command("avahi-publish", "-s", service, "_http._tcp", "8080")
	if err := cmd.Run(); err != nil {
		// close program if broadcast fails
		log.Fatal(err)
	}

}

func main() {

	fmt.Println("listenning")

	// broadcast will stop
	// on program shutdown
	go BroadCastMe(8080, "service-one")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
