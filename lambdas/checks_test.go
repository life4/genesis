package lambdas_test

import (
	"math"
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

func TestIsZero(t *testing.T) {
	is := is.New(t)
	is.True(lambdas.IsZero(int(0)))
	is.True(lambdas.IsZero(int32(0)))
	is.True(lambdas.IsZero(int64(0)))
	is.True(lambdas.IsZero(uint64(0)))
	is.True(lambdas.IsZero(float64(0)))
	is.True(lambdas.IsZero(float32(0)))

	is.True(!lambdas.IsZero(int(1)))
	is.True(!lambdas.IsZero(int64(1)))
	is.True(!lambdas.IsZero(float64(13)))
	is.True(!lambdas.IsZero(-1))
}

func TestIsNotZero(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.IsNotZero(int(0)))
	is.True(!lambdas.IsNotZero(int32(0)))
	is.True(!lambdas.IsNotZero(int64(0)))
	is.True(!lambdas.IsNotZero(uint64(0)))
	is.True(!lambdas.IsNotZero(float64(0)))
	is.True(!lambdas.IsNotZero(float32(0)))

	is.True(lambdas.IsNotZero(int(1)))
	is.True(lambdas.IsNotZero(int64(1)))
	is.True(lambdas.IsNotZero(float64(13)))
	is.True(lambdas.IsNotZero(-1))
}

func TestIsEmpty(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.IsEmpty([]any{1}))
	is.True(lambdas.IsEmpty([]any{}))
	is.True(lambdas.IsEmpty[[]any](nil))
}

func TestIsNotEmpty(t *testing.T) {
	is := is.New(t)
	is.True(lambdas.IsNotEmpty([]any{1}))
	is.True(!lambdas.IsNotEmpty([]any{}))
	is.True(!lambdas.IsNotEmpty[[]any](nil))
}

func TestIsDefault(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.IsDefault(1))
	is.True(!lambdas.IsDefault("hi"))

	is.True(lambdas.IsDefault(""))
	is.True(lambdas.IsDefault(0))
}

func TestIsNotDefault(t *testing.T) {
	is := is.New(t)
	is.True(lambdas.IsNotDefault(1))
	is.True(lambdas.IsNotDefault("hi"))

	is.True(!lambdas.IsNotDefault(""))
	is.True(!lambdas.IsNotDefault(0))
}

func TestIsNil(t *testing.T) {
	is := is.New(t)
	var v1 *int
	is.True(lambdas.IsNil(v1))
	var v2 int = 2
	is.True(!lambdas.IsNil(&v2))
	var v3 []int
	is.True(!lambdas.IsNil(&v3))
}

func TestIsNaN(t *testing.T) {
	is := is.New(t)
	is.True(!lambdas.IsNaN(1.3))
	is.True(!lambdas.IsNaN(1.3))
	is.True(!lambdas.IsNaN(1))
	is.True(!lambdas.IsNaN(int64(1)))
	is.True(lambdas.IsNaN(float64(math.NaN())))
}

func TestIsNotNaN(t *testing.T) {
	is := is.New(t)
	is.True(lambdas.IsNotNaN(1.3))
	is.True(lambdas.IsNotNaN(1.3))
	is.True(lambdas.IsNotNaN(1))
	is.True(lambdas.IsNotNaN(int64(1)))
	is.True(!lambdas.IsNotNaN(float64(math.NaN())))
}

func TestIsNotNil(t *testing.T) {
	is := is.New(t)
	var v1 *int
	is.True(!lambdas.IsNotNil(v1))
	var v2 int = 2
	is.True(lambdas.IsNotNil(&v2))
	var v3 []int
	is.True(lambdas.IsNotNil(&v3))
}
