package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateArray(t *testing.T) {
	testCases := []struct {
		name   string
		entity any

		wantVals []any
		wantErr  error
	}{
		{
			name:   "[]int",
			entity: [3]int{1, 2, 3},

			wantVals: []any{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vals, err := IterateArrayOrSlice(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVals, vals)
		})

	}
}

func TestIterateSlice(t *testing.T) {
	testCases := []struct {
		name   string
		entity any

		wantVals []any
		wantErr  error
	}{
		{
			name:   "[]int",
			entity: [3]int{1, 2, 3},

			wantVals: []any{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vals, err := IterateArrayOrSlice(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVals, vals)
		})

	}
}

func TestIterateMap(t *testing.T) {
	testCases := []struct {
		name string

		entity   any
		wantKeys []any
		wantVals []any
		wantErr  error
	}{
		{
			name: "map",
			entity: map[string]string{
				"A": "a",
				"B": "b",
			},
			wantKeys: []any{"A", "B"},
			wantVals: []any{"a", "b"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, values, err := IterateMap(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, len(tc.wantKeys), len(keys))
			assert.EqualValues(t, tc.wantKeys, keys)
			assert.EqualValues(t, tc.wantVals, values)

		})
	}
}
