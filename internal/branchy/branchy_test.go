package branchy

import "testing"

func TestToLowerCase(t *testing.T) {
	got := toLowerCase("ABC")
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestReplaceAllSpecialCharacters(t *testing.T) {
	got := replaceAllSpecialCharacters("this!is@some#text")
	want := "this-is-some-text"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestTrimDash(t *testing.T) {
	got := trimDash("-ABC-")
	want := "ABC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildBranchName(t *testing.T) {
	got := buildBranchName("feat", "ABC-1234", "some-text")
	want := "feat/ABC-1234/some-text"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
