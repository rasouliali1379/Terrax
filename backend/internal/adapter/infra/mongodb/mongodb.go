package mongodb

import (
	"context"
	"fmt"
	"github.com/rasouliali1379/terrax/config"
	"github.com/rasouliali1379/terrax/internal/pkg/terminator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func Init() (*mongo.Database, terminator.Func) {

	slog.Info("Initializing MongoDB database")

	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(fmt.Sprintf("mongodb://%s:%d/", config.C().MongoDB.Host, config.C().MongoDB.Port)).
			SetAuth(options.Credential{Username: config.C().MongoDB.User, Password: config.C().MongoDB.Password}),
	)
	if err != nil {
		panic(err)
	}

	defer slog.Info("MongoDB database initialized successfully")
	return client.Database(config.C().MongoDB.Database), func() error {
		return client.Disconnect(context.TODO())
	}
}
