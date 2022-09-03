package calendar

import (
	"time"
)

// BaseInitializer initializes the generator.
type BaseInitializer struct{}

// setDefaultBegin sets the begin.
func (i *BaseInitializer) setDefaultBegin(gen *Generator) {
	if gen.Begin.IsZero() {
		d := time.Now()
		m := time.Date(
			d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local,
		)
		gen.Begin = m
	}
}

// DayInitializer initializes the generator.
type DayInitializer struct {
	BaseInitializer
}

// Init initializes the generator.
func (i *DayInitializer) Init(gen *Generator) {
	i.setDefaultBegin(gen)
	gen.End = gen.Begin
}

// WeekInitializer initializes the generator.
type WeekInitializer struct {
	BaseInitializer
}

// Init initializes the generator.
func (i *WeekInitializer) Init(gen *Generator) {
	i.setDefaultBegin(gen)
	days := 6
	gen.End = gen.Begin.AddDate(0, 0, days)
}

// MonthInitializer initializes the generator.
type MonthInitializer struct {
	BaseInitializer
}

// Init initializes the generator.
func (i *MonthInitializer) Init(gen *Generator) {
	i.setDefaultBegin(gen)
	gen.End = gen.Begin.AddDate(0, 1, 0)
}

// YearInitializer initializes the generator.
type YearInitializer struct{}

// Init initializes the generator.
func (i *YearInitializer) Init(gen *Generator) {
	if gen.Begin.IsZero() {
		gen.Begin = time.Now()
	}
	gen.End = gen.Begin.AddDate(1, 0, 0)
}
