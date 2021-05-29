package server

import (
	"github.com/li-zeyuan/micro/micro.common.api/pb/profile"

	"context"
	"log"
)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
}

func (s *ProfileServer) Upsert(ctx context.Context, in *profile.UpsertReq) (*profile.UpsertResp, error) {
	log.Printf("Received: %v", in.GetUid())
	return &profile.UpsertResp{}, nil
}
