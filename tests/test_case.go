package tests

import (
	"github.com/goravel/framework/testing"

	"grandesdelfutbol/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
