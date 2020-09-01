//package multicast_todo
//
//import (
//	"github.com/Mstch/naruto/conf"
//	"net"
//	"sync"
//	"time"
//)
//
//var (
//	ConnectedMembers map[string]*Member
//	ok               = make(chan bool)
//	discoverTicker   = time.NewTicker(time.Second * 1)
//	cLock            sync.Mutex
//	self             *Member
//	localIp          = make(map[string]interface{}, 4)
//	discoverCounter  map[string]int
//)
//
//func Startup() {
//	config := conf.Conf
//	self = &Member{
//		Port:  config.Port,
//		UPort: config.UPort,
//	}
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		panic(err)
//	}
//	for _, address := range addrs {
//		// 检查是不是寄几
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				localIp[ipnet.IP.String()] = new(interface{})
//			}
//
//		}
//	}
//	ConnectedMembers = make(map[string]*Member, config.LaunchSize)
//	discoverCounter = make(map[string]int, config.LaunchSize)
//	go multicast_todo.handleDiscovery()
//	go multicast_todo.hereIam()
//	<-ok
//}