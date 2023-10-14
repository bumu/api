package bootstrap

import (
	"noah/modules/dbmod"

	apigwHandler "noah/apps/apigw/handlers"
	defaultHandler "noah/apps/default/handlers"
	sgwHandler "noah/apps/sgw/handlers"

	"go.uber.org/fx"
)

func FxOptions() fx.Option {
	return fx.Options(
		fx.Provide(dbmod.NewConn),
		fx.Provide(
			New,
			NewMux,
		),

		// Register handlers.
		fx.Invoke(defaultHandler.Register),
		fx.Invoke(apigwHandler.Register),
		fx.Invoke(sgwHandler.Register),
	)
}
