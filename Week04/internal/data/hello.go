package data

import (
	"Go-000/Week04/internal/entity"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/go-redis/redis"
)

// 有循环导入问题
//var _ biz.HelloRepo = (*HelloRepository)(nil)

type HelloRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewHelloRepository(db *gorm.DB, cache *redis.Client) *HelloRepository {
	return &HelloRepository{
		db:    db,
		cache: cache,
	}
}

func (d *HelloRepository) GetHello(id int64) (*entity.Hello, error) {
	var hello entity.Hello
	if err := d.db.Model(&hello).Where("id = ?", id).First(&hello).Error; err != nil {
		return nil, errors.Wrap(err, "get hello")
	}
	return &hello, nil
}
