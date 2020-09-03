//package multicast_todo
//
//import (
//	"github.com/Mstch/naruto/conf"
//	"github.com/Mstch/naruto/helper/logger"
//	"github.com/Mstch/naruto/helper/domain"
//)
//
//func Join(m *domain.Member) {
//	domain.cLock.Lock()
//	_, loaded := domain.ConnectedMembers[m.Address]
//	if !loaded {
//		domain.ConnectedMembers[m.Address] = m
//		logger.Info("与%s互相探测成功", m.Address)
//		for _, callback := range joinCallbacks {
//			go callback(m)
//		}
//		if len(domain.ConnectedMembers) >= int(conf.Conf.LaunchSize) {
//			go domain.Stale()
//			logger.Info("加入集群成功,停止广播")
//		}
//	}
//	domain.cLock.Unlock()
//}
//
//func Remove(m *domain.Member) {
//
//	logger.Info("节点%s被移出集群", m.Address)
//}
