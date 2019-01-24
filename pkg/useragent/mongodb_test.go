package useragent

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/activecm/rita/database"
	"github.com/activecm/rita/parser/parsetypes"
	"github.com/globalsign/mgo/dbtest"
)

// Server holds the dbtest DBServer
var Server dbtest.DBServer

// Set the test database
var testTargetDB = "tmp_test_db"

var testRepo Repository

var testUserAgent = &parsetypes.UserAgent{
	UserAgent: "Debian APT-HTTP/1.3 (1.2.24)",
	TimesUsed: 123,
}

func TestCreateIndexes(t *testing.T) {
	err := testRepo.CreateIndexes(testTargetDB)
	if err != nil {
		t.Errorf("Error creating useragent indexes")
	}
}

func TestUpsert(t *testing.T) {
	err := testRepo.Upsert(testUserAgent, testTargetDB)
	if err != nil {
		t.Errorf("Error upserting to useragent collection")
	}
}

// TestMain wraps all tests with the needed initialized mock DB and fixtures
func TestMain(m *testing.M) {
	// Store temporary databases files in a temporary directory
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// Set the main session variable to the temporary MongoDB instance
	ssn := Server.Session()

	db := database.DB{Session: ssn}

	testRepo = NewMongoRepository(&db)

	// Run the test suite
	retCode := m.Run()

	// Clean up test database and session
	ssn.DB(testTargetDB).DropDatabase()
	ssn.Close()

	// Shut down the temporary server and removes data on disk.
	Server.Stop()

	// call with result of m.Run()
	os.Exit(retCode)
}
