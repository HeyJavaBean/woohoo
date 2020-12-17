package main

import (
	"testing"
)

func TestWg(t *testing.T) {


	cha:= make(chan int)
	close(cha)
	close(cha)

}