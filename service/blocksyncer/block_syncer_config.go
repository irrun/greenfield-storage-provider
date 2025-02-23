package blocksyncer

import (
	"os"

	"github.com/bnb-chain/greenfield-storage-provider/model/errors"
)

type Config struct {
	Modules        []string
	Dsn            string
	DsnSwitched    string
	RecreateTables bool
	Workers        uint
	EnableDualDB   bool
}

func getDBConfigFromEnv(dsn string) (string, error) {
	dsnVal, ok := os.LookupEnv(dsn)
	if !ok {
		return "", errors.ErrDSNNotSet
	}
	return dsnVal, nil
}
