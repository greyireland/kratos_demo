package dao

import (
	"context"

	"github.com/greyireland/kratos-demo/internal/model"

	"github.com/bilibili/kratos/pkg/log"
)

var (
	addMessage  = "insert into `message` (session_id,uid,peer_uid,message) values (?,?,?,?)"
	getMessages = "select id,session_id,uid,peer_uid,message from `message` where uid = ? and peer_uid = ? limit 10"
)

func (d *Dao) AddMessage(ctx context.Context, msg model.MessageInfo) (int64, error) {
	ret, err := d.db.Exec(ctx, addMessage, msg.SessionID, msg.UID, msg.PeerUID, msg.Message)
	if err != nil {
		log.Error("add message error(%s)", err)
		return 0, err
	}
	return ret.LastInsertId()
}
func (d *Dao) GetMessages(ctx context.Context, uid, peerUid int64) ([]model.MessageInfo, error) {
	rows, err := d.db.Query(ctx, getMessages, uid, peerUid)
	if err != nil {
		log.Error("query  error(%v)", err)
		return nil, err
	}
	defer rows.Close()
	var result []model.MessageInfo
	for rows.Next() {
		var msg model.MessageInfo
		if err = rows.Scan(&msg.ID, &msg.SessionID, &msg.UID, &msg.PeerUID, &msg.Message); err != nil {
			log.Error("scan demo log error(%v)", err)
			return nil, err
		}
		result = append(result, msg)
	}
	return result, nil
}
