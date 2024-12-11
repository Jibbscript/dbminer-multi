package main

import (
	"context"
	"os"

	"github.com/jibbscript/dbminer-multi/dbminer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMiner struct {
	Host    string
	session *mongo.Client
}

func New(host string) (*MongoMiner, error) {
	m := MongoMiner{Host: host}
	err := m.connect()
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *MongoMiner) connect() error {
	mHostScheme := "mongodb://" + m.Host
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mHostScheme))
	if err != nil {
		return err
	}
	m.session = client
	return nil
}

func (m *MongoMiner) CollectionNames(dbName string) ([]string, error) {
	db := m.session.Database(dbName)

	filter := bson.D{}
	collections, err := db.ListCollectionNames(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func (m *MongoMiner) GetSchema() (*dbminer.Schema, error) {
	var s = new(dbminer.Schema)

	// Create a command to list databases
	dbnames, err := m.session.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for _, dbname := range dbnames {
		db := dbminer.Database{Name: dbname, Tables: []dbminer.Table{}}

		collections, err := m.CollectionNames(dbname)
		if err != nil {
			return nil, err
		}

		for _, collection := range collections {
			table := dbminer.Table{Name: collection, Columns: []string{}}

			// Get a sample document to extract fields
			var doc bson.M
			err := m.session.Database(dbname).Collection(collection).FindOne(context.Background(), bson.M{}).Decode(&doc)
			if err != nil {
				continue // Skip this collection if we can't get a sample document
			}

			// Extract field names from the document
			for fieldName := range doc {
				table.Columns = append(table.Columns, fieldName)
			}

			db.Tables = append(db.Tables, table)
		}
		s.Databases = append(s.Databases, db)
	}
	return s, nil
}

func main() {
	mm, err := New(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := dbminer.Search(mm); err != nil {
		panic(err)
	}
}
