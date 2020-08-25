package stupid

import (
	"github.com/Mstch/naruto/raft/msg"
	"testing"
)

func TestClientImpl_Notify(t *testing.T) {
	  (&ClientImpl{}).Notify("FUCK", &msg.HeartbeatReq{
		Id:           1,
		Term:         2,
		LeaderCommit: 3,
	})

}
