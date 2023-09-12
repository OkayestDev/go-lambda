package repositories

import (
	"golambda/src/structs"
)

var UserRepo = Repository[structs.User]("users")
