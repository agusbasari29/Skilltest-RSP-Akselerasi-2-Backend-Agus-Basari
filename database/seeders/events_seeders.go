package seeders

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/bxcodec/faker/v3"
)

type EventSeeders struct {
	TitleEvent  string  `faker:"sentence"`
	LinkWebinar string  `faker:"url"`
	Description string  `faker:"paragraph"`
	Banner      string  `faker:"word"`
	Price       float32 `faker:"amount"`
	Quantity    int     `faker:"oneof: 150, 200, 250, 300"`
	Status      string  `faker:"oneof: release, draft"`
}

type EventTimeFaker struct {
	Time1 int64 `faker:"unix_time"`
	Time2 int64 `faker:"unix_time"`
	Time3 int64 `faker:"unix_time"`
	Time4 int64 `faker:"unix_time"`
}

func EventsSeedersUp(number int) {
	seeder := EventSeeders{}
	event := entity.Event{}
	times := EventTimeFaker{}
	result := userRepo.GetUserRole("creator")
	for i := 0; i < number; i++ {
		var status entity.EventStatus
		j := rand.Intn(len(result))
		k := rand.Intn(3)
		err := faker.FakeData(&seeder)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", seeder)

		err = faker.FakeData(&times)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", times)
		switch k {
		case 1:
			status = "release"
		case 2:
			status = "draft"
		}

		event.TitleEvent = seeder.TitleEvent
		event.LinkWebinar = seeder.LinkWebinar
		event.Description = seeder.Description
		event.Banner = seeder.Banner
		event.Price = seeder.Price
		event.Quantity = seeder.Quantity
		event.Status = status
		event.CreatorId = int(result[j].ID)
		event.EventStartDate = convertTime(times.Time1)
		event.EventEndDate = convertTime(times.Time2)
		event.CampaignStartDate = convertTime(times.Time3)
		event.CampaignEndDate = convertTime(times.Time4)
		event.CreatedAt = time.Now()
		eventRepo.InsertEvent(event)
	}
}

func convertTime(unix int64) time.Time {
	i, err := strconv.ParseInt(strconv.Itoa(int(unix)), 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}
