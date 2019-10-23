package stringslice

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		input []string
		val   string
		want  []string
	}{
		"should add": {
			[]string{},
			"foo",
			[]string{"foo"},
		},
		"should add to nil slice": {
			nil,
			"foo",
			[]string{"foo"},
		},
		"should not add when exists": {
			[]string{"foo"},
			"foo",
			[]string{"foo"},
		},
	}

	for name, tc := range tests {
		name, tc := name, tc // https://github.com/golang/go/wiki/CommonMistakes
		t.Run(name, func(t *testing.T) {
			got := Add(tc.input, tc.val)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := map[string]struct {
		input []string
		val   string
		want  []string
	}{
		"should remove": {
			[]string{"foo"},
			"foo",
			[]string{},
		},
		"should remove correct element from several": {
			[]string{"foo", "bar", "baz"},
			"bar",
			[]string{"foo", "baz"},
		},
		"should not remove if not found": {
			[]string{"foo"},
			"bar",
			[]string{"foo"},
		},
		"should return nil if called on nil slice with any string": {
			nil,
			"foo",
			nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc // https://github.com/golang/go/wiki/CommonMistakes
		t.Run(name, func(t *testing.T) {
			got := Remove(tc.input, tc.val)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestHas(t *testing.T) {
	tests := map[string]struct {
		input []string
		val   string
		want  bool
	}{
		"should find": {
			[]string{"foo"},
			"foo",
			true,
		},
		"should not find": {
			[]string{"foo"},
			"bar",
			false,
		},
		"should work on empty array": {
			nil,
			"bar",
			false,
		},
		"should find with multiple elements": {
			[]string{"foo", "bar", "baz"},
			"baz",
			true,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc // https://github.com/golang/go/wiki/CommonMistakes
		t.Run(name, func(t *testing.T) {
			got := Has(tc.input, tc.val)
			if tc.want != got {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		input []string
		keep  func(val string) bool
		want  []string
	}{
		"should return input unchange": {
			[]string{"foo"},
			func(val string) bool { return true },
			[]string{"foo"},
		},
		"should return empty slice": {
			[]string{"foo"},
			func(val string) bool { return false },
			[]string{},
		},
		"should work on nil slice": {
			nil,
			func(val string) bool { return true },
			nil,
		},
		"should remove multiple elements": {
			[]string{"foo", "bar", "baz"},
			func(val string) bool { return val == "baz" },
			[]string{"baz"},
		},
	}

	for name, tc := range tests {
		name, tc := name, tc // https://github.com/golang/go/wiki/CommonMistakes
		t.Run(name, func(t *testing.T) {
			got := Filter(tc.input, tc.keep)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
