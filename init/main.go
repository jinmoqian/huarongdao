package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"huarongdao/game"
	"os"
	"sort"
)

func findAll(hor, ver, liter, limit int) []byte {
	var thisMaxDepth = 0
	g := game.New(hor, ver, liter)
	ite := g.Iterate()
	uint32Board := make(map[game.BoardHash]*game.Board)
	for ite.Next() {
		val := g.Hash(ite.Board)
		if oldVal, ok := uint32Board[val]; ok {
			panic(fmt.Sprintf("repeated uint32=%v oldVal=%v, newVal=%v", val.Bytes(), oldVal.Cells(), ite.Board.Cells()))
		}
		uint32Board[val] = ite.Board.Copy()
	}
	var head, tail *game.BoardLinkNode
	for uint32Val, board := range uint32Board {
		if board.Win() {
			n := &game.BoardLinkNode{
				Uint32Val: uint32Val,
				Board:     board,
			}
			if head == nil {
				head = n
				tail = n
			} else {
				tail.Next = n
				tail = n
			}
		}
	}
	iterated := make(map[game.BoardHash]*game.BoardLinkNode)
	for {
		if head == nil {
			break
		}
		if head.Depth > thisMaxDepth {
			thisMaxDepth = head.Depth
		}
		if _, ok := uint32Board[head.Uint32Val]; !ok {
			panic(g.StringBoard(head.Board))
		}
		if _, ok := iterated[head.Uint32Val]; !ok {
			iterated[head.Uint32Val] = head
			movs := g.Movable(head.Board)
			for _, mov := range movs {
				newBoards := g.Move(head.Board, mov)
				for i := range newBoards {
					newBoard := newBoards[i]
					nbUint32Val := g.Hash(newBoard)
					n := &game.BoardLinkNode{
						Uint32Val: nbUint32Val,
						Board:     newBoard,
						Depth:     head.Depth + 1,
					}
					head.Nexts = append(head.Nexts, nbUint32Val)
					tail.Next = n
					tail = n
					// fmt.Println(g.StringBoard(head.Board), "=>", g.StringBoard(newBoard))
				}
			}
		}
		head = head.Next
	}
	return game.SaveBoard(iterated)
}
func main() {
	bs := all()
	reduct(bs)
	// stat()
}
func stat() {
	df := &game.DataFile{}
	err := df.Open()
	if err != nil {
		panic(err)
	}

	var all = 0
	stat := make(map[byte]int)
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7-i; j++ {
			for k := 0; k <= 20-4-i*2-j*2-2; k++ {
				hs := df.All(game.FileHead{Hor: i, Ver: j, Little: k})
				for _, d := range hs {
					stat[d]++
					all++
				}
			}
		}
	}
	ks := make([]byte, 0)
	for k := range stat {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] > ks[j]
	})
	var a = 0
	fmt.Printf("\t Count\tPercentage\tAccum\t\tAccum Percentage\n")
	for _, k := range ks {
		a += stat[k]
		fmt.Printf("%d\t%d\t%f%%\t%d\t\t%f%%\n", k, stat[k], float64(stat[k])/float64(all)*100, a, float64(a)/float64(all)*100)
	}
}
func reduct(raw []byte) {
	df := &game.DataFile{}
	err := df.OpenBytes(raw)
	if err != nil {
		panic(err)
	}

	keys := make([]game.MapKey, 0)
	m := game.NewBytesToBytesMap(4, 4)
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7-i; j++ {
			for k := 0; k <= 20-4-i*2-j*2-2; k++ {
				nodes := make(map[game.BoardHash]*game.BoardLinkNode)
				fh := game.FileHead{Hor: i, Ver: j, Little: k}
				hs := df.All(fh)
				for bh, d := range hs {
					if d >= game.StoreLevel {
						nodes[bh] = &game.BoardLinkNode{Depth: int(d)}
					}
				}
				bs := game.SaveBoard(nodes)
				m.Put(fh, game.BytesValue(bs))
				keys = append(keys, fh)
			}
		}
	}
	bs := m.Bytes(keys)
	file, err := os.OpenFile("data/data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(bs)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
func all() []byte {
	keys := make([]game.MapKey, 0)
	m := game.NewBytesToBytesMap(4, 4)
	var counter = 0
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7-i; j++ {
			for k := 0; k <= 20-4-i*2-j*2-2; k++ {
				bs := findAll(i, j, k, 5000)
				key := game.FileHead{Hor: i, Ver: j, Little: k}
				m.Put(key, game.BytesValue(bs))
				keys = append(keys, key)
				fmt.Println("Hor=", i, "Ver=", j, "Litter=", k, "done, len(data)=", len(bs), "maxDepth=")
				counter++
				// if k == 2 {
				// break ALL
				// }
			}
		}
	}
	bs := m.Bytes(keys)
	fmt.Println("All bytes=", len(bs), "counter=", counter)
	dig := md5.Sum(bs)
	fmt.Println(hex.EncodeToString(dig[:]))
	// m2 := game.NewBytesToBytesMap[uint32](4, 4)
	// m2.FromBytes(bs)
	// bs, _ = m2.Get(game.FileHead{Hor: 0, Ver: 0, Little: 0})
	// fmt.Println("bs=", len(bs))
	// bs, _ = m2.Get(game.FileHead{Hor: 0, Ver: 0, Little: 1})
	// fmt.Println("bs=", len(bs))
	// bs, _ = m2.Get(game.FileHead{Hor: 0, Ver: 0, Little: 2})
	// fmt.Println("bs=", len(bs))

	// file, err := os.OpenFile("../data/data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = file.Write(bs)
	// if err != nil {
	// 	panic(err)
	// }
	// err = file.Close()
	// if err != nil {
	// 	panic(err)
	// }
	return bs
}
