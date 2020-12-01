package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

var (
	DataNotExist   = errors.New("db: data not exists from db")
	DBConnectError = errors.New("db: connect error")
)

type User struct {
	ID   int
	name string
}
type UserDao struct{}

func (u *UserDao) getUserIDFromDB(ID int) (*User, error) {
	log.Printf("user %d to find User in database", ID)
	return nil, sql.ErrNoRows
}

func (u *UserDao) getUserFromID(ID int) (*User, error) {
	user, err := u.getUserIDFromDB(ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, DataNotExist
		}

		return nil, DBConnectError
	}

	return user, nil
}

func BizUserDetail(ID int) (*User, error) {
	u := &UserDao{}
	user, err := u.getUserFromID(ID)
	if err != nil {
		return nil, errors.WithMessagef(err, "biz query user: %d ", ID)
	}
	return user, nil
}

func main() {
	user, err := BizUserDetail(1)
	if err != nil {
		if errors.Is(err, DataNotExist) {
			log.Printf("stack trace: \n%+v\n", err)
			return
		}

		log.Printf("query user detail failed: %+v\n", err)
		return
	}

	log.Printf("user info: %+v\n", user)
}
