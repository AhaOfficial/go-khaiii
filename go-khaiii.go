package go_khaiii

/*
#include <khaiiic.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
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

func (m *Model) Parse(line string, opt ...map[string]string) ([][]string, error) {
	var morphs_result [][]string

	c_line := C.CString(line)
	c_parsedString := C.Parse(m.model.tagger, c_line)

	if c_parsedString == nil {
		return nil, errors.New("[ Fail ] Parse Error!")
	}

	parsedString := C.GoString(c_parsedString)
	C.free(unsafe.Pointer(c_parsedString))

	morph_string_list := strings.Split(parsedString, " + ")

	for _, morph_string := range morph_string_list {
		morph_split := strings.Split(morph_string, "/")
		lex := strings.Join(morph_split[:len(morph_split)-1], "/")
		tag := morph_split[len(morph_split)-1]

		morphs_result = append(morphs_result, []string{lex, tag})
	}

	if len(opt) > 0 {
		if _, ok := opt[0]["--no-consonant"]; ok {
			morphs_result = removeConsonantMorphs(morphs_result, line)
		}
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

func removeConsonantMorphs(morphs [][]string, org string) [][]string {
	var morphs_result [][]string
	var keep_word []string

	include_consonant := false
	consonant_regexp, _ := regexp.Compile("^[ㄱ-ㅎ]")
	for _, word := range morphs {
		if consonant_regexp.MatchString(word[0]) {
			include_consonant = true
			break
		}
	}
	if !include_consonant {
		return morphs
	}

	temp := strings.ReplaceAll(org, " ", "")
	for _, word := range morphs {
		index := -1
		if keep_word == nil {
			index = strings.IndexAny(temp, word[0])
		} else {
			left_word, right_word := combineWord(keep_word[0], word[0])
			index = strings.IndexAny(temp, fmt.Sprintf("%s%s", left_word, right_word))

			keep_word = []string{left_word, keep_word[1]}
			word = []string{right_word, word[1]}
		}

		if index == 0 {
			temp = temp[len(word[0]):]
			if keep_word != nil {
				morphs_result = append(morphs_result, keep_word)
				keep_word = nil
			}
			if len(word[0]) > 0 {
				morphs_result = append(morphs_result, word)
			}
		} else {
			keep_word = word
		}
	}
	return morphs_result
}

func combineWord(a_str string, b_str string) (string, string) {
	a_str_tail_rune := []rune(a_str)[utf8.RuneCountInString(a_str)-1]
	a_str_head := a_str[:len(a_str)-len(string(a_str_tail_rune))]

	b_str_head_rune := []rune(b_str)[0]
	b_str_tail := b_str[len(string(b_str_head_rune)):]

	middle := string(combineChar(a_str_tail_rune, b_str_head_rune))

	return fmt.Sprintf("%s%s", a_str_head, middle), b_str_tail
}

func combineChar(cho_jung rune, jong rune) rune {
	temp := cho_jung - rune('가')
	cho := temp / 588
	jung := (temp - (cho * 588)) / 28
	jong = jong - rune('ㄱ') + 1
	return (588 * rune(cho)) + (28 * rune(jung)) + jong + rune('가')
}
