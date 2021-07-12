package auth

import "testing"

func TestSimpleAuth(t *testing.T) {
	authorizedUsers := []string{"mike", "fer", "richi"}

	simpleAuth := NewSimpleAuth(authorizedUsers)

	t.Run("Check users", func(t *testing.T) {
		assertError(t, false, simpleAuth.CheckUser("mike"))  // exists, should pass
		assertError(t, false, simpleAuth.CheckUser("fer"))   // exists, should pass
		assertError(t, false, simpleAuth.CheckUser("richi")) // exists, should pass
		assertError(t, true, simpleAuth.CheckUser("jur"))    // doesn't exists, should fail
	})
}

func assertError(t testing.TB, hasError bool, err error) {
	t.Helper()
	if !hasError && err != nil {
		t.Errorf("no error expected and got one %s", err)
	}
	if hasError && err == nil {
		t.Errorf("error expected and got none")
	}
}
