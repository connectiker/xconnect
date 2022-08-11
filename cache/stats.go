package cache

import "sync"

type Statistics interface {
	SetQuery(isHit bool)
	SetMutation(isHit bool)
	SetItems(totalItems int64)
	SetSize(bytes int64)
	GetStats() StatsReader
}

type StatsReader interface {
	Size() int64
	Items() int64
	QueryHits() int64
	QueryMiss() int64
	MutationHits() int64
	MutationMiss() int64
}

type Stats struct {
	lock *sync.RWMutex
	// Size of entries in bytes.
	size int64
	// Items represents the total number of entries stored in the cache.
	items int64

	// describes the type of operation you’re trying to do.
	// There is either query or mutation
	queries   *Query
	mutations *Mutation
}

// Query is any type of operation that does request for information on data.
type Query struct {
	lock *sync.RWMutex
	// queryHits represents the number of successfully query type operation
	queryHits int64
	// queryMisses represents the number of error query type operation
	queryMisses int64
}

// Mutation is any type of operation that does modify the data, such as insertion, updating, deleting or other forms
// of data manipulation.
type Mutation struct {
	lock *sync.RWMutex
	// mutationHits represents the number of successfully mutation type operation
	mutationHits int64
	// mutationMisses represents the number of error mutation type operation
	mutationMisses int64
}

// NewMetrics returns a new instance of stats
func NewMetrics() *Stats {
	return &Stats{
		lock: new(sync.RWMutex),
		queries: &Query{
			lock: new(sync.RWMutex),
		},
		mutations: &Mutation{
			lock: new(sync.RWMutex),
		},
	}
}

// GetStats returns some statistics about the current codec
func (s *Stats) GetStats() StatsReader {
	stats := NewMetrics()

	s.lock.RLock()
	stats.items = s.items
	stats.size = s.size
	s.lock.RUnlock()

	s.queries.lock.RLock()
	stats.queries = s.queries
	s.queries.lock.RUnlock()

	s.queries.lock.RLock()
	stats.mutations = s.mutations
	s.queries.lock.RUnlock()

	return stats
}

func (s *Stats) Size() int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.size
}

func (s *Stats) Items() int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.items
}

func (s *Stats) QueryHits() int64 {
	s.queries.lock.RLock()
	defer s.queries.lock.RUnlock()

	return s.queries.queryHits
}

func (s *Stats) QueryMiss() int64 {
	s.queries.lock.RLock()
	defer s.queries.lock.RUnlock()

	return s.queries.queryMisses
}

func (s *Stats) MutationHits() int64 {
	s.mutations.lock.RLock()
	defer s.mutations.lock.RUnlock()

	return s.mutations.mutationHits
}

func (s *Stats) MutationMiss() int64 {
	s.mutations.lock.RLock()
	defer s.mutations.lock.RUnlock()

	return s.mutations.mutationMisses
}

func (s *Stats) SetItems(totalItems int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = totalItems
}

func (s *Stats) SetSize(bytes int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.size = bytes
}

func (s *Stats) SetQuery(isHit bool) {
	s.queries.lock.Lock()
	defer s.queries.lock.Unlock()

	if isHit {
		s.queries.queryHits++
		return
	}

	s.queries.queryMisses++
}

func (s *Stats) SetMutation(isHit bool) {
	s.mutations.lock.Lock()
	defer s.mutations.lock.Unlock()

	if isHit {
		s.mutations.mutationHits++
		return
	}

	s.mutations.mutationMisses++
}

// IncStats is a helper method to increment statistics
func (s *Stats) IncStats(isQuery bool, isHit bool) {
	if isQuery {
		s.SetQuery(isHit)
	}
	s.SetMutation(isHit)
}
