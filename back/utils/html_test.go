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

func TestAddSchemeIfNotExists(t *testing.T) {
	var want string
	var got string

	url := "https://example.com"
	want = "https://example.com"
	got = AddSchemeIfNotExists(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "http://example.com"
	want = "http://example.com"
	got = AddSchemeIfNotExists(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "example.com"
	want = "https://example.com"
	got = AddSchemeIfNotExists(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "//example.com"
	want = "https://example.com"
	got = AddSchemeIfNotExists(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestIsUrlPDF(t *testing.T) {
	var want bool
	var got bool

	url := "https://example.com"
	want = false
	got = IsUrlPDF(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "https://example.com/test.pdf"
	want = true
	got = IsUrlPDF(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "https://example.com/test.pdf?test=1"
	want = true
	got = IsUrlPDF(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	url = "https://example.com/test.pdf?test=1&test2=2"
	want = true
	got = IsUrlPDF(url)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
