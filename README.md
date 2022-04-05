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
