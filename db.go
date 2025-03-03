package main

import (
	"encoding/json"
	"log"

	c "github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

const collectionName = "ArticlesCollection"

var db *c.DB

func InitDB() error {
	var err error
	var exist bool
	db, err = c.Open("clover-db")
	if err != nil {
		return err
	}
	exist, err = db.HasCollection(collectionName)
	if err != nil {
		return err
	}
	if !exist {
		err = db.CreateCollection(collectionName)
		if err != nil {
			return err
		}
	}
	log.Printf("Initialized CloverDB")
	return nil
}

func AddToDB(article Article) error {
	doc := document.NewDocumentOf(article)
	_, err := db.InsertOne(collectionName, doc)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetAllArticles() ([]Article, error) {
	docs, err := db.FindAll(query.NewQuery(collectionName))
	if err != nil {
		return nil, err
	}
	var result []Article
	for _, doc := range docs {
		article := &Article{}
		doc.Unmarshal(article)
		result = append(result, *article)
	}

	return result, nil
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
