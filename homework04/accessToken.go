package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll, FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func getAccessToken() {
	var url = "https://oauth.vk.com/authorize?client_id=" + clientID +
		"&redirect_uri=https://oauth.vk.com/blank.html" +
		"&scope=" + scope +
		"&response_type=token" +
		"&display=page"
	openBrowser(url)
}
