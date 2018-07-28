package hangulize

// Hangulize is the most simple and useful API of thie package. It transcribes
// a non-Korean word into Hangul, which is the Korean alphabet. For example, it
// will transcribe "Владивосто́к" in Russian into "블라디보스토크".
func Hangulize(lang string, word string) string {
	spec, ok := LoadSpec(lang)
	if !ok {
		// spec not found
		return word
	}

	h := NewHangulizer(spec)
	return h.Hangulize(word)
}

// Hangulizer provides the transcription logic for the underlying spec.
type Hangulizer struct {
	spec *Spec
}

// NewHangulizer creates a Hangulizer for a spec.
func NewHangulizer(spec *Spec) *Hangulizer {
	return &Hangulizer{spec}
}

// Spec returns the underlying spec.
func (h *Hangulizer) Spec() *Spec {
	return h.spec
}

// Hangulize transcribes a loanword into Hangul.
func (h *Hangulizer) Hangulize(word string) string {
	p := pipeline{h, nil}
	return p.forward(word)
}

// HangulizeTrace transcribes a loanword into Hangul
// and returns the traced internal events too.
func (h *Hangulizer) HangulizeTrace(word string) (string, []Trace) {
	var tr tracer
	p := pipeline{h, &tr}

	word = p.forward(word)

	return word, tr.Traces()
}