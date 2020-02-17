package main

import (
	"dashboard/api"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(requestlogger.Module),
			new(api.Module),
		},
	)
}