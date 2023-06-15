package src

import (
	"github.com/shinochiha/gotoko/app"
	"github.com/shinochiha/gotoko/src/callback"
	"github.com/shinochiha/gotoko/src/contact"
	"github.com/shinochiha/gotoko/src/paymentinvoice"
	"github.com/shinochiha/gotoko/src/paymentsubscription"
	"github.com/shinochiha/gotoko/src/product"
	// import : DONT REMOVE THIS COMMENT
)

func Migrator() *migratorUtil {
	if migrator == nil {
		migrator = &migratorUtil{}
		migrator.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			migrator.Run()
		}
		migrator.isConfigured = true
	}
	return migrator
}

var migrator *migratorUtil

type migratorUtil struct {
	isConfigured bool
}

func (*migratorUtil) Configure() {
	app.DB().RegisterTable("main", contact.Contact{})
	app.DB().RegisterTable("main", product.Product{})
	app.DB().RegisterTable("main", paymentinvoice.PaymentInvoice{})
	app.DB().RegisterTable("main", paymentinvoice.Item{})
	app.DB().RegisterTable("main", paymentsubscription.PaymentSubscription{})
	app.DB().RegisterTable("main", callback.Callback{})
	// RegisterTable : DONT REMOVE THIS COMMENT
}

func (*migratorUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().MigrateTable(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
