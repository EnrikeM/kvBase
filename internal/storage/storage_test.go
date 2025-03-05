package storage_test

import (
	"fmt"
	"testing"

	"github.com/EnrikeM/kvBase/internal/compute"
	"github.com/EnrikeM/kvBase/internal/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestUnitStoragePositive(t *testing.T) {

	ctr := gomock.NewController(t)
	defer ctr.Finish()

	comp := NewMockCompute(ctr)
	eng := NewMockEngine(ctr)
	srvc := storage.NewService(eng, comp, zaptest.NewLogger(t))

	require.NotNil(t, srvc)

	cases := map[string]struct {
		setupMocks func()
		wantError  bool
		query      string
		wantResp   string
	}{
		"default_set": {
			setupMocks: func() {
				gomock.InOrder(
					comp.EXPECT().HandleQuery("SET key val").Return(compute.Query{
						Method: compute.SET,
						Args:   []string{"key", "val"},
					}, nil),
					eng.EXPECT().Set([]string{"key", "val"}).Return("key key set for val val", nil),
				)
			},
			query:     "SET key val",
			wantResp:  "key key set for val val",
			wantError: false,
		},
		"bad_handle": {
			setupMocks: func() {
				gomock.InOrder(
					comp.EXPECT().HandleQuery(gomock.Any()).Return(compute.Query{}, fmt.Errorf("wrong num of args")),
					// eng.EXPECT().Set([]string{"key", "val"}).Return("key key set for val val", nil),
				)
			},
			query:     "SET key val val",
			wantResp:  "",
			wantError: true,
		},
		"bad_set": {
			setupMocks: func() {
				gomock.InOrder(
					comp.EXPECT().HandleQuery(gomock.Any()).Return(compute.Query{Method: compute.SET, Args: []string{
						"key", "val", "val",
					}}, nil),
					eng.EXPECT().Set([]string{"key", "val", "val"}).Return("", fmt.Errorf("wrong num of args")),
				)
			},
			query:     "SET key val",
			wantResp:  "",
			wantError: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc.setupMocks()
			data, err := srvc.Update(tc.query)
			if tc.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantResp, data)
			}
		})
	}
}
