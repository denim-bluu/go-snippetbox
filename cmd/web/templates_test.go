package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2024, 1, 2, 15, 4, 0, 0, time.UTC)

	hd := humanDate(tm)

	if hd != "2024 Jan 02 at 15:04" {
		t.Errorf("expected %q; got %q", "2024 Jan 02 at 15:04", hd)
	}
}

func TestNewTemplateCache(t *testing.T) {
	tc, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := tc["home.html"]; !ok {
		t.Error("expected home.html to be in template cache")
	}
}
