package file

import (
	"errors"
	"fmt"
	"github.com/Mstch/naruto/conf"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/member/domain"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Members struct {
	self            *domain.Member
	discoveringChan chan *domain.Member
	discoveredMap   map[uint32]*domain.Member
	discoverLock    sync.Mutex
}

func (m *Members) Discover() {
	waiter := sync.WaitGroup{}
	waiter.Add(int(conf.Conf.LaunchSize - 1))
	go func() {
		for {
			dm := <-m.discoveringChan
			{
				go func(dm *domain.Member) {
					var err error
					dm.Conn, err = net.Dial("tcp", dm.Address)
					if err != nil {
						if dm.Conn != nil {
							_ = dm.Conn.Close()
						}
						m.discoveringChan <- dm
						return
					}
					m.discoverLock.Lock()
					m.discoveredMap[dm.Id] = dm
					m.discoverLock.Unlock()
					waiter.Done()
				}(dm)
			}
		}
	}()
	waiter.Wait()
}

func (m *Members) GetMembers() map[uint32]*domain.Member {
	return m.discoveredMap
}

func (m *Members) Self() *domain.Member {
	return m.self
}

func NewFileMembers() *Members {
	l := len(conf.Conf.Members)
	if uint32(l) < conf.Conf.LaunchSize {
		panic(errors.New(fmt.Sprintf("members in config file not present,num of members [%d] less than launch_size [%d]", l, conf.Conf.LaunchSize)))
	}
	var self *domain.Member
	discoveringChan := make(chan *domain.Member, l)
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
		m := &domain.Member{
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
		discoverLock:    sync.Mutex{},
		discoveredMap:   make(map[uint32]*domain.Member, l),
	}

}
