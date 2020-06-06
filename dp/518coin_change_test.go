package main

import "testing"

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				amount: 3,
				coins:  []int{2},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				amount: 10,
				coins:  []int{10},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := change1(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change1() = %v, want %v", got, tt.want)
			}
		})
	}
}
