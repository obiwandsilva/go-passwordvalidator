package services_test

import (
	"fmt"
	"github.com/obiwandsilva/passwordvalidator/application/config"
	"github.com/obiwandsilva/passwordvalidator/domain/entities"
	"github.com/obiwandsilva/passwordvalidator/domain/services"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	passwordValidator := services.NewPasswordValidator(config.EnvironmentConfig{
		MinPasswordSize: 9,
		MaxPasswordSize: 32,
	})

	testCases := []struct {
		name     string
		input    string
		expected entities.Validation
	}{
		{
			name:  "should return false for passwords with less than 9 chars",
			input: "ABCabc1",
			expected: entities.Validation{
				IsValid: false,
				Errors: []string{
					fmt.Sprintf(
						"invalid length. Min: %d Max: %d",
						passwordValidator.EnvironmentConfig.MinPasswordSize,
						passwordValidator.EnvironmentConfig.MaxPasswordSize,
					),
				},
			},
		},
		{
			name:  "should return false for passwords with more than 32 chars",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1",
			expected: entities.Validation{
				IsValid: false,
				Errors: []string{
					fmt.Sprintf(
						"invalid length. Min: %d Max: %d",
						passwordValidator.EnvironmentConfig.MinPasswordSize,
						passwordValidator.EnvironmentConfig.MaxPasswordSize,
					),
				},
			},
		},
		{
			name:  "should return false for passwords missing at least one digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZa@",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"should have at least one digit"},
			},
		},
		{
			name:  "should return false for passwords missing at least one lower letter",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ1#",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"should have at least one lower letter"},
			},
		},
		{
			name:  "should return false for passwords missing at least one capital letter",
			input: "abcdefghijklmnopqrstuvwxyz1#",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"should have at least one capital letter"},
			},
		},
		{
			name:  "should return false for passwords missing at least one special character",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZabc1",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"should have at least one of the special characters: !@#$%^&*()-+"},
			},
		},
		{
			name:  "should return false for passwords with any invalid character. In this test: (white space)",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ abc1$",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"invalid character: (whitespace)"},
			},
		},
		{
			name:  "should return false for passwords with any invalid character. In this test: <",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ<abc1$",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"invalid character: <"},
			},
		},
		{
			name:  "should return a list of errors with missing lower letter and digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ%",
			expected: entities.Validation{
				IsValid: false,
				Errors:  []string{"should have at least one digit", "should have at least one lower letter"},
			},
		},
		{
			name:  "should return a list of errors with missing lower letter, special character and digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: entities.Validation{
				IsValid: false,
				Errors: []string{
					"should have at least one digit",
					"should have at least one lower letter",
					"should have at least one of the special characters: !@#$%^&*()-+",
				},
			},
		},
		{
			name:  "should return a list of errors with missing capital letter, special character and digit",
			input: "abcdefghijklmnopqrstuvwxyz",
			expected: entities.Validation{
				IsValid: false,
				Errors: []string{
					"should have at least one digit",
					"should have at least one capital letter",
					"should have at least one of the special characters: !@#$%^&*()-+",
				},
			},
		},
		{
			name:  "should return false for a password with repeated chars in sequence",
			input: "ABCDE$$abcde1(@)",
			expected: entities.Validation{
				IsValid: false,
				Errors: []string{"characters cannot repeat in conjunction: $$"},
			},
		},
		{
			name:  "should return true for a valid password",
			input: "ABCDEabcde1&",
			expected: entities.Validation{
				IsValid: true,
				Errors: []string{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validation := passwordValidator.IsValid(testCase.input)

			require.Equal(t, testCase.expected.IsValid, validation.IsValid)
			require.ElementsMatch(t, testCase.expected.Errors, validation.Errors)
		})
	}
}
