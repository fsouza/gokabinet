package kc

import (
	"strings"
	"testing"
)

func TestShouldBeAbleToSetAndGetIntegerRecord(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.SetInt("albums", 10)
	if value, _ := db.GetInt("albums"); value != 10 {
		t.Errorf("Should be able to set and get integer values")
	}
}

func TestSetIntShouldOverrideANonNumericRecords(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.Set("albums", "White Album")
	db.SetInt("albums", 10)
	if value, _ := db.GetInt("albums"); value != 10 {
		t.Errorf("SetInt should override non-numeric records")
	}
}

func TestSetIntShouldOverrideNumericRecords(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.SetInt("albums", 10)
	db.SetInt("albums", 11)
	if value, _ := db.GetInt("albums"); value != 11 {
		t.Errorf("SetInt should override numeric records")
	}
}

func TestSetIntShouldNotWorkInReadOnlyMode(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	db.Close()

	db, _ = Open(filepath, READ)
	defer db.Close()

	if err := db.SetInt("albums", 10); err == nil || !strings.Contains(err.Error(), "read-only mode") {
		t.Errorf("SetInt should not work in read-only mode")
	}
}

func TestGetIntShoulReturnAnErrorIfTheGivenKeyDoesNotExit(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	if _, err := db.GetInt("albums"); err == nil || !strings.Contains(err.Error(), "no record") {
		t.Errorf("GetInt: should return an error when there is not a record with the given key")
	}
}

func TestGetIntShouldRetornAnErrorIfTheGivenKeyIsANonNumericRecord(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.Set("name", "Mariah")

	if _, err := db.GetInt("name"); err == nil || !strings.Contains(err.Error(), "non-numeric record") {
		t.Errorf("GetInt should return an error when the given key refers to a non-numeric record")
	}
}

func TestIncrementShouldIncrementTheValueOfANumericRecord(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.SetInt("people", 1)
	db.Increment("people", 10)

	if v, _ := db.GetInt("people"); v != 11 {
		t.Errorf("Should increment 1 in 10 and get 11, but got %d", v)
	}
}

func TestIncrementShouldCrateTheRecordWhenItDoesNotExist(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.Increment("people", 100)

	if v, _ := db.GetInt("people"); v != 100 {
		t.Errorf("Should create the record with value 100, but the value is %d", v)
	}
}

func TestIncrementShouldReturnAnErrorIfTheIncrementedValueIsANonNumericRecord(t *testing.T) {
	filepath := "/tmp/musicians.kch"
	defer Remove(filepath)

	db, _ := Open(filepath, WRITE)
	defer db.Close()

	db.Set("name", "Francisco Souza")
	if err := db.Increment("name", 1); err == nil || !strings.Contains(err.Error(), "non-numeric record") {
		t.Errorf("Should return an error message when trying to increment a non-numeric record %s")
	}
}
