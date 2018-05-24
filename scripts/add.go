package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/google/uuid"
)

// Code for a single dimensions type
// Other models are taken from the package but the json tags
// for a code are different on import than output (the ID field)
type Code struct {
	ID    string           `json:"_id"`
	Code  string           `json:"code"    bson:"code"`
	Label string           `json:"label"   bson:"label"`
	Links models.CodeLinks `json:"links"   bson:"links"`
}

// CodeList differs from models.CodeList so json (output) of ID field has underscore: _id
type CodeList struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Links models.CodeListLink `json:"links"`
}

func main() {
	input := flag.String("input", "", "input file name")
	flag.Parse()

	if len(*input) == 0 {
		log.Fatal("-input argument required")
	}

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal("Could not open input file", err)
	}
	defer closeFile(f)

	csvr := csv.NewReader(f)
	recs, err := csvr.ReadAll()
	if err != nil {
		log.Fatal("Failed to read csv file", err)
	}

	header := recs[0]
	recs = recs[1:]

	listID := createCodeList(header)
	log.Print("Added row to codelists.json")

	createCodes(recs, listID)

}

func createCodes(recs [][]string, listID string) {
	filename := "codes/" + listID + ".json"

	if _, err := os.Stat(filename); os.IsExist(err) {
		if err = os.Rename(filename, filename+".old"); err != nil {
			log.Fatal("File already exists and cannot be renamed: ", err)
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to open code file: ", err)
	}
	defer closeFile(f)

	var wg sync.WaitGroup

	// read marshalled lines, write them to the file
	jsonLineChan := make(chan []byte)
	go func(f *os.File) {
		for msg := range jsonLineChan {
			if _, err := f.WriteString(string(msg) + "\n"); err != nil {
				log.Fatal("failed to append code to file: ", err)
			}
			wg.Done()
		}
	}(f)

	// marshal lines, send them to the file-writer above
	for _, v := range recs {
		id := uuid.New().String()
		code := Code{
			ID:    id,
			Code:  v[0],
			Label: v[1],
			Links: models.CodeLinks{
				CodeList: models.Link{
					ID:   listID,
					Href: "http://localhost:22400/code-lists/" + listID,
				},
				Self: models.Link{
					ID:   v[0],
					Href: "http://localhost:22400/code-lists/" + listID + "/codes/" + v[0],
				},
			},
		}

		wg.Add(1)
		go func(code Code) {
			b, err := json.Marshal(code)
			if err != nil {
				log.Fatal("cannot marshal json for code: ", err)
			}

			jsonLineChan <- b
		}(code)
	}

	wg.Wait()
	close(jsonLineChan)

	log.Printf("Added codes file: [%s]", filename)
}

func createCodeList(header []string) string {
	listID := uuid.New().String()
	headerListID := header[0]

	if headerListID != "" {
		listID = headerListID
		listID = strings.Replace(listID, " ", "", -1)
	}

	cl := CodeList{
		ID:   listID,
		Name: header[1],
		Links: models.CodeListLink{
			Self: models.Link{
				ID:   listID,
				Href: "http://localhost:22400/code-lists/" + listID,
			},
			Codes: models.Link{
				Href: "http://localhost:22400/code-lists/" + listID + "/codes",
			},
		},
	}

	b, err := json.Marshal(cl)
	if err != nil {
		log.Fatal("cannot marshal json for code-list line", err)
	}

	if err := appendToFile(b, "codelists.json"); err != nil {
		log.Fatal("failed to append code-list to file", err)
	}

	return listID
}

func appendToFile(b []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer closeFile(f)

	s := string(b) + "\n"

	if _, err := f.WriteString(s); err != nil {
		return err
	}

	return nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal("failed to close file: ", err)
	}
}
