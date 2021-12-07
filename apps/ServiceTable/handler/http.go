package handler

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/model"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

// Command : SetAdd key value1,value2,
// 集合添加数据
func SetAdd(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	values := utils.GetUrlArg(r, "values")

	model.SetAddAt(key, utils.SliceDelNullString(strings.Split(values, ",")))
	utils.OutSucceedBodyJsonP(w, "")
	return
}

// Command : SetAddExpire key value timeUnix
// 集合添加数据并指定过期时间
func SetAddExpire(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	value := utils.GetUrlArg(r, "value")
	timeUnix := utils.GetUrlArgInt64(r, "time")
	model.SetAddExpireAt(key, value, timeUnix)
	utils.OutSucceedBodyJsonP(w, "")
	return
}

// Command : SetValueExpire key value timeUnix
// 指定集合数据过期时间
func SetValueExpire(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	value := utils.GetUrlArg(r, "value")
	timeUnix := utils.GetUrlArgInt64(r, "time")
	rse := model.SetValueExpireAt(key, value, timeUnix)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetGet key
// 获取集合所有数据
func SetGet(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	rse := model.SetGet(key)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetDel key
// 删除指定集合
func SetDel(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	rse := model.SetDelAt(key)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetDelValue key value
// 删除指定集合的元素
func SetDelValue(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	value := utils.GetUrlArg(r, "value")
	rse := model.SetDelValueAt(key, value)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// key 是否存在
// Command : KeyHas key
func KeyHas(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	rse := model.Key.Has(key)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// key 模糊查询
// Command : KeyLike key
func KeyLike(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	rse := model.Key.Like(key)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// key 列表
// Command : KeyAll
func KeyAll(w http.ResponseWriter, r *http.Request) {
	rse := model.Key.GetAll()
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// KV
// 增,改
// Command : KVAdd key value
func KVAdd(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	value := utils.GetUrlArg(r, "value")
	model.KVAddAt(key, value)
	utils.OutSucceedBodyJsonP(w, "")
	return
}

// Command : KVAddExpire key value expire
func KVAddExpire(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	value := utils.GetUrlArg(r, "value")
	expire := utils.GetUrlArgInt64(r, "expire")
	model.KVAddExpireAt(key, value, expire)
	utils.OutSucceedBodyJsonP(w, "")
	return
}

// Command : KVExpire key expire
func KVExpire(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	expire := utils.GetUrlArgInt64(r, "expire")
	rse := model.KVExpireAt(key, expire)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// 删
// Command : KVDel key
func KVDel(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	rse := model.KVDelAt(key)
	utils.OutSucceedBodyJsonP(w, rse)
	return
}

// 查
// Command : KVGet key
func KVGet(w http.ResponseWriter, r *http.Request) {
	key := utils.GetUrlArg(r, "key")
	is, rse := model.KVGet(key)
	if !is {
		utils.OutSucceedBodyJsonP(w, "没有这个key的数据")
		return
	}
	utils.OutSucceedBodyJsonP(w, rse)
	return
}


