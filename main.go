package main

import "github.com/Mstch/naruto/helper/member"

func main() {
	member.Startup()
	<-member.OK
}
