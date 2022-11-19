package main

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	gotTime, err := getCurrentTime()

	if err != nil {
		t.Errorf("Got error from ntp lib: %s", err.Error())
	}

	if time.Now().Unix()-gotTime.Unix() > 2 {
		t.Errorf("The time from ntp has a discrepancy greater than 2 seconds")
	}

}
