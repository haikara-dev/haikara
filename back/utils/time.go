package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func HumanizeParseTime(str string, now time.Time) (time.Time, error) {
	var reg *regexp.Regexp
	var match []string

	reg = regexp.MustCompile("([0-9]+)日前")
	match = reg.FindStringSubmatch(str)
	if match != nil {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			duration := time.Duration(num) * time.Hour * 24
			return now.Add(-duration), nil
		}
	}

	reg = regexp.MustCompile("([0-9]+)時間前")
	match = reg.FindStringSubmatch(str)
	if match != nil {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			duration := time.Duration(num) * time.Hour
			return now.Add(-duration), nil
		}
	}

	reg = regexp.MustCompile("([0-9]+)分前")
	match = reg.FindStringSubmatch(str)
	if match != nil {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			duration := time.Duration(num) * time.Minute
			return now.Add(-duration), nil
		}
	}

	reg = regexp.MustCompile("([0-9]+)秒前")
	match = reg.FindStringSubmatch(str)
	if match != nil {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			duration := time.Duration(num) * time.Second
			return now.Add(-duration), nil
		}
	}

	reg = regexp.MustCompile("(今|たったいま|いま)")
	if reg.MatchString(str) {
		return now, nil
	}

	return time.Time{}, errors.New("not supported")
}
