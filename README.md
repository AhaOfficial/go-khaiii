# go-khaiii

## What is go-khaiii?

go-khaiii is [Khaiii](https://github.com/kakao/khaiii) binding for Golang.

## Sample code

- Create Model:
    - khaiii.Model.Create(rsc_dir string, opt_str string)
        - rsc_dir: a directory of dictionary resources (default: /usr/local/share/khaiii)
        - opt_str: options (JSON format)
- Morphological Analysis:
    - khaiii.Model.ParseV1(line string)
        - (input e.g.) 나는 회사에 간다
        - (output e.g.)
        [[나 NP] [는 JX] [회사 NNG] [에 JKB] [가 VV] [ㄴ다 EC]]
	- khaiii.Model.Parse(line string)
        - (input e.g.) 나는 회사에 간다
        - (output e.g.)
        [{나 NP} {는 JX} {회사 NNG} {에 JKB} {가 VV} {ㄴ다 EC}]
	- khaiii.Model.Nouns(line string)
        - (input e.g.) 나는 회사에 간다
        - (output e.g.)
        [나 회사]
	- khaiii.Model.Analyze(line string)
        - (input e.g.) 나는 회사에 간다
        - (output e.g.)
        [{나는 [{나 NP} {는 JX}]} {회사에 [{회사 NNG} {에 JKB}]} {간다 [{가 VV} {ㄴ다 EC}]}]
- Destroy Model:
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
	
	// ParseV1
	resultParseV1, err := model.ParseV1(sentence)
	if err != nil {
		fmt.Println("Parsing Error!")
		panic(err)
	}
	fmt.Println(resultParseV1)

	// Parse
	resultParse, err := model.Parse(sentence)
	if err != nil {
		fmt.Println("Parsing Error!")
		panic(err)
	}
	for _, m := range resultParse {
		fmt.Printf("(%s/%s) ", m.Lex, m.Tag)
	}
	fmt.Println()

	// Nouns
	resultNouns, err := model.Nouns(sentence)
	if err != nil {
		fmt.Println("Parsing Error!")
		panic(err)
	}
	fmt.Println(resultNouns)

	// Analyze
	resultAnalyze, err := model.Analyze(sentence)
	if err != nil {
		fmt.Println("Parsing Error!")
		panic(err)
	}
	for _, r := range resultAnalyze {
		fmt.Printf("%s\t", r.Word)
		for _, m := range r.Morphs {
			fmt.Printf("(%s/%s) ", m.Lex, m.Tag)
		}
		fmt.Println()
	}
}
```

# Install

Before using go-khaiii, need to install [Khaiii](https://github.com/kakao/khaiii).

Here is how to install go-khaiii with [Khaiii](https://github.com/kakao/khaiii) depending on OS (MacOS, Ubuntu Focal Fossa, and Amazon Linux).

### MacOS

```bash
$ brew install git cmake go

# Download go-khaiii
$ git clone https://github.com/AhaOfficial/go-khaiii.git

# Install khaiii
$ cd go-khaiii
$ sudo bash install_khaiii.sh

# Export variables
$ go get github.com/AhaOfficial/go-khaiii
```

### Ubuntu 20.04

```bash
$ sudo apt-get install -y git cmake golang-go language-pack-ko

# Download go-khaiii
$ git clone https://github.com/AhaOfficial/go-khaiii.git

# Install khaiii
$ cd go-khaiii
$ sudo bash install_khaiii.sh

# Export variables
$ export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"
$ go get github.com/AhaOfficial/go-khaiii
```

### Amazon Linux

```bash
$ yum install -y git golang gcc-c++ wget tar make python3 cmake3
$ ln -sf /usr/bin/cmake3 /usr/bin/cmake

# Download go-khaiii
$ git clone https://github.com/AhaOfficial/go-khaiii.git

# Install khaiii
$ cd go-khaiii
$ bash install_khaiii.sh

# Export variables
$ export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/lib"
$ go get github.com/AhaOfficial/go-khaiii
```
