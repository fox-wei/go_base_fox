package dynamic

import (
	"testing"
)

func TestKnapsack1(t *testing.T) {
	type args struct {
		weight []int
		n      int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"fox",
			args{[]int{2, 2, 4, 6, 3}, 5, 9},
			9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack2(tt.args.weight, tt.args.n, tt.args.w); got != tt.want {
				t.Errorf("Knapsack2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnapsack3(t *testing.T) {
	type args struct {
		weight []int
		values []int
		n      int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"van",
			args{[]int{3, 2, 1, 4, 5}, []int{25, 20, 15, 40, 50}, 5, 6},
			65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack3(tt.args.weight, tt.args.values, tt.args.n, tt.args.w); got != tt.want {
				t.Errorf("Knapsack3() = %v, want %v", got, tt.want)
			}
		})
	}
}
