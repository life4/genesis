package lambdas_test

import (
	"testing"

	"github.com/life4/genesis/lambdas"
	"github.com/matryer/is"
)

func TestEqualTo(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.EqualTo(2)(3))
	is.True(!lambdas.EqualTo(3)(2))
	is.True(!lambdas.EqualTo("ab")("cd"))

	is.True(lambdas.EqualTo(2)(2))
	is.True(lambdas.EqualTo("")(""))
	is.True(lambdas.EqualTo("ab")("ab"))
}

func TestLessThan(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.LessThan(2)(3))
	is.True(!lambdas.LessThan(2)(2))
	is.True(lambdas.LessThan(3)(2))
}

func TestNot(t *testing.T) {
	is := is.New(t)
	l := lambdas.Not(lambdas.EqualTo(2))
	is.True(!l(2))
	is.True(l(3))
	is.True(l(-1))
}

func TestIsDefault(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.IsDefault(1))
	is.True(!lambdas.IsDefault("hi"))

	is.True(lambdas.IsDefault(""))
	is.True(lambdas.IsDefault(0))
}
