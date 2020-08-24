package graph

import (
	"context"
	"fmt"

	"github.com/labstack/echo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func getEchoContext(ctx context.Context) (echo.Context, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return nil, err
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return nil, err
	}
	return ec, nil
}

type Resolver struct{}
