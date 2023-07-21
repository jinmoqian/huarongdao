package game

import (
	"bytes"
	"testing"
)

func Test_caocao(t *testing.T) {
	var b = Board{}
	var o = optionSet{}
	var n = b.search(2, 2, 0, &o)
	if n != 12 || !bytes.Equal(o[:], []byte{0, 1, 2, 4, 5, 6, 8, 9, 10, 12, 13, 14, 0, 0, 0, 0}) {
		t.Error(n, o)
	}
}
func Test_search_ver(t *testing.T) {
	var b = Board{}
	b.Put(2, 2, 9)
	var o = optionSet{}
	n := b.search(1, 2, 0, &o)
	if n != 10 || !bytes.Equal(o[:], []byte{0, 1, 2, 3, 4, 7, 8, 11, 12, 15, 0, 0, 0, 0, 0, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(2, 1, 9)
	o = optionSet{}
	n = b.search(1, 2, 0, &o)
	if n != 12 || !bytes.Equal(o[:], []byte{0, 1, 2, 3, 4, 7, 8, 11, 12, 13, 14, 15, 0, 0, 0, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(1, 2, 10)
	o = optionSet{}
	n = b.search(1, 2, 0, &o)
	if n != 13 || !bytes.Equal(o[:], []byte{0, 1, 2, 3, 4, 5, 7, 8, 9, 11, 12, 13, 15, 0, 0, 0}) {
		t.Error(n, o)
	}

	b.pop()
	b.Put(1, 1, 9)
	o = optionSet{}
	n = b.search(1, 2, 0, &o)
	if n != 14 || !bytes.Equal(o[:], []byte{0, 1, 2, 3, 4, 6, 7, 8, 10, 11, 12, 13, 14, 15, 0, 0}) {
		t.Error(n, o)
	}
}

func Test_search_hor(t *testing.T) {
	var b = Board{}
	var o = optionSet{}
	var n = b.search(2, 1, 0, &o)
	if n != 15 || !bytes.Equal(o[:], []byte{0, 1, 2, 4, 5, 6, 8, 9, 10, 12, 13, 14, 16, 17, 18, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(2, 2, 5)
	o = optionSet{}
	n = b.search(2, 1, 0, &o)
	if n != 9 || !bytes.Equal(o[:], []byte{0, 1, 2, 12, 13, 14, 16, 17, 18, 0, 0, 0, 0, 0, 0, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(2, 1, 9)
	o = optionSet{}
	n = b.search(2, 1, 0, &o)
	if n != 12 || !bytes.Equal(o[:], []byte{0, 1, 2, 4, 5, 6, 12, 13, 14, 16, 17, 18, 0, 0, 0, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(1, 2, 10)
	o = optionSet{}
	n = b.search(2, 1, 0, &o)
	if n != 11 || !bytes.Equal(o[:], []byte{0, 1, 2, 4, 5, 6, 8, 12, 16, 17, 18, 0, 0, 0, 0, 0}) {
		t.Error(n, o)
	}

	b = Board{}
	b.Put(1, 1, 10)
	o = optionSet{}
	n = b.search(2, 1, 0, &o)
	if n != 13 || !bytes.Equal(o[:], []byte{0, 1, 2, 4, 5, 6, 8, 12, 13, 14, 16, 17, 18, 0, 0, 0}) {
		t.Error(n, o)
	}
}
func Test_search_little(t *testing.T) {
	var b = Board{}
	b.Put(2, 2, 5)
	var o = optionSet{}
	var n = b.search(1, 1, 0, &o)
	if n != 16 || !bytes.Equal(o[:], []byte{0, 1, 2, 3, 4, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Error(n, o)
	}
}
func Test_Hash(t *testing.T) {
	g := New(3, 2, 4) // [2][2]uint32 [[2,4],[3,4]]
	b := &Board{}
	b.Put(2, 2, 0)  // 第0个
	b.Put(2, 1, 2)  // 第0个
	b.Put(2, 1, 6)  // 第0个
	b.Put(2, 1, 9)  // 第1个 0001
	b.Put(1, 2, 12) // 第2个 010
	b.Put(1, 2, 13) // 第1个 001
	b.Put(1, 1, 14) // 第2个 10
	b.Put(1, 1, 15) // 第2个 10
	b.Put(1, 1, 18) // 第2个 10
	b.Put(1, 1, 19) // 第2个 10
	v := g.Hash(b)
	if v.hash != 0x51AA {
		t.Errorf("%X", v.hash)
	}

	g = New(1, 1, 1) //[2][2]uint32 [[4,4],[4,4]]
	b = &Board{}
	b.Put(2, 2, 0)  // 0 0000
	b.Put(2, 1, 9)  // 3 0011
	b.Put(1, 2, 12) // 5 0101
	b.Put(1, 1, 6)  // 2 0010
	v = g.Hash(b)
	if v.hash != 0x352 {
		t.Errorf("%X", v.hash)
	}

	g = New(3, 2, 5) // [2][2]uint32 [[1,4],[3,4]]
	b = &Board{}
	b.Put(2, 2, 0)  // 第0个 0000
	b.Put(2, 1, 2)  // 第0个 0000
	b.Put(2, 1, 6)  // 第0个 0000
	b.Put(2, 1, 9)  // 第1个 0001
	b.Put(1, 2, 12) // 第2个 010
	b.Put(1, 2, 13) // 第1个 001
	b.Put(1, 1, 11) // 第1个 1
	b.Put(1, 1, 14) // 第1个 1
	b.Put(1, 1, 15) // 第1个 1
	b.Put(1, 1, 18) // 第1个 1
	b.Put(1, 1, 19) // 第1个 1
	v = g.Hash(b)
	if v.hash != 0xA3F {
		t.Errorf("%X", v.hash)
	}
}
func Test_UnBoardHash(t *testing.T) {
	g := New(2, 1, 0)
	b := g.EmptyBoard()
	b.Put(2, 2, 0)
	b.Put(2, 1, 2)
	b.Put(2, 1, 6)
	b.Put(1, 2, 8)
	bh := g.Hash(b)
	eb := g.EmptyBoard()
	g.UnBoardHash(bh, eb)
	if b.n != eb.n || !bytes.Equal(b.cell[:], eb.cell[:]) {
		t.Error(b, eb)
	}

	g = New(1, 4, 0)
	b = g.EmptyBoard()
	b.Put(2, 2, 1)
	b.Put(2, 1, 18)
	b.Put(1, 2, 3)
	b.Put(1, 2, 4)
	b.Put(1, 2, 10)
	b.Put(1, 2, 12)
	bh = g.Hash(b)
	eb = g.EmptyBoard()
	g.UnBoardHash(bh, eb)
	if b.n != eb.n || !bytes.Equal(b.cell[:], eb.cell[:]) {
		t.Error(b, eb)
	}

	g = New(0, 0, 0)
	b = g.EmptyBoard()
	b.Put(2, 2, 10)
	bh = g.Hash(b)
	eb = g.EmptyBoard()
	g.UnBoardHash(bh, eb)
	if b.n != eb.n || !bytes.Equal(b.cell[:], eb.cell[:]) {
		t.Error(b, eb)
	}

	g = New(5, 0, 0)
	b = g.EmptyBoard()
	b.Put(2, 2, 1)
	b.Put(2, 1, 8)
	b.Put(2, 1, 10)
	b.Put(2, 1, 12)
	b.Put(2, 1, 14)
	b.Put(2, 1, 16)
	bh = g.Hash(b)
	eb = g.EmptyBoard()
	g.UnBoardHash(bh, eb)
	if b.n != eb.n || !bytes.Equal(b.cell[:], eb.cell[:]) {
		t.Error(b, eb)
	}
}
func Test_win(t *testing.T) {
	b := &Board{}
	b.Put(2, 2, 14)
	if b.Win() {
		t.Error("false")
	}

	b = &Board{}
	b.Put(2, 2, 13)
	if !b.Win() {
		t.Error("true")
	}
}
func Test_Movable_1(t *testing.T) {
	g := New(0, 0, 0)
	b := &Board{}
	b.Put(2, 2, 0)
	r1 := g.Movable(b)
	if len(r1) != 1 {
		t.Error(len(r1))
	}
	if r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Piece != 0 {
		t.Error(r1[0])
	}

	g = New(0, 0, 0)
	b = &Board{}
	b.Put(2, 2, 2)
	r1 = g.Movable(b)
	if len(r1) != 1 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 {
		t.Error(r1[0])
	}

	g = New(0, 0, 0)
	b = &Board{}
	b.Put(2, 2, 12)
	r1 = g.Movable(b)
	if len(r1) != 1 {
		t.Error(len(r1))
	}
	if r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP == 0 || r1[0].Piece != 0 {
		t.Error(r1[0])
	}

	g = New(0, 0, 0)
	b = &Board{}
	b.Put(2, 2, 14)
	r1 = g.Movable(b)
	if len(r1) != 1 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&UP == 0 || r1[0].Piece != 0 {
		t.Error(r1[0])
	}

	g = New(0, 0, 0)
	b = &Board{}
	b.Put(2, 2, 5)
	r1 = g.Movable(b)
	if len(r1) != 1 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 {
		t.Error(r1[0])
	}

	g = New(0, 1, 0)
	b = &Board{}
	b.Put(2, 2, 5)
	b.Put(1, 2, 11)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0])
	}

	g = New(0, 1, 0)
	b = &Board{}
	b.Put(2, 2, 5)
	b.Put(1, 2, 0)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(1, 0, 0)
	b = &Board{}
	b.Put(2, 2, 10)
	b.Put(2, 1, 12)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(1, 0, 0)
	b = &Board{}
	b.Put(2, 2, 8)
	b.Put(2, 1, 10)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(1, 0, 0)
	b = &Board{}
	b.Put(2, 2, 9)
	b.Put(2, 1, 4)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(1, 0, 0)
	b = &Board{}
	b.Put(2, 2, 9)
	b.Put(2, 1, 4)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 1, 0)
	b = &Board{}
	b.Put(2, 2, 13)
	b.Put(1, 2, 6)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 1, 0)
	b = &Board{}
	b.Put(2, 2, 5)
	b.Put(1, 2, 13)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(1, 0, 0)
	b = &Board{}
	b.Put(2, 2, 5)
	b.Put(2, 1, 14)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 0, 1)
	b = &Board{}
	b.Put(2, 2, 9)
	b.Put(1, 1, 5)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 0, 1)
	b = &Board{}
	b.Put(2, 2, 10)
	b.Put(1, 1, 9)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 0, 1)
	b = &Board{}
	b.Put(2, 2, 10)
	b.Put(1, 1, 13)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP == 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}

	g = New(0, 0, 1)
	b = &Board{}
	b.Put(2, 2, 9)
	b.Put(1, 1, 6)
	r1 = g.Movable(b)
	if len(r1) != 2 {
		t.Error(len(r1))
	}
	if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
		r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 {
		t.Error(r1[0], r1[1])
	}
}
func Test_Movable_2(t *testing.T) {
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 2, 4)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 2, 8)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 2, 1)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 2, 2)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 2, 7)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 2, 11)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 2, 13)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 1, 0)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 2, 14)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 1, 13)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 1, 14)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 2)
		b.Put(2, 1, 9)
		b.Put(1, 1, 8)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT == 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 1, 5)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 1, 6)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 12)
		b.Put(2, 1, 9)
		b.Put(1, 1, 11)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(2, 0, 0)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 17)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
	{
		g := New(2, 0, 0)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 14)
		b.Put(1, 1, 17)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1])
		}
	}
}

func Test_Movable_3(t *testing.T) {
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 6)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 9)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 13)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 18)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 11)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 1, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 10)
		b.Put(1, 1, 15)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT != 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 2, 0)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 6)
		b.Put(1, 2, 14)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 2, 0)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 13)
		b.Put(1, 2, 14)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 2, 0)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 2, 9)
		b.Put(1, 2, 14)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN != 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
}
func Test_Movable_4(t *testing.T) {
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 6)
		b.Put(1, 1, 10)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT != 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN == 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT != 0 || r1[1].Direction&RIGHT == 0 || r1[1].Direction&UP == 0 || r1[1].Direction&DOWN != 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT == 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP != 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 9)
		b.Put(1, 1, 10)
		r1 := g.Movable(b)
		if len(r1) != 3 {
			t.Error(len(r1))
		}
		if r1[0].Direction&LEFT != 0 || r1[0].Direction&RIGHT == 0 || r1[0].Direction&UP != 0 || r1[0].Direction&DOWN != 0 || r1[0].Piece != 0 ||
			r1[1].Direction&LEFT == 0 || r1[1].Direction&RIGHT != 0 || r1[1].Direction&UP != 0 || r1[1].Direction&DOWN == 0 || r1[1].Piece != 1 ||
			r1[2].Direction&LEFT != 0 || r1[2].Direction&RIGHT == 0 || r1[2].Direction&UP == 0 || r1[2].Direction&DOWN == 0 || r1[2].Piece != 2 {
			t.Error(r1[0], r1[1], r1[2])
		}
	}
}
func Test_Move(t *testing.T) {
	{
		g := New(0, 0, 0)
		b := &Board{}
		b.Put(2, 2, 5)
		moves := g.Move(b, Movable{Piece: 0, Direction: UP | DOWN | LEFT | RIGHT})
		if len(moves) != 4 ||
			moves[0].cell[0]&valMask != 1 || moves[1].cell[0]&valMask != 9 || moves[2].cell[0]&valMask != 4 || moves[3].cell[0]&valMask != 6 {
			t.Error(moves[0].cell, moves[1].cell, moves[2].cell, moves[3].cell)
		}
	}
	{
		g := New(0, 0, 0)
		b := &Board{}
		b.Put(2, 2, 5)
		moves := g.Move(b, Movable{Piece: 0, Direction: 0})
		if len(moves) != 0 {
			t.Error(moves)
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 8)
		b.Put(1, 1, 9)
		moves := g.Move(b, Movable{Piece: 2, Direction: RIGHT})
		if len(moves) != 1 || moves[0].cell[2]&valMask != 10 {
			t.Error(moves)
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 8)
		b.Put(1, 1, 9)
		moves := g.Move(b, Movable{Piece: 2, Direction: DOWN})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 8 || moves[0].cell[2]&valMask != 13 {
			t.Error(moves)
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 8)
		b.Put(1, 1, 9)
		moves := g.Move(b, Movable{Piece: 2, Direction: UP})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 5 || moves[0].cell[2]&valMask != 8 {
			t.Error(moves)
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 8)
		b.Put(1, 1, 9)
		moves := g.Move(b, Movable{Piece: 1, Direction: UP})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 4 || moves[0].cell[2]&valMask != 9 {
			t.Error(moves)
		}
	}
	{
		g := New(0, 0, 2)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(1, 1, 8)
		b.Put(1, 1, 9)
		moves := g.Move(b, Movable{Piece: 1, Direction: DOWN})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 9 || moves[0].cell[2]&valMask != 12 {
			t.Error(moves)
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 12)
		moves := g.Move(b, Movable{Piece: 1, Direction: UP})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 9 || moves[0].cell[2]&valMask != 12 {
			t.Error(moves)
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 12)
		moves := g.Move(b, Movable{Piece: 1, Direction: DOWN})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 17 || moves[0].cell[2]&valMask != 12 {
			t.Error(moves)
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 12)
		moves := g.Move(b, Movable{Piece: 1, Direction: RIGHT})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 14 || moves[0].cell[2]&valMask != 12 {
			t.Error(moves)
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 12)
		moves := g.Move(b, Movable{Piece: 2, Direction: UP})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 13 || moves[0].cell[2]&valMask != 8 {
			t.Error(moves)
		}
	}
	{
		g := New(1, 0, 1)
		b := &Board{}
		b.Put(2, 2, 0)
		b.Put(2, 1, 13)
		b.Put(1, 1, 12)
		moves := g.Move(b, Movable{Piece: 2, Direction: DOWN})
		if len(moves) != 1 || moves[0].cell[1]&valMask != 13 || moves[0].cell[2]&valMask != 16 {
			t.Error(moves)
		}
	}
}
func Test_Boardhash_Bytes(t *testing.T) {
	g := New(0, 0, 0)
	b := &Board{game: &g}
	b.Put(2, 2, 2)
	bs := g.Hash(b).Bytes()
	if g.byteSize != 1 || !bytes.Equal(bs, []byte{2}) {
		t.Error(bs)
	}
	{
		b2 := &BoardHash{game: &g}
		b2.FromByte(bs)
		if b2.hash != 2 {
			t.Error(b2.hash)
		}
	}

	g = New(2, 2, 3)
	b = &Board{game: &g}
	b.Put(2, 2, 2)  // 第2个 0010
	b.Put(2, 1, 4)  // 第1个 0001
	b.Put(2, 1, 8)  // 第1个 0001
	b.Put(1, 2, 10) // 第0个 0000
	b.Put(1, 2, 11) // 第0个 0000
	b.Put(1, 1, 12) // 第2个 010
	b.Put(1, 1, 13) // 第2个 010
	b.Put(1, 1, 16) // 第2个 010
	bs = g.Hash(b).Bytes()
	if g.byteSize != 4 || !bytes.Equal(bs, []byte{0x04, 0x22, 0x00, 0x92}) {
		t.Error(bs)
	}
	{
		b2 := &BoardHash{game: &g}
		b2.FromByte(bs)
		if b2.hash != 0x04220092 {
			t.Error(b2.hash)
		}
	}
}
func Test_can(t *testing.T) {
	g := New(0, 0, 0)
	b := &Board{game: &g}
	b.Put(2, 2, 5)
	for i := 0; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 0 || i == 1 || i == 2 || i == 12 || i == 13 || i == 14 || i == 16 || i == 17 || i == 18)) ||
			(!c && ((3 <= i && i <= 11) || i == 15 || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 0 || i == 3 || i == 4 || i == 7 || i == 8 || i == 11 || (12 <= i && i <= 15))) ||
			(!c && (i == 1 || i == 2 || i == 5 || i == 6 || i == 9 || i == 10 || i >= 16))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && ((0 <= i && i <= 4) || i == 7 || i == 8 || (11 <= i))) ||
			(!c && (i == 5 || i == 6 || i == 9 || i == 10))) {
			t.Error(c, i)
		}
	}

	g = New(1, 0, 0)
	b = &Board{game: &g}
	b.Put(2, 2, 0)
	b.Put(2, 1, 13)
	for i := 8; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 8 || i == 9 || i == 10 || i == 14 || i == 16 || i == 17 || i == 18)) ||
			(!c && ((11 <= i && i <= 15) || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 8 || i == 11 || i == 12 || i == 15)) ||
			(!c && (i == 9 || i == 10 || i == 13 || i == 14 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && ((8 <= i && i <= 12) || (15 <= i && i <= 19))) ||
			(!c && (i == 13 || i == 14))) {
			t.Error(c, i)
		}
	}
	g.hor = 2
	b.Put(2, 1, 16)
	for i := 8; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 8 || i == 9 || i == 10 || i == 18)) ||
			(!c && ((11 <= i && i <= 17) || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 8 || i == 11 || i == 15)) ||
			(!c && (i == 9 || i == 10 || i == 12 || i == 13 || i == 14 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && ((8 <= i && i <= 12) || i == 15 || i == 18 || i == 19)) ||
			(!c && (i == 13 || i == 14 || i == 16 || i == 17))) {
			t.Error(c, i)
		}
	}

	g = New(0, 1, 0)
	b = &Board{game: &g}
	b.Put(2, 2, 0)
	b.Put(1, 2, 10)
	for i := 0; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 2 || i == 6 || i == 8 || i == 12 || (16 <= i && i <= 18))) ||
			(!c && (i == 0 || i == 1 || (3 <= i && i <= 5) || i == 7 || (9 <= i && i <= 11) || (13 <= i && i <= 15) || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 2 || i == 3 || i == 7 || i == 8 || i == 9 || i == 11 || i == 12 || i == 13 || i == 15)) ||
			(!c && (i == 0 || i == 1 || i == 4 || i == 5 || i == 6 || i == 10 || i == 14 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && (i == 2 || i == 3 || (6 <= i && i <= 9) || (11 <= i && i <= 13) || (15 <= i))) ||
			(!c && (i == 0 || i == 1 || i == 4 || i == 5 || i == 10 || i == 14))) {
			t.Error(c, i)
		}
	}
	g.ver = 2
	b.Put(1, 2, 12)
	for i := 0; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 2 || i == 6 || i == 8 || i == 17 || i == 18)) ||
			(!c && (i == 0 || i == 1 || (3 <= i && i <= 5) || i == 7 || (9 <= i && i <= 16) || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 2 || i == 3 || i == 7 || i == 9 || i == 11 || i == 13 || i == 15)) ||
			(!c && (i == 0 || i == 1 || i == 4 || i == 5 || i == 6 || i == 8 || i == 10 || i == 12 || i == 14 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && (i == 2 || i == 3 || (6 <= i && i <= 9) || 11 == i || i == 13 || 15 == i || 17 <= i)) ||
			(!c && (i == 0 || i == 1 || i == 4 || i == 5 || i == 10 || i == 12 || i == 14 || i == 16))) {
			t.Error(c, i)
		}
	}

	g = New(0, 0, 1)
	b = &Board{game: &g}
	b.Put(2, 2, 0)
	b.Put(1, 1, 10)
	for i := 0; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 2 || i == 6 || i == 8 || (12 <= i && i <= 14) || (16 <= i && i <= 18))) ||
			(!c && (i == 0 || i == 1 || (3 <= i && i <= 5) || i == 7 || (9 <= i && i <= 11) || i == 15 || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 2 || i == 3 || (7 <= i && i <= 9) || (11 <= i && i <= 15))) ||
			(!c && (i == 0 || i == 1 || (4 <= i && i <= 6) || i == 10 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && (i == 2 || i == 3 || (6 <= i && i <= 9) || 11 <= i)) ||
			(!c && i == 0 || i == 1 || i == 4 || i == 5 || i == 10)) {
			t.Error(c, i)
		}
	}
	g.little = 2
	b.Put(1, 1, 13)
	for i := 0; i < width*height; i++ {
		c := b.Can(2, 1, byte(i))
		if !((c && (i == 2 || i == 6 || i == 8 || i == 14 || (16 <= i && i <= 18))) ||
			(!c && (i == 0 || i == 1 || (3 <= i && i <= 5) || i == 7 || (9 <= i && i <= 13) || i == 15 || i == 19))) {
			t.Error(c, i)
		}

		c = b.Can(1, 2, byte(i))
		if !((c && (i == 2 || i == 3 || i == 7 || i == 8 || i == 11 || i == 12 || i == 14 || i == 15)) ||
			(!c && (i == 0 || i == 1 || (4 <= i && i <= 6) || i == 9 || i == 10 || i == 13 || 16 <= i))) {
			t.Error(c, i)
		}

		c = b.Can(1, 1, byte(i))
		if !((c && (i == 2 || i == 3 || (6 <= i && i <= 9) || i == 11 || i == 12 || 14 <= i)) ||
			(!c && i == 0 || i == 1 || i == 4 || i == 5 || i == 10 || i == 13)) {
			t.Error(c, i)
		}
	}
}
