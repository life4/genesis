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
