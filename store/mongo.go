package store

import (
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-code-list-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDataStore contains collections of codes and code lists
type MongoDataStore struct {
	Session             *mgo.Session
	db                  string
	codelistsCollection string
	codesCollection     string
}

// CreateMongoDataStore which can fetch and store data about codes and code lists
func CreateMongoDataStore(cfg config.MongoConfig) (*MongoDataStore, error) {
	session, err := mgo.Dial(cfg.BindAddr)
	if err != nil {
		return nil, err
	}
	return &MongoDataStore{
		Session:             session,
		db:                  cfg.Database,
		codelistsCollection: cfg.CodelistsCollection,
		codesCollection:     cfg.CodesCollection,
	}, nil
}

// GetCodes returns all codes which are stored in mongodb
func (s *MongoDataStore) GetCodes(codeListID string) (*models.CodeResults, error) {
	session := s.Session.Clone()
	defer session.Close()
	iter := session.DB(s.db).C(s.codesCollection).Find(bson.M{"links.code_list.id": codeListID}).Iter()
	defer iter.Close()
	var codes []models.Code
	err := iter.All(&codes)
	if err != nil {
		return nil, err
	}
	return &models.CodeResults{Items: codes}, nil
}

// GetCode returns a single code from a code list
func (s *MongoDataStore) GetCode(codeListID, codeID string) (*models.Code, error) {
	session := s.Session.Clone()
	defer session.Close()
	var result models.Code
	err := session.DB(s.db).C(s.codesCollection).Find(bson.M{"links.code_list.id": codeListID, "code": codeID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCodeLists returns all code lists which are stored in mongodb
func (s *MongoDataStore) GetCodeLists() (*models.CodeListResults, error) {
	session := s.Session.Clone()
	defer session.Close()
	iter := session.DB(s.db).C(s.codelistsCollection).Find(nil).Iter()
	defer iter.Close()
	var codeLists []models.CodeList
	err := iter.All(&codeLists)
	if err != nil {
		return nil, err
	}
	return &models.CodeListResults{Items: codeLists}, nil
}

// GetCodeList returns a single code list or a not found error
func (s *MongoDataStore) GetCodeList(codeListID string) (*models.CodeList, error) {
	session := s.Session.Clone()
	defer session.Close()
	var result models.CodeList
	err := session.DB(s.db).C(s.codelistsCollection).Find(bson.M{"_id": codeListID}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
