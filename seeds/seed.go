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
