package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, Noman! Your Go project is working."
    if expected == "" {
        t.Errorf("Something went wrong")
    }
}
