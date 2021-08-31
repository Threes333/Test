package gid

import (
	"sync"
	"time"
)

const (
	MachineId       = 1    //机器id
	MaxSequence     = 4095 //同一时间最大序列号
	TimeOffset      = 22   //时间戳偏移量
	MachineIdOffset = 12   //机器id偏移量
)

type IdMaker struct {
	Mu       sync.Mutex //	用于保证并发安全
	PreTime  int64      //	上一次生成id时间
	Machine  int64      //	机器id
	Sequence int64      //  序列号
}

//New一个IdMaker
func NewIdMaker() *IdMaker {
	return &IdMaker{Machine: MachineId, Sequence: -1}
}

//获得毫秒级别的当前时间
func getMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//生成一个新的id
func (i *IdMaker) NewId() int {
	i.Mu.Lock()
	defer i.Mu.Unlock()
	curTime := getMillisecond()
	if curTime == i.PreTime {
		//序列号大于当前时间所能表示的最大范围则等待下一时间的到来
		if i.Sequence > MaxSequence {
			time.Sleep(time.Millisecond)
			curTime = getMillisecond()
			i.Sequence = 0
		} else {
			i.Sequence++
		}
	} else {
		//记录这次生成id的时间
		i.PreTime = curTime
		i.Sequence = 0
	}
	return int(curTime<<TimeOffset | i.Machine<<MachineIdOffset | i.Sequence)
}

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(5000)
//	m := NewIdMaker()
//	for i:=0; i<5000; i++ {
//		go func() {
//			fmt.Println(m.NewId())
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}
