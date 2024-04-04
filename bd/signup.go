package bd

import (
	"fmt"

	"github.com/jorojas/gambituser/models"
	"github.com/jorojas/gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Begin Save()")

	err := DbConnect()

	if err != nil {
		return err
	}
	defer Db.Close()

	sqlQuery := fmt.Sprintf("INSERT INTO users(User_Email, User_UUID, User_DateAdd) VALUES ('%s', '%s', '%s')",
		sig.UserEmail, sig.UserUUID, tools.MySQLDate())

	fmt.Println(sqlQuery)

	_, err = Db.Exec(sqlQuery)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp record was successful")

	return nil
}
