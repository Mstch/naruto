package file

import (
	"errors"
	"fmt"
	"github.com/Mstch/naruto/conf"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/member"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Members struct {
	self            *member.Member
	discoveringChan chan *member.Member
	discoveredMap   map[uint32]*member.Member
}

func (m *Members) Discover() {
	waiter := sync.WaitGroup{}
	waiter.Add(int(conf.Conf.LaunchSize - 1))
	go func() {
		for {
			dm := <-m.discoveringChan
			{
				go func(dm *member.Member) {
					var err error
					dm.Conn, err = net.Dial("tcp", dm.Address)
					if err != nil {
						_ = dm.Conn.Close()
						m.discoveringChan <- dm
						return
					}
					m.discoveredMap[dm.Id] = dm
					waiter.Done()
				}(dm)
			}
		}
	}()
	waiter.Wait()
}

func (m *Members) GetMembers() map[uint32]*member.Member {
	return m.discoveredMap
}

func (m *Members) Self() *member.Member {
	return m.self
}

func NewFileMembers() *Members {
	l := len(conf.Conf.Members)
	if uint32(l) < conf.Conf.LaunchSize {
		panic(errors.New(fmt.Sprintf("members in config file not present,num of members [%d] less than launch_size [%d]", l, conf.Conf.LaunchSize)))
	}
	var self *member.Member
	discoveringChan := make(chan *member.Member, l)
	for id, address := range conf.Conf.Members {
		infoSplit := strings.Split(address, ":")
		if len(infoSplit) != 2 {
			logger.Error("配置文件中的节点信息不符合格式[host:port]:%s,忽略此节点配置", address)
			continue
		}
		port, err := strconv.Atoi(infoSplit[1])
		if err != nil {
			logger.Error("配置文件中的节点信息不符合格式[host:port]:%s,忽略此节点配置", address)
			continue
		}
		m := &member.Member{
			Id:      id,
			Host:    infoSplit[0],
			Port:    uint32(port),
			Address: address,
		}
		if id == conf.Conf.Id {
			self = m
		} else {
			discoveringChan <- m
		}
	}
	return &Members{
		self:            self,
		discoveringChan: discoveringChan,
		discoveredMap:   make(map[uint32]*member.Member, l),
	}

}
