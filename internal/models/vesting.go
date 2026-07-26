package models

import "time"

type Vesting struct {
	TotalTokens int
	StartDate   time.Time
	Months      int
}
