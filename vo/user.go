package vo

import (
	"github.com/jinzhu/gorm"
	"github.com/senseoki/iris_ex/entity"
)

// User is ...
type User struct {
	RDBTX *gorm.DB

	User *entity.User
}
