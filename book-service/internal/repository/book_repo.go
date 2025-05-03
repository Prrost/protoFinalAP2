package repository

import (
	"context"
	"fmt"

	"github.com/Prrost/protoFinalAP2/book-service/internal/model"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) List(ctx context.Context, offset, limit int) ([]model.Book, int64, error) {
	var (
		books []model.Book
		count int64
	)
	if err := r.db.Model(&model.Book{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Offset(offset).Limit(limit).Find(&books).Error; err != nil {
		return nil, 0, err
	}
	return books, count, nil
}

func (r *BookRepo) Get(ctx context.Context, id string) (*model.Book, error) {
	var b model.Book
	if err := r.db.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookRepo) Create(ctx context.Context, b *model.Book) error {
	return r.db.Create(b).Error
}

func (r *BookRepo) Update(ctx context.Context, b *model.Book) error {
	return r.db.Save(b).Error
}

func (r *BookRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(&model.Book{}, "id = ?", id).Error
}

func (r *BookRepo) AdjustQuantity(ctx context.Context, id string, delta int32) (*model.Book, error) {
	b, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	b.AvailableCopies += delta
	if b.AvailableCopies < 0 {
		return nil, fmt.Errorf("not enough copies")
	}
	if err := r.db.Save(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}
