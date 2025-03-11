package compute_test

import (
	"fmt"
	"testing"

	"github.com/EnrikeM/kvBase/internal/compute"
	"github.com/EnrikeM/kvBase/internal/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestUnitCompute(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	logger := zaptest.NewLogger(t)
	parser := NewMockParser(ctr)
	srvc := compute.NewService(logger, parser)
	require.NotNil(t, srvc)

	cases := map[string]struct {
		setupMocks func()
		query      string
		wantResp   domain.Query
		wantError  bool
	}{
		"parse_positive": {
			setupMocks: func() {
				gomock.InOrder(
					parser.EXPECT().Parse("SET key val").Return(domain.Query{
						Method: domain.SET,
						Args:   []string{"key", "val"},
					}, nil),
				)
			},
			query: "SET key val",
			wantResp: domain.Query{
				Method: domain.SET,
				Args:   []string{"key", "val"},
			},
			wantError: false,
		},
		"parse_negative": {
			setupMocks: func() {
				gomock.InOrder(
					parser.EXPECT().Parse(gomock.Any()).Return(domain.Query{}, fmt.Errorf("some error")),
				)
			},
			query:     "SET key val",
			wantResp:  domain.Query{},
			wantError: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc.setupMocks()
			data, err := srvc.HandleQuery(tc.query)
			require.Equal(t, tc.wantResp, data)
			if tc.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

			}
		})
	}
}
