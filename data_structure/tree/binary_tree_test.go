package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_Insert(t *testing.T) {
	type fields struct {
		root *BinaryNode
	}
	type args struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BinaryTree
	}{
		{
			name: "Happy path",
			fields: fields{
				root: nil,
			},

			args: args{
				data: []int{2, 1, -1, 5, 6},
			},
			want: &BinaryTree{
				root: &BinaryNode{
					left:  &BinaryNode{data: 1, left: &BinaryNode{data: -1}},
					right: &BinaryNode{data: 5, right: &BinaryNode{data: 6}},
					data:  2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				root: tt.fields.root,
			}

			var newTree *BinaryTree
			for _, d := range tt.args.data {
				newTree = b.Insert(d)
			}

			// print inserted data
			b.Print()

			assert.Equal(t, tt.want, newTree)
		})
	}
}
