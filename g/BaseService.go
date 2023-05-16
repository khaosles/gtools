package g

/*
   @File: BaseService.go
   @Author: khaosles
   @Time: 2023/5/9 21:01
   @Desc:
*/

type BaseService[T ModelInterface] interface {
	Create(obj T)
	Delete(obj T)
	DeleteByID(id string)
	GetByID(id string) (obj T)
	GetBy(filedName string, value ...any) (obj T)
	GetAll() (objs []T)
	GetPage(pagination Pagination) (objs []T)
	Update(obj T)
}
