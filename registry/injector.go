//+build wireinject


package registry

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/takkee/go-sample-api/domain"
	"github.com/takkee/go-sample-api/infrastructure/persistence/datastore"
	"github.com/takkee/go-sample-api/usecase"
	"github.com/takkee/go-sample-api/interface/controller"
	"github.com/takkee/go-sample-api/infrastructure/api/handler"
)

var wireSet = wire.NewSet(domain.WireSet, datastore.WireSet, usecase.WireSet, controller.WireSet, handler.WireSet)

func InitializeApp(conn *gorm.DB) (appHandler handler.AppHandler) {
	wire.Build(wireSet)
	return
}
