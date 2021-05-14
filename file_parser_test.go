package main

import (
	 "testing"
)

func TestFileNameLessThan80Chars(t *testing.T) {
	fileName := ParseFile("Microservice Design Patterns")
	if fileName >= 80 {
		t.Errorf("[ERROR] - filename '%s' length greater than 80 chars", fileName)
}
