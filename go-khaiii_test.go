package go_khaiii

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	var model Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Wrong Create")
	} else {
		defer model.Destroy()
	}
}

func TestParse(t *testing.T) {
	var model Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Create Error!")
		return
	}
	result, error := model.Parse("나는 회사에 간다")
	if error != nil {
		t.Error("Parse Error!")
		return
	} else {
		fmt.Println(result)
		defer model.Destroy()
	}
}

func TestDestroy(t *testing.T) {
	var model Model
	error := model.Create("", "")
	if error != nil {
		t.Error("Create Error!")
		return
	}
	error = model.Destroy()
	if error != nil {
		t.Error("Destroy Error!")
		return
	}
}
