package main

import (
	"fmt"
	"github.com/wuciyou/godoc/reflect/controller"
	"log"
	"reflect"
)

type ControllerInterface interface {
	Sayhello(str, name string)
}

type User struct {
	Username string
	Password string
}

func (u *User) String(str string) {
	fmt.Printf("Username:%s Password:%s \n", u.Username, u.Password)

}
func (u *User) Sayhello(str, name string) {
	fmt.Printf("Sayhello内容:%s \n", str)
}

func (u *User) Usernames(str string) {
	fmt.Printf("我的名字是:%s %s \n", u.Username, str)
}

func main() {
	user := &User{Username: "吴赐有", Password: "密码"}

	// t1()
	// t2()
	// t3(user)
	t4(user)

}
func ffewfew(name, age string) {
	fmt.Println(name)
}
func t4(user ControllerInterface) {
	// userReflectVal := reflect.ValueOf(user)

	SayhelloType := reflect.TypeOf(ffewfew)
	fmt.Println(SayhelloType.NumIn())
	fmt.Println(SayhelloType.In(0))
	// userReflectVal.Type().Method(1).Func.Call([]reflect.Value{reflect.ValueOf("aaaaa")})
}

func t3(user ControllerInterface) {

	reflectVal := reflect.ValueOf(user)
	rt := reflectVal.Type()
	rv := reflect.Indirect(reflectVal)

	fmt.Println(reflectVal.Type())
	fmt.Println(rv.Type())
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		// m.Func.Call([]reflect.Value{reflect.ValueOf("小小人物")})
		fmt.Printf("name:%s, PkgPath:%s Kind:%s \n", m.Name, m.PkgPath, m.Func.Kind())
	}
}

func t2() {
	user := &User{Username: "吴赐有", Password: "密码"}
	admin := &controller.Admin{Key: "key", Token: "Token", Version: 1}
	v := reflect.ValueOf(user)
	t := reflect.TypeOf(admin)
	fmt.Printf("Pointer:%d\n", v.Pointer())
	log.Println(t)
	log.Printf("pkgpath:%s \n", t.Elem().PkgPath())
}
func t1() {
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
	par := v.MethodByName("Usernames").Type().In(0)
	// fmt.Println(par.FieldByNameFunc(func(parstr string) bool {
	// 	fmt.Println(parstr)
	// 	return true
	// }))
	fmt.Println(par.PkgPath())
	// m.Call(nil)
	// m.Elem().FieldByNameFunc(func(str string) bool {
	// 	fmt.Println(str + "______")
	// 	return true
	// })
	user.String("few")
	// fmt.Println(v.Kind())

	// fmt.Println("fewfewfdsfs")
}
