package mapper

import (
	"github.com/pkg/errors"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"strings"
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

const (
	codeListURI = "/code-lists/%s"
)

func CodeLists(codeLists *models.CodeListResults, prefix string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		var label string
		for _, v := range r.Data[0].([]interface{}) {
			s := v.(string)
			if strings.Contains(s, prefix) {
				label = strings.Replace(s, prefix, "", -1)
				break
			}
		}

		codeLists.Items = append(codeLists.Items, models.CodeList{
			Links: models.CodeListLink{
				Self: &models.Link{
					Href: fmt.Sprintf(codeListURI, label),
					ID:   label,
				},
				Editions: &models.Link{
					Href: fmt.Sprintf(editionsURI, label),
				},
			},
		})
		return nil
	}
}

func CodeList(codeList *models.CodeList, id string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		if len(r.Data) == 0 {
			return datastore.NOT_FOUND
		}

		codeList.Links = models.CodeListLink{
			Self: &models.Link{
				Href: fmt.Sprintf(codeListURI, id),
				ID:   id,
			},
			Editions: &models.Link{
				Href: fmt.Sprintf(editionsURI, id),
			},
		}
		return nil
	}
}

// EditionCount is a result mapper for gett
func CodeListCount() (*int64, dpbolt.ResultMapper) {
	var count int64
	return &count, func(r *dpbolt.Result) error {
		var ok bool
		count, ok = r.Data[0].(int64)
		if !ok {
			return errors.New("extract row result error: failed to cast result to int64")
		}
		return nil
	}
}
