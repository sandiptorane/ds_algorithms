package linked_list

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleList_AddFront(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SingleList
	}{
		{
			name: "Happy path when list is empty",
			fields: fields{
				length: 0,
				Head:   nil,
			},
			args: args{
				data: 11,
			},
			want: &SingleList{
				length: 1,
				Head:   &Node{Data: 11},
			},
		},
		{
			name: "Happy path when list have data",
			fields: fields{
				length: 1,
				Head:   &Node{Data: 2},
			},
			args: args{
				data: 1,
			},
			want: &SingleList{
				length: 2,
				Head:   &Node{Data: 1, Next: &Node{Data: 2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			s.AddFront(tt.args.data)

			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSingleList_AddBack(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}

	type args struct {
		data int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SingleList
	}{
		{
			name: "Happy path when list is empty",
			fields: fields{
				length: 0,
				Head:   nil,
			},
			args: args{
				data: 11,
			},
			want: &SingleList{
				length: 1,
				Head:   &Node{Data: 11},
			},
		},
		{
			name: "Happy path when list have data",
			fields: fields{
				length: 2,
				Head:   &Node{Data: 1, Next: &Node{Data: 2}},
			},
			args: args{
				data: 3,
			},
			want: &SingleList{
				length: 3,
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewList()

			s.length = tt.fields.length
			s.Head = tt.fields.Head

			s.AddBack(tt.args.data)

			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSingleList_RemoveFront(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
		want    *SingleList
	}{
		{
			name: "Happy path when only  head present",
			fields: fields{
				length: 1,
				Head:   &Node{Data: 1},
			},
			wantErr: nil,
			want:    &SingleList{},
		},
		{
			name: "Happy path when two element present",
			fields: fields{
				length: 2,
				Head:   &Node{Data: 1, Next: &Node{Data: 2}},
			},
			wantErr: nil,
			want: &SingleList{length: 1,
				Head: &Node{Data: 2}},
		},
		{
			name: "should return error when empty list",
			fields: fields{
				length: 0,
				Head:   nil,
			},
			wantErr: errors.New("list is empty"),
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			err := s.RemoveFront()
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSingleList_RemoveBack(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
		want    *SingleList
	}{
		{
			name: "Happy path when only  head present",
			fields: fields{
				length: 1,
				Head:   &Node{Data: 1},
			},
			wantErr: nil,
			want:    &SingleList{},
		},
		{
			name: "Happy path when more elements present",
			fields: fields{
				length: 3,
				Head:   &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3}}},
			},
			wantErr: nil,
			want: &SingleList{
				length: 2,
				Head:   &Node{Data: 1, Next: &Node{Data: 2}},
			},
		},
		{
			name: "should return error when empty list",
			fields: fields{
				length: 0,
				Head:   nil,
			},
			wantErr: errors.New("removeBack : List is empty"),
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			err := s.RemoveBack()
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSingleList_Remove(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
		args    args
		want    *SingleList
	}{
		{
			name: "Happy path when only  head present",
			fields: fields{
				length: 1,
				Head:   &Node{Data: 1},
			},
			args:    args{value: 1},
			wantErr: nil,
			want:    &SingleList{},
		},
		{
			name: "Happy path when more element present",
			fields: fields{
				length: 4,
				Head:   &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4}}}},
			},
			wantErr: nil,
			args:    args{value: 3},
			want: &SingleList{
				length: 3,
				Head:   &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 4}}},
			},
		},
		{
			name: "should return error when empty list",
			fields: fields{
				length: 0,
				Head:   nil,
			},
			wantErr: errors.New("remove : List is empty"),
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			err := s.Remove(tt.args.value)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSingleList_Traverse(t *testing.T) {
	type fields struct {
		length int
		Head   *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "Happy path",
			fields: fields{
				length: 3,
				Head:   &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3}}},
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleList{
				length: tt.fields.length,
				Head:   tt.fields.Head,
			}

			assert.Equalf(t, tt.want, s.Traverse(), "Traverse()")
		})
	}
}
