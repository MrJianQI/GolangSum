package set

import (
	"sync"

	"github.com/dengjiawen8955/gostl/util/gosync"
)

// Options holds the Set's options
type Options struct {
	locker gosync.Locker
}

// Option is a function  type used to set Options
type Option func(option *Options)

func WithSync() Option {
	return func(option *Options) {
		option.locker = &sync.RWMutex{}
	}
}

type Set struct {
	//集合的底层用 map 实现
	setMap map[interface{}]struct{}
	//可选: 锁
	locker gosync.Locker
}

// 返回一个空的 set 对象
// opts 支持线程安全
func New(opts ...Option) (set *Set) {
	option := &Options{
		// 默认使用假锁,线程不安全
		locker: gosync.FakeLocker{},
	}
	//如果 opts 选项中有锁,将会在这里加锁
	for _, opt := range opts {
		opt(option)
	}
	return &Set{
		setMap: make(map[interface{}]struct{}),
		locker: option.locker,
	}
}

// 添加
func (s *Set) Add(element interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.setMap[element] = struct{}{}
}

// 删除
func (s *Set) Del(element interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()
	delete(s.setMap, element)
}

// 查看集合中是否含有元素
func (s *Set) Exists(element interface{}) (has bool) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	_, has = s.setMap[element]
	return
}

// 集合是否为空
func (s *Set) IsEmpty() (isEmpty bool) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	return len(s.setMap) == 0
}

// 返回集合长度
func (s *Set) Len() (length int) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	return len(s.setMap)
}

// 返回集合所有元素,乱序
func (s *Set) All() (elements []interface{}) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	elements = make([]interface{}, 0)
	for element, _ := range s.setMap {
		elements = append(elements, element)
	}
	return
}

// Inter 交集(默认返回线程不安全的集合)
func (s *Set) Inter(sets ...*Set) (resultSet *Set) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	resultSet = New()
	for e1, _ := range s.setMap {
		isInter := true
		for _, set := range sets {
			if !set.Exists(e1) {
				isInter = false
				break
			}
		}
		if isInter {
			resultSet.Add(e1)
		}
	}
	return
}

// Union 并集(默认返回线程不安全的集合)
// todo 使用迭代器
func (s *Set) Union(sets ...*Set) (resultSet *Set) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	resultSet = New()
	for e1 := range s.setMap {
		resultSet.Add(e1)
	}
	for _, set := range sets {
		for e2 := range set.setMap {
			resultSet.Add(e2)
		}
	}
	return
}

// Diff 差集(默认返回线程不安全的集合)
func (s *Set) Diff(sets ...*Set) (resultSet *Set) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	resultSet = New()
	for e1 := range s.setMap {
		isDiff := true
		for _, set := range sets {
			if set.Exists(e1) {
				isDiff = false
				break
			}
		}
		if isDiff {
			resultSet.Add(e1)
		}
	}
	return
}
