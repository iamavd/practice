package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Username string
	Password string
	Address  string
	Database string
}

func Init(ctx context.Context, cfg *Config) (*mongo.Database, error) {
	opts, err := createOptions(cfg)
	if err != nil {
		return nil, err
	}
	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err = checkConnection(ctx, c); err != nil {
		return nil, err
	}
	return c.Database(cfg.Database), nil
}

func createOptions(cfg *Config) (*options.ClientOptions, error) {
	url := "mongodb://"
	if cfg.Username != "" {
		url += cfg.Username
	}
	if cfg.Password != "" {
		url += ":" + cfg.Password + "@"
	}
	url += cfg.Address
	opts := options.Client().ApplyURI(url)
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	return opts, nil
}

func checkConnection(ctx context.Context, c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return c.Ping(ctx, readpref.Primary())
}
