// golife is a library of functions that implement Conway's Game of Life, and
// display it using termbox-go.
//
// Unlike traditional Game Of Life, it keeps track of 'health points' (or HP) on
// a per-cell basis. Cells maintain their maximum health with the usual
// Game-of-Life rules, but when they die, they don't die immediately, instead
// they start to loose HP at a rate of 1 per update.
//
// Assumes a that the board wraps around at the edges.
package golife

// Board keeps the state of a single board.
type Board struct {
	state [][]int
	w, h  int
	maxHP int
}

// newBoard allocates a new board.
func newBoard(w, h, maxHP int) *Board {
	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
	}
	return &Board{state: s, w: w, h: h, maxHP: maxHP}
}

// getHP gets the HP at (x,y), wrapping around
func (b *Board) getHP(x, y int) int {
	for ; x < 0; x += b.w {
	}
	for ; y < 0; y += b.h {
	}
	return b.state[y%b.h][x%b.w]
}

// liveAround(x,y) counts the number of cells that are live around a certain
// cell in the board.
func (b *Board) liveAround(x, y int) (cnt int) {
	for i := (x - 1); i <= x+1; i++ {
		for j := (y - 1); j <= y+1; j++ {
			if (x == i) && (y == j) {
				continue
			}
			if b.getHP(i, j) == b.maxHP {
				cnt++
			}
		}
	}
	return
}

// next determines the next value for the cell
func (b *Board) next(x, y int) int {
	cnt := b.liveAround(x, y)
	if cnt == 3 {
		return b.maxHP
	}
	if cnt == 2 && b.getHP(x, y) == b.maxHP {
		return b.maxHP
	}
	hp := b.getHP(x, y)
	if hp <= 0 {
		return 0
	}
	return hp - 1
}

// update sets a cell of the board to a value
func (b *Board) update(x, y int, next int) {
	b.state[y][x] = next
}

// Life keeps the state of the game. Maintains two boards, and swaps them at
// every update.
type Life struct {
	board, buffer *Board
	w, h          int
}

func NewLife(w, h, maxHP int) *Life {
	return &Life{
		board:  newBoard(w, h, maxHP),
		buffer: newBoard(w, h, maxHP),
		w:      w,
		h:      h,
	}
}

// Step steps the game of life, swapping the buffers
func (l *Life) Step() {
	for x := 0; x < l.w; x++ {
		for y := 0; y < l.h; y++ {
			l.buffer.update(x, y, l.board.next(x, y))
		}
	}
	l.buffer, l.board = l.board, l.buffer
}
