package main

import "testing"

func Test_sss(t *testing.T) {
	type args struct {
		operations []int
		n          int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				operations: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				n:          66,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sss(tt.args.operations, tt.args.n)
		})
	}
}

func Test_cal73(t *testing.T) {
	type args struct {
		seats int
		arr   []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				seats: 66,
				arr:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal73(tt.args.seats, tt.args.arr)
		})
	}
}
