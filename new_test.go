package gonita

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want *BPMClient
	}{
		{
			name: "",
			args: args{
				username: "",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				username: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
