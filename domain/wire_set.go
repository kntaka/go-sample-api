package domain

import (
	"github.com/google/wire"
	"github.com/knwoop/go-sample-api/domain/service"
)

var WireSet = wire.NewSet(
	service.NewProfileService,
	service.NewUserService,
)
