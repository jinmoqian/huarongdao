package game

import (
	"bytes"
	"encoding/binary"
	"testing"
)

type dummyKey uint32

func (k dummyKey) Bytes() []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(k))
	return buf[:]
}
func (k dummyKey) Hash() int {
	return int(k) % 10
}

func Test_All(t *testing.T) {
	m := NewBytesToBytesMap(4, 2)
	v, ok := m.Get(dummyKey(100))
	if ok {
		t.Error(v)
	}
	bs := m.Bytes([]MapKey{dummyKey(100)})
	if !bytes.Equal(bs, []byte{0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0}) {
		t.Error(bs)
	}

	dk := dummyKey(10)
	m.Put(dk, BytesValue([]byte{10, 11, 12}))
	if len(m.kvs) != 1 || len(m.kvs[0]) != 1 || m.kvs[0][0].Key != dk || !bytes.Equal(m.kvs[0][0].Value.Bytes(), []byte{10, 11, 12}) {
		t.Error(len(m.kvs))
	}
	v, ok = m.Get(dk)
	if !ok || !bytes.Equal(v.Bytes(), []byte{10, 11, 12}) {
		t.Error(v)
	}
	v, ok = m.Get(dummyKey(100))
	if ok {
		t.Error(v)
	}
	bs = m.Bytes([]MapKey{dummyKey(10)})
	if !bytes.Equal(bs, []byte{0, 0, 0, 10, 0, 0, 0, 12, 0, 0, 0, 3, 10, 11, 12}) {
		t.Error(bs)
	}

	dk = dummyKey(10)
	m.Put(dk, BytesValue([]byte{19, 18, 17}))
	if len(m.kvs) != 1 || len(m.kvs[0]) != 1 || m.kvs[0][0].Key != dk || !bytes.Equal(m.kvs[0][0].Value.Bytes(), []byte{19, 18, 17}) {
		t.Error(len(m.kvs))
	}
	v, ok = m.Get(dk)
	if !ok || !bytes.Equal(v.Bytes(), []byte{19, 18, 17}) {
		t.Error(v)
	}
	bs = m.Bytes([]MapKey{dummyKey(10), dummyKey(99)})
	if !bytes.Equal(bs, []byte{0, 0, 0, 10, 0, 0, 0, 24, 0, 0, 0, 3, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 0, 19, 18, 17}) {
		t.Error(bs)
	}

	dk2 := dummyKey(11)
	m.Put(dk2, BytesValue([]byte{11, 12, 13}))
	if len(m.kvs) != 2 ||
		len(m.kvs[0]) != 1 || m.kvs[0][0].Key != dk || !bytes.Equal(m.kvs[0][0].Value.Bytes(), []byte{19, 18, 17}) ||
		len(m.kvs[1]) != 1 || m.kvs[1][0].Key != dk2 || !bytes.Equal(m.kvs[1][0].Value.Bytes(), []byte{11, 12, 13}) {
		t.Error(len(m.kvs))
	}
	v, ok = m.Get(dk2)
	if !ok || !bytes.Equal(v.Bytes(), []byte{11, 12, 13}) {
		t.Error(v)
	}
	bs = m.Bytes([]MapKey{dummyKey(10), dummyKey(11)})
	if !bytes.Equal(bs, []byte{0, 0, 0, 10, 0, 0, 0, 24, 0, 0, 0, 3, 0, 0, 0, 11, 0, 0, 0, 27, 0, 0, 0, 3, 19, 18, 17, 11, 12, 13}) {
		t.Error(bs)
	}

	dk3 := dummyKey(21)
	m.Put(dk3, BytesValue([]byte{21, 22}))
	if len(m.kvs) != 2 ||
		len(m.kvs[0]) != 1 || m.kvs[0][0].Key != dk || !bytes.Equal(m.kvs[0][0].Value.Bytes(), []byte{19, 18, 17}) ||
		len(m.kvs[1]) != 2 || m.kvs[1][0].Key != dk2 || !bytes.Equal(m.kvs[1][0].Value.Bytes(), []byte{11, 12, 13}) ||
		m.kvs[1][1].Key != dk3 || !bytes.Equal(m.kvs[1][1].Value.Bytes(), []byte{21, 22}) {
		t.Error(len(m.kvs))
	}
	bs = m.Bytes([]MapKey{dummyKey(10), dk3, dummyKey(11)})
	if !bytes.Equal(bs, []byte{
		0, 0, 0, 10, 0, 0, 0, 36, 0, 0, 0, 3,
		0, 0, 0, 21, 0, 0, 0, 39, 0, 0, 0, 2,
		0, 0, 0, 11, 0, 0, 0, 41, 0, 0, 0, 3,
		19, 18, 17, 21, 22, 11, 12, 13}) {
		t.Error(bs)
	}

	dk4 := dummyKey(13)
	m.Put(dk4, BytesValue([]byte{33}))
	if len(m.kvs) != 3 ||
		len(m.kvs[0]) != 1 || m.kvs[0][0].Key != dk || !bytes.Equal(m.kvs[0][0].Value.Bytes(), []byte{19, 18, 17}) ||
		len(m.kvs[1]) != 2 || m.kvs[1][0].Key != dk2 || !bytes.Equal(m.kvs[1][0].Value.Bytes(), []byte{11, 12, 13}) ||
		m.kvs[1][1].Key != dk3 || !bytes.Equal(m.kvs[1][1].Value.Bytes(), []byte{21, 22}) ||
		len(m.kvs[3]) != 1 || m.kvs[3][0].Key != dk4 || !bytes.Equal(m.kvs[3][0].Value.Bytes(), []byte{33}) {
		t.Error(len(m.kvs))
	}
	bs = m.Bytes([]MapKey{dummyKey(10), dk3, dummyKey(11), dk4})
	if !bytes.Equal(bs, []byte{
		0, 0, 0, 10, 0, 0, 0, 48, 0, 0, 0, 3,
		0, 0, 0, 21, 0, 0, 0, 51, 0, 0, 0, 2,
		0, 0, 0, 11, 0, 0, 0, 53, 0, 0, 0, 3,
		0, 0, 0, 13, 0, 0, 0, 56, 0, 0, 0, 1,
		19, 18, 17, 21, 22, 11, 12, 13, 33}) {
		t.Error(bs)
	}

}
