// 泛型实现

package store

import (
	"context"

	"adtmp/internal/domain/repositories"

	"gorm.io/gorm"
)

// 基础 Repository 实现
type baseRepo[E repositories.Entity] struct {
	db *gorm.DB
}

func newBaseRepo[E repositories.Entity](db *gorm.DB) *baseRepo[E] {
	return &baseRepo[E]{db: db}
}

func (r *baseRepo[E]) Create(ctx context.Context, e *E) error {
	// var entity E
	// // 这里需要将 form 转换为 entity，可以使用 mapstruct 或其他映射工具
	// if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
	// 	return nil, err
	// }

	return gorm.G[E](r.db).Create(ctx, e)
}

func (r *baseRepo[E]) Destroy(ctx context.Context, id uint) error {
	// var entity E
	// return r.db.WithContext(ctx).Delete(&entity, id).Error
	_, err := gorm.G[E](r.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (r *baseRepo[E]) Update(ctx context.Context, id uint, e *E) error {
	// var entity E
	// return r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).Updates(&entity).Error
	_, err := gorm.G[E](r.db).Where("id = ?", id).Updates(ctx, *e)
	return err
}

func (r *baseRepo[E]) GetById(ctx context.Context, id uint) (E, error) {
	// var entity E
	// if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
	// 	return nil, err
	// }
	return gorm.G[E](r.db).Where("id = ?", id).First(ctx)
}

func (r *baseRepo[E]) List(ctx context.Context, limit, offset int) ([]E, error) {
	var entities []E
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&entities).Error
	return entities, err
}
