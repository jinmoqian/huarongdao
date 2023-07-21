package game

import (
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
)

const StoreLevel = 29

type BoardLinkNode struct {
	Uint32Val BoardHash
	Board     *Board
	Next      *BoardLinkNode
	Depth     int
	Nexts     []BoardHash
}
type TreeNode struct {
	Uint32Val BoardHash
	Left      *TreeNode
	Right     *TreeNode
}

func makeTree(keys []BoardHash) *TreeNode {
	l := len(keys)
	if l == 0 {
		return nil
	}
	depth := math.Log2(float64(l))
	depth = math.Floor(depth) + 1
	deepest := l - int(math.Pow(2, depth-1)-1)
	var left int
	if deepest > int(math.Pow(2, depth-1)/2) {
		left = int(math.Pow(2, depth-1)) - 1
	} else {
		right := int(math.Pow(2, depth-2)) - 1
		left = l - 1 - right
	}
	return &TreeNode{
		Uint32Val: keys[left],
		Left:      makeTree(keys[0:left]),
		Right:     makeTree(keys[left+1:]),
	}
}
func treeArray(n *TreeNode) []BoardHash {
	ret := make([]BoardHash, 0)
	if n == nil {
		return ret
	}
	type LinkNode struct {
		node *TreeNode
		next *LinkNode
	}
	head := &LinkNode{
		node: n,
	}
	tail := head
	for {
		if head == nil {
			break
		}
		if head.node.Left != nil {
			tail.next = &LinkNode{node: head.node.Left}
			tail = tail.next
		}
		if head.node.Right != nil {
			tail.next = &LinkNode{node: head.node.Right}
			tail = tail.next
		}
		ret = append(ret, head.node.Uint32Val)
		head = head.next
	}
	return ret
}

type boardHashMapKey struct {
	boardHash BoardHash
}

func (b *boardHashMapKey) Bytes() []byte {
	r := b.boardHash.Bytes()
	return r[:]
}
func (b *boardHashMapKey) Hash() int {
	var h [8]byte
	r := b.boardHash.Bytes()
	copy(h[2:], r[:])
	return int(binary.BigEndian.Uint64(h[:]))
}

func SaveBoard(nodes map[BoardHash]*BoardLinkNode) []byte {
	keys := make([]BoardHash, 0)
	for key := range nodes {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].hash < keys[j].hash
	})
	keys = treeArray(makeTree(keys))
	buf := bytes.NewBuffer([]byte{})
	for _, key := range keys {
		buf.Write(key.Bytes())
		buf.Write([]byte{byte(nodes[key].Depth)})
	}
	return buf.Bytes()
}

type FileHead struct {
	Hor    int
	Ver    int
	Little int
}

func (fh *FileHead) UnBytes(buf []byte) {
	var value uint32 = binary.BigEndian.Uint32(buf)
	fh.Hor = int(value) / 1000
	fh.Ver = (int(value) - fh.Hor*1000) / 100
	fh.Little = int(value) % 100
}

func (fh FileHead) Bytes() []byte {
	var buf [4]byte
	h := fh.Hash()
	binary.BigEndian.PutUint32(buf[:], uint32(h))
	return buf[:]
}
func (fh FileHead) Hash() int {
	return fh.Hor*1000 + fh.Ver*100 + fh.Little
}

type DataFile struct {
	fh  *os.File
	lck *sync.Mutex
	mp  *BytesToBytesMap
}

func (df *DataFile) OpenBytes(bs []byte) error {
	df.lck = &sync.Mutex{}
	df.mp = NewBytesToBytesMap(4, 4)
	return df.mp.Load(strings.NewReader(string(bs)))
}
func (df *DataFile) Open() error {
	fh, err := os.OpenFile("data/data_compressed", os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	df.fh = fh
	runtime.SetFinalizer(df, func(obj *DataFile) {
		df.Close()
	})
	df.lck = &sync.Mutex{}
	df.mp = NewBytesToBytesMap(4, 4)
	err = df.mp.Load(df.fh)
	return err
}
func (df *DataFile) Close() error {
	if df.fh != nil {
		err := df.fh.Close()
		df.fh = nil
		return err
	}
	return nil
}

type BytesDepth struct {
	data []byte
	game *Game
}

func (df *BytesDepth) All() (ret map[BoardHash]byte) {
	ret = make(map[BoardHash]byte)
	var position, index int
	size := df.game.BoardHashSize()
	for {
		position = index * (size + 1)
		if position >= len(df.data) {
			return
		}
		k := df.data[position : position+size]
		depth := df.data[position+size]
		bh := BoardHash{game: df.game}
		bh.FromByte(k)
		ret[bh] = depth
		index++
	}
}
func (df *BytesDepth) Depth(bh BoardHash) byte {
	var position, index int
	size := df.game.BoardHashSize()
	for {
		position = index * (size + 1)
		if position >= len(df.data) {
			return NoDepth
		}
		k := df.data[position : position+size]
		var buf [8]byte
		copy(buf[8-size:], k)
		kv := binary.BigEndian.Uint64(buf[:])
		copy(buf[8-size:], bh.Bytes())
		bv := binary.BigEndian.Uint64(buf[:])
		if kv == bv {
			return df.data[position+size]
		} else if bv < kv {
			index = index*2 + 1
		} else {
			index = index*2 + 2
		}
	}
}
func (df *DataFile) Search(fh FileHead, b *Board) []*Board {
	value, exists := df.mp.Get(fh)
	if !exists {
		return nil
	}
	bs := value.Bytes()
	g := New(fh.Hor, fh.Ver, fh.Little)
	return Search(&BytesDepth{data: bs, game: &g}, &g, b)
}

type Data struct {
	BoardHash BoardHash
	Nexts     []BoardHash
}

func (df *DataFile) All(fh FileHead) map[BoardHash]byte {
	value, exists := df.mp.Get(fh)
	if !exists {
		return nil
	}
	bs := value.Bytes()
	if len(bs) == 0 {
		return nil
	}
	g := New(fh.Hor, fh.Ver, fh.Little)
	return (&BytesDepth{data: bs, game: &g}).All()
}
