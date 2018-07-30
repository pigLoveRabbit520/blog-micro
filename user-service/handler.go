package main

import (
	pb "blog-micro/user-service/proto"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// 定义服务
type handler struct {
	repo         Repository
	tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User) (resp *pb.Response, err error) {
	// 哈希处理用户输入的密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPwd)
	if err := h.repo.Create(req); err != nil {
		return nil, err
	}
	resp = &pb.Response{}
	resp.User = req
	return
}

func (h *handler) Get(ctx context.Context, req *pb.User) (resp *pb.Response, err error) {
	u, err := h.repo.Get(req.Uid)
	if err != nil {
		return nil, err
	}
	resp.User = u
	return
}

func (h *handler) Auth(ctx context.Context, req *pb.User) (resp *pb.Token, err error) {
	// 在 part3 中直接传参 &pb.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// 即 req.Password 与 u.Password 都是加密后的密码
	// 将无法通过验证
	u, err := h.repo.Get(req.Uid)
	if err != nil {
		return nil, err
	}

	// 进行密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	t, err := h.tokenService.Encode(u)
	if err != nil {
		return nil, err
	}
	resp.Token = t
	return
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token) (resp *pb.Token, err error) {
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.User.Uid <= 0 {
		return nil, errors.New("invalid user")
	}

	resp.Valid = true
	return
}
