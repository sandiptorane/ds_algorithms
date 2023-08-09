package stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		data []string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
	}{
		{
			name: "push happy path",
			args: args{
				value: "A",
			},
			wantSize: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()

			s.Push(tt.args.value)

			assert.Equal(t, s.Size(), tt.wantSize)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		data []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr error
	}{
		{
			name: "Happy path",
			fields: fields{
				data: []string{"A", "B"},
			},
			want:    "B",
			wantErr: nil,
		},
		{
			name: "Should fail when stack is empty",
			fields: fields{
				data: nil,
			},
			want:    "",
			wantErr: errors.New("pop error : Stack is empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			// add test data
			for _, d := range tt.fields.data {
				s.Push(d)
			}

			got, err := s.Pop()
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
