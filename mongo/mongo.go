package mongo

import (
	"github.com/ONSdigital/dp-code-list-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DATABASE = "codelists"
const CODE_LISTS = "codelists"
const CODES = "codes"

//MongoDataStore
type MongoDataStore struct {
	session *mgo.Session
}

func CreateMongoDataStore(url string) (*MongoDataStore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDataStore{session: session}, nil
}

// GetCodes
func (mds *MongoDataStore) GetCodes(codeListID string) (*models.CodeResults, error) {
	s := mds.session.Clone()
	defer s.Clone()
	iter := s.DB(DATABASE).C(CODES).Find(nil).Iter()
	defer iter.Close()
	var codes []models.Code
	err := iter.All(&codes)
	if err != nil {
		return nil, err
	}
	return &models.CodeResults{Items: codes}, nil
}

// GetCode
func (mds *MongoDataStore) GetCode(codeListID, CodeID string) (*models.Code, error) {
	s := mds.session.Clone()
	defer s.Clone()
	var result models.Code
	err := s.DB(DATABASE).C(CODES).Find(bson.M{"_id": codeListID, "code": CodeID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCodeLists
func (mds *MongoDataStore) GetCodeLists() (*models.CodeListResults, error) {
	s := mds.session.Clone()
	defer s.Clone()
	iter := s.DB(DATABASE).C(CODE_LISTS).Find(nil).Iter()
	defer iter.Close()
	var codeLists []models.CodeList
	err := iter.All(&codeLists)
	if err != nil {
		return nil, err
	}
	return &models.CodeListResults{Items: codeLists}, nil
}

// GetCodeList
func (mds *MongoDataStore) GetCodeList(codeListID string) (*models.CodeList, error) {
	s := mds.session.Clone()
	defer s.Clone()
	var result models.CodeList
	err := s.DB(DATABASE).C(CODE_LISTS).Find(bson.M{"_id": codeListID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
