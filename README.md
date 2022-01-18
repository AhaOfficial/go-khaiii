# go-khaiii

## What is go-khaiii?

go-khaiii is Khaiii binding for Golang.

## Sample code

- 모델 생성:
    - khaiii.Model.Create(rsc_dir string, opt_str string)
- 형태소 분석:
    - khaiii.Model.Parse(line string)
        - (입력 예시) 나는 회사에 간다
        - (출력 예시)
        [ ["나", "NP"], ["는", "JX"],
           ["회사", "NNG"], ["에", "JKB"], 
           ["가", "VV"], ["ㄴ다", "EC"] ]
- 모델 소멸:
    - khaiii.Model.Destroy()

```go
package main

import (
	"fmt"

	khaiii "github.com/AhaOfficial/go-khaiii"
)

func main() {
	var model khaiii.Model

	err := model.Create("", "")
	if err != nil {
		fmt.Println("Create Error!")
		panic(err)
	}
	defer model.Destroy()

	var sentence string = "나는 회사에 간다"
	parsedSentence, err := model.Parse(sentence)
	if err != nil {
		fmt.Println("Parsing Error!")
		panic(err)
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

# move libraries to /usr/local/lib
$ mv libkhaiii*.* /usr/local/lib

# Export Variables
$ export CGO_LDFLAGS="-L/usr/local/lib -lkhaiiic"
$ go get github.com/AhaOfficial/go-khaiii
```