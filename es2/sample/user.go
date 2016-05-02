package sample

import (
    "github.com/d-smith/go-examples/es2"
)

type User struct {
    *es2.Aggregate
}

func NewUser() *User {
    var user = new(User)
    user.Aggregate = es2.NewAggregate()
    return user
}

