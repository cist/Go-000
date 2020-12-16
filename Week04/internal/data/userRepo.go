package data

import (
	"Week04/internal/biz"
	"log"
	"math/rand"
)

var _ biz.UserRepo = new(userRepo)

func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

type userRepo struct{}

func (r *userRepo) Save(u *biz.User) int32 {
	uid := rand.Int31()
	log.Printf("save user, name: %s, age: %d, id: %d", u.Name, u.Age, uid)
	return uid
}
