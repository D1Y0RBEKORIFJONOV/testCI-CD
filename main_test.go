package main

import "testing"

func Test_SayHello(t *testing.T) {
	name := "Bob"
	want := "Hello, qwer"
	if got := sayHelloWorld(name); got != want {
		t.Errorf("sayHelloWorld() = %v, want %v", got, want)
	}
}
