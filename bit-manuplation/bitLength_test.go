package bitmanuplation

import "testing"

func Test_bitLength(t *testing.T) {
	type args struct {
		number int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{3},
			want: 2,
		},
		{
			name: "test 2",
			args: args{16},
			want: 5,
		},
		{
			name: "test 3",
			args: args{5},
			want: 3,
		},
		{
			name: "test 4",
			args: args{6},
			want: 3,
		},
		{
			name: "test 5",
			args: args{8},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitLength(tt.args.number); got != tt.want {
				t.Errorf("bitLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
