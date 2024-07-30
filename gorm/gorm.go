package gorm

import (
	"time"

	"github.com/nextgis/go-sessions"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

type Store interface {
	sessions.Store
}

func NewStore(d *gorm.DB, expiredSessionCleanup bool, maxKeyLength int, keyPairs ...[]byte) Store {
	s := gormstore.New(d, keyPairs...)
	if expiredSessionCleanup {
		quit := make(chan struct{})
		go s.PeriodicCleanup(1*time.Hour, quit)
	}
	s.MaxLength(maxKeyLength)
	return &store{s}
}

type store struct {
	*gormstore.Store
}

func (s *store) Options(options sessions.Options) {
	s.Store.SessionOpts = options.ToGorillaOptions()
}
