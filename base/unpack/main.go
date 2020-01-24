package unpack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {
	//类比net.Conn
	bytesBuffer := bytes.NewBuffer([]byte{})

	//粘包
	if err := Encode(bytesBuffer, "hello world 0!!!"); err != nil {
		panic(err)
	}
	if err := Encode(bytesBuffer, "hello world 1!!!"); err != nil {
		panic(err)
	}
	//拆包
	for {
		if bt, err := Decode(bytesBuffer); err == nil {
			fmt.Println(string(bt))
			continue
		}
		break
	}
}

func Encode(bytesBuffer io.Writer, content string) error {
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
