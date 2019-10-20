package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}
type MessageInfo struct {
	ID        int64  `json:"id"`
	SessionID string `form:"session_id" json:"session_id"`
	UID       int64  `form:"uid" json:"uid" validate:"required"`
	PeerUID   int64  `form:"peer_uid" json:"peer_uid" validate:"required"`
	Message   string `form:"message" json:"message"`
}
