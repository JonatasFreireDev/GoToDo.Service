package entities

type Status struct {
	Value string `json:"value" validate:"oneof=completed pending"`
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
