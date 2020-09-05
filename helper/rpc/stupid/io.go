package stupid

import (
	"encoding/binary"
	"errors"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/sbuf"
	"github.com/Mstch/naruto/helper/util"
	"github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	goio "io"
	"net"
	"sync"
)

var (
	UseRWPool  = false
	writerPool = sync.Pool{New: func() interface{} {
		return NewUint32DelimitedWriter(nil, binary.BigEndian)
	}}
	readerPool = sync.Pool{New: func() interface{} {
		return NewUint32DelimitedReader(nil, binary.BigEndian, 10*1000*1000)
	}}

	UseBufPool = false
	lenPool    = sync.Pool{New: func() interface{} {
		return make([]byte, 4)
	}}
	BufPool = sync.Pool{New: func() interface{} {
		return sbuf.NewTree()
	}}
)

func serveConn(handlers map[string]*handler, conn net.Conn) {
	sb := BufPool.Get().(*sbuf.BufferTree)
	lb := lenPool.Get().([]byte)
	var lastName string
	var lastBody []byte
	for {
		name, msgBody, err := read(sb, lb, conn)
		if err != nil {
			if err == goio.EOF {
				logger.Info("%s close connection", conn.RemoteAddr())
			} else {
				logger.Error("conn read failed:%s", err)
			}
			break
		}
		lastName = name
		lastBody = msgBody
		if lastName != "Test" {
			panic(lastBody)
		}
		go func(name string, msgBody []byte) {
			if handler, ok := handlers[name]; ok {
				factory := MsgFactoryRegisterInstance().factoryMap[handler.argId]
				arg := factory.produce()
				_, err := arg.(marshaler).MarshalTo(msgBody)
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
	BufPool.Put(sb)
	lenPool.Put(lb)
}

func read(sb *sbuf.BufferTree, lb []byte, conn net.Conn) (string, []byte, error) {
	var err error
	if UseBufPool {
		_, err := goio.ReadFull(conn, lb)
		if err != nil {
			return "", nil, err
		}
		nLen := lb[0]
		lb[0] = 0
		msgLen := util.BytesToInt32(lb)
		if msgLen != 1645 {
			panic(msgLen)
		}
		msgBytes := sb.Take(int(msgLen + int32(nLen)))
		_, err = goio.ReadFull(conn, msgBytes)
		if err != nil {
			return "", nil, err
		}
		return string(msgBytes[:nLen]), msgBytes[nLen:], nil
	} else {
		msg := &StupidMsg{}
		if UseRWPool {
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
}

func write(conn net.Conn, name string, res proto.Message) error {
	var err error
	if UseBufPool {
		r := res.(marshaler)
		rLen := r.Size()
		nLen := len(name)
		if rLen > 0xffffff {
			return errors.New("msg to write too large")
		}
		b := BufPool.Get().(*sbuf.BufferTree)
		msgBytes := b.Take(4 + nLen + rLen)
		util.WriteUInt32ToBytes(uint32(rLen), msgBytes)
		msgBytes[0] = byte(nLen)
		copy(msgBytes[4:nLen+4], name)
		_, err = r.MarshalTo(msgBytes[nLen+4:])
		if err != nil {
			return err
		}
		_, err = conn.Write(msgBytes)
		if err != nil {
			return err
		}
		BufPool.Put(b)
	} else {
		msg := &StupidMsg{}
		msg.Name = name
		msg.Body, err = proto.Marshal(res)
		if err != nil {
			return err
		}
		if UseRWPool {
			writer := writerPool.Get().(*uint32Writer)
			writer.w = conn
			err = writer.WriteMsg(msg)
			writerPool.Put(writer)
		} else {
			writer := io.NewUint32DelimitedWriter(conn, binary.BigEndian)
			err = writer.WriteMsg(msg)
		}
	}
	return err
}
