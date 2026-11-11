package main

type contact struct {
	age int32
	sendingLimit int32
	userID string
}

type perms struct {
	canSend 		bool
	canReceive 		bool
	canManage 		bool
	permissionLevel int
}