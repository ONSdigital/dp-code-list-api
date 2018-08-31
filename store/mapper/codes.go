package mapper

import (
	"fmt"
	"strconv"

	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
)

var errCodeNotFound = errors.New("code not found")

const (
	codesURI = "/code-lists/%s/editions/%s/codes"
	codeURI  = "/code-lists/%s/editions/%s/codes/%s"
)

//Codes returns a dpbolt.ResultMapper mapper which converts dpbolt.Result to models.CodeResults
func (m *Mapper) Codes(results *models.CodeResults, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		if len(r.Data) == 0 {
			return errCodeNotFound
		}

		var err error
		var node graph.Node
		if node, err = getNode(r.Data[0]); err != nil {
			return err
		}

		id := node.NodeIdentity

		var codeVal string
		if codeVal, err = getStringProperty("value", node.Properties); err != nil {
			return err
		}

		var rel graph.Relationship
		if rel, err = getRelationship(r.Data[1]); err != nil {
			return err
		}

		var codeLabel string
		if codeLabel, err = getStringProperty("label", rel.Properties); err != nil {
			return err
		}

		results.Items = append(results.Items, models.Code{
			ID:    strconv.FormatInt(id, 10),
			Code:  codeVal,
			Label: codeLabel,
			Links: models.CodeLinks{
				Self:     *models.CreateLink("", fmt.Sprintf(codeURI, codeListID, edition, codeVal), m.Host),
				Datasets: *models.CreateLink("", fmt.Sprintf(datasetsURI, codeListID, edition, codeVal), m.Host),
				CodeList: *models.CreateLink("", fmt.Sprintf(codeListURI, codeListID), m.Host),
			},
		})
		return nil
	}
}

//Code returns a dpbolt.ResultMapper which converts a dpbolt.Result to models.Code
func (m *Mapper) Code(codeModel *models.Code, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		if len(r.Data) == 0 {
			return errCodeNotFound
		}

		var err error
		var node graph.Node
		if node, err = getNode(r.Data[0]); err != nil {
			return err
		}

		id := node.NodeIdentity

		var codeVal string
		if codeVal, err = getStringProperty("value", node.Properties); err != nil {
			return err
		}

		var rel graph.Relationship
		if rel, err = getRelationship(r.Data[1]); err != nil {
			return err
		}

		var codeLabel string
		if codeLabel, err = getStringProperty("label", rel.Properties); err != nil {
			return err
		}

		codeModel.ID = strconv.FormatInt(id, 10)
		codeModel.Code = codeVal
		codeModel.Label = codeLabel
		codeModel.Links = models.CodeLinks{
			Self:     *models.CreateLink("", fmt.Sprintf(codeURI, codeListID, edition, codeVal), m.Host),
			Datasets: *models.CreateLink("", fmt.Sprintf(datasetsURI, codeListID, edition, codeVal), m.Host),
			CodeList: *models.CreateLink("", fmt.Sprintf(codeListURI, codeListID), m.Host),
		}
		return nil
	}
}
