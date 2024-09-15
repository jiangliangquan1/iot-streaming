package main

import (
	"flag"
)

func main() {
	//
	//s1, _ := commons.AESEncrypt("1234546", "jlqiot-streaming")
	//s2, _ := commons.EncryptAES("1234546", "jlqiot-streaming")
	//
	//if s1 == s2 {
	//
	//}
	//
	//commons.AESDecrypt(s1, "jlqiot-streaming")
	//commons.AESDecrypt(s2, "jlqiot-streaming")

	configfile := flag.String("config.file", "./configs/application.yml", "")
	flag.Parse()

	app := InitializeApp(*configfile)
	if app != nil {
		app.Run()
	}
}
