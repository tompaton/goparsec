package parsec

import "testing"
import "reflect"

func TestStrEOF(t *testing.T) {
	word := String()
	Y := Many(
		func(ns []ParsecNode) ParsecNode {
			return ns
		},
		word)

	input := `"alpha" "beta" "gamma"`
	s := NewScanner([]byte(input))

	root, _ := Y(s)
	nodes := root.([]ParsecNode)
	ref := []ParsecNode{"\"alpha\"", "\"beta\"", "\"gamma\""}
	if !reflect.DeepEqual(nodes, ref) {
		t.Fatal(nodes)
	}
}

func TestMany(t *testing.T) {
	w := Token("\\w+", "W")
	m := Many(nil, w)
	s := NewScanner([]byte("one two stop"))
	v, e := m(s)
	if v == nil {
		t.Errorf("Many() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 3 {
		t.Errorf("Many() didn't match all words %q", v)
	}
}

func TestManyUntilNoStop(t *testing.T) {
	w := Token("\\w+", "W")
	u := Token("stop", "S")
	m := ManyUntil(nil, w, u)
	s := NewScanner([]byte("one two three"))
	v, e := m(s)
	if v == nil {
		t.Errorf("ManyUntil() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 3 {
		t.Errorf("ManyUntil() didn't match all words %q", v)
	}
}

func TestManyUntilStop(t *testing.T) {
	w := Token("\\w+", "W")
	u := Token("stop", "S")
	m := ManyUntil(nil, w, u)
	s := NewScanner([]byte("one two stop"))
	v, e := m(s)
	if v == nil {
		t.Errorf("ManyUntil() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 2 {
		t.Errorf("ManyUntil() didn't stop %q", v)
	}
}

func TestManyUntilNoStopSep(t *testing.T) {
	w := Token("\\w+", "W")
	u := Token("stop", "S")
	z := Token("z", "Z")
	m := ManyUntil(nil, w, z, u)
	s := NewScanner([]byte("one z two z three"))
	v, e := m(s)
	if v == nil {
		t.Errorf("ManyUntil() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 3 {
		t.Errorf("ManyUntil() didn't match all words %q", v)
	}
}

func TestManyUntilStopSep(t *testing.T) {
	w := Token("\\w+", "W")
	u := Token("stop", "S")
	z := Token("z", "Z")
	m := ManyUntil(nil, w, z, u)
	s := NewScanner([]byte("one z two z stop"))
	v, e := m(s)
	if v == nil {
		t.Errorf("ManyUntil() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 2 {
		t.Errorf("ManyUntil() didn't stop %q", v)
	}
}

func TestOptionalMissing(t *testing.T) {
	w1 := Token("one", "W1")
	w2 := Token("two", "W2")
	x := Token("x", "X")
	m := And(nil, w1, Optional(nil, x, Terminal{"X", "x", 0}), w2)
	s := NewScanner([]byte("one two"))
	v, e := m(s)
	if v == nil {
		t.Errorf("Optional() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 3 {
		t.Errorf("Optional() didn't match %q", v)
	}
}

func TestOptionalPresent(t *testing.T) {
	w1 := Token("one", "W1")
	w2 := Token("two", "W2")
	x := Token("x", "X")
	m := And(nil, w1, Optional(nil, x, Terminal{"X", "x", 0}), w2)
	s := NewScanner([]byte("one x two"))
	v, e := m(s)
	if v == nil {
		t.Errorf("Optional() didn't match %q", e)
	} else if len(v.([]ParsecNode)) != 3 {
		t.Errorf("Optional() didn't match %q", v)
	}
}
