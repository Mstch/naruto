package member

import (
	"github.com/Mstch/raft/conf"
	"github.com/Mstch/raft/helper/logger"
)

func Join(m *Member) {
	cLock.Lock()
	_, loaded := ConnectedMembers[m.Address]
	if !loaded {
		ConnectedMembers[m.Address] = m
		logger.Info("与%s互相探测成功", m.Address)
		for _, callback := range joinCallbacks {
			go callback(m)
		}
		if len(ConnectedMembers) >= conf.Conf.LaunchSize {
			go Stale()
			logger.Info("加入集群成功,停止广播")
		}
	}
	cLock.Unlock()
}

func Remove(m *Member) {

	logger.Info("节点%s被移出集群", m.Address)
}
