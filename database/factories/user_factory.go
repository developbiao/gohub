package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/user"
	"gohub/pkg/helpers"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	// Set unique value
	faker.SetGenerateUniqueValues(false)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}
	return objs
}
