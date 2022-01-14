package khaiii_test

import (
	"fmt"
	"testing"
	"tokenizer-khaiii/khaiii"
)

func TestCreate(t *testing.T) {
	var model khaiii.Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Wrong Create")
	} else {
		model.Destroy()
	}
}

func TestParse(t *testing.T) {
	var model khaiii.Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Create Error!")
	}
	result, error := model.Parse("나는 학교에 간다")
	if error != nil {
		t.Error("Parse Error!")
	} else {
		fmt.Println(result)
		model.Destroy()
	}
}

func TestDestroy(t *testing.T) {
	var model khaiii.Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Create Error!")
	}
	error = model.Destroy()
	if error != nil {
		t.Error("Destroy Error!")
	}
}
