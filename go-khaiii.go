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

type WordC struct {
	sentence string
	word     *C.khaiii_word_t
}

type Word struct {
	Word   string
	Morphs []Morph
}

type Morph struct {
	Lex string
	Tag string
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

func (m *Model) Analyze(sentence string) ([]Word, error) {
	var parsedWords []Word
	wordResults, _ := m.analyzeRaw(sentence)

	for _, w := range wordResults {
		word := Word{
			Word:   w.sentence[w.word.begin : w.word.begin+w.word.length],
			Morphs: []Morph{},
		}
		morphsPtr := w.word.morphs
		for morphsPtr != nil {
			word.Morphs = append(
				word.Morphs,
				Morph{
					Lex: C.GoString(morphsPtr.lex),
					Tag: C.GoString(morphsPtr.tag),
				},
			)
			morphsPtr = morphsPtr.next
		}
		parsedWords = append(parsedWords, word)
	}
	return parsedWords, nil
}

func (m *Model) Nouns(sentence string) ([]string, error) {
	var parsedNouns []string
	wordResults, _ := m.analyzeRaw(sentence)

	for _, w := range wordResults {
		morphs := w.word.morphs
		for morphs != nil {
			if C.GoString(morphs.tag)[0] == 'N' {
				parsedNouns = append(parsedNouns, C.GoString(morphs.lex))
			}
			morphs = morphs.next
		}
	}
	return parsedNouns, nil
}

func (m *Model) Parse(sentence string) ([]Morph, error) {
	var parsedMorphs []Morph
	wordResults, _ := m.analyzeRaw(sentence)

	for _, w := range wordResults {
		morphs := w.word.morphs
		for morphs != nil {
			parsedMorphs = append(parsedMorphs, Morph{
				Lex: C.GoString(morphs.lex),
				Tag: C.GoString(morphs.tag),
			})
			morphs = morphs.next
		}
	}
	return parsedMorphs, nil
}

func (m *Model) ParseV1(sentence string) ([][]string, error) {
	var parsedSentence [][]string
	wordResults, _ := m.analyzeRaw(sentence)

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

func (m *Model) analyzeRaw(sentence string) ([]WordC, error) {
	if m.words != nil {
		m.freeResults()
	}
	var wordResults []WordC

	c_sentence := C.CString(sentence)
	wordPoint := C.khaiii_analyze(m.handle, c_sentence, C.CString(""))
	m.words = wordPoint

	for wordPoint != nil {
		w := WordC{sentence: sentence, word: wordPoint}
		wordResults = append(wordResults, w)
		wordPoint = wordPoint.next
	}

	return wordResults, nil
}
