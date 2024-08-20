package main

import "testing"

func Test_calculator(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				"1+2*3+1",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculator(tt.args.str); got != tt.want {
				t.Errorf("calculator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				"--1*2",
			},
			want: false,
		},
		{
			name: "t2",
			args: args{
				"-11+2",
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				"-11+2a",
			},
			want: false,
		},
		{
			name: "t4",
			args: args{
				"-11+2-",
			},
			want: false,
		},
		{
			name: "t5",
			args: args{
				"-11+2-1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.str); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
