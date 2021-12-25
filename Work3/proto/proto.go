package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(msg string) ([]byte, error) {
	length := int64(len(msg))
	// 创建一个数据包
	pkg := new(bytes.Buffer)
	// 写入数据包头，表示消息体的长度
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息体
	err = binary.Write(pkg, binary.BigEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
func Decode(reader *bufio.Reader) (string, error) {
	// 读取前8个字节的数据（表示数据包长度的信息）
	lengthByte, _ := reader.Peek(8)
	// 将前8个字节数据读入字节缓冲区
	lengthBuf := bytes.NewBuffer(lengthByte)
	var dataLen int64
	// 读取数据包长度
	err := binary.Read(lengthBuf, binary.BigEndian, &dataLen)
	if err != nil {
		return "", err
	}
	// 判断数据包的总长度是否合法
	if int64(reader.Buffered()) < dataLen + 8{
		return "", err
	}
	pack := make([]byte, 8 + dataLen)
	// 读取整个数据包
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[8:]), nil
}
