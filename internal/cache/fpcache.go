package cache

import "sync"

type fingerPrintsCache struct {
	Fingerprints []string
	sync.Mutex
}

var fpc *fingerPrintsCache

func init() {
	fpc = newFingerPrintsCache()
}

func newFingerPrintsCache() *fingerPrintsCache {
	return &fingerPrintsCache{}
}

func GetFingerPrintsCache() []string {
	fpc.Lock()
	defer fpc.Unlock()
	return fpc.Fingerprints
}

func AddFingerPrintsCache(fingerprints string) {
	fpc.Lock()
	defer fpc.Unlock()
	fpc.Fingerprints = append(fpc.Fingerprints, fingerprints)
}

func ClearFingerPrintsCache() {
	fpc.Lock()
	defer fpc.Unlock()
	fpc.Fingerprints = make([]string, 0)
}
