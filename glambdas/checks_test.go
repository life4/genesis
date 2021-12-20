package glambdas_test

import (
	"testing"

	"github.com/life4/genesis/glambdas"
	"github.com/stretchr/testify/require"
)

func TestEqualTo(t *testing.T) {
	is := require.New(t)
	is.False(glambdas.EqualTo(2)(3))
	is.False(glambdas.EqualTo(3)(2))
	is.False(glambdas.EqualTo("ab")("cd"))

	is.True(glambdas.EqualTo(2)(2))
	is.True(glambdas.EqualTo("")(""))
	is.True(glambdas.EqualTo("ab")("ab"))
}

func TestLessThan(t *testing.T) {
	is := require.New(t)
	is.False(glambdas.LessThan(2)(3))
	is.False(glambdas.LessThan(2)(2))
	is.True(glambdas.LessThan(3)(2))
}

func TestNot(t *testing.T) {
	is := require.New(t)
	l := glambdas.Not(glambdas.EqualTo(2))
	is.False(l(2))
	is.True(l(3))
	is.True(l(-1))
}

func TestIsDefault(t *testing.T) {
	is := require.New(t)
	is.False(glambdas.IsDefault(1))
	is.False(glambdas.IsDefault("hi"))

	is.True(glambdas.IsDefault(""))
	is.True(glambdas.IsDefault(0))
}
