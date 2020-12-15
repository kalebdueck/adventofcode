package day11

import "testing"

func Test_countOccupiedSightLines(t *testing.T) {
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
				row: 20,
				col: 2,
				seating: [][]rune{
					[]rune{'#', '.', '#', '#', '#'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'#', '.', '#', '.', '#'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
					[]rune{'#', '#', '#', '#', '#'},
				},
			},
			wantAdjacent: 8,
		},
		{
			name: "full house, middle",
			args: args{
				row: 2,
				col: 2,
				seating: [][]rune{
					[]rune{'#', '.', '#', '#', '#'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'#', '.', '#', '.', '#'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'#', '#', '#', '#', '#'},
				},
			},
			wantAdjacent: 8,
		},
		{
			name: "full house, top right",
			args: args{
				row: 0,
				col: 4,
				seating: [][]rune{
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'#', '#', '#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
		{
			name: "full house, corner",
			args: args{
				row: 4,
				col: 4,
				seating: [][]rune{
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'#', '#', '#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
		{
			name: "full house, bottom right",
			args: args{
				row: 4,
				col: 0,
				seating: [][]rune{
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'#', '#', '#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},

		{
			name: "full house, rando",
			args: args{
				row: 4,
				col: 2,
				seating: [][]rune{
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '#', '.', '.'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '#', '#', '#'},
				},
			},
			wantAdjacent: 4,
		},
		{
			name: "full house, midish",
			args: args{
				row: 3,
				col: 1,
				seating: [][]rune{
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '#', '.', '.'},
					[]rune{'#', '.', '.', '.', '#'},
					[]rune{'.', '.', '.', '.', '.'},
					[]rune{'.', '.', '#', '#', '#'},
				},
			},
			wantAdjacent: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdjacent := countOccupiedSightLines(tt.args.row, tt.args.col, tt.args.seating); gotAdjacent != tt.wantAdjacent {
				t.Errorf("countOccupiedSightLines() = %v, want %v", gotAdjacent, tt.wantAdjacent)
			}
		})
	}
}
