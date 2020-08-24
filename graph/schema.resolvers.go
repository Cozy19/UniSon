package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/rekksson/UniSon/graph/generated"
	"github.com/rekksson/UniSon/graph/model"
)

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var links []*model.Order

	ct, err := getEchoContext(ctx)
	if err != nil {
		return nil, err
	}

	user := (ct).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["sub"].(string)

	links = append(links, &model.Order{ID: "User", Name: name})
	return links, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
