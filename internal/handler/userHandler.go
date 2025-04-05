package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gotemplateproject/internal/db"
	"gotemplateproject/internal/model"
)

// Get all users
func GetUsers(c *gin.Context) {
	users1, err := db.ExecuteQuery("select userName, userEmail , deleted, userId from user where deleted = 2")

	if err != nil {
		fmt.Println("Unable to fetch users [Error: %s]", err)
	}
	var users []model.User
	for _, record := range users1 {
		// user := User{
		// 	UserId:    int(record["userId"]),
		// 	UserName:  record["userName"].string(),
		// 	UserEmail: record["userEmail"].string(),
		// 	Deleted:   int(record["deleted"]),
		// }
		user := model.User{}
		userID, ok := record["userId"].(int64) // Usually, numbers come as float64 from JSON
		if !ok {
			fmt.Println("Error: userId is not a valid number")
		}

		intUserID := int(userID)
		userName, ok := record["userName"].(string)
		if !ok {
			fmt.Println("Error: userName is not a valid string")
		}
		userEmail, ok := record["userEmail"].(string)
		if !ok {
			fmt.Println("Error: userEmail is not a valid string")
		}

		userdeleted, ok := record["deleted"].(int64)
		if !ok {
			fmt.Println("Error: userEmail is not a valid string")
		}
		userDeleted := int(userdeleted)
		user.UserId = intUserID
		user.UserName = userName
		user.UserEmail = userEmail
		user.Deleted = userDeleted

		users = append(users, user)
	}

	fmt.Println("UserModel", users)

	c.JSON(200, users)
	// jsondata, err := json.Marshal(users)

	// if err != nil {
	// 	errors.New("Error while marshal data")
	// 	c.JSON(500, gin.H{"error": "Internal server error"})
	// }
	// c.Data(200, "application/json", jsondata)
}
