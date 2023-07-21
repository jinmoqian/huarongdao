package game

import (
	"fmt"
	"testing"
)

func boardLineGraph() ([5][5]*BoardLinkNode, map[BoardHash]*BoardLinkNode) {
	g := New(0, 0, 2)
	boards := [5][5]*BoardLinkNode{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board := &Board{game: &g}
			no1 := byte((i*5+j)/8) + 8
			no2 := byte((i*5+j)%8) + no1 + 1
			board.Put(2, 2, 0)
			board.Put(1, 1, no1)
			board.Put(1, 1, no2)
			fmt.Println("(", i, j, ")", no1, no2)
			boards[i][j] = &BoardLinkNode{
				Uint32Val: g.Hash(board),
				Board:     board,
				Depth:     i*5 + j + 1,
			}
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i > 0 {
				boards[i][j].Nexts = append(boards[i][j].Nexts, g.Hash(boards[i-1][j].Board))
			}
			if i < 4 {
				boards[i][j].Nexts = append(boards[i][j].Nexts, g.Hash(boards[i+1][j].Board))
			}
			if j > 0 {
				boards[i][j].Nexts = append(boards[i][j].Nexts, g.Hash(boards[i][j-1].Board))
			}
			if j < 4 {
				boards[i][j].Nexts = append(boards[i][j].Nexts, g.Hash(boards[i][j+1].Board))
			}
		}
	}
	ret := make(map[BoardHash]*BoardLinkNode)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ret[boards[i][j].Uint32Val] = boards[i][j]
		}
	}
	return boards, ret
}
func Test_reduct(t *testing.T) {
	var batch float64 = 1.0
	batchAdjuster := func(incr bool, ratio float64, _ int) float64 {
		if incr {
			batch++
			return batch
		} else {
			if batch > 1 {
				batch--
			}
			return batch
		}
	}
	boards, grpah := boardLineGraph()
	edges := Reduct(grpah, 24.0/25.0, batch,
		func(_ map[BoardHash]struct{}) *BoardHash {
			return &boards[0][0].Uint32Val
		}, batchAdjuster)
	if len(edges) != 22 {
		t.Error(len(edges))
	}
	for k := range edges {
		if k == boards[0][0].Uint32Val || k == boards[0][1].Uint32Val || k == boards[1][0].Uint32Val {
			t.Error(k)
		}
	}

	var index int
	batch = 1.0
	boards, grpah = boardLineGraph()
	edges = Reduct(grpah, 20.0/25.0, batch,
		func(_ map[BoardHash]struct{}) *BoardHash {
			defer func() { index++ }()
			return []*BoardHash{&boards[0][0].Uint32Val, &boards[4][4].Uint32Val}[index]
		}, batchAdjuster)
	if len(edges) != 19 {
		t.Error(len(edges))
	}
	for k := range edges {
		if k == boards[0][0].Uint32Val || k == boards[0][1].Uint32Val || k == boards[1][0].Uint32Val ||
			k == boards[4][4].Uint32Val || k == boards[4][3].Uint32Val || k == boards[3][4].Uint32Val {
			t.Error(k)
		}
	}

	index = 0
	batch = 1.0
	boards, grpah = boardLineGraph()
	edges = Reduct(grpah, 14.0/25.0, batch,
		func(_ map[BoardHash]struct{}) *BoardHash {
			defer func() { index++ }()
			return []*BoardHash{&boards[0][0].Uint32Val, &boards[4][4].Uint32Val, &boards[2][2].Uint32Val}[index]
		}, batchAdjuster)
	if len(edges) != 14 {
		t.Error(len(edges))
	}
	for k := range edges {
		if k == boards[0][0].Uint32Val || k == boards[0][1].Uint32Val || k == boards[1][0].Uint32Val ||
			k == boards[4][4].Uint32Val || k == boards[4][3].Uint32Val || k == boards[3][4].Uint32Val ||
			k == boards[1][2].Uint32Val || k == boards[2][1].Uint32Val || k == boards[2][2].Uint32Val || k == boards[2][3].Uint32Val || k == boards[3][2].Uint32Val {
			t.Error(k)
		}
	}
	index = 0
	batch = 1.0
	boards, grpah = boardLineGraph()
	edges = Reduct(grpah, 1.0/25.0, batch,
		func(_ map[BoardHash]struct{}) *BoardHash {
			defer func() { index++ }()
			if index >= 30 {
				return nil
			} else if index >= 5 {
				return &boards[2][0].Uint32Val
			}
			return []*BoardHash{&boards[0][0].Uint32Val, &boards[4][4].Uint32Val, &boards[2][2].Uint32Val, &boards[0][4].Uint32Val, &boards[4][0].Uint32Val}[index]
		}, batchAdjuster)
	if len(edges) != 8 {
		t.Error(len(edges))
	}
	for k := range edges {
		if k == boards[0][0].Uint32Val || k == boards[0][1].Uint32Val || k == boards[1][0].Uint32Val ||
			k == boards[4][4].Uint32Val || k == boards[4][3].Uint32Val || k == boards[3][4].Uint32Val ||
			k == boards[1][2].Uint32Val || k == boards[2][1].Uint32Val || k == boards[2][2].Uint32Val || k == boards[2][3].Uint32Val || k == boards[3][2].Uint32Val ||
			k == boards[0][3].Uint32Val || k == boards[0][4].Uint32Val || k == boards[1][4].Uint32Val ||
			k == boards[3][0].Uint32Val || k == boards[4][0].Uint32Val || k == boards[4][1].Uint32Val {
			t.Error(k)
		}
	}
}

// func Test_Search(t *testing.T) {
// 	g := New(0, 0, 0)
// 	b := &Board{game: &g}
// 	b.Put(2, 2, 0)
// 	sb := &Board{game: &g}
// 	sb.Put(2, 2, 1)
// 	Search(
// 		map[BoardHash]byte{g.Hash(b): 4},
// 		&g, sb)
// }
