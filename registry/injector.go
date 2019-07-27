//+build wireinject


package registry

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/kntaka/go-sample-api/domain"
	"github.com/kntaka/go-sample-api/infrastructure/persistence/datastore"
	"github.com/kntaka/go-sample-api/usecase"
	"github.com/kntaka/go-sample-api/interface/controller"
	"github.com/kntaka/go-sample-api/infrastructure/api/handler"
)

var wireSet = wire.NewSet(domain.WireSet, datastore.WireSet, usecase.WireSet, controller.WireSet, handler.WireSet)

func InitializeApp(conn *gorm.DB) (appHandler handler.AppHandler) {
	wire.Build(wireSet)
	return
}
