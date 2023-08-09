package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy  path found at start",
			args: args{
				arr: []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 124},
				key: 1,
			},
			want: true,
		},
		{
			name: "Happy  path found in between",
			args: args{
				arr: []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 124},
				key: 20,
			},
			want: true,
		},
		{
			name: "Happy  path found at end",
			args: args{
				arr: []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 124},
				key: 124,
			},
			want: true,
		},
		{
			name: "Not found",
			args: args{
				arr: []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 124},
				key: 21,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.args.arr, 0, len(tt.args.arr)-1, tt.args.key)
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
