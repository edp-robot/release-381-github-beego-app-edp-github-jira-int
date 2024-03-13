package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Data["json"] = "Hello, World!"
	c.ServeJSON()
}

func getMessage() string {
	return "Hello, World!"
}
