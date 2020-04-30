//+build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/knwoop/go-sample-api/domain"
	"github.com/knwoop/go-sample-api/infrastructure/api/handler"
	"github.com/knwoop/go-sample-api/infrastructure/persistence/datastore"
	"github.com/knwoop/go-sample-api/interface/controller"
	"github.com/knwoop/go-sample-api/usecase"
)

var wireSet = wire.NewSet(domain.WireSet, datastore.WireSet, usecase.WireSet, controller.WireSet, handler.WireSet)

func InitializeApp(conn *gorm.DB) (appHandler handler.AppHandler) {
	wire.Build(wireSet)
	return
}
