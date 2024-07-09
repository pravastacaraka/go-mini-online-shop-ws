package utils

import "testing"

func TestIDR(t *testing.T) {
	type args struct {
		amount any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive test case",
			args: args{
				amount: 10435,
			},
			want: "Rp10.435",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IDR(tt.args.amount); got != tt.want {
				t.Errorf("IDR() = %v, want %v", got, tt.want)
			}
		})
	}
}
