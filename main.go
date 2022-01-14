package main

import (
	"fmt"
	"os"

	"go-khaiii/khaiii"
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
