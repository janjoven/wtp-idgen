package idgen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Generator struct {
	TransactionType int
	ProviderId      int
}

func New(transactionType int, providerId int) *Generator {
	return &Generator{
		TransactionType: transactionType,
		ProviderId:      providerId,
	}
}
func (gen *Generator) IdGen() (int64, error) {
	var id string
	timeunix := fmt.Sprint(time.Now().Unix())
	transactionType := fmt.Sprint(gen.TransactionType)
	providerId := fmt.Sprintf("%03d", gen.ProviderId)
	// dateNow := time.Now()
	// dateformated := fmt.Sprint(dateNow.Format("20060201"))
	randLastNumber := randGenerator()
	id = transactionType + providerId + timeunix + randLastNumber
	i, err := strconv.ParseInt(id, 10, 64)
	return i, err
}
func randGenerator() string {
	b := 99999
	a := 1
	rand.Seed(time.Now().UnixNano())
	n := a + rand.Intn(b-a+1)
	return fmt.Sprintf("%05d", n)
}
