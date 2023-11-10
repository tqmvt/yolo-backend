package seeds

import (
	"github.com/go-faker/faker/v4"

	"saas/models"
	"saas/repository"
)

func SeedUsers() error {
	length := 20

	for i := 0; i < length; i++ {
		user := models.User{
			Id:            i + 1,
			FirstName:     faker.FirstName(),
			LastName:      faker.LastName(),
			Email:         faker.Email(),
			WalletAddress: "cosmos84afvm0yfs27a00hnr85064r69lpg46zjaz5udp",
		}

		repository.Save(&user)
	}

	return nil
}

func SeedRounds() error {
	length := 30

	for i := 0; i < length; i++ {
		round := models.Round{
			Id:          i + 1,
			Round:       i + 1,
			MaxDuration: 300,
			MaxPlayers:  50,
			Winner:      "cosmos84afvm0yfs27a00hnr85064r69lpg46zjaz5udp",
		}

		repository.Save(&round)
	}

	return nil
}

func SeedPlayers() error {
	length := 40

	for i := 0; i < length; i++ {
		fakeRound, err := faker.RandomInt(1, 30, 1)
		if err != nil {
			return err
		}

		fakeUser, err := faker.RandomInt(1, 20, 1)
		if err != nil {
			return err
		}

		fakeAmount, err := faker.RandomInt(1, 100, 1)
		if err != nil {
			return err
		}

		player := models.Player{
			Id:            i + 1,
			RoundId:       fakeRound[0],
			UserId:        fakeUser[0],
			WalletAddress: "cosmos84afvm0yfs27a00hnr85064r69lpg46zjaz5udp",
			BidAmount:     float32(fakeAmount[0]) / 1000,
			Color:         faker.UUIDDigit()[:6],
		}

		repository.Save(&player)
	}

	return nil
}
