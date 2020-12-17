package it_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/obiwandsilva/passwordvalidator/domain/entities"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

// Some values used in the tests are defined on environment variables on docker-compose file

func TestPasswordValidator(t *testing.T) {
	host := "http://localhost:7000"
	httpClient := http.Client{
		Timeout: 200 * time.Millisecond,
	}

	t.Run("should return isValid equal false when validating a short password", func(t *testing.T) {
		requestBody, err := json.Marshal(map[string]string{
			"password": "ABC",
		})
		require.NoError(t, err)
		request, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/validate", host),
			bytes.NewBuffer(requestBody),
		)
		require.NoError(t, err)

		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()

		var validation entities.Validation

		err = json.NewDecoder(resp.Body).Decode(&validation)
		require.NoError(t, err)

		require.False(t, validation.IsValid)
		require.ElementsMatch(t, validation.Errors, []string{"invalid length. Min: 9 Max: 32"})
	})

	t.Run("should return isValid equal false when validating a password missing digits", func(t *testing.T) {
		requestBody, err := json.Marshal(map[string]string{
			"password": "ABCDEFGHIJKLabcdefghijk@",
		})
		require.NoError(t, err)
		request, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/validate", host),
			bytes.NewBuffer(requestBody),
		)
		require.NoError(t, err)

		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()

		var validation entities.Validation

		err = json.NewDecoder(resp.Body).Decode(&validation)
		require.NoError(t, err)

		require.False(t, validation.IsValid)
		require.ElementsMatch(t, validation.Errors, []string{"should have at least one digit"})
	})

	t.Run("should return isValid equal true when validating a valid password", func(t *testing.T) {
		requestBody, err := json.Marshal(map[string]string{
			"password": "ABCDEFGHIJKLabcdefghijk*(2020)",
		})
		require.NoError(t, err)
		request, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/validate", host),
			bytes.NewBuffer(requestBody),
		)
		require.NoError(t, err)

		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()

		var validation entities.Validation

		err = json.NewDecoder(resp.Body).Decode(&validation)
		require.NoError(t, err)

		require.True(t, validation.IsValid)
		require.ElementsMatch(t, validation.Errors, []string{})
	})

	t.Run("should return status code 400 for any invalid request body", func(t *testing.T) {
		requestBody, err := json.Marshal(map[string]string{
			"wordpass": "ABCDEFGHIJKLabcdefghijk*(2020)",
		})
		require.NoError(t, err)
		request, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/validate", host),
			bytes.NewBuffer(requestBody),
		)
		require.NoError(t, err)

		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
