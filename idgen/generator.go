package idgen

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

// func New(UserID int32, providerId int) *Generator {
// 	return &Generator{
// 		UserID:     UserID,
// 		ProviderId: providerId,
// 	}
// }
func IdGenRound(providerID int64) (int64, error) {
	mx.Lock()
	defer mx.Unlock()
	if RoundSerialNumber == 0 {
		id, _ := fileRead("dat/roundid.txt")
		RoundSerialNumber = id + 1
	}
	current_time := time.Now()
	Identifier := "1"
	userID := fmt.Sprintf("%03d", providerID)
	time := current_time.Format("060201")
	serialNumber := fmt.Sprintf("%09d", RoundSerialNumber)
	format := Identifier + userID + time + serialNumber
	id, err := strconv.ParseInt(format, 10, 64)
	if err != nil {
		writeRoundID(id)
		return 0, err
	}
	writeRoundID(id)
	RoundSerialNumber += 1
	return id, nil
}
func IdGenTransaction(providerID int64) (int64, error) {
	mx.Lock()
	defer mx.Unlock()
	if TransactionSerialNumber == 0 {
		id, _ := fileRead("dat/transactionid.txt")
		TransactionSerialNumber = id + 1
	}
	current_time := time.Now()
	Identifier := "2"
	userID := fmt.Sprintf("%03d", providerID)
	time := current_time.Format("060201")
	serialNumber := fmt.Sprintf("%09d", TransactionSerialNumber)
	format := Identifier + userID + time + serialNumber
	id, err := strconv.ParseInt(format, 10, 64)
	if err != nil {
		writeTransactionID(id)
		return 0, err
	}
	writeTransactionID(id)
	TransactionSerialNumber += 1
	return id, nil
}
func fileRead(pathname string) (int64, error) {
	content, err := ioutil.ReadFile(pathname)
	if err != nil {
		emptyData := []byte("")
		if err := ioutil.WriteFile(pathname, emptyData, 0777); err != nil {
			log.Println(err)
		}
	}
	stringSplit := strings.Split(string(content), "")
	if len(stringSplit) == 0 {
		return 0, nil
	}
	idString := stringSplit[10:]
	id, err := strconv.ParseInt(strings.Join(idString, ""), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil

}
func writeRoundID(roundID int64) {
	roundIDData := fmt.Sprintf("%v", roundID)
	rountIDByte := []byte(roundIDData)
	if err := ioutil.WriteFile("dat/roundid.txt", rountIDByte, 0777); err != nil {
		log.Println(err)
	}
}
func writeTransactionID(transactionID int64) {
	transactionIDData := fmt.Sprintf("%v", transactionID)
	transactionIDByte := []byte(transactionIDData)
	if err := ioutil.WriteFile("dat/transactionid.txt", transactionIDByte, 0777); err != nil {
		log.Println(err)
	}
}
