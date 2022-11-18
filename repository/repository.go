package repository

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

// FindOptions ...
type FindOptions struct {
	Limit int
	Page  int
}

// Repository ...
type Repository[T any] interface {
	NewSession(ctx context.Context) *gorm.DB
	GetDB() *gorm.DB
	Create(ctx context.Context, entity *T) error
	CreateInTransaction(ctx context.Context, db *gorm.DB, entity *T) error
	Find(ctx context.Context, entity interface{}, options *FindOptions) (list []*T, err error)
	FindInTransaction(ctx context.Context, db *gorm.DB, entity interface{}, options *FindOptions) (list []*T, err error)
	Count(ctx context.Context, entity interface{}) (count int64, err error)
	CountInTransaction(ctx context.Context, db *gorm.DB, entity interface{}) (count int64, err error)
	Get(ctx context.Context, id string) (entity *T, err error)
	GetInTransaction(ctx context.Context, db *gorm.DB, id string) (entity *T, err error)
	Update(ctx context.Context, entity *T) error
	UpdateInTransaction(ctx context.Context, db *gorm.DB, entity *T) error
	Delete(ctx context.Context, id string) (deleted bool, err error)
	DeleteInTransaction(ctx context.Context, db *gorm.DB, id string) (deleted bool, err error)
}

// Repository ...
type repository[T any] struct {
	db *gorm.DB
}

// NewSession ...
func (r *repository[T]) NewSession(ctx context.Context) *gorm.DB {
	return r.db.Session(&gorm.Session{}).Model(new(T)).WithContext(ctx)
}

// GetDB ...
func (r *repository[T]) GetDB() *gorm.DB {
	return r.db
}

// Create ...
func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	db := r.NewSession(ctx)

	return r.CreateInTransaction(ctx, db, entity)
}

// CreateInTransaction ...
func (r *repository[T]) CreateInTransaction(ctx context.Context, db *gorm.DB, entity *T) error {
	result := db.Create(&entity)
	if result.Error != nil {
		log.WithContext(ctx).Errorf("error on insert the result into model: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		err := errors.New("error on insert not inserted")
		log.WithContext(ctx).Error(err.Error())
		return err
	}

	return nil
}

// Find ...
func (r *repository[T]) Find(ctx context.Context, filters interface{}, options *FindOptions) (list []*T, err error) {
	db := r.NewSession(ctx)

	return r.FindInTransaction(ctx, db, filters, options)
}

// FindInTransaction ...
func (r *repository[T]) FindInTransaction(ctx context.Context, db *gorm.DB, filters interface{}, options *FindOptions) (list []*T, err error) {
	if options != nil {
		limit := 10
		if options.Limit != 0 {
			limit = options.Limit
		}
		db = db.Limit(limit)
		if options.Page != 0 {
			db = db.Offset((options.Page - 1) * limit)
		}
	}

	result := db.Find(&list, filters)

	err = result.Error
	if err != nil {
		log.WithContext(ctx).Errorf("Error on find: %v", err)
		return
	}

	return
}

// Count ...
func (r *repository[T]) Count(ctx context.Context, filters interface{}) (count int64, err error) {
	db := r.NewSession(ctx)

	return r.CountInTransaction(ctx, db, filters)
}

// CountInTransaction ...
func (r *repository[T]) CountInTransaction(ctx context.Context, db *gorm.DB, filters interface{}) (count int64, err error) {
	count = 0

	result := db.Where(filters).Count(&count)

	err = result.Error
	if err != nil {
		log.WithContext(ctx).Errorf("Error on count: %v", err)
		return
	}

	return
}

// Get ...
func (r *repository[T]) Get(ctx context.Context, id string) (entity *T, err error) {
	db := r.NewSession(ctx)

	return r.GetInTransaction(ctx, db, id)
}

// GetInTransaction ...
func (r *repository[T]) GetInTransaction(ctx context.Context, db *gorm.DB, id string) (entity *T, err error) {
	result := db.First(&entity, "id = ?", id)

	err = result.Error
	if err != nil {
		log.WithContext(ctx).Errorf("Error on find one result into collection: %v", err)
		return
	}

	return
}

// Update ...
func (r *repository[T]) Update(ctx context.Context, entity *T) error {
	db := r.NewSession(ctx)

	return r.UpdateInTransaction(ctx, db, entity)
}

// UpdateInTransaction ...
func (r *repository[T]) UpdateInTransaction(ctx context.Context, db *gorm.DB, entity *T) (err error) {
	result := db.Updates(entity)

	err = result.Error
	if err != nil {
		log.WithContext(ctx).Errorf("Error on update into collection: %v", err)
		return
	}

	return
}

// Delete ...
func (r *repository[T]) Delete(ctx context.Context, id string) (deleted bool, err error) {
	db := r.NewSession(ctx)

	return r.DeleteInTransaction(ctx, db, id)
}

// DeleteInTransaction ...
func (r *repository[T]) DeleteInTransaction(ctx context.Context, db *gorm.DB, id string) (deleted bool, err error) {
	entity, err := r.GetInTransaction(ctx, db, id)

	if err != nil {
		log.WithContext(ctx).Errorf("Error on get before delete: %v", err)
		return
	}

	if entity == nil {
		deleted = true
		return
	}

	result := db.Model(entity).Delete(entity)

	err = result.Error
	if err != nil {
		log.Errorf("Error on delete from collection: %v", err)
		return
	}

	deleted = result.RowsAffected == 1

	return
}
