package gerror

/*
   @File: redis.go
   @Author: khaosles
   @Time: 2023/3/5 11:53
   @Desc:
*/

var (
	RedisKeyNotFoundException = Exception("RedisKeyNotFound")
)
