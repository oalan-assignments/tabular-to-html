package parser

import (
	"bufio"
	"encoding/csv"
	"errors"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Workbook struct {
	Header   []string
	Contacts []Contact
}

type Contact struct {
	Name        string
	Address     string
	Postcode    string
	Phone       string
	CreditLimit string
	Birthday    string
}

type WorkbookParser interface {
	parse(reader io.Reader) (Workbook, error)
}

type Csv struct{}
type Prn struct{}
type Unknown struct{}

func ParseWorkbook(file string) (Workbook, error) {
	workBookFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return Workbook{}, errors.New("file could not be read")
	}
	defer workBookFile.Close()
	readerWithDecoder := charmap.ISO8859_1.NewDecoder().Reader(workBookFile)
	fileExtension := filepath.Ext(workBookFile.Name())
	parser := getParser(fileExtension)
	return parser.parse(readerWithDecoder)
}

func getParser(fileExtension string) WorkbookParser {
	switch fileExtension {
	case ".prn":
		return &Prn{}
	case ".csv":
		return &Csv{}
	default:
		return &Unknown{}
	}
}

func (Unknown) parse(reader io.Reader) (Workbook, error) {
	msg := "file is available in directory listing but not a recognized format"
	log.Println(msg)
	return Workbook{}, errors.New(msg)
}

func (p *Csv) parse(readerWithDecoder io.Reader) (Workbook, error) {
	reader := csv.NewReader(bufio.NewReader(readerWithDecoder))
	header, error := reader.Read()
	if error != nil {
		log.Println(error)
		return Workbook{}, errors.New("could not read the header")
	}
	contacts, error := readContacts(reader)
	return Workbook{
		Header:   header,
		Contacts: contacts,
	}, error
}

func readContacts(reader *csv.Reader) ([]Contact, error) {
	var contacts []Contact
	for {
		contactInfoAsArray, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return []Contact{}, errors.New("could not read contacts")

		}
		contacts = append(contacts, getContactFrom(contactInfoAsArray))
	}
	return contacts, nil
}

func (p *Prn) parse(readerWithDecoder io.Reader) (Workbook, error) {
	var header []string
	var contacts []Contact
	scanner := bufio.NewScanner(readerWithDecoder)
	if scanner.Scan() {
		headerLine := scanner.Text()
		header = getColumnValuesFrom(headerLine)
	}
	for scanner.Scan() {
		line := scanner.Text()
		contactInfoAsArray := getColumnValuesFrom(line)
		contact := getContactFrom(contactInfoAsArray)
		contacts = append(contacts, contact)
	}
	return Workbook{Header: header, Contacts: contacts}, nil
}

func getColumnValuesFrom(line string) []string {
	var columns []string
	runes := []rune(line)
	columns = append(columns, strings.TrimSpace(string(runes[0:16])))  //name
	columns = append(columns, strings.TrimSpace(string(runes[16:38]))) //address
	columns = append(columns, strings.TrimSpace(string(runes[38:47]))) //postcode
	columns = append(columns, strings.TrimSpace(string(runes[47:61]))) //phone
	columns = append(columns, strings.TrimSpace(string(runes[61:74]))) //credit limit
	columns = append(columns, strings.TrimSpace(string(runes[74:])))   //birthday
	return columns
}

func getContactFrom(contactInfoAsArray []string) Contact {
	return Contact{
		Name:        contactInfoAsArray[0],
		Address:     contactInfoAsArray[1],
		Postcode:    contactInfoAsArray[2],
		Phone:       contactInfoAsArray[3],
		CreditLimit: contactInfoAsArray[4],
		Birthday:    contactInfoAsArray[5],
	}
}
