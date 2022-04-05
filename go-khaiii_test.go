package go_khaiii

import (
	"fmt"
	"testing"
)

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
	}
	fmt.Println(result)
	model.Destroy()
}
