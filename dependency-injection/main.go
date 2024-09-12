package main

import "fmt"

// Database is a service
type Database interface {
	GetData() string
}

type MockPostgreDB struct{}

func (md *MockPostgreDB) GetData() string {
	return "Some data from pg mock"
}

type MockMongoDB struct{}

func (md *MockMongoDB) GetData() string {
	return "Some data from mongo mock"
}

type Client struct {
	db Database
}

// NewClient is an injector
func NewClient(db Database) *Client {
	return &Client{db: db}
}

func main() {
	pg := &MockPostgreDB{}
	mongo := &MockMongoDB{}

	// Injector provides services to client
	pgClient := NewClient(pg)
	fmt.Println(pgClient.db.GetData())

	// Injector provides services to client
	mongoClient := NewClient(mongo)
	fmt.Println(mongoClient.db.GetData())
}
