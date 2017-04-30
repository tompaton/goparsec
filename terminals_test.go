package parsec

import "testing"

func TestEnd(t *testing.T) {
	p := And(nil, Token("test", "T"), End())
	s := NewScanner([]byte("test"))
	v, e := p(s)
	if v == nil {
		t.Errorf("End() didn't match %q", e)
	}
}

func TestNotEnd(t *testing.T) {
	p := And(nil, Token("test", "T"), End())
	s := NewScanner([]byte("testing"))
	v, _ := p(s)
	if v != nil {
		t.Errorf("End() shouldn't have matched %q", v)
	}
}

func TestNoEnd(t *testing.T) {
	p := And(nil, Token("test", "T"), NoEnd())
	s := NewScanner([]byte("testing"))
	v, e := p(s)
	if v == nil {
		t.Errorf("NoEnd() didn't match %q", e)
	}
}

func TestNotNoEnd(t *testing.T) {
	p := And(nil, Token("test", "T"), NoEnd())
	s := NewScanner([]byte("test"))
	v, _ := p(s)
	if v != nil {
		t.Errorf("NoEnd() shouldn't have matched %q", v)
	}
}

func TestCaseInsensitive(t *testing.T) {
	in := []string{"abc", "def"}
	want := []string{"(?i)abc", "(?i)def"}
	out := case_insensitive(in)
	if out[0] != want[0] || out[1] != want[1] {
		t.Errorf("case_insensitive(%q) == %q, want %q", in, out, want)
	}
}
