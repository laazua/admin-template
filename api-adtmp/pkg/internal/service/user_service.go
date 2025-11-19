package service

import (
	"context"

	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/entities/dto"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/domain/entities/mapper"
	"adtmp/pkg/internal/domain/repositories"
	"adtmp/pkg/security"
)

type userService struct {
	pwdProvider security.PwdProvider
	// userRepo    repositories.UserRepository
	userRepo repositories.Repo[entities.User]
}

func NewUserService(pwdProvider security.PwdProvider, userRepo repositories.Repo[entities.User]) UserService {
	return &userService{pwdProvider: pwdProvider, userRepo: userRepo}
}

func (userService *userService) CreateUser(ctx context.Context, user *form.UserCreate) error {
	dbUser := mapper.ToUserDb(user)
	passwordHash, err := userService.pwdProvider.HashPwd(dbUser.Password)
	if err != nil {
		return err
	}
	dbUser.Password = passwordHash

	return userService.userRepo.Create(ctx, dbUser)
}

func (userService *userService) DestroyUser(ctx context.Context, id uint) error {
	return userService.userRepo.Destroy(ctx, id)
}

func (userService *userService) UpdateUser(ctx context.Context, id uint, user *form.UserUpdate) error {
	dbUser := form.ToUserDb(user)
	passwordHash, err := userService.pwdProvider.HashPwd(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	return userService.userRepo.Update(ctx, id, dbUser)
}

func (userService *userService) GetById(ctx context.Context, id uint) (*dto.UserResponse, error) {
	repoUser, err := userService.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToDtoResp(&repoUser).(*dto.UserResponse), nil
}

func (userService *userService) ListUser(ctx context.Context, limit, offset int) ([]*dto.UserResponse, error) {
	usersRepo, err := userService.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	userDtos := mapper.ToDtoListResp(usersRepo, func(e entities.User) *dto.UserResponse {
		return &dto.UserResponse{
			ID: e.ID, Name: e.Name, Email: e.Email,
		}
	})
	return userDtos, nil
}
