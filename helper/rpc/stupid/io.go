package stupid

import (
	"encoding/binary"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	goio "io"
	"net"
	"sync"
)

var (
	UsePool    = false
	writerPool = sync.Pool{New: func() interface{} {
		return NewUint32DelimitedWriter(nil, binary.BigEndian)
	}}
	readerPool = sync.Pool{New: func() interface{} {
		return NewUint32DelimitedReader(nil, binary.BigEndian, 10*1000*1000)
	}}
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
			if handler, ok := handlers[name]; ok {
				factory := MsgFactoryRegisterInstance().factoryMap[handler.argId]
				var arg proto.Message
				arg = factory.produce()
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
			}
		}(name, msgBody)
	}
}

func read(conn net.Conn) (string, []byte, error) {
	msg := &StupidMsg{}

	var err error
	if UsePool {
		reader := readerPool.Get().(*uint32Reader)
		reader.r = conn
		err = reader.ReadMsg(msg)
		readerPool.Put(reader)
	} else {
		reader := io.NewUint32DelimitedReader(conn, binary.BigEndian, 10*1000*1000)
		err = reader.ReadMsg(msg)
	}
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

	if UsePool {
		writer := writerPool.Get().(*uint32Writer)
		writer.w = conn
		err = writer.WriteMsg(msg)
		writerPool.Put(writer)
	} else {
		writer := io.NewUint32DelimitedWriter(conn, binary.BigEndian)
		err = writer.WriteMsg(msg)
	}
	return err
}
