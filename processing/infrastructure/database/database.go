package database

import (
	"path/filepath"

	"github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/config"
	"github.com/cryptix-network/cryptix-graph-inspector/processing/infrastructure/logging"
	"github.com/cryptix-network/cryptixd/infrastructure/db/database"
	"github.com/cryptix-network/cryptixd/infrastructure/db/database/ldb"
)

const (
	databaseDirectoryName = "database"
	levelDBCacheSizeMiB   = 256
)

var (
	log = logging.Logger()
)

func Open(config *config.Config) (database.Database, error) {
	databaseDirectory := filepath.Join(config.AppDir, databaseDirectoryName)
	log.Infof("Loading database from '%s'", databaseDirectory)
	return ldb.NewLevelDB(databaseDirectory, levelDBCacheSizeMiB)
}
