package stupid

import (
	"errors"
	"github.com/Mstch/naruto/helper/util"
	"github.com/gogo/protobuf/proto"
	"net"
)

type ClientImpl struct {
	conn net.Conn
}

func (c ClientImpl) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c ClientImpl) Call(name string, arg proto.Message) (res proto.Message, err error) {
	panic("implement me")
}

func (c ClientImpl) AsyncCall(name string, arg proto.Message) (resC chan proto.Message, err error) {
	panic("implement me")
}

func (c ClientImpl) Notify(name string, arg proto.Message) error {
	body, err := proto.Marshal(arg)
	if err != nil {
		return err
	}
	nameBytes := []byte(name)
	nameLen := len(nameBytes)
	bodyLen := len(body)
	total := 4 + nameLen + bodyLen
	packet := make([]byte, total)
	util.WriteUInt32ToBytes(uint32(nameLen), packet)
	copy(packet[4:nameLen+4], nameBytes)
	util.WriteUInt32ToBytes(uint32(nameLen), packet[4+nameLen:8+nameLen])
	copy(packet[8+nameLen:], body)
	/**
	packet: |{4}bytes{name-len}|{name-len}bytes{name}|{4}bytes{body-len}|{body-len}bytes{body}
	*/
	l, err := c.conn.Write(packet)
	if err != nil {
		return err
	}
	if total != l {
		return errors.New("The length of the written data does not match the calculated\n")
	}
	return nil
}
