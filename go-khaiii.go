package go_khaiii

// #cgo LDFLAGS: -lkhaiii
// #include <khaiii/khaiii_api.h>
import "C"
import (
	"errors"
)

type Model struct {
	handle C.int
	words  *C.khaiii_word_t
}

type Word struct {
	sentence string
	word     *C.khaiii_word_t
}

func (m *Model) Create(rsc_dir string, opt_str string) error {
	handle := C.khaiii_open(C.CString(rsc_dir), C.CString(opt_str))
	if handle < 0 {
		return errors.New("[Fail] Create Model Error")
	}

	m.handle = handle

	return nil
}

func (m *Model) Destroy() {
	if m.words != nil {
		m.freeResults()
	}
	C.khaiii_close(m.handle)
}

func (m *Model) Parse(sentence string) ([][]string, error) {
	var parsedSentence [][]string
	wordResults, _ := m.analyze(sentence)

	for _, w := range wordResults {
		morphs := w.word.morphs
		for morphs != nil {
			lex := C.GoString(morphs.lex)
			tag := C.GoString(morphs.tag)
			mPair := []string{lex, tag}

			parsedSentence = append(parsedSentence, mPair)

			morphs = morphs.next
		}
	}
	return parsedSentence, nil
}

func (m *Model) freeResults() {
	C.khaiii_free_results(m.handle, m.words)
	m.words = nil
}

func (m *Model) analyze(sentence string) ([]Word, error) {
	if m.words != nil {
		m.freeResults()
	}
	var wordResults []Word

	c_sentence := C.CString(sentence)
	wordPoint := C.khaiii_analyze(m.handle, c_sentence, C.CString(""))
	m.words = wordPoint

	for wordPoint != nil {
		w := Word{sentence: sentence, word: wordPoint}
		wordResults = append(wordResults, w)
		wordPoint = wordPoint.next
	}

	return wordResults, nil
}
