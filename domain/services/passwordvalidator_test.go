package services_test

import (
	"fmt"
	"github.com/obiwandsilva/passwordvalidator/appplication/config"
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
		expected func() (bool, []string)
	}{
		{
			name:  "should return false for passwords with less than 9 chars",
			input: "ABCabc1",
			expected: func() (bool, []string) {
				missingRule := fmt.Sprintf(
					"invalid length. Min: %d Max: %d",
					passwordValidator.EnvironmentConfig.MinPasswordSize,
					passwordValidator.EnvironmentConfig.MaxPasswordSize,
				)
				return false, []string{missingRule}
			},
		},
		{
			name:  "should return false for passwords with more than 32 chars",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1",
			expected: func() (bool, []string) {
				missingRule := fmt.Sprintf(
					"invalid length. Min: %d Max: %d",
					passwordValidator.EnvironmentConfig.MinPasswordSize,
					passwordValidator.EnvironmentConfig.MaxPasswordSize,
				)
				return false, []string{missingRule}
			},
		},
		{
			name:  "should return false for passwords missing at least one digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZa@",
			expected: func() (bool, []string) {
				return false, []string{"should have at least one digit"}
			},
		},
		{
			name:  "should return false for passwords missing at least one lower letter",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ1#",
			expected: func() (bool, []string) {
				return false, []string{"should have at least one lower letter"}
			},
		},
		{
			name:  "should return false for passwords missing at least one capital letter",
			input: "abcdefghijklmnopqrstuvwxyz1#",
			expected: func() (bool, []string) {
				return false, []string{"should have at least one capital letter"}
			},
		},
		{
			name:  "should return false for passwords missing at least one special character",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZabc1",
			expected: func() (bool, []string) {
				return false, []string{"should have at least one of the special characters: !@#$%^&*()-+"}
			},
		},
		{
			name:  "should return false for passwords with any invalid character. In this test: (white space)",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ abc1$",
			expected: func() (bool, []string) {
				return false, []string{"invalid character: (whitespace)"}
			},
		},
		{
			name:  "should return false for passwords with any invalid character. In this test: <",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ<abc1$",
			expected: func() (bool, []string) {
				return false, []string{"invalid character: <"}
			},
		},
		{
			name:  "should return a list of errors with missing lower letter and digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ%",
			expected: func() (bool, []string) {
				return false, []string{"should have at least one digit", "should have at least one lower letter"}
			},
		},
		{
			name:  "should return a list of errors with missing lower letter, special character and digit",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: func() (bool, []string) {
				return false, []string{
					"should have at least one digit",
					"should have at least one lower letter",
					"should have at least one of the special characters: !@#$%^&*()-+",
				}
			},
		},
		{
			name:  "should return a list of errors with missing capital letter, special character and digit",
			input: "abcdefghijklmnopqrstuvwxyz",
			expected: func() (bool, []string) {
				return false, []string{
					"should have at least one digit",
					"should have at least one capital letter",
					"should have at least one of the special characters: !@#$%^&*()-+",
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedIsValid, expectedErrors := testCase.expected()
			isValid, errors := passwordValidator.IsValid(testCase.input)

			require.Equal(t, expectedIsValid, isValid)
			require.ElementsMatch(t, expectedErrors, errors)
		})
	}
}
