package main

import "time"

// Classification option for Image
type Classification struct {
	ID        int64
	CreatedAt time.Time
	Name      string
	Hotkey    string
}
