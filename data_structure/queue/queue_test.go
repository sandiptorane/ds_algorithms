package queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
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
			name: "Happy path",
			fields: fields{
				data: nil,
			},
			args: args{
				value: "A",
			},
			wantSize: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewQueue()

			s.Enqueue(tt.args.value)

			assert.Equal(t, tt.wantSize, s.Size())
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		data []string
	}
	tests := []struct {
		name     string
		fields   fields
		want     string
		wantSize int
		wantErr  error
	}{
		{
			name: "Happy Path",
			fields: fields{
				data: []string{"A", "B"},
			},
			want:     "A",
			wantSize: 1,
			wantErr:  nil,
		},
		{
			name: "Should return error when queue is empty",
			fields: fields{
				data: nil,
			},
			want:     "",
			wantSize: 0,
			wantErr:  errors.New("dequeue error : Queue is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				data: tt.fields.data,
			}

			got, err := q.Dequeue()
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Equalf(t, tt.want, got, "Dequeue()")
			assert.Equal(t, tt.wantSize, q.Size())
		})
	}
}
