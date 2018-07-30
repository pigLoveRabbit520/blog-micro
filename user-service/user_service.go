package main

import (
	"github.com/jinzhu/gorm"
	pb "blog-micro/user-service/proto"
)

type Repository interface {
	Get(id int32) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(*pb.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Get(id int32) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("uid = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(u *pb.User) error {
	if err := repo.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByUsername(username string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("username = ? ", username).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}