package server

import (
	"testing"

	"context"
	"github.com/karux/liletestone"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &liletestone.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
