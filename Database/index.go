package Database

import (
	"context"
	"github.com/rapido/db"
)

type PrismaDB struct {
	Client *db.PrismaClient
	Context context.Context
}


var PrismaClient *PrismaDB

func InitializeDB() (*PrismaDB, error){
	client := db.NewClient()

	if err := client.Connect(); err != nil {
		return nil, err
	}

	PrismaClient = &PrismaDB{
		Client: client,
		Context: context.Background(),
	}

	PrismaClient.Client.Connect()

	return PrismaClient, nil
}
