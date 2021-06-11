package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"hello/models"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

type DynamicsController struct {
	//beego.Controller
	CommonsController
}

type res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserInfo struct {
	ID       uint64
	Username string
	Nbf      int64
	Iat      int64
	Exp      int64
}

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

func (this *DynamicsController) Get() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	// profile := new(models.Profile)
	// profile.Age = 30

	user := new(models.User)
	//user.Profile = profile
	user.Name = "slene"
	f, _ := o.Insert(user)
	fmt.Println(f)

	var users []*User
	num, err := o.Raw("SELECT id,name FROM user").QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num)
	}

	r := new(res)
	r.Code = 1001
	r.Msg = "success"

	//r.Data = f
	r.Data = users

	beego.Debug("this is debug")

	this.Data["json"] = &r
	this.ServeJSON()
	return

	// rj,err := json.Marshal(r)
	// if err != nil {
	// 	r.Msg = "failed"
	// 	c.Ctx.WriteString(string(rj[:]))
	// }
	// c.Ctx.WriteString(string(rj[:]))
}

func (this *DynamicsController) Post() {
	this.Ctx.WriteString("hello2")
}

func (this *DynamicsController) GenToken() {
	u := new(UserInfo)
	u.ID = 99
	u.Username = "gold"

	token, _ := CreateToken(u)
	this.Ctx.WriteString(token)

}

func (this *DynamicsController) GetAnnouceList() {
	jsoninfo := this.GetString("token")
	if jsoninfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		return
	}

	f, err := ParseToken(jsoninfo)
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("failed")
		return
	}

	fmt.Println(f)

	this.Ctx.WriteString("ok")
	return
}

func (this *DynamicsController) ParseFormData() {
	u := user{}
	if err := this.ParseForm(&u); err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("failed")
		return
	}

	fmt.Println(u)

	this.Ctx.WriteString("ok")
	return
}

func (this *DynamicsController) PostBody() {
	var ob res
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	this.Ctx.WriteString("success")
	return
}

func (this *DynamicsController) PostUpload() {
	f, h, err := this.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		this.SaveToFile("uploadname", "G:/www/"+h.Filename)
	}

	this.Ctx.WriteString("success")
	return
}

func (this *DynamicsController) UpdateFood() {
	this.Ctx.WriteString("put success")
	return
}

func (this *DynamicsController) DeleteFood() {
	this.Ctx.WriteString("delete success")
	return
}
