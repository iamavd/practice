package mongodb

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Init(t *testing.T) {
	tt := []struct {
		name   string
		cfg    *Config
		expErr error
	}{
		{
			name:   "create options error",
			cfg:    &Config{},
			expErr: errors.New("error parsing uri: must have at least 1 host"),
		},
		{
			name: "connect error",
			cfg: &Config{
				Address: "10.0.0.10:8080",
			},
			expErr: errors.New("context deadline exceeded"),
		},
		{
			name: "all ok",
			cfg: &Config{
				Address:  "127.0.0.1:27017",
				Database: "MyDataBase",
			},
		},
		{
			name: "create options error",
			cfg: &Config{
				Username: "some username",
				Password: "some password",
			},
			expErr: errors.New("error parsing uri: must have at least 1 host"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			db, err := Init(ctx, tc.cfg)
			if tc.expErr != nil {
				assert.EqualError(t, tc.expErr, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
			}
		})
	}
}
