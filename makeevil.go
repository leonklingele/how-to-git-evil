package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	file     = "./evil"
	backchar = "\b"
)

func main() {
	evil := []byte(`console.log("EVIL");//`)
	sep := bytes.Repeat([]byte(backchar), len(evil))
	good := []byte(`console.log("Good");`)
	pl := fmt.Sprintf("%s%s%s\n", evil, sep, good)

	if err := ioutil.WriteFile(file, []byte(pl), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s written\n", file)
}
