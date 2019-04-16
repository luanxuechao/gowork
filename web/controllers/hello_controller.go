package controllers

import (
	"errors"
	"github.com/kataras/iris/mvc"
)
type HelloController struct {}

var helloView = mvc.View{
	Name:"hello/index.html",
	Data:map[string] interface{}{
		"Title":" hello page",
		"MyMessage":"12123131",
	},
}

func (c *HelloController) Get()mvc.Result {
	return helloView
}
var errBadName = errors.New("bad name")
var badName = mvc.Response{Err: errBadName, Code: 400}
func (c *HelloController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
		// or
		// GetBy(name string) (mvc.Result, error) {
		//	return nil, errBadName
		// }
	}

	// return mvc.Response{Text: "Hello " + name} OR:
	return mvc.View{
		Name: "hello/name.html",
		Data: name,
	}
}