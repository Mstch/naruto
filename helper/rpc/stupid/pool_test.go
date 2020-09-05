package stupid_test

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/helper/sbuf"
	"github.com/Mstch/naruto/helper/util"
	"github.com/Mstch/naruto/raft/msg"
	goio "io"
	"net"
	"sync"
	"testing"
)

var (
	lenPool = sync.Pool{New: func() interface{} {
		return make([]byte, 4)
	}}
	bufPool = sync.Pool{New: func() interface{} {
		return &sbuf.BufferTree{}
	}}
)

func TestWrite(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8888")
	println("connect to server")
	if err != nil {
		panic(err)
	}
	waiter := sync.WaitGroup{}
	waiter.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				m := &msg.HeartbeatReq{
					Term:         2,
					LeaderCommit: 1,
				}
				b := bufPool.Get().(*sbuf.BufferTree)
				//b.Reset()
				bodyBytes := b.Take(m.Size())
				l, err := m.MarshalTo(bodyBytes)
				if err != nil {
					panic(err)
				}
				if l != m.Size() {
					panic("marshal length and size not match")
				}
				s := &stupid.StupidMsg{
					Seq:  0,
					Name: "",
					Body: bodyBytes,
				}

				sbytes := b.Take(s.Size() + 4)
				l, err = s.MarshalTo(sbytes[4:])
				if err != nil {
					panic(err)
				}
				util.WriteInt32ToBytes(int32(s.Size()), sbytes[:4])
				conn.Write(sbytes)
				bufPool.Put(b)
			}
			println("write success")
			waiter.Done()
		}()
	}
	waiter.Wait()
}

func TestRead(t *testing.T) {
	l, _ := net.Listen("tcp", ":8888")
	for {
		c, _ := l.Accept()
		go func(conn net.Conn) {
			for {
				lenBytes := lenPool.Get().([]byte)
				_, err := conn.Read(lenBytes)
				if err != nil {
					if err == goio.EOF {
						logger.Info("%s close connection", conn.RemoteAddr())
					} else {
						logger.Error("conn read failed:%s", err)
					}
					break
				}
				if lenBytes[3] != 6 || lenBytes[0] != 0 {
					panic("lenbytes error")
				}
				slen := int(util.BytesToInt32(lenBytes))
				b := bufPool.Get().(*sbuf.BufferTree)
				//b.Reset()
				sbytes := b.Take(slen)
				_, err = conn.Read(sbytes)
				if err != nil {
					if err == goio.EOF {
						logger.Info("%s close connection", conn.RemoteAddr())
					} else {
						logger.Error("conn read failed:%s", err)
					}
					break
				}
				smsg := &stupid.StupidMsg{}
				err = smsg.Unmarshal(sbytes)
				if err != nil {
					panic(err)
				}
				hmsg := &msg.HeartbeatReq{}
				err = hmsg.Unmarshal(smsg.Body)
				if err != nil {
					panic(err)
				}
				bufPool.Put(b)
				lenPool.Put(lenBytes)
			}
		}(c)
	}
}