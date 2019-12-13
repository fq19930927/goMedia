package bench

import "testing"

func Print1to20() int  {
	n:=0
	for i:=0;i<20 ;i++  {
		n += i
	}
	return n
}

func TestMain(m *testing.M)  {
	m.Run()
}

func Benchmark(b *testing.B)  {
	for n:=0;n<b.N ;n++  {
		Print1to20()
	}
}
