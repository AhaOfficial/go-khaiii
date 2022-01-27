package go_khaiii

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Wrong Create")
		panic(err)
	} else {
		err = model.Destroy()
		if err != nil {
			t.Error("Destroy Error!")
			panic(err)
		}
	}
}

func TestParse(t *testing.T) {
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
	} else {
		fmt.Println(result)
		err = model.Destroy()
		if err != nil {
			t.Error("Destroy Error!")
			panic(err)
		}
	}
}

func TestParseCustom_without_consonant(t *testing.T) {
	var model Model
	sentence := "나는 회사에 갔습니다"
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.Parse(sentence, map[string]string{"--no-consonant": ""})
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	} else {
		fmt.Println(result)
		err = model.Destroy()
		if err != nil {
			t.Error("Destroy Error!")
			panic(err)
		}
	}
}

func TestParseCustom(t *testing.T) {
	var model Model
	sentence := "나는 회사에 간다"
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		panic(err)
	}
	result, err := model.Parse(sentence, map[string]string{"--no-consonant": ""})
	if err != nil {
		t.Error("Parse Error!")
		panic(err)
	} else {
		var resultSentence string
		for _, morphs := range result {
			resultSentence += morphs[0]
		}
		if resultSentence != strings.ReplaceAll(sentence, " ", "") {
			panic(errors.New("Result is Incorrect!"))
		}
		fmt.Println(result)
		err = model.Destroy()
		if err != nil {
			t.Error("Destroy Error!")
			panic(err)
		}
	}
}

func TestDestroy(t *testing.T) {
	var model Model
	err := model.Create("", "")
	if err != nil {
		t.Error("Create Error!")
		return
	}
	err = model.Destroy()
	if err != nil {
		t.Error("Destroy Error!")
		return
	}
}
