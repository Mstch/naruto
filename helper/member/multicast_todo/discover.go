//package multicast_todo
//
//import (
//	"github.com/Mstch/naruto/helper/logger"
//	"github.com/Mstch/naruto/helper/member"
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
//	udpConn := getConn("255.255.255.255:" + strconv.Itoa(int(member.self.UPort)))
//	sbuf := util.UInt32ToBytes(member.self.Port)
//	for range member.discoverTicker.C {
//		go func(sbuf []byte) {
//			_, err := udpConn.Write(sbuf)
//			if err != nil {
//				logger.Error("发送UDP广播包到%s失败", udpConn.RemoteAddr().String())
//			}
//		}(sbuf)
//	}
//	udpConn.Close()
//}
//
//func iKnowYouThere(addr string) {
//	udpConn := getConn(addr)
//	sbuf := util.Int32ToBytes(-1 * int32(member.self.Port))
//	_, err := udpConn.Write(sbuf)
//	if err != nil {
//		logger.Error("发送UDP响应包到%s失败", udpConn.RemoteAddr().String())
//	}
//	udpConn.Close()
//}
//
//func handleDiscovery() {
//	pc, err := net.ListenPacket("udp4", ":"+strconv.Itoa(int(member.self.UPort)))
//	logger.Info("listen udp port :%d", member.self.UPort)
//	if err != nil {
//		panic(err)
//	}
//	defer pc.Close()
//	for {
//		sbuf := make([]byte, 8)
//		n, addr, err := pc.ReadFrom(sbuf)
//		if err != nil {
//			panic(err)
//		}
//		go func(addr net.Addr, sbuf []byte, n int) {
//			port := util.BytesToUInt32(sbuf[:n])
//			ip := strings.Split(addr.String(), ":")[0]
//			if _, isSelf := member.localIp[ip]; !isSelf {
//				logger.Debug("收到来自%s的发现udp包:%d", ip, port)
//				if port > 0 {
//					iKnowYouThere(ip + ":" + strconv.Itoa(int(member.self.UPort)))
//				} else if port < 0 {
//					m := &member.Member{
//						Host:      ip,
//						Port:    -port,
//						UPort:   member.self.UPort,
//						Address: ip + ":" + strconv.Itoa(int(-port)),
//					}
//					if _, ok := member.discoverCounter[m.Address]; !ok {
//						member.discoverCounter[m.Address] = 0
//					}
//					member.discoverCounter[m.Address]++
//					if member.discoverCounter[m.Address] > 2 {
//						Join(m)
//					}
//				}
//			}
//		}(addr, sbuf, n)
//	}
//}
