package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"sunspa/common"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if c.Param("id") == "" || clientToken == "" {
			logrus.Errorf("Authorization token was not provided")
			c.JSON(common.NewUnauthorizedResponse("Authorization Token is required"))
			c.Abort()
			return
		}

		//claims := jwt.MapClaims{}
		extractedToken := strings.Split(clientToken, "Bearer ")

		// Verify if the format of the token is correct
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			logrus.Errorf("Incorrect Format of Authn Token")
			c.JSON(common.NewUnauthorizedResponse("Incorrect Format of Authorization Token"))
			c.Abort()
			return
		}

		//foundInBlacklist := IsBlacklisted(extractedToken[1])
		//
		//if foundInBlacklist == true {
		//	logger.Logger.Infof("Found in Blacklist")
		//	c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
		//	c.Abort()
		//	return
		//}
		//
		// Parse the claims
		//var user = &models.User{}
		//_, err := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
		//	var userValid = &models.User{
		//		UserId:   "NV001",
		//		FullName: "Luu Duc Hoang",
		//		Phone:    "0966 936 938",
		//		Role:     models.Role{Id: 1, Name: "admin"},
		//	}
		//	return userValid, nil
		//})
		//
		//if err != nil {
		//	logrus.Errorf("ParseWithClaims Error: %v", err)
		//}
		//
		//if user.UserId != c.Param("id") {
		//	c.JSON(common.NewUnauthorizedResponse("Incorrect Token with User Info"))
		//	c.Abort()
		//	return
		//}

		//
		//if err != nil {
		//	if err == jwt.ErrSignatureInvalid {
		//		logrus.Errorf("Invalid Token Signature")
		//		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token Signature"})
		//		c.Abort()
		//		return
		//	}
		//	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad Request"})
		//	c.Abort()
		//	return
		//}
		//
		//if !parsedToken.Valid {
		//	logrus.Errorf("Invalid Token")
		//	c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}
