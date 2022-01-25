package go_khaiii

/*
#include <khaiiic.h>
*/
import "C"
import (
	"runtime"
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

	C.free(unsafe.Pointer(c_rsc_dir))
	C.free(unsafe.Pointer(c_opt_str))

	return nil
}

func (m *Model) Parse(line string) ([][]string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var morphs_result [][]string

	c_line := C.CString(line)
	c_parsedString := C.Parse(m.model.tagger, c_line)
	parsedString := C.GoString(c_parsedString)
	C.free(unsafe.Pointer(c_parsedString))

	if parsedString == "" {
		return nil, nil
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
	C.Destroy(m.model)
	return nil
}
