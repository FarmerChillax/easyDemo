package gorm_utils

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrModelNotBeNil         = errors.New("model not be nil")
	ErrIdLessThanZero        = errors.New("id not less than zero")
	ErrModelsNotBeEmptySlice = errors.New("models not be empty slice")
	ErrConditionNotBeNil     = errors.New("condition not be nil")
	ErrIdsNotBeEmpty         = errors.New("ids not be empty array")
)

type ModelT interface {
	TableName() string
}

type CommonRepo[M ModelT] struct {
	db *gorm.DB
}

func (r *CommonRepo[M]) Create(model *M) error {
	db := r.db.Create(model)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (r *CommonRepo[M]) CreateBatch(models []*M) error {
	db := r.db.Create(models)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (r *CommonRepo[M]) Update(model *M, condition *M, fields ...string) error {
	db := r.db
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	db = db.Where(condition).Updates(model)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (r *CommonRepo[M]) UpdateById(model *M, id int64) error {
	db := r.db.Where("id = ?", id).Updates(model)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// more see: https://gorm.io/zh_CN/docs/create.html#Upsert-%E5%8F%8A%E5%86%B2%E7%AA%81
func (r *CommonRepo[M]) Upsert(models *[]M) error {
	db := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(models)

	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// 通过条件删除记录
func (r *CommonRepo[M]) Delete(condition *M, deleteUser ...string) error {
	if condition == nil {
		return ErrConditionNotBeNil
	}

	var deleteUser0 string
	if len(deleteUser) > 0 {
		deleteUser0 = deleteUser[0]
	}

	db := r.db.Where("deleted_at = ?", 0).Where(condition).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": deleteUser0,
	})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// 通过 IDs 删除记录
func (r *CommonRepo[M]) DeleteByIds(ids []int64, deleteUser ...string) error {
	if len(ids) <= 0 {
		return ErrIdsNotBeEmpty
	}

	var deleteUser0 string
	if len(deleteUser) > 0 {
		deleteUser0 = deleteUser[0]
	}

	db := r.db.Where(`deleted_at IS NULL AND id IN ?`, ids).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": deleteUser0,
	})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// 通过 ID 删除记录
func (r *CommonRepo[M]) DeleteById(id int64, deleteUser ...string) error {
	if id <= 0 {
		return ErrIdLessThanZero
	}

	err := r.DeleteByIds([]int64{id}, deleteUser...)
	if err != nil {
		return err
	}
	return nil
}

// 通过条件查询记录
func (r *CommonRepo[M]) Get(condition *M, fields ...string) (*M, error) {
	result, err := r.GetExist(condition, fields...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// 通过 ID 查询记录，如果不存在则会报错
func (r *CommonRepo[M]) GetExist(condition *M, fields ...string) (*M, error) {
	var record M
	db := r.db.Where(condition)
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	db = db.First(&record)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// 通过 ID 查询记录
func (r *CommonRepo[M]) GetById(id int64, fields ...string) (*M, error) {
	result, err := r.GetExistById(id, fields...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// 通过 ID 查询记录(包括被软删除的数据)
func (r *CommonRepo[M]) GetExistUnscopedById(id int64, fields ...string) (*M, error) {
	if id <= 0 {
		return nil, ErrIdLessThanZero
	}
	var record M
	db := r.db.Session(&gorm.Session{})
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	db = db.Unscoped().First(&record, id)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// 通过 ID 查询记录，如果不存在则会报错
func (r *CommonRepo[M]) GetExistById(id int64, fields ...string) (*M, error) {
	if id <= 0 {
		return nil, ErrIdLessThanZero
	}
	var record M
	db := r.db
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	db = db.First(&record, id)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// 通过条件查询是否存在记录
func (r *CommonRepo[M]) Exist(condition *M) (bool, error) {
	_, err := r.GetExist(condition)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// 通过条件查询是否存在记录
func (r *CommonRepo[M]) ExistById(id int64) (bool, error) {
	_, err := r.GetExistById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// 通过条件查询是否存在记录，排除指定 IDs
func (r *CommonRepo[M]) ExistExcludeIds(condition *M, ids []int64) (bool, error) {
	if len(ids) <= 0 {
		return false, ErrIdsNotBeEmpty
	}

	var record M
	db := r.db.Where("id NOT IN ?", ids).Where(condition)
	db = db.First(&record)
	if err := db.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// 通过条件查询是否存在记录，排除指定 ID
func (r *CommonRepo[M]) ExistExcludeId(condition *M, id int64) (bool, error) {
	if id <= 0 {
		return false, ErrIdLessThanZero
	}
	return r.ExistExcludeIds(condition, []int64{id})
}

// 通过条件查询记录列表
func (r *CommonRepo[M]) List(condition *M, fields ...string) ([]*M, error) {
	var records []*M
	db := r.db.Where(condition)
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	db = db.Find(&records)
	if err := db.Error; err != nil {
		return records, err
	}
	return records, nil
}
