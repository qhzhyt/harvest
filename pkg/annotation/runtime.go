package annotation

import (
    "sync"
)

var (
    annotationMap       = map[ID]*Annotation{}
    targetMap           = map[TargetID]*Target{}
    targetsByAnnotation = map[Name][]*Target{}
    targetMapRWMutex    = &sync.RWMutex{}
)

func RegisterTargets(targets []*Target) {
    targetMapRWMutex.Lock()
    defer targetMapRWMutex.Unlock()

    for _, target := range targets {
        targetMap[target.ID] = target.Clone()
        for _, ann := range target.Annotations {
            targetsByAnnotation[ann.Name] = append(targetsByAnnotation[ann.Name], targetMap[target.ID])
            annotationMap[ann.ID] = ann
        }
    }
}

func GetTargetByID(id TargetID) *Target {
    targetMapRWMutex.RLock()
    defer targetMapRWMutex.RUnlock()

    return targetMap[id].Clone()
}

func GetTargetsByAnnotation(annName Name) []*Target {
    targetMapRWMutex.RLock()
    defer targetMapRWMutex.RUnlock()

    var result []*Target
    if targets, ok := targetsByAnnotation[annName]; ok {
        for _, target := range targets {
            result = append(result, target.Clone())
        }
    }

    return result
}

func GetAnnotationByID(id ID) *Annotation {
    targetMapRWMutex.RLock()
    defer targetMapRWMutex.RUnlock()

    return annotationMap[id].Clone()
}
