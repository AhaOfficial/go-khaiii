package go_khaiii

import (
	"fmt"
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
