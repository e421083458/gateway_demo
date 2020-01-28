package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "12345678"

func Encode(bytesBuffer io.Writer, content string) error {
	//msg_header+content_len+content
	//8+4+content_len
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err
}
