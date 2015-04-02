package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Password string
}

func (u *User) string() *User {
	fmt.Printf("Username:%s Password:%s \n", u.Username, u.Password)
	return u
}

func (u *User) Usernames() {
	fmt.Printf("我的名字是:%s \n", u.Username)

}

func main() {
	user := &User{Username: "吴赐有", Password: "密码"}
	v := reflect.ValueOf(user)
	elem := v.Elem()
	username := elem.FieldByName("Username")
	username.SetString("程序修改的名字")
	password := elem.FieldByName("Password")
	password.SetString("程序修改的密码 ")
	// v.Elem()
	elem.FieldByNameFunc(func(funname string) bool {
		fmt.Println(funname + "fewfew")
		return true
	})
	elem.MethodByName("Usernames").Call([]reflect.Value{})
	user.string()
	// fmt.Println(v.Kind())

	// fmt.Println("fewfewfdsfs")
}
