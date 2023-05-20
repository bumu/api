package repos

import (
	"context"

	"apigw/apps/apigw/data/schema"

	"gorm.io/gorm"
)

type KeyRepo struct {
	Conn *gorm.DB
}

func NewKeyRepo(conn *gorm.DB) *KeyRepo {
	return &KeyRepo{conn}
}

func (r KeyRepo) Create(ctx context.Context, entity *schema.Key) error {
	return r.Conn.Create(entity).Error
}

func (r KeyRepo) First(ctx context.Context, token string) (*schema.Key, error) {
	dst := &schema.Key{}
	err := r.Conn.Where("token", token).First(dst).Error
	return dst, err
}

func (r KeyRepo) List(ctx context.Context, limit, offset int) ([]*schema.Key, error) {
	dst := []*schema.Key{}
	err := r.Conn.Debug().Limit(limit).Offset(offset).Find(&dst).Error
	return dst, err
}
