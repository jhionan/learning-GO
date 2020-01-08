package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("from Me")
	want := "Hello GOrld, from Me"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
