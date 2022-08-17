package main

import (
	"fmt"
	"os"
	"net/http"

	"testing"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func setup(t *testing.T) (*selenium.Service, selenium.WebDriver, *http.Server) {

	fsport := ":8090"
	server := http.Server{
		Addr : fsport,
		Handler : http.FileServer(http.Dir(".")),
	}
	
	go server.ListenAndServe()

	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "drivers/selenium-server.jar"
		chromePath = "drivers/chromedriver"
		port            = 8080
	)

	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(chromePath),
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		t.Fatal(err) // panic is used only as an example and is not otherwise recommended.
	}

	chrCaps := chrome.Capabilities{
		Path: "./drivers/chrome-linux/chrome",
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	caps.AddChrome(chrCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		t.Fatal(err)
	}


	return service, wd, &server
}

func TestSample(t *testing.T) {

	service, wd, server := setup(t)

	defer service.Stop()
	defer wd.Quit()
	defer server.Close()

	t.Log("Loading webpage")
	if err := wd.Get("http://localhost:8090/index.html"); err != nil {
		t.Fatal(err)
	}

	button, err := wd.FindElement(selenium.ByCSSSelector, "#action")
	if err != nil {
		t.Fatal(err)
	}

	err = button.Click()

	if err != nil {
		t.Fatal(err)
	}

	elem, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		t.Fatal(err)
	}

	expecting := "Hello World"

	if text,err := elem.Text(); text != expecting {
		
		if err != nil {
			t.Fatal(err)
		}

		t.Errorf("Unexpected output found. expected %v got %v",
			expecting, text)
	}

}
