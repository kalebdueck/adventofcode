package day11

import (
	"testing"
)

func Test_countOccupiedAdjacency(t *testing.T) {
	type args struct {
		row     int
		col     int
		seating [][]rune
	}
	tests := []struct {
		name         string
		args         args
		wantAdjacent int
	}{
		{
			name: "full house, middle",
			args: args{
				row: 1,
				col: 1,
				seating: [][]rune{
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
				},
			},
			wantAdjacent: 8,
		},
		{
			name: "full house, top left",
			args: args{
				row: 0,
				col: 0,
				seating: [][]rune{
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
		{
			name: "full house, bottom left",
			args: args{
				row: 2,
				col: 0,
				seating: [][]rune{
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
		{
			name: "full house, bottom right",
			args: args{
				row: 2,
				col: 0,
				seating: [][]rune{
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
					[]rune{'#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
		{
			name: "empty house, middle",
			args: args{
				row: 2,
				col: 0,
				seating: [][]rune{
					[]rune{'L', 'L', 'L'},
					[]rune{'L', 'L', 'L'},
					[]rune{'L', 'L', 'L'},
				},
			},
			wantAdjacent: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdjacent := countOccupiedAdjacency(tt.args.row, tt.args.col, tt.args.seating); gotAdjacent != tt.wantAdjacent {
				t.Errorf("countOccupiedAdjacency() = %v, want %v", gotAdjacent, tt.wantAdjacent)
			}
		})
	}
}
