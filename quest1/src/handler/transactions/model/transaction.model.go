package model

type TransactionRequest struct {
	SenderID    string `json:"sender_id"`
	ReceiverID  string `json:"receiver_id"`
	Amount      int64  `json:"amount"`
	ReferenceID string `json:"reference_id"`
}
