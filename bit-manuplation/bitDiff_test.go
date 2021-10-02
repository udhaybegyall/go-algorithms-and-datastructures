package bitmanuplation

import "testing"

func Test_bitDiff(t *testing.T) {
	type args struct {
		number_one int
		number_two int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{5, 4},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{3, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitDiff(tt.args.number_one, tt.args.number_two); got != tt.want {
				t.Errorf("bitDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
