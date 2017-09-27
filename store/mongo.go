package store

import (
	"github.com/ONSdigital/dp-code-list-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DATABASE = "codelists"
const CODE_LISTS = "codelists"
const CODES = "codes"

// MongoDataStore contains collections of codes and code lists
type MongoDataStore struct {
	Session *mgo.Session
}

// CreateMongoDataStore which can fetch and store data about codes and code lists
func CreateMongoDataStore(url string) (*MongoDataStore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDataStore{Session: session}, nil
}

// GetCodes returns all codes which are stored in mongodb
func (mds *MongoDataStore) GetCodes(codeListID string) (*models.CodeResults, error) {
	s := mds.Session.Clone()
	defer s.Clone()
	iter := s.DB(DATABASE).C(CODES).Find(bson.M{"links.code_list.id": codeListID}).Iter()
	defer iter.Close()
	var codes []models.Code
	err := iter.All(&codes)
	if err != nil {
		return nil, err
	}
	return &models.CodeResults{Items: codes}, nil
}

// GetCode returns a single code from a code list
func (mds *MongoDataStore) GetCode(codeListID, CodeID string) (*models.Code, error) {
	s := mds.Session.Clone()
	defer s.Clone()
	var result models.Code
	err := s.DB(DATABASE).C(CODES).Find(bson.M{"links.code_list.id": codeListID, "code": CodeID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCodeLists returns all code lists which are stored in mongodb
func (mds *MongoDataStore) GetCodeLists() (*models.CodeListResults, error) {
	s := mds.Session.Clone()
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

// GetCodeList returns a single code list or a not found error
func (mds *MongoDataStore) GetCodeList(codeListID string) (*models.CodeList, error) {
	s := mds.Session.Clone()
	defer s.Clone()
	var result models.CodeList
	err := s.DB(DATABASE).C(CODE_LISTS).Find(bson.M{"_id": codeListID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
