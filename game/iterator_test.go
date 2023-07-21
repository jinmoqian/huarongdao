package game

import (
	"bytes"
	"testing"
)

func compareInit(t *testing.T, ite *Iterator) bool {
	return (ite.states[0].n == 12 && bytes.Equal(ite.states[0].ops[:], []byte{0, 1, 2, 4, 5, 6, 8, 9, 10, 12, 13, 14, 0, 0, 0, 0})) &&
		(ite.states[1].n == 11 && bytes.Equal(ite.states[1].ops[:], []byte{2, 6, 8, 9, 10, 12, 13, 14, 16, 17, 18, 0, 0, 0, 0, 0})) &&
		(ite.states[2].n == 10 && bytes.Equal(ite.states[2].ops[:], []byte{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[3].n == 8 && bytes.Equal(ite.states[3].ops[:], []byte{7, 8, 9, 11, 12, 13, 14, 15, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[4].n == 6 && bytes.Equal(ite.states[4].ops[:], []byte{8, 9, 12, 13, 14, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[5].n == 4 && bytes.Equal(ite.states[5].ops[:], []byte{9, 13, 14, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[6].n == 6 && bytes.Equal(ite.states[6].ops[:], []byte{14, 15, 16, 17, 18, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[7].n == 5 && bytes.Equal(ite.states[7].ops[:], []byte{15, 16, 17, 18, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[8].n == 4 && bytes.Equal(ite.states[8].ops[:], []byte{16, 17, 18, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) &&
		(ite.states[9].n == 3 && bytes.Equal(ite.states[9].ops[:], []byte{17, 18, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
func Test_iterator_init(t *testing.T) {
	var g = New(1, 4, 4)
	ite := g.Iterate()
	if !ite.init() || ite.all != 10 || !compareInit(t, ite) {
		t.Error(ite.all)
	}
}

func Test_iterator_next(t *testing.T) {
	var g = New(1, 4, 4)
	ite := g.Iterate()
	if !ite.Next() || !compareInit(t, ite) {
		t.Fatal("Next")
	}
	if !ite.Next() || !(ite.states[9].n == 3 && ite.states[9].cur == 1 && bytes.Equal(ite.states[9].ops[:], []byte{17, 18, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) {
		t.Error(ite.states[9].n, ite.states[9].ops[:])
	}
	if !ite.Next() || !(ite.states[9].cur == 2) {
		t.Error(ite.states[9].n, ite.states[9].ops[:])
	}
	if !ite.Next() ||
		!(ite.states[8].n == 4 && ite.states[8].cur == 1 && bytes.Equal(ite.states[8].ops[:4], []byte{16, 17, 18, 19})) ||
		!(ite.states[9].n == 2 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Error(ite.states[8].n, ite.states[8].ops[:])
		t.Fatal(ite.states[9].n, ite.states[9].ops[:])
	}
	if !ite.Next() ||
		!(ite.states[9].n == 2 && ite.states[9].cur == 1 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Fatal(ite.states[9].n, ite.states[9].ops[:])
	}
	if !ite.Next() ||
		!(ite.states[8].n == 4 && ite.states[8].cur == 2 && bytes.Equal(ite.states[8].ops[:4], []byte{16, 17, 18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}

	if !ite.Next() ||
		!(ite.states[7].n == 5 && ite.states[7].cur == 1 && bytes.Equal(ite.states[7].ops[:5], []byte{15, 16, 17, 18, 19})) ||
		!(ite.states[8].n == 3 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:3], []byte{17, 18, 19})) ||
		!(ite.states[9].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[9].n == 2 && ite.states[9].cur == 1 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[8].n == 3 && ite.states[8].cur == 1 && bytes.Equal(ite.states[8].ops[:3], []byte{17, 18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[7].n == 5 && ite.states[7].cur == 2 && bytes.Equal(ite.states[7].ops[:5], []byte{15, 16, 17, 18, 19})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 1 && bytes.Equal(ite.states[6].ops[:6], []byte{14, 15, 16, 17, 18, 19})) ||
		!(ite.states[7].n == 4 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:4], []byte{16, 17, 18, 19})) ||
		!(ite.states[8].n == 3 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:3], []byte{17, 18, 19})) ||
		!(ite.states[9].n == 2 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[9].n == 2 && ite.states[9].cur == 1 && bytes.Equal(ite.states[9].ops[:2], []byte{18, 19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[8].n == 3 && ite.states[8].cur == 1 && bytes.Equal(ite.states[8].ops[:3], []byte{17, 18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[7].n == 4 && ite.states[7].cur == 1 && bytes.Equal(ite.states[7].ops[:4], []byte{16, 17, 18, 19})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	if !ite.Next() ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 2 && bytes.Equal(ite.states[6].ops[:6], []byte{14, 15, 16, 17, 18, 19})) ||
		!(ite.states[7].n == 3 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:3], []byte{17, 18, 19})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}

	if !ite.Next() ||
		!(ite.states[5].n == 4 && ite.states[5].cur == 1 && bytes.Equal(ite.states[5].ops[:4], []byte{9, 13, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 0 && bytes.Equal(ite.states[6].ops[:6], []byte{9, 14, 15, 16, 18, 19})) ||
		!(ite.states[7].n == 5 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:5], []byte{14, 15, 16, 18, 19})) ||
		!(ite.states[8].n == 4 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:4], []byte{15, 16, 18, 19})) ||
		!(ite.states[9].n == 3 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:3], []byte{16, 18, 19})) {
		t.Error(ite.all)
	}
	// C(6,2)=15
	for i := 0; i < 15-2; i++ {
		ite.Next()
	}
	if !ite.Next() ||
		!(ite.states[5].n == 4 && ite.states[5].cur == 1 && bytes.Equal(ite.states[5].ops[:4], []byte{9, 13, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 2 && bytes.Equal(ite.states[6].ops[:6], []byte{9, 14, 15, 16, 18, 19})) ||
		!(ite.states[7].n == 3 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:3], []byte{16, 18, 19})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{18, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	for i := 0; i < 15-1; i++ {
		ite.Next()
	}
	if !ite.Next() ||
		!(ite.states[5].n == 4 && ite.states[5].cur == 2 && bytes.Equal(ite.states[5].ops[:4], []byte{9, 13, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 2 && bytes.Equal(ite.states[6].ops[:6], []byte{9, 13, 15, 16, 17, 19})) ||
		!(ite.states[7].n == 3 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:3], []byte{16, 17, 19})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{17, 19})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{19})) {
		t.Error(ite.all)
	}
	for i := 0; i < 15-1; i++ {
		ite.Next()
	}
	if !ite.Next() ||
		!(ite.states[5].n == 4 && ite.states[5].cur == 3 && bytes.Equal(ite.states[5].ops[:4], []byte{9, 13, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 2 && bytes.Equal(ite.states[6].ops[:6], []byte{9, 13, 14, 16, 17, 18})) ||
		!(ite.states[7].n == 3 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:3], []byte{16, 17, 18})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{17, 18})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{18})) {
		t.Error(ite.all)
	}

	if !ite.Next() ||
		!(ite.states[4].n == 6 && ite.states[4].cur == 1 && bytes.Equal(ite.states[4].ops[:6], []byte{8, 9, 12, 13, 14, 15})) ||
		!(ite.states[5].n == 3 && ite.states[5].cur == 0 && bytes.Equal(ite.states[5].ops[:3], []byte{12, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 0 && bytes.Equal(ite.states[6].ops[:6], []byte{8, 14, 15, 17, 18, 19})) ||
		!(ite.states[7].n == 5 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:5], []byte{14, 15, 17, 18, 19})) ||
		!(ite.states[8].n == 4 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:4], []byte{15, 17, 18, 19})) ||
		!(ite.states[9].n == 3 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:3], []byte{17, 18, 19})) {
		t.Error(ite.all)
	}
	for i := 0; i < 3*15-1; i++ {
		ite.Next()
	}
	if !ite.Next() ||
		!(ite.states[4].n == 6 && ite.states[4].cur == 2 && bytes.Equal(ite.states[4].ops[:6], []byte{8, 9, 12, 13, 14, 15})) ||
		!(ite.states[5].n == 3 && ite.states[5].cur == 0 && bytes.Equal(ite.states[5].ops[:3], []byte{13, 14, 15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 0 && bytes.Equal(ite.states[6].ops[:6], []byte{8, 9, 14, 15, 18, 19})) ||
		!(ite.states[7].n == 5 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:5], []byte{9, 14, 15, 18, 19})) ||
		!(ite.states[8].n == 4 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:4], []byte{14, 15, 18, 19})) ||
		!(ite.states[9].n == 3 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:3], []byte{15, 18, 19})) {
		t.Error(ite.all)
	}
	for i := 0; i < 2*3*15-2; i++ {
		ite.Next()
	}
	if !ite.Next() ||
		!(ite.states[4].n == 6 && ite.states[4].cur == 4 && bytes.Equal(ite.states[4].ops[:6], []byte{8, 9, 12, 13, 14, 15})) ||
		!(ite.states[5].n == 1 && ite.states[5].cur == 0 && bytes.Equal(ite.states[5].ops[:1], []byte{15})) ||
		!(ite.states[6].n == 6 && ite.states[6].cur == 2 && bytes.Equal(ite.states[6].ops[:6], []byte{8, 9, 12, 13, 16, 17})) ||
		!(ite.states[7].n == 3 && ite.states[7].cur == 0 && bytes.Equal(ite.states[7].ops[:3], []byte{13, 16, 17})) ||
		!(ite.states[8].n == 2 && ite.states[8].cur == 0 && bytes.Equal(ite.states[8].ops[:2], []byte{16, 17})) ||
		!(ite.states[9].n == 1 && ite.states[9].cur == 0 && bytes.Equal(ite.states[9].ops[:1], []byte{17})) {
		t.Error(ite.all)
	}
	var counter = 0
	for ite.Next() {
		counter++
	}
	// t.Error(counter)
}
func Test_iterator_sum(t *testing.T) {
	var g = New(0, 0, 2)
	ite := g.Iterate()
	var counter int
	for ite.Next() {
		counter++
	}
	if counter != 12*16*15/2 {
		t.Error(counter)
	}

	counter = 0
	g = New(0, 0, 3)
	ite = g.Iterate()
	for ite.Next() {
		counter++
	}
	if counter != 12*16*15*14/3/2/1 {
		t.Error(counter)
	}

	counter = 0
	g = New(0, 0, 5)
	ite = g.Iterate()
	for ite.Next() {
		counter++
	}
	if counter != 12*16*15*14*13*12/5/4/3/2/1 {
		t.Error(counter)
	}
}
