package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(test *testing.T, result, expect string) {
		test.Helper()

		if result != expect {
			test.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("Should to say Hello for peoples", func(test *testing.T) {
		result := Hello("BYS")

		const expect string = "Hello BYS"

		verifyMessage(test, result, expect)
	})

	t.Run("Should to say Hello Anybody while empty", func(test *testing.T) {
		result := Hello("")

		const expect string = "Hello Anybody"

		verifyMessage(test, result, expect)
	})

}

// For tests
/*
	archive must have name with _test.go
	The test function must to start with Test
	The test function get a only arg (test *testing.T)

	test.Helper is a auxiliary method for tests
*/

// while var to use ":="
// while const to use "="
