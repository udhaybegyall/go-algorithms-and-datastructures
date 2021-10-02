package bitmanuplation

import "testing"

func Test_countSetBits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{20},
			want: 2,
		},
		{
			name: "test 2",
			args: args{3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSetBits(tt.args.number); got != tt.want {
				t.Errorf("countSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
