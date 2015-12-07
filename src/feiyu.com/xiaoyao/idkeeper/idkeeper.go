package xyidkeeper

import (
	"time"
)

// 定义标识管理器

//生成新的64位唯一标识
//|----------------------|-------|--------|
//         42 bit             10bits   12 bit
//
//高42位 时间戳，粒度为毫秒级别。137年后时间戳溢出
//中10位 服务节点id，其中的高4位为数据中心标识（gateway id）；低6位为节点标识（nodeid）。最多可以有16个gateway，每个gateway下可以挂64个nodeid
//低12位为计数器 最大支持每毫秒 4096个id分配
func NewUint64Id(timestamp, dcid, nodeid, counter int64) (newid uint64) {
	return uint64(timestamp<<22 + dcid<<18 + nodeid<<12 + counter)
}

//id管理器的时间戳起点是2014-08-14 14:37:18
//用这个时间是因为代码是这个时间点写的~ o(∩_∩)o
//单位是毫秒
var DefBeginTimeStamp = IdKeeperBeginTimeStamp()

func IdKeeperBeginTimeStamp() int64 {
	t, err := time.Parse("2006-01-02 15:04:05", "2014-08-14 14:37:18")
	if err != nil {
		return 0
	}

	return t.UnixNano() / int64(time.Millisecond)
}

type IdKeeper struct {
	m         sync.Mutex
	name      string //业务名称
	from      int64  //起始时间
	dcid      int64  //数据中心标识
	nodeid    int64  //节点标识
	counter   int64  //计数器
	timestamp int64  //计数器有效时间戳
	//-------------------------------
	idsum          int64 //分配的id总数
	begintimestamp int64 //启动分配的时间戳
}

func NewIdKeeper(dcid, nodeid int64, name string) *IdKeeper {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	return &IdKeeper{
		name:           name,
		from:           DefBeginTimeStamp,
		dcid:           dcid,
		nodeid:         nodeid,
		timestamp:      now,
		begintimestamp: now,
		counter:        0,
	}
}

//生成新的整型(uint64) id
func (ig *IdGenerater) NewID() (id uint64) {

	//如果时间戳变化，重置计数和时间戳
	now := xyutil.CurTimeMs()
	ig.m.Lock()
	if ig.timestamp < now {
		ig.counter = 0
		ig.timestamp = now
	}

	id = NewUint64Id(ig.timestamp-ig.from, ig.dcid, ig.nodeid, ig.counter)

	ig.counter++
	ig.m.Unlock()

	return
}
