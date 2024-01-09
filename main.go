package main

import (
	"example.com/v2/router"
)

func main() {

	r := router.Router()

	//defer recover panic
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常：", err)
	//	}
	//}()
	r.Run(":1212")
}
