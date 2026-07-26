package main

import (
	"fmt"
	"token-vesting-calculator/internal/calculator"
)

func main() {

	schedule := calculator.DefaultSchedule()

	fmt.Println("Unlocked:", schedule.Unlocked())

	fmt.Println("Locked:", schedule.Locked())

	fmt.Printf("Progress: %.2f%%\n", schedule.Progress())

	schedule.ExportCalendar()
}
