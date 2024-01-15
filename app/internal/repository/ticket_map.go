package repository

import (
	"app/app/internal"
	"app/app/internal/loader"
	"context"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
//
//	func NewRepositoryTicketMap(dbFile string, lastId int) *RepositoryTicketMap {
//		return &RepositoryTicketMap{
//			dbFile: dbFile,
//			lastId: lastId,
//		}
//	}
func NewRepositoryTicketMap(loader *loader.LoaderTicketCSV) *RepositoryTicketMap {

	db, err := loader.Load()
	if err != nil {
		db = make(map[int]internal.TicketAttributes)
	}

	// get the last id
	lastId := 0
	for k := range db {
		if k > lastId {
			lastId = k
		}
	}

	return &RepositoryTicketMap{
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}
