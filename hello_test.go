package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(test *testing.T, result, expect string) {
		test.Helper()

		if result != expect {
			test.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("Should to say Hello for brazilian peoples", func(test *testing.T) {
		result := Hello("BYS", "brazilian")

		const expect string = "Ola BYS"

		verifyMessage(test, result, expect)
	})

	t.Run("Should to say Hello Anybody while empty", func(test *testing.T) {
		result := Hello("", "english")

		const expect string = "Hello Anybody"

		verifyMessage(test, result, expect)
	})

}
