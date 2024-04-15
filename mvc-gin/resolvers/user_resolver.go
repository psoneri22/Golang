package resolvers

import (
	"context"
	"mvc-gin/models"
)

type Resolver struct{}

func (r *Resolver) Users(ctx context.Context) ([]*models.User, error) {
	users := models.GetAllUsers()
	var result []*models.User
	for i := range users {
		result = append(result, &users[i])
	}
	return result, nil
}

func (r *Resolver) CreateUser(ctx context.Context, args struct {
	Name string
	Age  int
}) (*models.User, error) {
	user := &models.User{Username: args.Name}
	// Call CreateUser function from models package
	if err := models.CreateUser(user); err != nil {
		return nil, err
	}
	// Return the created user
	return user, nil
}
