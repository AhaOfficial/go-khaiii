package go_khaiii

import (
	"fmt"
	"testing"
)

func TestAnalyze(t *testing.T) {
	fmt.Printf("\nAnalyze\n====================\n")
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.Analyze("나는 회사에 간다")
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	}
	fmt.Println(result)
	for _, r := range result {
		fmt.Printf("%s\t", r.Word)
		for _, m := range r.Morphs {
			fmt.Printf("(%s/%s) ", m.Lex, m.Tag)
		}
		fmt.Println()
	}
	model.Destroy()
}

func TestNouns(t *testing.T) {
	fmt.Printf("\nNouns\n====================\n")
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.Nouns("나는 회사에 간다")
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	}
	fmt.Println(result)
	model.Destroy()
}

func TestParse(t *testing.T) {
	fmt.Printf("\nParse\n====================\n")
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.Parse("나는 회사에 간다")
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	}
	fmt.Println(result)
	for _, m := range result {
		fmt.Printf("(%s/%s) ", m.Lex, m.Tag)
	}
	fmt.Println()
	model.Destroy()
}

func TestParseV1(t *testing.T) {
	fmt.Printf("\nParseV1\n====================\n")
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.ParseV1("나는 회사에 간다")
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	}
	fmt.Println(result)
	model.Destroy()
}
