package idgen

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Generator struct {
	UserID     int32
	ProviderId int
}

var RoundSerialNumber int64
var TransactionSerialNumber int64
var mx sync.Mutex

func New(UserID int32, providerId int) *Generator {
	return &Generator{
		UserID:     UserID,
		ProviderId: providerId,
	}
}
func (gen *Generator) IdGenRound(providerID int64, serverID int) (int64, error) {
	mx.Lock()
	defer mx.Unlock()
	current_time := time.Now()
	Identifier := "1"
	serverNumber := fmt.Sprint(serverID)
	userID := fmt.Sprintf("%02d", providerID)
	time := current_time.Format("060201")
	serialNumber := fmt.Sprintf("%09d", RoundSerialNumber)
	format := Identifier + serverNumber + userID + time + serialNumber
	fmt.Println(format)
	id, err := strconv.ParseInt(format, 10, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println(id)
	RoundSerialNumber += 1
	return id, nil
}
func (gen *Generator) IdGenTransaction(providerID int64, serverID int) (int64, error) {
	mx.Lock()
	defer mx.Unlock()
	current_time := time.Now()
	Identifier := "2"
	serverNumber := fmt.Sprint(serverID)
	userID := fmt.Sprintf("%02d", providerID)
	time := current_time.Format("060201")
	serialNumber := fmt.Sprintf("%09d", RoundSerialNumber)
	format := Identifier + serverNumber + userID + time + serialNumber
	fmt.Println(format)
	id, err := strconv.ParseInt(format, 10, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println(id)
	RoundSerialNumber += 1
	return id, nil
}
