package comma

import "testing"

func TestComma(t *testing.T) {
	t.Run("Check number is segmented with commas", func(t *testing.T) {
		got := Comma(12345)
		want := "12,345"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("Check number is segmented with commas", func(t *testing.T) {
		got := Comma(44444444)
		want := "44,444,444"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
