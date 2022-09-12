package utils

import (
	"testing"
	"time"
)

func TestHumanizeParseTime(t *testing.T) {
	var err error
	var want time.Time
	var got time.Time

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	want = now.Add(-3 * time.Hour * 24)
	got, err = HumanizeParseTime("3日前", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now.Add(-1 * time.Hour)
	got, err = HumanizeParseTime("1時間前", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now.Add(-2 * time.Minute)
	got, err = HumanizeParseTime("2分前", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now.Add(-3 * time.Second)
	got, err = HumanizeParseTime("3秒前", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now
	got, err = HumanizeParseTime("たったいま", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now
	got, err = HumanizeParseTime("いま", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	want = now
	got, err = HumanizeParseTime("今", now)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	got, err = HumanizeParseTime("1ヶ月後", now)

	if "not supported" != err.Error() {
		t.Errorf("want %v, got %v", "not supported", got)
	}

}
