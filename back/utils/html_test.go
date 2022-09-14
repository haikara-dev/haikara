package utils

import (
	"testing"
)

func TestCreateSelectorOnChildrenScopeFeatureSupport(t *testing.T) {
	var want string
	var got string

	rootSelector := ".pc_only .fp_media_tile"

	titleSelector := " > a > div"
	linkSelector := " > a"
	dateSelector := " > div > span"

	want = ".pc_only .fp_media_tile > a > div"
	got = CreateSelectorOnChildrenScopeFeatureSupport(
		titleSelector,
		rootSelector,
	)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = ".pc_only .fp_media_tile > a"
	got = CreateSelectorOnChildrenScopeFeatureSupport(
		linkSelector,
		rootSelector,
	)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = ".pc_only .fp_media_tile > div > span"
	got = CreateSelectorOnChildrenScopeFeatureSupport(
		dateSelector,
		rootSelector,
	)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}