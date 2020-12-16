package biz

type User struct {
	Name string
	Age  int32
}

type UserRepo interface {
	Save(*User) int32
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

type UserUsecase struct {
	repo UserRepo
}

func (s *UserUsecase) SaveUser(u *User) int32 {
	return s.repo.Save(u)
}
