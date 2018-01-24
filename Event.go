package main

import (
   "time"
)

type Event struct {
   Title string         `json:"title"`
   Description string   `json:"desc, omitempty"`
   Start time.Time      `json:"start"`
   End time.Time        `json:"end"`
   Location string      `json:"loc"`
}
