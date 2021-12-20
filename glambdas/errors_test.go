package glambdas_test

import (
	"errors"
	"testing"

	"github.com/life4/genesis/glambdas"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	is := require.New(t)
	f := func() (int, error) { return 13, nil }
	res := glambdas.Must(f())
	is.Equal(res, 13)
}

func TestSafe(t *testing.T) {
	is := require.New(t)
	f1 := func() (int, error) { return 13, nil }
	is.Equal(glambdas.Safe(f1()), 13)
	f2 := func() (int, error) { return 13, errors.New("hi") }
	is.Equal(glambdas.Safe(f2()), 0)
}

func TestDefaultTo(t *testing.T) {
	is := require.New(t)
	f1 := func() (int, error) { return 13, nil }
	is.Equal(glambdas.DefaultTo(4)(f1()), 13)
	f2 := func() (int, error) { return 13, errors.New("hi") }
	is.Equal(glambdas.DefaultTo(4)(f2()), 4)
}
