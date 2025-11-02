package main

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{	
			name: "1+2=3",
			args: args{x: 1, y:2},
			want: 3, 
		},
		{	
			name: "2+3=5",
			args: args{x: 2, y:3},
			want: 5, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		y    int
		want float32
	}{
		// TODO: Add test cases.
		{	
			name: "1/2=0.5",
			x: 1,
			y: 2,
			want: 0.5,
		},
		{	
			name: "2/0=0",
			x: 2,
			y: 0,
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Divide(tt.x, tt.y)
			// compare got with expected value
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
