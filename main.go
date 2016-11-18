package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Emails []string
}

func main() {
	filename := os.Args[1]
	var config Config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	hostFileHandler, err2 := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err2 != nil {
		panic(err2)
	}
	defer hostFileHandler.Close()

	for _, email := range config.Emails {
		s := fmt.Sprintf("216.35.221.76 %s\r\n", email)
		_, er := hostFileHandler.WriteString(s)
		if er != nil {
			panic(er)
		}
	}

}
