package go_khaiii

/*
#include <khaiiic.h>
*/
import "C"
import (
	"errors"
	"strings"
	"unsafe"
)

type Model struct {
	model *C.khaiii_model_t
}

type Morph struct {
	lex string
	tag string
}

func (m *Model) Create(rsc_dir string, opt_str string) error {
	c_rsc_dir := C.CString(rsc_dir)
	c_opt_str := C.CString(opt_str)

	m.model = C.Create(c_rsc_dir, c_opt_str)
	if m.model == nil {
		return errors.New("[ Fail ] Create Model Error!")
	}

	C.free(unsafe.Pointer(c_rsc_dir))
	C.free(unsafe.Pointer(c_opt_str))

	return nil
}

func (m *Model) Parse(line string) ([][]string, error) {
	var morphs_result [][]string

	c_line := C.CString(line)
	c_parsedString := C.Parse(m.model.tagger, c_line)
	parsedString := C.GoString(c_parsedString)
	C.free(unsafe.Pointer(c_parsedString))

	if line != "" && parsedString == "" {
		return nil, errors.New("[ Fail ] Parse Error!")
	}

	morph_string_list := strings.Split(parsedString, " + ")

	for _, morph_string := range morph_string_list {
		morph_split := strings.Split(morph_string, "/")
		lex := strings.Join(morph_split[:len(morph_split)-1], "/")
		tag := morph_split[len(morph_split)-1]

		morphs_result = append(morphs_result, []string{lex, tag})
	}

	return morphs_result, nil
}

func (m *Model) Destroy() error {
	is_destroyed := int(C.Destroy(m.model))
	if is_destroyed == 0 {
		return errors.New("[ Fail ] Destroy Model Error!")
	} else {
		return nil
	}
}
