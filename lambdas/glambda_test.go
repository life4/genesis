package lambdas_test

import (
	"testing"

	"github.com/life4/genesis/lambdas"
	"github.com/matryer/is"
)

func TestAbs(t *testing.T) {
	is := is.New(t)
	is.Equal(lambdas.Abs(2), 2)
	is.Equal(lambdas.Abs(-2), 2)
	is.Equal(lambdas.Abs(0), 0)
	is.Equal(lambdas.Abs(-1.2), 1.2)
}

func TestMin(t *testing.T) {
	is := is.New(t)
	is.Equal(lambdas.Min(2, 3), 2)
	is.Equal(lambdas.Min(3, 2), 2)
	is.Equal(lambdas.Min(-2, 3), -2)
	is.Equal(lambdas.Min(2, -3), -3)
	is.Equal(lambdas.Min(2, 2), 2)
}

func TestMax(t *testing.T) {
	is := is.New(t)
	is.Equal(lambdas.Max(2, 3), 3)
	is.Equal(lambdas.Max(3, 2), 3)
	is.Equal(lambdas.Max(-2, 3), 3)
	is.Equal(lambdas.Max(2, -3), 2)
	is.Equal(lambdas.Max(2, 2), 2)
}

func TestDefault(t *testing.T) {
	is := is.New(t)
	is.Equal(lambdas.Default(3), 0)
	is.Equal(lambdas.Default(int32(3)), int32(0))
	is.Equal(lambdas.Default(int64(3)), int64(0))
	is.Equal(lambdas.Default(0), 0)
	is.Equal(lambdas.Default(3.5), 0.0)
	is.Equal(lambdas.Default("hi"), "")
}
