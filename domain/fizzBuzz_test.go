package main

import (
    "testing"
)

func TestFizzBuzz(t *testing.T) {
    ans := fizzbuzz(1)
    if fizzbuzz(1) != "1" {
        t.Fatal("failed test", ans, "1")
    }
}

// 参考
// https://elliotchance.medium.com/go-data-driven-or-parameterized-tests-cd2b62b21c8a

func TestParameterizedTest(t *testing.T) {
	parameters := []struct {
		input int
        expected string
	}{
		{1, "1"},
	}

	for i := range parameters {
		actual := fizzbuzz(parameters[i].input)
		if actual != parameters[i].expected {
			t.Logf("expected%s: , actual:%s", parameters[i].expected, actual)
			t.Fail()
		}
	}
}
