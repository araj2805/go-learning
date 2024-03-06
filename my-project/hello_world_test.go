package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ankit")
	want := "Hello Ankit"
	
	if got != want {
	
		t.Errorf("got %q want %q",got, want)
	
	}



}