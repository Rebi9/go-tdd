package test

import (
  "testing"
)

func TestHelloWorld(t *testing.T) {
  actual := fuga()
  expected := "fuga"
  if actual != expected {
    t.Errorf("actual %v\nwant %v", actual, expected)
  }
}