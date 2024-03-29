package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type path struct {
	Name string
	Path string
}

var paths = []path{
	{
		Name: "Security Levels",
		Path: "security_levels",
	},
	{
		Name: "Postions",
		Path: "positions",
	},
	{
		Name: "Users",
		Path: "users",
	},
}

func GetRoot(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api-root.tmpl", gin.H{
		"paths": paths,
	})
}
