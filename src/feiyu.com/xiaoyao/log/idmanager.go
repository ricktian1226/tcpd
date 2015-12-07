package xylog

import (
	"feiyu.com/xiaoyao/cache"
)

const (
	DefaultLogId      = ""
	DEBUGUSER_SUBJECT = "debuguserreload"
)

type empty struct{}

type MapId map[interface{}]empty

func (m *MapId) Reset() {
	*m = make(MapId, 0)
}

func NewMapId() *MapId {
	m := make(MapId, 0)
	return &m
}

type IdManager struct {
	caches [2]*MapId
	xycache.CacheBase
}

//重新加载id列表信息
// ids []interface{} 列表信息
func (im *IdManager) Load(ids []interface{}) {
	secondary := im.caches[int(im.Secondary())]
	secondary.Reset()
	for _, id := range ids {
		(*secondary)[id] = empty{}
	}

	im.Switch()

	//im.Print()
}

//判断id是否存在
func (im *IdManager) IsIdExist(id interface{}) bool {

	if _, ok := (*im.caches[int(im.Major())])[id]; ok {
		return true
	}

	return false
}

func (im *IdManager) Print() {
	major := im.caches[int(im.Major())]

	DebugNoId("---------- xylog Ids begin ------------")
	for k, _ := range *major {
		DebugNoId("%v", k)
	}
	DebugNoId("---------- xylog Ids end ------------")
}

func (im *IdManager) String() (str string) {
	//major := im.caches[int(im.Major())]

	//str = "xylog.Ids :["
	//for k, _ := range *major {
	//	//str += fmt.Sprintf("%v,", k)
	//}
	//str += "]"
	return
}

func NewIdManager() (im *IdManager) {
	im = new(IdManager)
	for i := 0; i < 2; i++ {
		im.caches[i] = NewMapId()
	}
	return
}

var DefIdManager = NewIdManager()
