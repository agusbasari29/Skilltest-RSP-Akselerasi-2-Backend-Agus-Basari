package seeders

import (
	"fmt"
	"math/rand"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

type UserSeeders struct {
	Username string `faker:"username"`
	Fullname string `faker:"name"`
	Email    string `faker:"email"`
}

func UsersSeedersUp(number int) {
	seeder := UserSeeders{}
	users := entity.Users{}
	for i := 0; i < number; i++ {
		var userRole entity.UserRole
		j := rand.Intn(3)
		err := faker.FakeData(&seeder)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		switch j {
		case 1:
			userRole = "admin"
		case 2:
			userRole = "creator"
		case 3:
			userRole = "participant"
		}
		password, _ := bcrypt.GenerateFromPassword([]byte("sangatrahasia"), bcrypt.DefaultCost)
		users.Role = userRole
		users.Username = seeder.Username
		users.Fullname = seeder.Fullname
		users.Email = seeder.Email
		users.Password = string(password)
		userRepo.InsertUser(users)
	}
}
