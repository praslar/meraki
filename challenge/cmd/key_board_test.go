package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{
				S: "CrCellBax",
			},
			want: "Relax",
		},
		{
			name: "Case 2",
			args: args{
				S: "CgCoodlBClCuck",
			},
			want: "GoodLuck",
		},
		{
			name: "Case 3",
			args: args{
				S: "aCaBBCCab",
			},
			want: "AB",
		},
		{
			name: "Case 4",
			args: args{
				S: "aBB",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
