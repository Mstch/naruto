package main

import "github.com/Mstch/naruto/raft"

func main() {
	raft.StartupStatDB()
	raft.StartupKVDB()
	raft.StartLogDB()
	raft.StartupServer()
	raft.StartupClient()
	raft.StartupTimer()
	select {}
}
