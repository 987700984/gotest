package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"os"
	"time"
)

type ImageController struct {
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (ImageController) Upload(c *gin.Context) {
	file, _ := c.FormFile("img")
	//将文件保存至本项目根目录中
	h := md5.New()
	io.WriteString(h, generateRandomString(16))
	sum := h.Sum(nil)
	idstr := hex.EncodeToString(sum[:])
	timeStr := time.Now().Format("2006-01-02")
	ImageHost := "./upload/" + timeStr + "/"
	if _, err := os.Stat(ImageHost); os.IsNotExist(err) {
		err = os.MkdirAll(ImageHost, 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir '%s'error: %s", ImageHost, err))
		}
	}
	destImage := ImageHost + idstr + ".jpg"
	err := c.SaveUploadedFile(file, destImage)
	if err != nil {
		fmt.Println("save err:")
		fmt.Println(err)
		Common{}.ReturnError(c, 1, "是u吧、、、")
	} else {

		Common{}.ReturnSuccess(c, 0, "登录成功", destImage, 1)
	}

}
