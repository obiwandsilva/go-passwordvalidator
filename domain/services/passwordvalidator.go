package services

import (
	"fmt"
	"github.com/obiwandsilva/passwordvalidator/appplication/config"
)

type PasswordValidator struct {
	EnvironmentConfig config.EnvironmentConfig
}

func NewPasswordValidator(EnvironmentConfig config.EnvironmentConfig) *PasswordValidator {
	return &PasswordValidator{
		EnvironmentConfig: EnvironmentConfig,
	}
}

func (pvs *PasswordValidator) IsValid(password string) (bool, []string) {
	errors := pvs.validatePassword(password)

	if len(errors) > 0 {
		return false, errors
	}

	return true, nil
}

func (pvs *PasswordValidator) validatePassword(password string) []string {
	errors := make([]string, 0)

	if !validateLength(password, pvs.EnvironmentConfig.MinPasswordSize, pvs.EnvironmentConfig.MaxPasswordSize) {
		errorMessage := fmt.Sprintf(
			"invalid length. Min: %d Max: %d",
			pvs.EnvironmentConfig.MinPasswordSize,
			pvs.EnvironmentConfig.MaxPasswordSize,
		)
		return append(errors, errorMessage)
	}

	missingCharsRules, err := validateRules(password)
	if err!= nil {
		errors = append(errors, err.Error())
		return errors
	}

	for rule, _ := range missingCharsRules {
		switch rule {
		case RuleOneDigit:
			errors = append(errors, "should have at least one digit")
		case RuleCapitalLetter:
			errors = append(errors, "should have at least one capital letter")
		case RuleLowerLetter:
			errors = append(errors, "should have at least one lower letter")
		case RuleSpecialCharacter:
			errors = append(errors, "should have at least one of the special characters: !@#$%^&*()-+")
		default:
		}
	}

	return errors
}

func validateLength(password string, min, max int) bool {
	if len(password) < min || len(password) > max {
		return false
	}

	return true
}

func validateRules(password string) (map[Rule]struct{}, error) {
	missing := copyRuleMap(RULES)
	var previousChar uint8 = 0

	for i := 0; i < len(password); i++ {
		char := password[i]

		if char == previousChar {
			return nil, fmt.Errorf("characters cannot repeat in conjunction: %c%c", char, previousChar)
		}

		switch {
		case char > 47 && char < 58:
			delete(missing, RuleOneDigit)
		case char > 64 && char < 91:
			delete(missing, RuleCapitalLetter)
		case char > 96 && char < 123:
			delete(missing, RuleLowerLetter)
		case validateSpecialCharacter(rune(char)):
			delete(missing, RuleSpecialCharacter)
		default:
			if char == 32 {
				return nil, fmt.Errorf("invalid character: (whitespace)")
			}

			return nil, fmt.Errorf("invalid character: %c", char)
		}

		previousChar = char
	}

	return missing, nil
}

func validateSpecialCharacter(char rune) bool {
	asciiSpecialChars := map[rune]string {
		33: "!",
		35: "#",
		36: "$",
		37: "%",
		38: "&",
		40: "(",
		41: ")",
		42: "*",
		43: "+",
		45: "-",
		64: "@",
		94: "^",
	}

	_, ok := asciiSpecialChars[char]

	return ok
}

func copyRuleMap(source map[Rule]struct{}) map[Rule]struct{} {
	dest := make(map[Rule]struct{}, 0)

	for k, v := range source {
		dest[k] = v
	}

	return dest
}
