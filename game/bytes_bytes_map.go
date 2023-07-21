package game

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"unsafe"
)

type MapKey interface {
	Bytes() []byte
	Hash() int
}
type MapValue interface {
	Bytes() []byte
}
type BytesValue []byte

func (vb BytesValue) Bytes() []byte {
	return []byte(vb)
}

type pair struct {
	Key   MapKey
	Value MapValue
}
type BytesToBytesMap struct {
	size    int
	posSize int
	kvs     map[int][]pair
}
type FileBytesValue struct {
	reader         io.ReadSeeker
	size, position uint64
	err            error
}

func (fb *FileBytesValue) Bytes() []byte {
	_, err := fb.reader.Seek(int64(fb.position), io.SeekStart)
	if err != nil {
		fb.err = err
		return nil
	}
	data := make([]byte, fb.size)
	_, err = fb.reader.Read(data)
	if err != nil {
		fb.err = err
		return nil
	}
	return data
}

func NewBytesToBytesMap(size, posSize int) *BytesToBytesMap {
	var p uint32
	return &BytesToBytesMap{size: size, posSize: int(unsafe.Sizeof(p)), kvs: make(map[int][]pair)}
}
func (m *BytesToBytesMap) Load(reader io.ReadSeeker) error {
	var counter int = -1
	var index int
	for {
		key := make([]byte, m.size)
		_, err := reader.Read(key)
		if err != nil {
			return err
		}
		var buf [8]byte
		_, err = reader.Read(buf[8-m.posSize:])
		if err != nil {
			return err
		}
		var position uint64 = binary.BigEndian.Uint64(buf[:])
		_, err = reader.Read(buf[8-m.posSize:])
		if err != nil {
			return err
		}
		var size uint64 = binary.BigEndian.Uint64(buf[:])
		if counter == -1 {
			counter = int(position / (uint64(m.size + 2*m.posSize)))
		}
		fh := FileHead{}
		(&fh).UnBytes(key)
		m.Put(fh, &FileBytesValue{reader: reader, size: size, position: position})
		index++
		if index == counter {
			break
		}
	}
	return nil
}
func (m *BytesToBytesMap) Put(key MapKey, value MapValue) {
	h := key.Hash()
	for i, k := range m.kvs[h] {
		if bytes.Equal(k.Key.Bytes(), key.Bytes()) {
			m.kvs[h][i].Value = value
			return
		}
	}
	m.kvs[h] = append(m.kvs[h], pair{Key: key, Value: value})
}
func (m *BytesToBytesMap) Get(key MapKey) (value MapValue, exists bool) {
	h := key.Hash()
	for _, k := range m.kvs[h] {
		if bytes.Equal(k.Key.Bytes(), key.Bytes()) {
			value = k.Value
			exists = true
			return
		}
	}
	return
}
func (m *BytesToBytesMap) headNodesize() int {
	return m.size + 2*m.posSize
}
func (m *BytesToBytesMap) Bytes(seqKeys []MapKey) []byte {
	headNodeSize := m.headNodesize()
	var position uint64 = uint64(headNodeSize * len(seqKeys))
	ret := make([]byte, position)
	for idx, key := range seqKeys {
		kbs := key.Bytes()
		if len(kbs) != m.size {
			panic(fmt.Sprintf("key.Bytes()[%+v] of key[%+v] not match size[%d", kbs, key, m.size))
		}
		copy(ret[idx*headNodeSize:], kbs)
		value, exists := m.Get(key)
		var pos uint64
		if !exists {
			value = BytesValue([]byte{})
		} else {
			pos = position
		}
		ret = append(ret, value.Bytes()...)
		from := ret[idx*headNodeSize+m.size:]
		switch m.posSize {
		case 1:
			from[0] = uint8(pos)
			from[1] = uint8(len(value.Bytes()))
		case 2:
			binary.BigEndian.PutUint16(from, uint16(pos))
			binary.BigEndian.PutUint16(from[m.posSize:], uint16(len(value.Bytes())))
		case 4:
			binary.BigEndian.PutUint32(from, uint32(pos))
			binary.BigEndian.PutUint32(from[m.posSize:], uint32(len(value.Bytes())))
		case 8:
			binary.BigEndian.PutUint64(from, pos)
			binary.BigEndian.PutUint64(from[m.posSize:], uint64(len(value.Bytes())))
		default:
		}
		position += uint64(len(value.Bytes()))
	}
	return ret
}

type ByteSliceMapKey []byte

func (u ByteSliceMapKey) Bytes() []byte {
	return []byte(u)
}
func (u ByteSliceMapKey) Hash() int {
	var buf [8]byte
	l := len(u)
	copy(buf[8-l:], u)
	return int(binary.BigEndian.Uint64(buf[:]))
}
func (m *BytesToBytesMap) FromBytes(bs []byte) {
	var buf [8]byte
	copy(buf[8-m.posSize:], bs[m.size:m.size+m.posSize])
	end := binary.BigEndian.Uint64(buf[:])
	for pos := uint64(0); pos < end; {
		var key, position, size [8]byte
		copy(key[8-m.size:], bs[pos:pos+uint64(m.size)])
		copy(position[8-m.posSize:], bs[pos+uint64(m.size):pos+uint64(m.size+m.posSize)])
		copy(size[8-m.posSize:], bs[pos+uint64(m.size)+uint64(m.posSize):pos+uint64(m.size+m.posSize+m.posSize)])
		p := binary.BigEndian.Uint64(position[:])
		e := p + binary.BigEndian.Uint64(size[:])
		m.Put(ByteSliceMapKey(key[8-m.size:]), BytesValue(bs[p:e]))
		pos += uint64(m.size + m.posSize*2)
	}
}
