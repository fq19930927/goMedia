package test

import (
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T)  {
	fmt.Println("test1")
	testDemo2(t)
}

func testDemo2(t *testing.T)  {
	fmt.Println("test2")
}