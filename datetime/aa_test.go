package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	inRange, err := RandomTimeInRange(start, end)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inRange)
}
