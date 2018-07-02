package p2


import "testing"


var testcases = []struct {
	description string
	values []int
	expected []int
}{
	{
		description: "1",
		values: []int{2,3,4,5},
		expected: []int{60,40,30,24},
	},

}
func TestBest_price(t *testing.T) {
	for _,tc :=range testcases{
		actual := Product_all_expect_at_index(tc.values)
		for i,val := range actual {
			if val != tc.expected[i] {
				t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
			}
		}
		t.Logf("PASS: %s", tc.description)


	}
}