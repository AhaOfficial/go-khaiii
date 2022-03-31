# go-khaiii

## What is go-khaiii?

go-khaiii is [Khaiii](https://github.com/kakao/khaiii) binding for Golang.

## Sample code

- 모델 생성:
    - khaiii.Model.Create(rsc_dir string, opt_str string)
        - rsc_dir: 리소스 사전의 경로 (default: /usr/local/share/khaiii)
        - opt_str: 옵션 (JSON 포맷)
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

	err := model.Create("", "") // "": default
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

### MacOS

```bash
$ brew install git cmake qt go

# Install for go-khaiii libraries
$ git clone https://github.com/AhaOfficial/go-khaiii.git
$ cd go-khaiii
$ make
$ sudo make install

# Export variables
$ export CGO_LDFLAGS="-L/usr/local/lib -lkhaiiic"
$ go get github.com/AhaOfficial/go-khaiii
```

### Ubuntu 20.04

```bash
$ sudo apt-get install -y git cmake qt5-default golang-go language-pack-ko

# Install for go-khaiii libraries
$ git clone https://github.com/AhaOfficial/go-khaiii.git
$ cd go-khaiii
$ make
$ sudo make install

# Export variables
$ export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"
$ export CGO_LDFLAGS="-L/usr/local/lib -lkhaiiic"
$ go get github.com/AhaOfficial/go-khaiii
```

### Amazon Linux

```bash
$ yum install -y git qt-devel golang gcc-c++ wget tar make python3
$ wget https://github.com/Kitware/CMake/releases/download/v3.23.0/cmake-3.23.0-linux-$(echo `arch`).sh
$ bash ./cmake-3.23.0-linux-$(echo `arch`).sh --skip-license
$ ln -s /usr/bin/sh /bin/sh

# Install for go-khaiii libraries
$ git clone https://github.com/AhaOfficial/go-khaiii.git
$ cd go-khaiii
$ make
$ make install

# Export variables
$ export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"
$ export CGO_LDFLAGS="-L/usr/local/lib -lkhaiiic"
$ go get github.com/AhaOfficial/go-khaiii
```
