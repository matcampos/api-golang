package userRoute

import (
	controller "../../controllers/user"
)

func Routes(router *Routes) {
	router.HandleFunc("/users", controller.create).Methods("POST")
}
