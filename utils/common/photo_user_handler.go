package common

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/gin-gonic/gin"
)

func FileVerifyHandler(c *gin.Context) (dto.VerifyUser, error) {
	user := c.PostForm("user")
	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		return dto.VerifyUser{}, err
	}
	defer file.Close()

	var ext = [3]string{".jpg", ".jpeg", ".png"}
	valid := checkExtention(header.Filename, ext)
	if !valid {
		return dto.VerifyUser{}, fmt.Errorf("ekstensi file harus %s", ext)
	}

	fileName := fmt.Sprintf("%v_user_%s", rand.New(rand.NewSource(time.Now().UnixNano())).Int(), filepath.Ext(header.Filename))
	fileLocation := filepath.Join("uploads", fileName)

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return dto.VerifyUser{}, err
	}

	if err := c.SaveUploadedFile(header, fileLocation); err != nil {
		return dto.VerifyUser{}, err
	}

	var userCredential dto.VerifyUser

	json.Unmarshal([]byte(user), &userCredential)
	userCredential.Photo = fileLocation

	return userCredential, nil

}

func checkExtention(filename string, ext [3]string) bool {
	e := filepath.Ext(filename)
	for _, a := range ext {
		if a == e {
			return true
		}
	}
	return false
}
