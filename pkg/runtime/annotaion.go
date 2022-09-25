package runtime

import (
	"github.com/qhzhyt/harvest/pkg/annotation"
	"sync"
)

var (
	targetMap        = map[annotation.TargetID]*annotation.Target{}
	targetMapRWMutex = &sync.RWMutex{}
)

func RegisterTargets(targets []*annotation.Target) {
	targetMapRWMutex.Lock()
	defer targetMapRWMutex.Unlock()

	for _, target := range targets {
		targetMap[target.ID] = target.Clone()
	}
}

func GetTargetByID(id annotation.TargetID) *annotation.Target {
	targetMapRWMutex.RLock()
	defer targetMapRWMutex.RUnlock()
	return targetMap[id].Clone()
}
