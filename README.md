# go-khaiii

## What is go-khaiii?

go-khaiii is Khaiii binding for Golang.

## Sample code

```go
package main

import (
	"fmt"
	"os"

	khaiii "github.com/AhaOfficial/go-khaiii"
)

func main() {
	var model khaiii.Model

	error := model.Create("", "")
	if error != nil {
		fmt.Println("Create Error!")
		os.Exit(1)
	}
	defer model.Destroy()

	// input: 나는 학교에 간다
	// output: [["나", "NP"], ["는", "JX"], ["회사", "NNG"], ["에", "JKB"], ["가", "VV"], ["ㄴ다", "EC"]]
	var sentence string = "나는 회사에 간다"
	parsedSentence, error := model.Parse(sentence)
	if error != nil {
		fmt.Println("Parsing Error!")
	}
	fmt.Println(sentence)
	fmt.Println(parsedSentence)
}
```

# Install

```bash
$ brew install git cmake qt go # for MacOS

# Install for building C Library
$ git clone https://github.com/AhaOfficial/go-khaiii.git
$ cd go-khaiii
$ .github/install_base.sh
$ .github/install_libc_for_go.sh

# move libkhaiiic.* to /path/to/lib
$ mv libkhaiiic.* /path/to/lib

# move khaiiic.h to /path/to/include
$ mv khaiiic.h /path/to/include

# Export Variables
$ export CGO_LDFLAGS="-L/path/to/lib -lkhaiiic"
$ export CGO_CFLAGS="-I/path/to/include"
$ go get github.com/AhaOfficial/go-khaiii
```