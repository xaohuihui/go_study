package geeCache

// author: songyanhui
// datetime: 2022/1/18 16:30:11
// software: GoLand

/*  抽象出一个只读数据结构 ByteView 用来表示缓存值 */

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// Len returns the view`s length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
