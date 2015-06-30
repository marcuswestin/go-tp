package tokens

import (
	"bytes"
	"errors"

	"github.com/cespare/go-smaz"
)

type StringIndex uint32

type StringStore struct {
	offsets  []int
	buffer   bytes.Buffer
	compress bool
}

func (c *StringStore) Append(s string) (index StringIndex, err error) {
	indexOffset := c.buffer.Len()
	bstr := []byte(s)
	if c.compress {
		bstr = smaz.Compress([]byte(s))
	}
	_, err = c.buffer.Write(bstr)
	if err != nil {
		return
	}
	index = StringIndex(len(c.offsets))
	c.offsets = append(c.offsets, indexOffset)
	return
}

func (c *StringStore) Read(stringIndex StringIndex) (str string, err error) {
	index := int(stringIndex)
	if index < 0 || index >= len(c.offsets) {
		return "", errors.New("index out of range")
	}
	indexOffset := c.offsets[int(index)]
	var endOffset int
	if index+1 < len(c.offsets) {
		endOffset = c.offsets[int(index)+1]
	} else {
		endOffset = c.buffer.Len()
	}
	bstr := c.buffer.Bytes()[indexOffset:endOffset]
	if c.compress {
		bstr, err = smaz.Decompress(bstr)
		if err != nil {
			return
		}
	}
	str = string(bstr)
	return
}
