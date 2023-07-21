package game

type state struct {
	ops optionSet
	cur int
	n   int
}

type Iterator struct {
	game   *Game
	Board  *Board
	states [15]state
	all    int
	pieces [15][2]int
	firsts [4]int
}

func (i *Iterator) init() bool {
	i.Board = &Board{game: i.game}
	i.all = 1 + i.game.hor + i.game.ver + i.game.little
	var pieceIdx int
	for iteIdx, ite := range []struct {
		w int
		h int
		n int
	}{
		{w: 2, h: 2, n: 1},
		{w: 2, h: 1, n: i.game.hor},
		{w: 1, h: 2, n: i.game.ver},
		{w: 1, h: 1, n: i.game.little},
	} {
		i.firsts[iteIdx] = pieceIdx
		for c := 0; c < ite.n; c++ {
			i.pieces[pieceIdx][0] = ite.w
			i.pieces[pieceIdx][1] = ite.h
			pieceIdx++
		}
	}
	for j := 0; j < pieceIdx; j++ {
		i.states[j].n = i.Board.search(i.pieces[j][0], i.pieces[j][1], 0, &i.states[j].ops)
		if i.states[j].n == 0 {
			return false
		}
		i.Board.Put(i.pieces[j][0], i.pieces[j][1], i.states[j].ops[0])
	}
	return true
}
func (i *Iterator) Next() bool {
	if i.Board == nil {
		return i.init()
	}

	cur := i.all - 1
	var willPop bool = true
	for {
		if willPop {
			i.Board.pop()
		}
		if i.states[cur].cur == i.states[cur].n-1 {
			if cur == 0 {
				return false
			}
			i.states[cur].cur = -1
			cur--
			willPop = true
		} else {
			c := i.states[cur].cur + 1
			i.Board.Put(i.pieces[cur][0], i.pieces[cur][1], i.states[cur].ops[c])
			i.states[cur].cur = c

			if cur == i.all-1 {
				break
			}
			cur++
			var from byte
			if cur == i.firsts[3] || cur == i.firsts[2] || cur == i.firsts[1] {
				from = 0
			} else {
				t := i.states[cur-1].cur
				from = i.states[cur-1].ops[t]
			}
			i.states[cur].n = i.Board.search(i.pieces[cur][0], i.pieces[cur][1], from, &i.states[cur].ops)
			i.states[cur].cur = -1
			if i.states[cur].n == 0 {
				cur--
				willPop = true
			} else {
				willPop = false
			}
		}
	}
	return true
}
