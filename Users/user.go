package users

import (
	"fmt"
	"git/models"
	"time"
)

func SaveUser() {
	User := new(models.User)
	User.AddUser(1, "andrea", time.Now(), true)
	fmt.Println(User)
}
