package g

/*
   @File: BaseServiceImpl.go
   @Author: khaosles
   @Time: 2023/5/9 21:17
   @Desc:
*/

type BaseServiceImpl[T ModelInterface] struct {
}

func (service BaseServiceImpl[T]) Create(obj T) {
}

func (service BaseServiceImpl[T]) DeleteByID(id string) {

}

func (service BaseServiceImpl[T]) Delete(obj T) {

}

func (service BaseServiceImpl[T]) GetByID(id string) (obj T) {

	return
}

func (service BaseServiceImpl[T]) GetList() (objs []T) {

	return
}

func (service BaseServiceImpl[T]) GetPage(pagination Pagination) (objs []T) {

	return
}

func (service BaseServiceImpl[T]) Update(obj T) {

}
