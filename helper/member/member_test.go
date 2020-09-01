package member

import (
	"github.com/Mstch/naruto/helper/util"
	"testing"
)

func TestStartup(t *testing.T) {
	buf := make([]byte, 4)
	util.WriteInt32ToBytes(-1*1234, buf)
	println(uint32(-util.BytesToInt32(buf[:4])))
}
