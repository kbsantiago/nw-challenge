package util

import "testing"
import util "github.com/kbsantiago/nw-challenge/src/util"

const path = "../src/dao/dataset/"

func TestCsvReaderWhenLoadCatalogFile(t *testing.T) {
	filename := "q1_catalog.csv"
	absolutePath := path + filename
	hasHeader := true
	in := make(chan string)
	
	go util.CsvReader(absolutePath, hasHeader, in)
	
	if in == nil {
		t.Errorf("FAIL")
	}
}

func TestCsvReaderWhenLoadClientDataFile(t *testing.T) {
	filename := "q2_clientData.csv"
	absolutePath := path + filename
	hasHeader := true
	in := make(chan string)
	
	go util.CsvReader(absolutePath, hasHeader, in)
	
	if in == nil {
		t.Errorf("FAIL")
	}
}