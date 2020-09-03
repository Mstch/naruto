package member

import (
	"github.com/Mstch/naruto/helper/member/file"
	"github.com/Mstch/naruto/helper/member/domain"
)

type Members interface {
	//block until discover complete
	Discover()
	GetMembers() map[uint32]*domain.Member
	Self() *domain.Member
}

func Default() Members {
	return file.NewFileMembers()
}
