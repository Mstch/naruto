package stupid

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/sbuf"
	"github.com/Mstch/naruto/helper/util"
	"github.com/gogo/protobuf/proto"
	goio "io"
	"net"
	"sync"
)

var (
	byteBufPool = sync.Pool{New: func() interface{} {
		return sbuf.NewBuffer()
	}}
	lenBufPool = sync.Pool{New: func() interface{} {
		return make([]byte, 4)
	}}
)

func serveConn(handlers map[string]*handler, conn net.Conn) {
	for {
		name, resp, err := read(handlers, conn)
		if err != nil {
			if err == goio.EOF {
				logger.Info("%s close connection", conn.RemoteAddr())
				break
			} else {
				logger.Error("conn read failed:%s", err)
			}
			continue
		}
		if resp != nil {
			err := write(conn, name, resp)
			if err != nil {
				logger.Error("error when resp: %s", err)
			}
		}
	}
}

func read(handlers map[string]*handler, conn net.Conn) (string, proto.Message, error) {
	msg := &StupidMsg{}
	lenBuf := lenBufPool.Get().([]byte)
	_, err := conn.Read(lenBuf)
	if err != nil {
		lenBufPool.Put(lenBuf)
		return "", nil, err
	}
	length := int(util.BytesToUInt32(lenBuf))
	lenBufPool.Put(lenBuf)
	buf := byteBufPool.Get().(*sbuf.Buffer)
	buf.Reset()
	msgBuf := buf.Take(length)
	l, err := conn.Read(msgBuf)
	if l != 1036 {
		logger.Error("%d : %s", l,msgBuf)
	}
	if err != nil {
		byteBufPool.Put(buf)
		return "", nil, err
	}
	err = msg.Unmarshal(msgBuf)
	byteBufPool.Put(buf)
	if err != nil {
		return "", nil, err
	}
	name := msg.Name
	if handler, ok := handlers[name]; ok {
		factory := MsgFactoryRegisterInstance().factoryMap[handler.argId]
		var arg proto.Message
		if factory.usePool {
			arg = factory.pool.Get().(proto.Message)
		} else {
			arg = factory.produce()
		}
		err := proto.Unmarshal(msg.Body, arg)
		if factory.usePool {
			factory.pool.Put(arg)
		}
		if err != nil {
			return "", nil, err
		}
		resPb := handlers[name].handleFunc(arg)
		return msg.Name, resPb, nil
	}
	return "", nil, nil
}

func write(conn net.Conn, name string, res proto.Message) error {
	msg := &StupidMsg{
		Name: name,
	}
	var err error
	resMarshaler := res.(marshaler)
	buf := byteBufPool.Get().(*sbuf.Buffer)
	buf.Reset()
	msg.Body = buf.Take(resMarshaler.Size())
	_, err = resMarshaler.MarshalTo(msg.Body)
	if err != nil {
		byteBufPool.Put(buf)
		return err
	}
	msgBuf := buf.Take(msg.Size() + 4)
	util.WriteUInt32ToBytes(uint32(msg.Size()), msgBuf)
	_, err = msg.MarshalTo(msgBuf[4 : 4+msg.Size()])
	if err != nil {
		byteBufPool.Put(buf)
		return err
	}
	_, err = conn.Write(msgBuf)
	byteBufPool.Put(buf)
	return err
}

type marshaler interface {
	Size() int
	MarshalTo(data []byte) (int, error)
}
