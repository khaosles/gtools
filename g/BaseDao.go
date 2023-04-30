package g

import (
	"gorm.io/gorm"
)

/*
   @File: BaseDao.go
   @Author: khaosles
   @Time: 2023/4/29 22:33
   @Desc:
*/

type BaseDao[T any] struct {
	db *gorm.DB
}

func (dao *BaseDao[T]) SetDB(db *gorm.DB) *BaseDao[T] {
	dao.db = db
	return dao
}

// Create 创建
func (dao BaseDao[T]) Create(obj *T) (err error) {
	err = dao.db.Create(obj).Error
	return
}

// GetByID 根据id查询
func (dao BaseDao[T]) GetByID(id string) (obj T, err error) {
	err = dao.db.Where("id = ?", id).First(&obj).Error
	return
}

// GetByColumn 根据唯一值列查询
func (dao BaseDao[T]) GetByColumn(column, value string) (obj T, err error) {
	err = dao.db.Where(column+" = ?", value).First(&obj).Error
	return
}

// GetByFilter 根据条件查询
func (dao BaseDao[T]) GetByFilter(filter string, value ...any) (objs []T, err error) {
	err = dao.db.Where(filter, value).Find(&objs).Error
	return
}

// GetByIDs 根据多个id查询
func (dao BaseDao[T]) GetByIDs(ids ...string) (objs []T, err error) {
	err = dao.db.Where("id in (?)", ids).Find(&objs).Error
	return
}

// GetByPage 分页查询
func (dao BaseDao[T]) GetByPage(pagination Pagination, filter string, args ...any) (objs []T, totalPages int64, err error) {
	var totalCount int64
	// 计算总记录数
	if err = dao.db.Model(objs).Count(&totalCount).Error; err != nil {
		return
	}
	// 获取总页数
	totalPages = totalCount / int64(pagination.PageSize)
	if totalCount%int64(pagination.PageSize) > 0 {
		totalPages++
	}
	// 当前页
	pageIndex := (pagination.Page - 1) * pagination.PageSize
	err = dao.db.Order(pagination.Sort).
		Offset(pageIndex).
		Limit(pagination.PageSize).
		Where(filter, args).
		Find(&objs).
		Error
	return
}

// GetByPageALL 分页查询全部
func (dao BaseDao[T]) GetByPageALL(pagination Pagination) (objs []T, totalPages int64, err error) {
	var totalCount int64
	// 计算总记录数
	if err = dao.db.Model(objs).Count(&totalCount).Error; err != nil {
		return
	}
	// 获取总页数
	totalPages = totalCount / int64(pagination.PageSize)
	if totalCount%int64(pagination.PageSize) > 0 {
		totalPages++
	}
	// 获取当前页索引
	pageIndex := (pagination.Page - 1) * pagination.PageSize
	err = dao.db.Order(pagination.Sort).
		Offset(pageIndex).
		Limit(pagination.PageSize).
		Find(&objs).
		Error
	return
}

// GetList 获取多条数据
func (dao BaseDao[T]) GetList(filter string, value ...any) (objs []T, err error) {
	err = dao.db.Where(filter, value).Find(&objs).Error
	return
}

// GetAll 获取全部数据
func (dao BaseDao[T]) GetAll() (objs []T, err error) {
	err = dao.db.Find(&objs).Error
	return
}

// DeleteByIDs 通过一个或者多个id删除
func (dao BaseDao[T]) DeleteByIDs(ids ...string) (err error) {
	var objs []T
	err = dao.db.Where("id in (?)", ids).Find(&objs).Error
	if err != nil {
		return
	}
	err = dao.DeleteByObjs(objs...)
	return
}

// DeleteByIDs 通过一个或者多个id删除
func (dao BaseDao[T]) DeleteByColumns(column string, value ...string) (err error) {
	var objs []T
	err = dao.db.Where(column+" in (?)", value).Find(&objs).Error
	if err != nil {
		return
	}
	err = dao.DeleteByObjs(objs...)
	return
}

// DeleteByObjs 根据对象删除
func (dao BaseDao[T]) DeleteByObjs(objs ...T) (err error) {
	err = dao.db.Delete(&objs).Error
	return
}

// UpdateByObj 通过对象更新
func (dao BaseDao[T]) UpdateByObj(obj *T) (err error) {
	err = dao.db.Save(obj).Error
	return err
}

// UpdateByColumn 通过列更新
func (dao BaseDao[T]) UpdateByColumn(obj *T, value any) (err error) {
	err = dao.db.Model(obj).Updates(value).Error
	return err
}

// RawQuery 原始查询
func (dao BaseDao[T]) RawQuery(sql string, values ...any) (objs []T, err error) {
	err = dao.db.Exec(sql, values).Scan(&objs).Error
	return
}

// Raw 原始cud
func (dao BaseDao[T]) Raw(sql string, values ...any) (err error) {
	err = dao.db.Exec(sql, values).Error
	return
}
