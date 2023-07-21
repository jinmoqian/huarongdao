package game

import (
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"strings"
)

const width = 4
const height = 5

// 第0位为1表示水平2格，第1位为1表示垂直2格
const horMask = 0x80
const verMask = 0x40
const valMask = 0x3F

func Value(cell byte) byte {
	return cell & valMask
}

type BoardHash struct {
	hash uint64
	game *Game
}

func (b *BoardHash) Push(shift uint32, value uint32) {
	v := (uint64(b.hash) << uint64(shift)) | uint64(value)
	b.hash = v
}

func (b BoardHash) Bytes() (ret []uint8) {
	var buf [8]uint8
	binary.BigEndian.PutUint64(buf[:], uint64(b.hash))
	ret = make([]uint8, b.game.byteSize)
	copy(ret, buf[8-b.game.byteSize:])
	return
}
func (b *BoardHash) FromByte(bs []uint8) {
	var buf [8]uint8
	copy(buf[8-b.game.byteSize:], bs)
	v := binary.BigEndian.Uint64(buf[:])
	b.hash = v
}

type Board struct {
	game *Game
	n    int
	cell [width*height - 2 - 4 + 1]byte
}

func (b *Board) Cells() [width*height - 2 - 4 + 1]byte {
	var v [width*height - 2 - 4 + 1]byte
	for i, c := range b.cell {
		v[i] = c & valMask
	}
	return v
}

func (b *Board) Copy() *Board {
	nb := *b
	return &nb
}
func (b *Board) Can(w, h int, pos byte) bool {
	var i int
	px, py := toxy(pos)
	px2 := px + byte(w) - 1
	py2 := py + byte(h) - 1
	if px2 >= width {
		return false
	}
	if py2 >= height {
		return false
	}
	var ret, confliced bool
	b.game.loop(func(hor, ver, n int) {
		if confliced || i >= b.n {
			return
		}
		for j := 0; j < n && i < b.n; j++ {
			x, y := toxy(b.cell[i] & valMask)
			x2 := x + byte(hor) - 1
			y2 := y + byte(ver) - 1
			bx1 := px <= x && x <= px2
			bx2 := px <= x2 && x2 <= px2
			by1 := py <= y && y <= py2
			by2 := py <= y2 && y2 <= py2
			ret = !((bx1 || bx2) && (by1 || by2))
			if !ret {
				confliced = true
				break
			}
			i++
		}
	})
	return ret
}
func (b *Board) Put(w, h int, pos byte) {
	if w == 2 {
		pos |= horMask
	}
	if h == 2 {
		pos |= verMask
	}
	b.cell[b.n] = pos
	b.n++
}
func (b *Board) pop() {
	b.n--
}
func (b *Board) Win() bool {
	return b.n > 0 && (b.cell[0]&valMask) == 13
}

type optionSet [16]byte

func (b *Board) search(w, h int, from byte, out *optionSet) int {
	// 这个函数里没有判断曹操
	var pos int
	for i := byte(from); i < width*height; i++ {
		var addw, addh byte
		if w == 2 {
			if i%width == 3 {
				continue
			}
			addw = i + 1
		}
		if h == 2 {
			if i/width > 3 {
				break
			}
			addh = i + width
		}
		var can bool = true
		for j := 0; j < b.n; j++ {
			val := b.cell[j] & valMask
			if (val == i) || (addw > 0 && addw == val) || (addh > 0 && addh == val) {
				can = false
				break
			}
			var bw = false
			if b.cell[j]&horMask != 0 {
				if val+1 == i || (addh > 0 && addh == val+1) {
					can = false
					break
				}
				bw = true
			}
			if b.cell[j]&verMask != 0 {
				if val+width == i || (addw > 0 && addw == val+width) {
					can = false
					break
				}
				if bw {
					if val+width+1 == i {
						can = false
						break
					}
				}
			}
		}
		if can {
			out[pos] = i
			pos++
		}
	}
	return pos
}

type Game struct {
	hor         int
	ver         int
	little      int
	uint32Bytes [2][2]uint32
	byteSize    int
}

func New(hor, ver, lit int) Game {
	g := Game{
		hor:    hor,
		ver:    ver,
		little: lit,
		uint32Bytes: [2][2]uint32{
			{
				uint32(math.Ceil(math.Log2(float64(width*height - 4 - hor*2 - ver*2 - lit + 1)))),
				uint32(math.Ceil(math.Log2(float64(width*height - 4 - hor*2 + 1)))),
			},
			{
				uint32(math.Ceil(math.Log2(float64(width*height - 4 - hor*2 - ver*2 + 1)))),
				uint32(math.Ceil(math.Log2((width - 1) * (height - 1)))),
			},
		},
	}
	allBytes := float64(int(g.uint32Bytes[1][1]) + g.hor*int(g.uint32Bytes[0][1]) + g.ver*int(g.uint32Bytes[1][0]) + int(g.uint32Bytes[0][0])*g.little)
	g.byteSize = int(math.Ceil(allBytes / 8.0))
	return g
}
func (g *Game) Pieces() int {
	return 1 + g.hor + g.ver + g.little
}
func (g *Game) EmptyBoard() *Board {
	return &Board{game: g}
}
func (g *Game) BoardHashSize() int {
	return g.byteSize
}
func (g *Game) AllShift() uint32 {
	return g.uint32Bytes[1][1] + uint32(g.little)*g.uint32Bytes[0][0] + uint32(g.hor)*2*g.uint32Bytes[0][1] + uint32(g.ver)*g.uint32Bytes[1][0]
}
func (g *Game) Iterate() *Iterator {
	return &Iterator{
		game: g,
	}
}
func (g *Game) loop(f func(hor, ver, n int)) {
	for _, p := range []struct {
		hor int
		ver int
		n   int
	}{
		{hor: 2, ver: 2, n: 1},
		{hor: 2, ver: 1, n: g.hor},
		{hor: 1, ver: 2, n: g.ver},
		{hor: 1, ver: 1, n: g.little},
	} {
		f(p.hor, p.ver, p.n)
	}
}
func (g *Game) Hash(b *Board) BoardHash {
	var i int
	var nb = Board{}
	var ret BoardHash
	g.loop(func(hor, ver, n int) {
		var out optionSet
		var from byte
		for j := 0; j < n; j++ {
			n := nb.search(hor, ver, from, &out)
			for k := uint32(0); k < uint32(n); k++ {
				if out[k] == (b.cell[i] & valMask) {
					(&nb).Put(hor, ver, out[k])
					shift := g.uint32Bytes[ver-1][hor-1]
					(&ret).Push(shift, k)
					// fmt.Println("ret=", ret, "shift=", shift, "k=", k)
					i++
					break
				}
			}
		}
	})
	ret.game = g
	return ret
}
func (g *Game) UnBoardHash(bh BoardHash, b *Board) {
	ns := [2][2]int{{g.little, g.ver}, {g.hor, 1}}
	vs := make([]uint64, 0)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			b := g.uint32Bytes[j][i]
			mask := uint64(1) << b
			mask--
			for n := 0; n < ns[i][j]; n++ {
				v := bh.hash & mask
				vs = append(vs, v)
				bh.hash = bh.hash >> b
			}
		}
	}
	var cursor = len(vs) - 1
	g.loop(func(hor, ver, n int) {
		var out optionSet
		for i := 0; i < n; i++ {
			b.search(hor, ver, 0, &out)
			b.Put(hor, ver, out[vs[cursor]])
			cursor--
		}
	})
}

const (
	UP = 1 << iota
	DOWN
	LEFT
	RIGHT
)

type Movable struct {
	Piece     int
	Direction byte
}

func toxy(p byte) (byte, byte) {
	x := p % width
	y := p / width
	return x, y
}
func (g *Game) Movable(b *Board) []Movable {
	var used [width * height]bool
	var i int = 0

	toidx := func(x, y byte) byte {
		return y*width + x
	}
	g.loop(func(hor, ver, n int) {
		if i >= b.n {
			return
		}
		for j := 0; j < n; j++ {
			x, y := toxy(b.cell[i] & valMask)
			for k := byte(0); k < byte(hor); k++ {
				for l := byte(0); l < byte(ver); l++ {
					np := (y+l)*width + (x + k)
					used[np] = true
				}
			}
			i++
		}
	})
	checkMove := func(horOrver int, x, y byte, dx, dy byte, dir, newDir byte) byte {
		var move = true
		for k := 0; k < horOrver; k++ {
			if used[toidx(x+byte(k)*dx, y+byte(k)*dy)] {
				move = false
				break
			}
		}
		if move {
			dir |= newDir
		}
		return dir
	}
	i = 0
	ret := make([]Movable, 0)
	g.loop(func(hor, ver, n int) {
		for j := 0; j < n; j++ {
			var dir byte
			x, y := toxy(b.cell[i] & valMask)
			if x > 0 {
				dir = checkMove(ver, x-1, y, 0, +1, dir, LEFT)
			}
			if x < byte(width-hor) {
				dir = checkMove(ver, x+byte(hor), y, 0, +1, dir, RIGHT)
			}
			if y > 0 {
				dir = checkMove(hor, x, y-1, +1, 0, dir, UP)
			}
			if y < byte(height-ver) {
				dir = checkMove(hor, x, y+byte(ver), +1, 0, dir, DOWN)
			}
			if dir != 0 {
				ret = append(ret, Movable{Piece: i, Direction: dir})
			}
			i++
		}
	})
	return ret
}
func (g Game) Move(b *Board, mov Movable) []*Board {
	var leftEdge int
	var rightEdge int = 1
	if mov.Direction&UP != 0 || mov.Direction&DOWN != 0 {
		if mov.Piece != 0 {
			leftEdge = rightEdge
			rightEdge = rightEdge + g.hor
			if leftEdge <= mov.Piece && mov.Piece < rightEdge {
				goto EDGECHECKED
			} else {
				leftEdge = rightEdge
				rightEdge = rightEdge + g.ver
				if leftEdge <= mov.Piece && mov.Piece < rightEdge {
					goto EDGECHECKED
				} else {
					leftEdge = rightEdge
					rightEdge = rightEdge + g.little
				}
			}
		}
	}
EDGECHECKED:

	ret := make([]*Board, 0)
	if mov.Direction&UP != 0 {
		nb := b.Copy()
		nb.cell[mov.Piece] -= width
		if leftEdge != 0 && leftEdge <= mov.Piece && mov.Piece <= rightEdge {
			x := nb.cell[leftEdge:rightEdge]
			sort.Slice(nb.cell[leftEdge:rightEdge], func(i, j int) bool { return x[i] < x[j] })
		}
		ret = append(ret, nb)
	}
	if mov.Direction&DOWN != 0 {
		nb := b.Copy()
		nb.cell[mov.Piece] += width
		if leftEdge != 0 && leftEdge <= mov.Piece && mov.Piece <= rightEdge {
			x := nb.cell[leftEdge:rightEdge]
			sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
		}
		ret = append(ret, nb)
	}
	if mov.Direction&LEFT != 0 {
		nb := b.Copy()
		nb.cell[mov.Piece]--
		ret = append(ret, nb)
	}
	if mov.Direction&RIGHT != 0 {
		nb := b.Copy()
		nb.cell[mov.Piece]++
		ret = append(ret, nb)
	}
	return ret
}

func (g Game) StringBoard(b *Board) string {
	ss := make([]string, 0)
	var pos int
	for _, n := range []int{1, g.hor, g.ver, g.little} {
		s := make([]string, 0)
		for i := 0; i < n; i++ {
			s = append(s, fmt.Sprintf("%d", b.cell[pos]&valMask))
			pos++
		}
		ss = append(ss, "("+strings.Join(s, ",")+")")
	}
	return "(" + strings.Join(ss, ",") + ")"
}
