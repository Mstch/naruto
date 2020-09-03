//package multicast_todo
//
//import (
//	"github.com/Mstch/naruto/helper/logger"
//	"github.com/Mstch/naruto/helper/domain"
//	"github.com/Mstch/naruto/helper/util"
//	"net"
//	"strconv"
//	"strings"
//)
//
//func getConn(remoteAddress string) *net.UDPConn {
//	local, err := net.ResolveUDPAddr("udp4", ":")
//	if err != nil {
//		panic(err)
//	}
//	remote, err := net.ResolveUDPAddr("udp4", remoteAddress)
//	if err != nil {
//		panic(err)
//	}
//	udpConn, err := net.DialUDP("udp4", local, remote)
//	if err != nil {
//		panic(err)
//	}
//	return udpConn
//}
//
//func hereIam() {
//	udpConn := getConn("255.255.255.255:" + strconv.Itoa(int(domain.self.UPort)))
//	buf := util.UInt32ToBytes(domain.self.Port)
//	for range domain.discoverTicker.C {
//		go func(buf []byte) {
//			_, err := udpConn.Write(buf)
//			if err != nil {
//				logger.Error("发送UDP广播包到%s失败", udpConn.RemoteAddr().String())
//			}
//		}(buf)
//	}
//	udpConn.Close()
//}
//
//func iKnowYouThere(addr string) {
//	udpConn := getConn(addr)
//	buf := util.Int32ToBytes(-1 * int32(domain.self.Port))
//	_, err := udpConn.Write(buf)
//	if err != nil {
//		logger.Error("发送UDP响应包到%s失败", udpConn.RemoteAddr().String())
//	}
//	udpConn.Close()
//}
//
//func handleDiscovery() {
//	pc, err := net.ListenPacket("udp4", ":"+strconv.Itoa(int(domain.self.UPort)))
//	logger.Info("listen udp port :%d", domain.self.UPort)
//	if err != nil {
//		panic(err)
//	}
//	defer pc.Close()
//	for {
//		buf := make([]byte, 8)
//		n, addr, err := pc.ReadFrom(buf)
//		if err != nil {
//			panic(err)
//		}
//		go func(addr net.Addr, buf []byte, n int) {
//			port := util.BytesToUInt32(buf[:n])
//			ip := strings.Split(addr.String(), ":")[0]
//			if _, isSelf := domain.localIp[ip]; !isSelf {
//				logger.Debug("收到来自%s的发现udp包:%d", ip, port)
//				if port > 0 {
//					iKnowYouThere(ip + ":" + strconv.Itoa(int(domain.self.UPort)))
//				} else if port < 0 {
//					m := &domain.Member{
//						Host:      ip,
//						Port:    -port,
//						UPort:   domain.self.UPort,
//						Address: ip + ":" + strconv.Itoa(int(-port)),
//					}
//					if _, ok := domain.discoverCounter[m.Address]; !ok {
//						domain.discoverCounter[m.Address] = 0
//					}
//					domain.discoverCounter[m.Address]++
//					if domain.discoverCounter[m.Address] > 2 {
//						Join(m)
//					}
//				}
//			}
//		}(addr, buf, n)
//	}
//}
