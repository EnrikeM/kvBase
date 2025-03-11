package engine_test

import (
	"testing"

	"github.com/EnrikeM/kvBase/internal/domain"
	"github.com/EnrikeM/kvBase/internal/storage/engine"
	"github.com/stretchr/testify/require"
)

func TestUnitEngine(t *testing.T) {
	engine := engine.NewEngine()
	require.NotNil(t, engine)

	cases := map[string]struct {
		method    domain.Method
		args      []string
		wantResp  string
		wantError bool
	}{
		"set_positive": {
			method:    domain.SET,
			args:      []string{"key", "val"},
			wantResp:  "key key set for val val",
			wantError: false,
		},

		"get_positive": {
			method:    domain.GET,
			args:      []string{"key"},
			wantResp:  "val",
			wantError: false,
		},

		"del_positive": {
			method:    domain.DEL,
			args:      []string{"key"},
			wantResp:  "key deleted",
			wantError: false,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			var data string
			var err error
			switch tc.method {
			case domain.SET:
				data, err = engine.Set(tc.args)
			case domain.GET:
				data, err = engine.Get(tc.args)
			case domain.DEL:
				data, err = engine.Del(tc.args)
			}

			require.Equal(t, tc.wantResp, data)
			if tc.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

			}
		})
	}
}
