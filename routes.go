package main

import (
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Responds with the request value, the second return paramter specifies that value
// is set and non-empty (false otherwise).
func getRequestValue(ctx *fasthttp.RequestCtx, name string, args ...bool) (string, bool) {
	value := strings.TrimSpace(
		string(ctx.URI().QueryArgs().Peek(name)),
	)
	verbose := true
	if len(args) > 0 {
		verbose = args[0]
	}

	// fail if a wrong user name is provided
	if value == "" {
		if verbose {
			logger.Error("empty name")
			ctx.Error("empty name", fasthttp.StatusBadRequest)
		}
		return "", true
	}
	return value, false
}

// Handles GET /hello
func Hello(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello Page")
}

// Creates LogUser event (POST to /user)
func CreateLogUser(ctx *fasthttp.RequestCtx) {
	v, empty := getRequestValue(ctx, "name")
	if empty {
		return
	}

	// logfile backend
	err := logfile.Writef("%s: %s\n", v, logfile.TimeNow())

	// orm backend
	if ormer != nil {
		user := User{Name: v}
		_, err = ormer.Insert(&user)
	}

	// fail
	if err != nil {
		logger.Error(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// GetUser handles GET requests to /user
func GetLogUser(ctx *fasthttp.RequestCtx) {
	if ormer == nil {
		ctx.WriteString("Not Found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	v, empty := getRequestValue(ctx, "name")
	if empty {
		return
	}

	var user User
	err := ormer.QueryTable("user").OrderBy("-id").Filter("Name", v).Limit(1).One(&user)

	if err == orm.ErrNoRows {
		ctx.WriteString("Not Found")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	} else if err != nil {
		logger.Error(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	data, _ := json.Marshal(&user)
	ctx.Write(data)
}

// Initialize router endpoints
func NewRouter() *router.Router {
	router := router.New()
	router.POST("/user", CreateLogUser)
	router.GET("/user", GetLogUser)
	router.GET("/hello", Hello)

	// use custom registry
	handler := promhttp.HandlerFor(metricsRegistry, promhttp.HandlerOpts{})
	router.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(handler))
	return router
}
