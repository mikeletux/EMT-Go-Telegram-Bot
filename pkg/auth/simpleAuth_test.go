package auth

import "testing"

func TestSimpleAuth(t *testing.T) {
	authorizedUsers := []string{"mike", "fer", "richi"}

	simpleAuth := NewSimpleAuth(authorizedUsers)

	t.Run("Check users", func(t *testing.T) {
		assertBool(t, true, simpleAuth.CheckUser("mike"))  // exists, should pass
		assertBool(t, true, simpleAuth.CheckUser("fer"))   // exists, should pass
		assertBool(t, true, simpleAuth.CheckUser("richi")) // exists, should pass
		assertBool(t, false, simpleAuth.CheckUser("jur"))  // doesn't exists, should fail
	})
}

func assertBool(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("error. Want %t, got %t", want, got)
	}
}
