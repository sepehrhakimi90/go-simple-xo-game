package main

import "testing"

func Test_checkLinearFinish(t *testing.T) {
	type args struct {
		b   []string
		lineIdx int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name: "Line 1 PC Winner",
			args:args{
				b:   []string{"X","X","X","O"," "," "," "," ","O"},
				lineIdx: 0,
			},
			want: true,
			want1: PC,
		}, {
			name: "Line 1 USER Winner",
			args:args{
				b:   []string{"O","O","O","X"," "," "," "," ","X"},
				lineIdx: 0,
			},
			want: true,
			want1: USER,
		},{
			name: "Game did not finish",
			args:args{
				b:   []string{"O","O"," ","X"," "," "," X"," ","X"},
				lineIdx: 0,
			},
			want: false,
			want1: NONE,
		},
		{
			name: "Line 2 PC won. passing the second IDX of row",
			args:args{
				b:   []string{"O","O"," ","X","X","X"," "," "," "},
				lineIdx: 1,
			},
			want: true,
			want1: PC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkLinearFinish(tt.args.b, tt.args.lineIdx)
			if got != tt.want {
				t.Errorf("checkLinearFinish() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkLinearFinish() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkColumnarFinish(t *testing.T) {
	type args struct {
		b      []string
		colIdx int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name: "column 1 PC Winner",
			args:args{
				b:   []string{"X","O"," ","X","O"," ","X"," ","O"},
				colIdx: 0,
			},
			want: true,
			want1: PC,
		},
		{
			name: "column 2 USER Winner",
			args:args{
				b:   []string{"X","O"," ","X","O"," "," ","O"," "},
				colIdx: 1,
			},
			want: true,
			want1: USER,
		},
		{
			name: "column 3 USER Winner",
			args:args{
				b:   []string{"X"," ","O","X"," ","O"," "," ","O"},
				colIdx: 2,
			},
			want: true,
			want1: USER,
		},
		{
			name: "did not finished",
			args:args{
				b:   []string{"X"," "," ","X"," ","O"," "," ","O"},
				colIdx: 2,
			},
			want: false,
			want1: NONE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkColumnarFinish(tt.args.b, tt.args.colIdx)
			if got != tt.want {
				t.Errorf("checkColumnarFinish() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkColumnarFinish() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_checkDiagonalFinish(t *testing.T) {
	type args struct {
		b []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name: "PC Winner",
			args:args{
				b:   []string{"X","O"," ","O","X"," ","O"," ","X"},
			},
			want: true,
			want1: PC,
		},
		{
			name: "USER Winner",
			args:args{
				b:   []string{"X"," ","O"," ","O"," ","O"," ","X"},
			},
			want: true,
			want1: USER,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkDiagonalFinish(tt.args.b)
			if got != tt.want {
				t.Errorf("checkDiagonalFinish() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkDiagonalFinish() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}