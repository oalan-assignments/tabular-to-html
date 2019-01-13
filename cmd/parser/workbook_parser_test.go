package parser

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	expectedHeader = []string{"Name", "Address", "Postcode", "Phone", "Credit Limit", "Birthday"}
)

func TestWorkbookCsvParser(t *testing.T) {
	var tests = []struct {
		input    int //row number
		expected Contact
	}{
		{0,
			Contact{"Johnson, John",
				"Voorstraat 32",
				"3122gg",
				"020 3849381",
				"10000",
				"01/01/1987"}},
		{1,
			Contact{"Anderson, Paul",
				"Dorpsplein 3A",
				"4532 AA",
				"030 3458986",
				"109093",
				"03/12/1965"}},
		{2,
			Contact{"Wicket, Steve",
				"Mendelssohnstraat 54d",
				"3423 ba",
				"0313-398475",
				"934",
				"03/06/1964"}},
		{3,
			Contact{"Benetar, Pat",
				"Driehoog 3zwart",
				"2340 CC",
				"06-28938945",
				"54",
				"04/09/1964"}},
		{4,
			Contact{"Gibson, Mal",
				"Vredenburg 21",
				"3209 DD",
				"06-48958986",
				"54.5",
				"09/11/1978"}},
		{5,
			Contact{"Friendly, User",
				"Sint Jansstraat 32",
				"4220 EE",
				"0885-291029",
				"63.6",
				"10/08/1980"}},
		{6,
			Contact{"Smith, John",
				"Børkestraße 32",
				"87823",
				"+44 728 889838",
				"9898.3",
				"20/09/1999"}},
	}

	workbook, _ := ParseWorkbook("../../Workbook2.csv")
	assert(tests, workbook, t)
}

func TestWorkbookPrnParser(t *testing.T) {
	var tests = []struct {
		input    int //row number
		expected Contact
	}{
		{0,
			Contact{"Johnson, John",
				"Voorstraat 32",
				"3122gg",
				"020 3849381",
				"1000000",
				"19870101"}},
		{1,
			Contact{"Anderson, Paul",
				"Dorpsplein 3A",
				"4532 AA",
				"030 3458986",
				"10909300",
				"19651203"}},
		{2,
			Contact{"Wicket, Steve",
				"Mendelssohnstraat 54d",
				"3423 ba",
				"0313-398475",
				"93400",
				"19640603"}},
		{3,
			Contact{"Benetar, Pat",
				"Driehoog 3zwart",
				"2340 CC",
				"06-28938945",
				"5400",
				"19640904"}},
		{4,
			Contact{"Gibson, Mal",
				"Vredenburg 21",
				"3209 DD",
				"06-48958986",
				"5450",
				"19781109"}},
		{5,
			Contact{"Friendly, User",
				"Sint Jansstraat 32",
				"4220 EE",
				"0885-291029",
				"6360",
				"19800810"}},
		{6,
			Contact{"Smith, John",
				"Børkestraße 32",
				"87823",
				"+44 728 889838",
				"989830",
				"19990920"}},
	}

	workbook, _ := ParseWorkbook("../../Workbook2.prn")
	assert(tests, workbook, t)
}

func assert(tests []struct {
	input    int
	expected Contact
}, workbook Workbook, t *testing.T) {
	for _, test := range tests {
		if output := workbook.Contacts[test.input]; output != test.expected {
			t.Error("Test failed {} inputted, {} expected, received: {}",
				test.input, test.expected, output)
		}
	}
	actualHeader := workbook.Header

	if !reflect.DeepEqual(actualHeader, expectedHeader) {
		t.Error("Test failed expected: {}, received: {}",
			expectedHeader, actualHeader)
	}
}

func TestFileIsNotThere(t *testing.T) {
	workbook, e := ParseWorkbook("test.test")
	if !reflect.DeepEqual(Workbook{}, workbook) {
		t.Error("Workbook should be empty")
	}
	if e == nil {
		t.Error("Error message should be: no such file or directory")
	}
	fmt.Println(e.Error())
}

func TestEmptyCsv(t *testing.T) {
	_, e := ParseWorkbook("../../test/Empty.csv")
	if e == nil {
		t.Error("Error should not be nil")
	}
}

func TestNoContactsCsv(t *testing.T) {
	workbook, e := ParseWorkbook("../../test/NoContacts.csv")
	if workbook.Contacts != nil {
		t.Error("Contacts should be empty")
	}
	if !reflect.DeepEqual(expectedHeader, workbook.Header) {
		t.Error("Header should have been read")
	}
	if e != nil {
		t.Error("Error should be nil")
	}
}

func TestMalformedCsv(t *testing.T) {
	workbook, e := ParseWorkbook("../../test/Malformed.csv")
	if len(workbook.Contacts) != 0 {
		t.Error("Contacts should be empty")
	}
	if !reflect.DeepEqual(expectedHeader, workbook.Header) {
		t.Error("Header should have been read")
	}
	if e == nil {
		t.Error("There must be error: could not read contacts")
	}
	fmt.Println(e.Error())
}
