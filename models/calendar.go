package models

import "time"

// CalendarSlot is a model
type CalendarSlot struct {
	Begin  time.Time `json:"begin"`
	End    time.Time `json:"end"`
	IsOpen bool      `json:"is_open"`
}

// CalendarDay is a model
type CalendarDay struct {
	Day       time.Time       `json:"day"`
	IsWeekend bool            `json:"is_weekend"`
	Slots     []*CalendarSlot `json:"slots"`
}
