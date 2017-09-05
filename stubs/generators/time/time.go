package time

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
)

type codeList struct {
	ID    string               `json:"id"`
	Name  string               `json:"name"`
	Codes map[string]codeEntry `json:"codes"`
}

type codeEntry struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Label string `json:"label"`
}

func main() {
	b, _ := ioutil.ReadFile("time.txt")
	s := string(b)
	i := strings.Split(s, "\n")
	o := make(map[string]codeEntry)
	for _, v := range i {
		id := uuid.New().String()
		o[id] = codeEntry{
			ID:    id,
			Code:  v,
			Label: v,
		}
	}

	cl := codeList{
		ID:    uuid.New().String(),
		Name:  "Time",
		Codes: o,
	}

	b, _ = json.MarshalIndent(&cl, "", " ")
	fmt.Println(string(b))
}
