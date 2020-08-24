package main

import (
	"context"
	"io/ioutil"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/rekksson/UniSon/graph"
	"github.com/rekksson/UniSon/graph/generated"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomContext struct {
	echo.Context
	ctx context.Context
}

func ProcessEchoHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c)
		c.SetRequest(c.Request().WithContext(ctx))

		cc := &CustomContext{c, ctx}

		return next(cc)
	}
}

func main() {
	e := echo.New()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	e.Use(ProcessEchoHeader)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	publicKeyData, err := ioutil.ReadFile("key.pub.pem")
	if err != nil {
		panic(err)
	}

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyData))

	g := e.Group("/query")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    publicKey,
		SigningMethod: "RS256",
		ContextKey:    "user",
	}))

	g.POST("", echo.WrapHandler(srv))

	e.Start(":3000")
}
