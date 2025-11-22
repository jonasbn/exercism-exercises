package lsproduct

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	//s := []string{"1", "2", "3"}
	s := "123"
	p, err := product(s)

	if err != nil {
		t.Fatalf("TestProduct(%s) returned error %q, Error not expected.", s, err)
	} else {
		fmt.Printf("We got product: %d", p)
	}
}

func TestLargestSeriesProduct(t *testing.T) {
	for _, test := range tests {
		p, err := LargestSeriesProduct(test.digits, test.span)
		if test.ok {
			// we do not expect error
			if err != nil {
				t.Fatalf("LargestSeriesProduct(%s, %d) returned error %q.  "+
					"Error not expected.",
					test.digits, test.span, err)
			}

			if int64(p) != test.product {
				t.Fatalf("LargestSeriesProduct(%s, %d) = %d, want %d",
					test.digits, test.span, p, test.product)
			}
		} else { // expect error
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("LargestSeriesProduct(%s, %d) = %d, %v."+
					"  Expected error got nil",
					test.digits, test.span, p, err)
			}
		}
	}
}

func BenchmarkLargestSeriesProduct(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			LargestSeriesProduct(test.digits, test.span)
		}
	}
}
