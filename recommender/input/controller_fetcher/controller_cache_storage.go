package controller_fetcher

import (
	autoscalingapi "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sync"
	"time"
)

type scaleCacheKey struct {
	namespace     string
	groupResource schema.GroupResource
	name          string
}

type scaleCacheEntry struct {
	refreshAfter time.Time
	deleteAfter  time.Time
	resource     *autoscalingapi.Scale
	err          error
}

// Cache for responses to get queries on controllers. Thread safe.
// Usage:
// - `Get` cached response. If there is one use it, otherwise make query and
// - `Insert` the response you got into the cache.
// When you create a `controllerCacheStorage` you should start two go routines:
// - One for refreshing cache entries, which calls `GetKeysToRefresh` then for
//   each key makes query to the API server and calls `Refresh` to update
//  content of the cache.
// - Second for removing stale entries which periodically calls `RemoveExpired`
// Each entry is refreshed after duration
// `validityTime` * (1 + `jitterFactor`)
// passes and is removed if there are no reads for it for more than `lifeTime`.
//
// Sometimes refreshing might take longer than refreshAfter (for example when
// VPA is starting in a big cluster and tries to fetch all controllers). To
// handle such situation lifeTime should be longer than refreshAfter so the main
// VPA loop can do its work quickly, using the cached information (instead of
// getting stuck on refreshing the cache).
// TODO(jbartosik): Add a way to detect when we don't refresh cache frequently
// enough.
// TODO(jbartosik): Add a way to learn how long we keep entries around so we can
// decide if / how we want to optimize entry refreshes.
type controllerCacheStorage struct {
	cache        map[scaleCacheKey]scaleCacheEntry
	mux          sync.Mutex
	validityTime time.Duration
	jitterFactor float64
	lifeTime     time.Duration
}
