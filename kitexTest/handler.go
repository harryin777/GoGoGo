package main

import (
	"context"
	"math/rand"
	"test1/kitexTest/kitex_gen/snow"
)

// SnowImpl implements the last service interface defined in the IDL.
type SnowImpl struct{}

// Wanted implements the SnowImpl interface.
func (s *SnowImpl) Wanted(ctx context.Context, req *snow.SnowRequest) (resp *snow.SnowResponse, err error) {
	// TODO: Your code here...
	r := rand.New(rand.NewSource(req.Wanted))
	resp = &snow.SnowResponse{
		SnowCount: r.Int63(),
	}

	return
}
