package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Service struct {
	Name        string
	Address     string
	Port        string
	AddressType string
}

func main() {

	services, err := GetServices()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(services)

}

func GetServices() ([]Service, error) {

	//avahi-browse -t _http._tcp -v -r -p
	cmd := exec.Command("avahi-browse", "-t", "_http._tcp", "-v", "-r", "-p")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	strOutput := string(stdoutStderr)
	rows := strings.Split(strOutput, "=")

	result := []Service{}

	for i,v := range rows {

		if i == 0 {
			continue
		}

		row := strings.Split(v, ";")
		service := Service{
			AddressType : row[2],
			Name : row[3],
			Address : row[7],
			Port : row[8],
		}
		
		result = append(result, service)

	}


	return  result, nil

}
