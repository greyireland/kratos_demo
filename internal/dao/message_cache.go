package dao

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/greyireland/kratos-demo/internal/model"
)

var (
	msgKey = "m_%v_%v"
)

func getKey(pre string, arg ...interface{}) string {
	return fmt.Sprintf(pre, arg)
}
func (d *Dao) AddMessageCache(ctx context.Context, msg model.MessageInfo) error {
	buf, _ := json.Marshal(msg)
	conn := d.redis.Get(ctx)
	defer conn.Close()
	_, err := conn.Do("SET", getKey(msgKey, msg.UID, msg.PeerUID), string(buf))
	return err
}
func (d *Dao) GetMessageCache(ctx context.Context, uid, peerUid int64) (*model.MessageInfo, error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	buf, err := redis.Bytes(conn.Do("GET", getKey(msgKey, uid, peerUid)))
	if err != nil {
		return nil, err
	}
	var ret model.MessageInfo
	json.Unmarshal(buf, &ret)
	return &ret, err
}
