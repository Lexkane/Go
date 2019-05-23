package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type Nichego struct {
	Str []byte
	idx int
}

func (ni *Nichego) Read(b []byte) (n int, err error) {
	if ni.idx == len(ni.Str) {
		return 0, io.EOF
	}
	//if len(b) < ... не влезает?,то возращаем error
	b[0] = ni.Str[ni.idx]
	ni.idx++
	return 1, nil
}

func main() {
	r := &Nichego{Str: []byte("string")}
	num, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(num))
}
