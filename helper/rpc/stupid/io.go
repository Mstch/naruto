package stupid

import (
	"encoding/binary"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	goio "io"
	"net"
)

func serveConn(handlers map[string]*handler, conn net.Conn) {
	for {
		name, msgBody, err := read(conn)
		if err != nil {
			if err == goio.EOF {
				logger.Info("%s close connection", conn.RemoteAddr())
			} else {
				logger.Error("conn read failed:%s", err)
			}
			break
		}
		go func(name string, msgBody []byte) {
			handler := handlers[name]
			factory := DefaultRegisterInstance().factoryMap[handler.argId]
			var arg proto.Message
			if factory.pool != nil {
				arg = factory.pool.Get().(proto.Message)
			} else {
				arg = factory.produce()
			}
			err := proto.Unmarshal(msgBody, arg)
			if err != nil {
				logger.Error("Unmarshal failed:%s", err)
				return
			}
			resPb := handlers[name].handleFunc(arg)
			if resPb != nil {
				err = write(conn, name, resPb)
				if err != nil {
					logger.Error("error when response :%s", err)
				}
			}
		}(name, msgBody)
	}
}

func read(conn net.Conn) (string, []byte, error) {
	msg := &StupidMsg{}
	reader := io.NewUint32DelimitedReader(conn, binary.BigEndian, 10*1000*1000)
	err := reader.ReadMsg(msg)
	if err != nil {
		return "", nil, err
	}
	return msg.GetName(), msg.GetBody(), nil
}

func write(conn net.Conn, name string, res proto.Message) error {
	msg := &StupidMsg{}
	msg.Name = name
	var err error
	msg.Body, err = proto.Marshal(res)
	if err != nil {
		return err
	}
	writer := io.NewUint32DelimitedWriter(conn, binary.BigEndian)
	err = writer.WriteMsg(msg)
	return err
}
