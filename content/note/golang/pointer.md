---
title: "pointer.md"
date: 1919-08-10T11:45:14Z
draft: false
categories: ["golang"]
---



---

## Pointer

```go
package main

import (
	"jd_workout_golang/app/models"
	"jd_workout_golang/lib/database"
	db "jd_workout_golang/lib/database"
	"jd_workout_golang/lib/file"
	"jd_workout_golang/lib/redis"
)

func init() {
	file.LoadConfigAndEnv()
	database.InitDatabase()
	redis.InitRedis()

}

func main() {
	println("test go")
	var user models.User
	db.Connection.Where("id = ?", 1).Last(&user)
	println(user.Username)
	userToBeLastUser(&user)
	println(user.Username)
}

// both in/out function, can't print out id = 2
func userToBeLastUser_v1(user *models.User) {
	db.Connection.Where("id = ?", 2).Last(&user)
	println("in function,  user name is " + user.Username)
}

// can only print out id = 2 in function
func userToBeLastUser_v2(user *models.User) {
	var newUser models.User
	db.Connection.Where("id = ?", 2).Last(&newUser)
	user = &newUser
	println("in function,  user name is " + user.Username)
}

// can print out id = 2 in/out function
func userToBeLastUser(user *models.User) {
	var newUser models.User
	db.Connection.Where("id = ?", 2).Last(&newUser)
	*user = newUser
	println("In function, user name is " + (*user).Username)
}
```



---

