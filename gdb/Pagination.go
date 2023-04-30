package gdb

/*
   @File: Pagination.go
   @Author: khaosles
   @Time: 2023/4/30 09:07
   @Desc: 分页查询选择器
*/

type Pagination struct {
	PageSize int    `json:"pageSize" uri:"pageSize"`
	Page     int    `json:"page" uri:"page"`
	Sort     string `json:"sort" uri:"sort"`
}

func NewPagination(pageSize, page int, sort string) Pagination {
	if pageSize < 1 {
		pageSize = 1
	}
	if page < 1 {
		page = 1
	}
	return Pagination{
		PageSize: pageSize,
		Page:     page,
		Sort:     sort,
	}
}

//type PaginationDto[T any] struct {
//	TotalPages int64 `json:"totalPages" uri:"totalPages"`
//	TotalCount int64 `json:"totalCount" uri:"totalCount"`
//	Objs       T     `json:"objs" uri:"objs"`
//}
