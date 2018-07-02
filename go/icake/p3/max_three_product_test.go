package p3


import "testing"


var testcases = []struct {
	description string
	values []int
	expected int
}{
	{
		description: "1",
		values: []int{2,3,4,5},
		expected: 60,
	},
	{
		description: "2",
		values: []int{-8,-4,2,3,4,5},
		expected: 160,
	},

}
func TestBest_price(t *testing.T) {
	for _,tc :=range testcases{
		actual := Get_max_three_product(tc.values)
		if actual != tc.expected {
				t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
			}
		t.Logf("PASS: %s", tc.description)

	}
}