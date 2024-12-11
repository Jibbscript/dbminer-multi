package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	CCNum  string  `bson:"ccnum"`
	Date   string  `bson:"date"`
	Amount float64 `bson:"amount"`
	Cvv    string  `bson:"cvv"`
	Exp    string  `bson:"exp"`
}

func main() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri: "mongodb://127.0.0.1:27017",
	})
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err = client.Close(ctx); err != nil {
			log.Panicln(err)
		}
	}()

	db := client.Database("test")
	coll := db.Collection("transactions")

	results := make([]Transaction, 0)
	filter := bson.M{}
	if err := coll.Find(ctx, filter).All(&results); err != nil {
		log.Panicln(err)
	}
	for _, txn := range results {
		fmt.Println(txn.CCNum, txn.Date, txn.Amount, txn.Cvv, txn.Exp)
	}
}
