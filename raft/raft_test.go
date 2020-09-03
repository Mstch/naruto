package raft

import "testing"

func TestFirstElection(t *testing.T) {
	StartupStatDB()
	StartupKVDB()
	StartupTimer()
	StartLogDB()
	StartupServer()
	StartupClient()
}
