package user

type UserService struct {
	userRepo UserRepository
}

type UserServiceParam struct {
	UserRepo UserRepository
}

func NewUserService(param UserServiceParam) *UserService {
	return &UserService{
		userRepo: param.UserRepo,
	}
}
