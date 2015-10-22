package controller

import (
	"log"
)

type Admin struct {
	Key     string
	Token   string
	Version int
}

func (a *Admin) Get() {
	log.Println("调用 admin.Get func")

}
func (a *Admin) Set() {
	log.Println("调用 admin.Set func")
}
