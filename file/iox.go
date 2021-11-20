
package util_file

import (
	io "io"
)

type devNull struct{}

func (devNull)Write(bt []byte)(n int, err error){
	return len(bt), nil
}

func (devNull)Read([]byte)(n int, err error){
	return 0, nil
}

var DevNull io.ReadWriter = devNull{}

type HandleWriter func([]byte)(int, error)

func (writer HandleWriter)Write(bts []byte)(int, error){
	return writer(bts)
}

type WrapWriteHandle func(w *WrapWriter, src []byte)(encoded []byte, err error)
type WrapWriterCloseHandle func(w *WrapWriter)(err error)
type WrapReadHandle func(r *WrapReader, src []byte)(decoded []byte, err error)
type WrapReaderCloseHandle func(r *WrapReader)(err error)

type WrapWriter struct{
	W io.Writer
	wc WrapWriteHandle
	cc WrapWriterCloseHandle
}

func NewWrapWriter(writer io.Writer, wc WrapWriteHandle, _cc ...WrapWriterCloseHandle)(*WrapWriter){
	var cc WrapWriterCloseHandle = nil
	if len(_cc) > 0 {
		cc = _cc[0]
	}
	return &WrapWriter{
		W: writer,
		wc: wc,
		cc: cc,
	}
}

func (w *WrapWriter)Write(src []byte)(n int, err error){
	if w.wc != nil {
		src, err = w.wc(w, src)
		if err != nil { return }
	}
	return w.W.Write(src)
}

func (w *WrapWriter)Close()(err error){
	if w.cc != nil {
		err = w.cc(w)
		if err != nil { return }
	}
	if c, ok := w.W.(io.Closer); ok {
		err = c.Close()
	}
	return
}

type WrapReader struct{
	R io.Reader
	rc WrapWriteHandle
	cc WrapReaderCloseHandle
}

func NewWrapReader(reader io.Reader, rc WrapReadHandle, _cc ...WrapReaderCloseHandle)(*WrapReader){
	var cc WrapReaderCloseHandle = nil
	if len(_cc) > 0 {
		cc = _cc[0]
	}
	return &WrapReader{
		R: reader,
		rc: rc,
		cc: cc,
	}
}

func (r *WrapReader)Read(src []byte)(n int, err error){
	if r.rc != nil {
		src, err = r.wc(r, src)
		if err != nil { return }
	}
	return r.R.Read(src)
}

func (r *WrapReader)Close()(err error){
	if r.cc != nil {
		err = r.cc(r)
		if err != nil { return }
	}
	if c, ok := r.R.(io.Closer); ok {
		err = c.Close()
	}
	return
}


