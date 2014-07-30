package rethink

import "testing"

var db DB

func TestConnect(t *testing.T) {
	var err error
	db, err = Connect("localhost:28015", "test")

	if err != nil {
		t.Errorf("expected nil error but got %s", err)
	}

	if len(db.collections) != 0 {
		t.Errorf("new DB should not have any collections.")
	}
}

func TestNewCollection(t *testing.T) {
	TestCollection := db.NewCollection("test")

	if TestCollection.db != &db {
		t.Errorf("expected TestCollection.db to equal global DB")
	}

	if len(db.collections) != 1 {
		t.Errorf("new DB should have one collection.")
	}

	if db.collections[0] != TestCollection {
		t.Errorf("collections should have TestCollection")
	}

}
