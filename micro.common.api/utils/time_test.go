package utils

import "testing"

func TestTimeStamp2Time(t *testing.T) {
	t.Log(TimeStamp2Time(-1))
}

func TestParseDay(t *testing.T) {
	date, err := ParseDay("1006-01-02")

	t.Log(date)
	t.Log(err)
}
