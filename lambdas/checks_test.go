package lambdas_test

import (
	"testing"

	"github.com/life4/genesis/lambdas"
	"github.com/stretchr/testify/require"
)

func TestEqualTo(t *testing.T) {
	is := require.New(t)
	is.False(lambdas.EqualTo(2)(3))
	is.False(lambdas.EqualTo(3)(2))
	is.False(lambdas.EqualTo("ab")("cd"))

	is.True(lambdas.EqualTo(2)(2))
	is.True(lambdas.EqualTo("")(""))
	is.True(lambdas.EqualTo("ab")("ab"))
}

func TestLessThan(t *testing.T) {
	is := require.New(t)
	is.False(lambdas.LessThan(2)(3))
	is.False(lambdas.LessThan(2)(2))
	is.True(lambdas.LessThan(3)(2))
}

func TestNot(t *testing.T) {
	is := require.New(t)
	l := lambdas.Not(lambdas.EqualTo(2))
	is.False(l(2))
	is.True(l(3))
	is.True(l(-1))
}

func TestIsDefault(t *testing.T) {
	is := require.New(t)
	is.False(lambdas.IsDefault(1))
	is.False(lambdas.IsDefault("hi"))

	is.True(lambdas.IsDefault(""))
	is.True(lambdas.IsDefault(0))
}
