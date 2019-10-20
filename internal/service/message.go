package service

import (
	"context"

	"github.com/greyireland/kratos-demo/internal/model"
)

func (s *Service) AddMessage(ctx context.Context, msg model.MessageInfo) (int64, error) {
	data, err := s.dao.AddMessage(ctx, msg)
	if err != nil {
		return 0, err
	}
	msg.ID = data
	err = s.dao.AddMessageCache(ctx, msg)
	if err != nil {
		return 0, err
	}
	return data, err

}
func (s *Service) GetMessages(ctx context.Context, uid, peerUid int64) ([]model.MessageInfo, error) {
	msg, err := s.dao.GetMessageCache(ctx, uid, peerUid)
	if msg != nil && err == nil {
		return []model.MessageInfo{*msg}, nil
	}
	return s.dao.GetMessages(ctx, uid, peerUid)
}
