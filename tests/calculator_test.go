package tests

import (
	"testing"
	"token-vesting-calculator/internal/calculator"
)

func TestUnlocked(t *testing.T) {

	s := calculator.DefaultSchedule()

	if s.Unlocked() <= 0 {
		t.Fail()
	}
}
