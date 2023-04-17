package model

import "time"

type Status struct {
	Value     string    `json:"value" validate:"oneof=completed pending"`
	UpdatedAt time.Time `json:"updatedAt" validate:"datetime=2006-01-02"`
}

// func NewStatus(value string) *Status {
// 	s := &Status{}
// 	s.Set(value)
// 	return s
// }

// func (s *Status) Get() string {
// 	return s.key
// }

// func (s *Status) Set(value string) {
// 	switch value {

// 	case "completed", "pending":
// 		s.key = value
// 	}
// }
