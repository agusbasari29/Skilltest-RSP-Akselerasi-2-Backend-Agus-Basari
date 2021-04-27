package seeders

import (
	"math/rand"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/bxcodec/faker/v3"
)

type TransactionsSeeders struct {
	StatusPayment string `faker:"oneof: passed, failed"`
}

func TransactionsSeedersUp(number int) {
	seeder := TransactionsSeeders{}
	trx := entity.Transaction{}
	creator, _ := userRepo.GetUserByRole(entity.Users{Role: "creator"})
	participant, _ := userRepo.GetUserByRole(entity.Users{Role: "participant"})
	events := eventRepo.GetEventByStatus(entity.Event{Status: "release"})
	var payment entity.StatusPayment
	for i := 0; i < number; i++ {
		j := rand.Intn(len(creator))
		k := rand.Intn(len(participant))
		l := rand.Intn(len(events))
		m := rand.Intn(3)
		err := faker.FakeData(&seeder)
		if err != nil {
			panic(err)
		}
		switch m {
		case 1:
			payment = "passed"
		case 2:
			payment = "failed"
		}
		trx.StatusPayment = payment
		trx.Amount = events[l].Price
		trx.EventId = int(events[l].ID)
		trx.CreatorId = int(creator[j].ID)
		trx.ParticipantId = int(participant[k].ID)
		trx.CreatedAt = time.Now()
		trxRepo.InsertTransaction(trx)
	}
}
