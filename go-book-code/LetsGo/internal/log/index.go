package log

import (
	"errors"
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size())
	if err = os.Truncate(
		f.Name(), int64(c.Segment.MaxIndexBytes),
	); err != nil {
		return nil, err
	}
	if idx.mmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}
	return idx, nil
}

func (i *index) Close() error {
	// make sure the memory-mapped file has synced its data to the persisted file
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	// make sure the persisted file has flushed its content to stable storage
	if err := i.file.Sync(); err != nil {
		return err
	}
	// truncate the persisted file to the amount of data that's actually in it then closes the file
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

// takes in an offset then returns the associated record's position of the store. The returning params are
// - out:
func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}
	if in < -1 {
		return 0, 0, errors.New("invalid input: must be > -1")
	}

	// get the offset of index's record
	var recordOffset uint64
	if in == -1 {
		recordOffset = (i.size / entWidth) - 1
	} else {
		recordOffset = uint64(in)
	}

	// the byte offset of record, calc by multiply record'offset with record's size
	byteOffset := recordOffset * entWidth
	// check if the index file has enough space to append new record
	if byteOffset+entWidth > i.size {
		return 0, 0, io.EOF
	}
	out = enc.Uint32(i.mmap[byteOffset : byteOffset+offWidth])
	pos = enc.Uint64(i.mmap[byteOffset+offWidth : byteOffset+entWidth])
	return out, pos, nil
}

func (i *index) Write(off uint32, pos uint64) error {
	if i.size+entWidth > uint64(len(i.mmap)) {
		return io.EOF
	}
	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += entWidth
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
