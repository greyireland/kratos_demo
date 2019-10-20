package service

import (
	"context"

	"github.com/greyireland/kratos-demo/internal/model"
)

func (s *Service) AddMessage(ctx context.Context, msg model.MessageInfo) (int64, error) {
	return s.dao.AddMessage(ctx, msg)
}
func (s *Service) GetMessages(ctx context.Context, uid, peerUid int64) ([]model.MessageInfo, error) {
	return s.dao.GetMessages(ctx, uid, peerUid)
}
