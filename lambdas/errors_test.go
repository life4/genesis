package lambdas_test

import (
	"errors"
	"testing"

	"github.com/life4/genesis/lambdas"
	"github.com/matryer/is"
)

func TestMust(t *testing.T) {
	is := is.New(t)
	f := func() (int, error) { return 13, nil }
	res := lambdas.Must(f())
	is.Equal(res, 13)
}

func TestEnsure(t *testing.T) {
	f := func() error { return nil }
	lambdas.Ensure(f())
}

func TestSafe(t *testing.T) {
	is := is.New(t)
	f1 := func() (int, error) { return 13, nil }
	is.Equal(lambdas.Safe(f1()), 13)
	f2 := func() (int, error) { return 13, errors.New("hi") }
	is.Equal(lambdas.Safe(f2()), 0)
}

func TestDefaultTo(t *testing.T) {
	is := is.New(t)
	f1 := func() (int, error) { return 13, nil }
	is.Equal(lambdas.DefaultTo(4)(f1()), 13)
	f2 := func() (int, error) { return 13, errors.New("hi") }
	is.Equal(lambdas.DefaultTo(4)(f2()), 4)
}
