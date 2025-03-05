package compute_test

import (
	"fmt"
	"testing"

	compute "github.com/EnrikeM/kvBase/internal/compute"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestUnitComputePostitive(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	logger := zaptest.NewLogger(t)
	parse := NewMockParser(ctr)
	srvc := compute.NewService(logger, parse)

	parse.EXPECT().Parse("SET key val").Return(compute.Query{
		Method: compute.SET,
		Args:   []string{"key", "val"},
	}, nil)

	query, err := srvc.HandleQuery("SET key val")
	require.NoError(t, err)
	require.Equal(t, query, compute.Query{
		Method: compute.SET,
		Args:   []string{"key", "val"},
	})
}

func TestUnitComputeNegative(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	logger := zaptest.NewLogger(t)
	parse := NewMockParser(ctr)
	srvc := compute.NewService(logger, parse)

	parse.EXPECT().Parse(gomock.Any()).Return(compute.Query{}, fmt.Errorf("bad query"))

	_, err := srvc.HandleQuery(gomock.Any().String())
	require.Error(t, err)
}
