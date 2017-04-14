package golife

import (
	"reflect"
	"testing"
)

func TestGetHP(t *testing.T) {
	board := Board{
		state: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		w:     3,
		h:     3,
	}

	tests := []struct {
		name string
		x, y int
		want int
	}{
		{
			name: "simple get",
			x:    1,
			y:    1,
			want: 5,
		}, {
			name: "wrap left",
			x:    -1,
			y:    0,
			want: 3,
		}, {
			name: "super wrap left",
			x:    -101,
			y:    0,
			want: 2,
		}, {
			name: "wrap right",
			x:    5,
			y:    0,
			want: 3,
		}, {
			name: "wrap top",
			x:    0,
			y:    -1,
			want: 7,
		}, {
			name: "wrap bottom",
			x:    0,
			y:    3,
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := board.getHP(test.x, test.y)
			if got != test.want {
				t.Errorf("%v.getHP(%d, %d) = %d, want %d", board, test.x, test.y, got, test.want)
			}
		})
	}
}

func TestLiveAround(t *testing.T) {
	board := Board{
		state: [][]int{
			{1, 2, 0, 0},
			{2, 2, 2, 0},
			{0, 1, 2, 2},
			{2, 1, 0, 2}},
		w:     4,
		h:     4,
		maxHP: 2,
	}

	tests := []struct {
		name string
		x, y int
		want int
	}{
		{
			name: "count at edge",
			x:    0,
			y:    0,
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := board.liveAround(test.x, test.y)
			if got != test.want {
				t.Errorf("%v.liveAround(%d, %d) = %d, want %d", board, test.x, test.y, got, test.want)
			}
		})
	}
}

func TestLifeStep(t *testing.T) {
	life := NewLife(5, 5, 2)
	life.board.state = [][]int{
		{0, 2, 0, 0, 0},
		{0, 0, 2, 0, 0},
		{2, 2, 2, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// wantStates is a sequence of states we should expect to see if we repeatedly
	// update. This sequence is a single cycle of a glider.
	wantStates := [][][]int{
		{
			{0, 1, 0, 0, 0},
			{2, 0, 2, 0, 0},
			{1, 2, 2, 0, 0},
			{0, 2, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}, {
			{0, 0, 0, 0, 0},
			{1, 0, 2, 0, 0},
			{2, 1, 2, 0, 0},
			{0, 2, 2, 0, 0},
			{0, 0, 0, 0, 0},
		}, {
			{0, 0, 0, 0, 0},
			{0, 2, 1, 0, 0},
			{1, 0, 2, 2, 0},
			{0, 2, 2, 0, 0},
			{0, 0, 0, 0, 0},
		}, {
			{0, 0, 0, 0, 0},
			{0, 1, 2, 0, 0},
			{0, 0, 1, 2, 0},
			{0, 2, 2, 2, 0},
			{0, 0, 0, 0, 0},
		},
	}

	for i, want := range wantStates {
		life.Step()
		if !reflect.DeepEqual(want, life.board.state) {
			t.Errorf("Unexpected state %d, want %v, got %v", i, want, life.board.state)
		}
	}
}
