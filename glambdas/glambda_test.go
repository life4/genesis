package glambdas_test

import (
	"testing"

	"github.com/life4/genesis/glambdas"
	"github.com/stretchr/testify/require"
)

func TestAbs(t *testing.T) {
	is := require.New(t)
	is.Equal(glambdas.Abs(2), 2)
	is.Equal(glambdas.Abs(-2), 2)
	is.Equal(glambdas.Abs(0), 0)
	is.Equal(glambdas.Abs(-1.2), 1.2)
}

func TestMin(t *testing.T) {
	is := require.New(t)
	is.Equal(glambdas.Min(2, 3), 2)
	is.Equal(glambdas.Min(3, 2), 2)
	is.Equal(glambdas.Min(-2, 3), -2)
	is.Equal(glambdas.Min(2, -3), -3)
	is.Equal(glambdas.Min(2, 2), 2)
}

func TestMax(t *testing.T) {
	is := require.New(t)
	is.Equal(glambdas.Max(2, 3), 3)
	is.Equal(glambdas.Max(3, 2), 3)
	is.Equal(glambdas.Max(-2, 3), 3)
	is.Equal(glambdas.Max(2, -3), 2)
	is.Equal(glambdas.Max(2, 2), 2)
}

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
