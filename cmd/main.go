package main

import "flag"

func main() {

	configfile := flag.String("config.file", "./configs/application.yml", "")
	flag.Parse()

	app := InitializeApp(*configfile)
	if app != nil {
		app.Run()
	}
}
