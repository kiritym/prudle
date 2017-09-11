package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type HttpReq struct {
	ApiPath     string
	ContentType string
	CharSet     string
	StatusCode  string
	Payload     string
}

func connection() *bolt.DB {
	db, err := bolt.Open("http.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("HttpBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db
}

func (hr *HttpReq) save(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("HttpBucket"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(hr)
		if err != nil {
			return err
		}
		return b.Put([]byte(hr.ApiPath), encoded)
	})
	return err
}

func find(db *bolt.DB, apiPath string) (HttpReq, error) {
	var hr HttpReq

	return hr, db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("HttpBucket"))
		u := b.Get([]byte(apiPath))
		if err := json.Unmarshal(u, &hr); err != nil {
			//panic(err)
			fmt.Println("error")
		}
		return nil
	})
}

func findAll(db *bolt.DB) ([]string, error) {
	hrList := make([]string, 0)
	return hrList, db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("HttpBucket"))
		fmt.Println("hi")
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s\n", string(k))
			fmt.Printf("value=%s\n", string(v))
			hrList = append(hrList, string(k))
			return nil
		})
		return nil
	})
}
