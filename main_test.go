package main

import "testing"

func TestHello(t *testing.T) {
	if HelloWorld() != "Hello World, Drone Workshop" {
		t.Error("Testing error")
	}
}
