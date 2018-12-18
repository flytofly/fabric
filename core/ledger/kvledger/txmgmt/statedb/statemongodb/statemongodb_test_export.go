package statemongodb

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
)

type TestDBEnv struct {
	t          testing.TB
	DBProvider statedb.VersionedDBProvider
}

func NewTestDBEnv(t testing.TB) *TestDBEnv {
	t.Logf("Creating new TestDBEnv")
	versionedDBProvider, _ := NewVersionedDBProvider()
	return &TestDBEnv{t: t, DBProvider: versionedDBProvider}
}

func (env *TestDBEnv) Cleanup(dbName string) {
	versionedDBProvider, _ := NewVersionedDBProvider()
	versionedDBProvider.session.DB(dbName).DropDatabase()
	versionedDBProvider.session.Close()
}