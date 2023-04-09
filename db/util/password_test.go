package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	hashedPass1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass1)

	err = CheckPassword(hashedPass1, password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(hashedPass1, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPass2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPass2)
	require.NotEqual(t, hashedPass1, hashedPass2)
}
