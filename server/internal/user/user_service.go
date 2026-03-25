package user

import (
	"context"
	"strconv"
	"time"

	"github.com/musishere/chat-app/internal/util"
)

type service struct {
	Repository Repository
	timeout    time.Duration
}

func NewService(repo Repository) Service {
	return &service{
		Repository: repo,
		timeout:    time.Duration(2) * time.Millisecond,
	}
}

func (s *service) CreateUser(c context.Context, req CreateUserReq) (CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)

	// hash the user password\
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return CreateUserRes{}, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return CreateUserRes{}, err
	}

	res := CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	defer cancel()
	return res, nil
}
