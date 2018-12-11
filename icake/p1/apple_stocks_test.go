package p1

import "testing"


var testcases = []struct {
	description string
	prices []int
	expected int
}{
	{
		description: "1",
		prices : []int{10,7,5,8,11,9},
		expected: 6,
	},

	{
		description: "2",
		prices: []int{8,11,20,2,22,14},
		expected: 20,

	},
}
func TestBest_price(t *testing.T) {
	for _,tc :=range testcases{
		val := Best_price(tc.prices)
		if val != tc.expected{
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, val)
		}

		t.Logf("PASS: %s", tc.description)


	}
}