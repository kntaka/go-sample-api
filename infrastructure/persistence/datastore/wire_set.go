package datastore

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewUserRepository,
	NewUserActiveRepository,
	NewUserDetailRepository,
	NewUserLeaveRepository,
)
