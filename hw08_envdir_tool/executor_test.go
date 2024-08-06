package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	// Place your code here
	dir, err := os.Getwd()
	require.Nil(t, err)

	testCmd := []string{filepath.Join(dir, "/testdata/echo.sh"), "arg1=1", "arg2=2"}

	t.Run("simple test", func(t *testing.T) {
		env := make(Environment)
		expectedCode := 0
		cmdCode := RunCmd(testCmd, env)
		require.Equal(t, expectedCode, cmdCode)
	})

	t.Run("error 'ErrEmptyCmd'", func(t *testing.T) {
		env := make(Environment)
		expectedCode := 1
		testCmd := make([]string, 0)
		cmdCode := RunCmd(testCmd, env)
		require.Equal(t, expectedCode, cmdCode)
	})

	t.Run("set environment", func(t *testing.T) {
		env := make(Environment)
		expectedCode := 0
		envName := "TEST_ENV_NAME"
		envValue := "TEST_ENV_VALUE"
		env[envName] = EnvValue{envValue, false}
		cmdCode := RunCmd(testCmd, env)
		require.Equal(t, expectedCode, cmdCode)

		val, ok := os.LookupEnv(envName)
		require.True(t, ok)
		require.Equal(t, val, envValue)
	})

	t.Run("unset environment", func(t *testing.T) {
		env := make(Environment)
		expectedCode := 0
		envName := "TEST_ENV_NAME"
		envValueForUnset := "TEST_ENV_VALUE"
		envValue := ""
		err := os.Setenv(envName, envValueForUnset)
		require.Nil(t, err)

		env[envName] = EnvValue{envValue, true}
		cmdCode := RunCmd(testCmd, env)
		require.Equal(t, expectedCode, cmdCode)

		val, ok := os.LookupEnv(envName)
		require.False(t, ok)
		require.Equal(t, val, envValue)
	})
}
