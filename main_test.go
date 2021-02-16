package main

import "testing"

func benchmarkmain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		main()
	}
}
