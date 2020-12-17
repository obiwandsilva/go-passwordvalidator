package entities

type Validation struct {
	IsValid bool     `json:"isValid"`
	Errors  []string `json:"errors"`
}
