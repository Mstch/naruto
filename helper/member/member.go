package member

import (
	"github.com/Mstch/naruto/conf"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Member struct {
	Ip     string
	Port   uint32
	Client rpc.Client
	UPort  uint32
	Meta   sync.Map
	//ip:port
	Address string
}

var (
	ConnectedMembers map[string]*Member
	ok               = make(chan bool)
	discoverTicker   = time.NewTicker(time.Second * 1)
	cLock            sync.Mutex
	self             *Member
	localIp          = make(map[string]interface{}, 4)
	discoverCounter  map[string]int
)

func Startup() {
	config := conf.Conf
	self = &Member{
		Port:  config.Port,
		UPort: config.UPort,
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		// 检查是不是寄几
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIp[ipnet.IP.String()] = new(interface{})
			}

		}
	}
	ConnectedMembers = make(map[string]*Member, config.LaunchSize)
	discoverCounter = make(map[string]int, config.LaunchSize)
	go handleDiscovery()
	go hereIam()
	<-ok
}

func Stale() {
	discoverTicker.Stop()
	ok <- true
}
