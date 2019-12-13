package test

import (
	"fmt"
	"testing"
)

/**
测试TestMain
 */

func TestMain(m *testing.M)  {
	fmt.Println("init...")
	m.Run()
}

func TestPrint(t *testing.T)  {
	t.Run("test1", func(t *testing.T) {
		testPrint1(t)
	})
	t.Run("test1", func(t *testing.T) {
		testPrint2(t)
	})
}

func testPrint1(t *testing.T)  {
	fmt.Println("print1")
}

func testPrint2(t *testing.T)  {
	fmt.Println("print2")
}
