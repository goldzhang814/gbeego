package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"hello/models"
	//"github.com/astaxie/beego/toolbox"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	//  beego.AppConfig.String("mysqluser")
	// beego.AppConfig.String("mysqlpass")
	// beego.AppConfig.String("mysqlurls")
	// beego.AppConfig.String("mysqldb")

	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}
func main() {
	// tk1 := toolbox.NewTask("tk1", "* * * * * *", func() error { fmt.Println("tk1"); return nil })
	// toolbox.AddTask("tk1", tk1)
	// toolbox.StartTask()
	// defer toolbox.StopTask()

	beego.Run()
	 //o := orm.NewOrm()
	// o.Using("default")
	// p := new(models.Profile)
	// p.Age = 30

	// user := new(models.User)
	// users := []models.User{
	// 	{Name: "slene"},
	// 	{Name: "astaxie"},
	// 	{Name: "unknown"},
	// }
	// //user.Profile = p
	// user.Name = "gold"

	// fmt.Println(o.Insert(user))
	// //fmt.Println(o.Insert(user))
	// i := len(users)
	// successNums, ok := o.InsertMulti(i, users)
	// if ok!=nil{
	// 	fmt.Println("insert failed!")
	// } else{
	// 	fmt.Println(successNums)
	// }

	// type User struct {
	// 	Id       int	`orm:"column(Id)"`
	// 	Name     string `orm:"column(Name)"`
	// }
	
	// var users []*User
	// num,err := o.Raw("SELECT Id,Name FROM user").QueryRows(&users)
	// if err == nil  {
	// 	fmt.Println("user nums: ", num)
	// }
	// fmt.Println(users)


	// for _,v := range users {	


	// 	 fmt.Printf("users data is:: %d\n",v.Id)
	// 	 fmt.Println("users data is::",v.Name)
	// 	}
}





