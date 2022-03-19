package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// Launch browser
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()
	defer pw.Stop()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	// Go to web
	if _, err := page.Goto("http://example.com", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// LOGIN PAGE
	// Obtain login frame
	el, err := page.WaitForSelector("html > frameset > frame:nth-child(2)")
	if err != nil {
		log.Fatalf("Frame not found: %s", err)
	}
	frame, err := el.ContentFrame()
	if err != nil {
		log.Fatalf("Element is not a frame: %s", err)
	}
	// Login: user > psswd > accept
	err = frame.Click("#txtusuario")
	if err != nil {
		log.Fatalf("could not click: %s", err.Error())
	}
	err = frame.Type("#txtusuario", "<YOUR USER>")
	if err != nil {
		log.Fatalf("could not type: %s", err.Error())
	}
	err = frame.Click("#txtContraseña")
	if err != nil {
		log.Fatalf("could not click: %s", err.Error())
	}
	err = frame.Type("#txtContraseña", "<YOUR PASSWORD>")
	if err != nil {
		log.Fatalf("could not type: %s", err.Error())
	}
	err = frame.Press("#btnAcpetar", "Enter")
	if err != nil {
		log.Fatalf("could not press: %s", err.Error())
	}

	// HOME PAGE
	// Go to machines state
	frame.Click("#lblMenu > ul > li:nth-child(5) > a")

	frame.WaitForSelector("#form1")
	txt, _ := frame.Content()
	fmt.Println(txt)
}
