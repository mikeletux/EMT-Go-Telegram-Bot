package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAuth(t *testing.T) {
	authorizedUsers := []string{"mike", "fer", "richi"}

	simpleAuth := NewSimpleAuth(authorizedUsers)

	t.Run("Check users", func(t *testing.T) {
		assert.NoError(t, simpleAuth.CheckUser("mike"))  // exists, should pass
		assert.NoError(t, simpleAuth.CheckUser("fer"))   // exists, should pass
		assert.NoError(t, simpleAuth.CheckUser("richi")) // exists, should pass
		assert.Error(t, simpleAuth.CheckUser("hell"))    // doesn't exists, should fail
	})

	t.Run("Register a new user", func(t *testing.T) {
		// Check if foo user exist before adding it. Should fail
		assert.Error(t, simpleAuth.CheckUser("foo"), "user foo didn't exist but check passed")

		// Add foo to simpleAuth
		assert.NoError(t, simpleAuth.Register("foo"), "user didn't exist but failed when adding")

		// Try to Add foo to simpleAuth again and see it failing
		assert.Error(t, simpleAuth.Register("foo"), "user did exist but was added to the registry")

		// Check again if foo user exists. Shoud pass
		assert.NoError(t, simpleAuth.CheckUser("foo"), "user foo exists but check didn't pass")

		// Try to add an empty user
		assert.Error(t, simpleAuth.Register(""), "a user with an empty username was added")
	})
}
