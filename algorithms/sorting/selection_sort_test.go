package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Sort array",
			args: args{
				arr: []int{54, 35, 46, 75, 2, 12},
			},
			want: []int{2, 12, 35, 46, 54, 75},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectionSort(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
