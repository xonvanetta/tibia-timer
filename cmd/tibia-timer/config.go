package main

import (
	"fmt"
	"time"
)

var (
	FirstAlarm  = ""
	SecondAlarm = ""
	Interval    = ""
	WindowTitle = ""
)

type Config struct {
	FirstAlarm  time.Duration
	SecondAlarm time.Duration
	Interval    time.Duration
	Volume      float64
}

func (c *Config) Validate() error {
	if c.FirstAlarm == 0 {
		return fmt.Errorf("missing first alarm")
	}

	if c.Interval == 0 {
		return fmt.Errorf("missing interval")
	}

	return nil
}

func MustParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Errorf("must parse duration: %w", err))
	}
	return d
}
