package test

import "testing"

type Testable interface {
	Test(t *testing.T)
	Name() string
}
