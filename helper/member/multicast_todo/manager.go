//package multicast_todo
//
//import (
//	"github.com/Mstch/naruto/conf"
//	"github.com/Mstch/naruto/helper/logger"
//	"github.com/Mstch/naruto/helper/member"
//)
//
//func Join(m *member.Member) {
//	member.cLock.Lock()
//	_, loaded := member.ConnectedMembers[m.Address]
//	if !loaded {
//		member.ConnectedMembers[m.Address] = m
//		logger.Info("与%s互相探测成功", m.Address)
//		for _, callback := range joinCallbacks {
//			go callback(m)
//		}
//		if len(member.ConnectedMembers) >= int(conf.Conf.LaunchSize) {
//			go member.Stale()
//			logger.Info("加入集群成功,停止广播")
//		}
//	}
//	member.cLock.Unlock()
//}
//
//func Remove(m *member.Member) {
//
//	logger.Info("节点%s被移出集群", m.Address)
//}
