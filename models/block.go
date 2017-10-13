package models

type Block struct {
	Index        int
	Timestamp    int64
	Data         Data
	PreviousHash []byte
	Hash         []byte
}
