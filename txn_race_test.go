//go:build memdbrace
// +build memdbrace

package memdb

import "testing"

func TestTxn_Insert_Dupe(t *testing.T) {
	db := testDB(t)
	txn := db.Txn(true)

	obj := testObj()
	err := txn.Insert("main", obj)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	raw, err := txn.First("main", "id", obj.ID)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if raw != obj {
		t.Fatalf("bad: %#v %#v", raw, obj)
	}

	recovered := false

	func() {

		defer func() {
			if err := recover(); err != nil {
				recovered = true
				t.Logf("caught panic: %v", err)
			}
		}()
		err := txn.Insert("main", obj)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
	}()

	if !recovered {
		t.Fatalf("expected panic")
	}
}
