package main

import "testing"

func TestSayHello(t *testing.T) {
	hello := sayHello()

	if hello != "Hello world!" {
		t.Fatal("sayHello() should say 'Hello world!'")
	}
}
