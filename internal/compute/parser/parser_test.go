package parser_test

import (
	"testing"

	"github.com/EnrikeM/kvBase/internal/compute/parser"
	"github.com/EnrikeM/kvBase/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestUnitParser(t *testing.T) {
	parser := parser.NewParser()
	require.NotNil(t, parser)

	cases := map[string]struct {
		query     string
		wantResp  domain.Query
		wantError bool
	}{
		"set_positive": {
			query: "SET key val",
			wantResp: domain.Query{
				Method: domain.SET,
				Args:   []string{"key", "val"},
			},
			wantError: false,
		},
		"get_positive": {
			query: "GET key",
			wantResp: domain.Query{
				Method: domain.GET,
				Args:   []string{"key"},
			},
			wantError: false,
		},
		"del_positive": {
			query: "DEL key",
			wantResp: domain.Query{
				Method: domain.DEL,
				Args:   []string{"key"},
			},
			wantError: false,
		},

		"wrong_len": {
			query:     "",
			wantResp:  domain.Query{},
			wantError: true,
		},
		"set_negative": {
			query:     "SET too many args",
			wantResp:  domain.Query{},
			wantError: true,
		},
		"get_negative": {
			query:     "GET too many args",
			wantResp:  domain.Query{},
			wantError: true,
		},
		"del_negative": {
			query:     "DEL too many args",
			wantResp:  domain.Query{},
			wantError: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			data, err := parser.Parse(tc.query)
			require.Equal(t, tc.wantResp, data)
			if tc.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

			}
		})
	}
}
