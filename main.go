package main

import (
  "fmt"
  "bufio"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "os"
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

  err  = yaml.Unmarshal(source, &config)
  if err != nil {
    panic(err)
  }
  // fmt.Println(config.Emails)

  hostFileHandler, err2 := os.OpenFile("/etc/hosts", os.O_APPEND, 0666)
  writer

}
