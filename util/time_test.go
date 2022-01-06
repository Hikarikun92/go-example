package util

import (
	"testing"
	. "time"
)

func TestTimeToIso(t *testing.T) {
	time := Date(2022, 1, 5, 23, 22, 42, 0, UTC)
	str := TimeToIso(time)
	expected := "2022-01-05T23:22:42"

	if str != expected {
		t.Errorf("Got %v, wanted %v", str, expected)
	}
}
