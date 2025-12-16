package starter

import (
	"capital-gains/src/application/handlers"
	"capital-gains/src/driven/capitalgains"
	"capital-gains/src/driven/operations"
	"capital-gains/src/driver/console"
)

type Dependencies struct {
	CalculateCapitalGain console.CalculateCapitalGain
}

func NewDependencies() Dependencies {
	operationsRepository := operations.NewRepository()
	capitalGainsRepository := capitalgains.NewRepository()

	registerBuyHandler := handlers.NewRegisterBuyHandler(operationsRepository)
	registerSellHandler := handlers.NewRegisterSellHandler(operationsRepository)
	calculateCapitalGainHandler := handlers.NewCalculateCapitalGainHandler(
		operationsRepository,
		capitalGainsRepository,
	)

	defaultConsole := console.NewDefaultConsole()
	calculateCapitalGain := console.NewCalculateCapitalGain(
		defaultConsole,
		registerBuyHandler,
		registerSellHandler,
		capitalGainsRepository,
		calculateCapitalGainHandler,
	)

	return Dependencies{
		CalculateCapitalGain: *calculateCapitalGain,
	}
}
