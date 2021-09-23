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