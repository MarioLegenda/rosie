package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

type user struct {
	UUID     string `faker:"uuid_digit"`
	Name     string `faker:"name"`
	LastName string `faker:"last_name"`
	Email    string `faker:"email"`

	Min int
	Max int

	urls []string
}

func createUsers(args arguments) []user {
	users := make([]user, 0)
	usrNum := args.users
	if usrNum == 0 {
		usrNum = 5
	}

	for i := 0; i < usrNum; i++ {
		u := user{}
		err := faker.FakeData(&u)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to fake user data with faker: %s", err.Error()))
		}

		u.Min = args.intervalMin
		u.Max = args.intervalMax
		u.urls = args.links

		users = append(users, u)
	}

	return users
}
