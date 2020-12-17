package services

const (
	RuleOneDigit = iota
	RuleCapitalLetter
	RuleLowerLetter
	RuleSpecialCharacter
)

type Rule int

var RULES = map[Rule]struct{}{
	RuleOneDigit: {},
	RuleCapitalLetter: {},
	RuleLowerLetter: {},
	RuleSpecialCharacter: {},
}
