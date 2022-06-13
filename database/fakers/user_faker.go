package fakers

import(
	"github.com/rizkyprawirap/Toko/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID: 			uuid.New().String(),
		Addresses		nil,
		FirstName: 		fakers.FirstName(),
		LastName: 		fakers.LastName(),
		Email: 			fakers.Email(),
		Password: 		"password", //password
		RememberToken: 	"",
		CreatedAt: 		time.Time{},
		UpdatedAt: 		time.Time{},
		DeletedAt: 		gorm.DeletedAt{},
	}
}