package explodedDNS

import (
	"github.com/activecm/rita/parser/parsetypes"
)

// Repository for uconn collection
type Repository interface {
	CreateIndexes(targetDB string) error
	Upsert(explodedDNS *parsetypes.ExplodedDNS, targetDB string) error
}