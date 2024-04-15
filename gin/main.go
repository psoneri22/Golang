package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// go:embed public/*
// var f embed.FS

// type TimeoffRequest struct {
// 	// Date   time.Time `form:"date" binding"-" time_format:"2006-01-02"`
// 	Date   time.Time `json:"date" form:"date" binding:"required,future" time_format:"2006-01-02"`
// 	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
// }

// var validatorFuture validator.Func = func(fl validator.FieldLevel) bool {
// 	date, ok := fl.Field().Interface().(time.Time)
// 	if ok {
// 		return date.After(time.Now())
// 	}
// 	return true
// }

func main() {

	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(myErrorLog)
	router.Use(gin.CustomRecovery(myRecoveryFunc))
	// router.Run()
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("future", validatorFuture)
	// }
	// router.StaticFile("/", "./public/index.html")
	// router.Static("/public", "./public")
	// router.StaticFS("/fs", http.FileSystem(http.FS(f)))
	// router.GET("/employee", func(ctx *gin.Context) {
	// 	ctx.File("./public/employee.html")
	// })
	// router.POST("/employee", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "New Request Posted Successfully!!")
	// })

	// router.GET("/employee/:username/*rest", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"username": ctx.Param("username"),
	// 		"rest":     ctx.Param("rest"),
	// 	})
	// })

	adminGroup := router.Group("/adminGrp", Benchmark)
	adminGroup.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "page administer users")
	})
	adminGroup.GET("/roles", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "page administer roles")
	})
	adminGroup.GET("/policies", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "page administer policies")
	})

	router.GET("/errors", func(ctx *gin.Context) {
		err := &gin.Error{
			Err:  errors.New("Something went wrong"),
			Type: gin.ErrorTypeRender | gin.ErrorTypePublic,
			Meta: "this error was intentional",
		}
		ctx.Error(err)
	})

	router.GET("/panic", func(ctx *gin.Context) {
		panic("a Go program should almostnever call 'panic'")
	})
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "Hello World")
	// })

	// router.GET("/*rest", func(ctx *gin.Context) {
	// 	url := ctx.Request.URL.String()
	// 	headers := ctx.Request.Header
	// 	cookies := ctx.Request.Cookies()

	// 	ctx.IndentedJSON(http.StatusOK, gin.H{
	// 		"url":     url,
	// 		"headers": headers,
	// 		"cookies": cookies,
	// 	})
	// })

	// router.GET("/query/*rest", func(ctx *gin.Context) {
	// 	username := ctx.Query("username")
	// 	year := ctx.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
	// 	months := ctx.QueryArray("month")

	// 	ctx.IndentedJSON(http.StatusOK, gin.H{
	// 		"url":    username,
	// 		"year":   year,
	// 		"months": months,
	// 	})
	// })

	// router.POST("/employee", func(ctx *gin.Context) {
	// 	date := ctx.PostForm("date")
	// 	amount := ctx.PostForm("amount")
	// 	username := ctx.DefaultPostForm("username", "Soneri")
	// 	ctx.IndentedJSON(http.StatusOK, gin.H{
	// 		"date":     date,
	// 		"amount":   amount,
	// 		"username": username,
	// 	})
	// })

	// router.POST("/employee", func(ctx *gin.Context) {
	// 	var timeoffRequest TimeoffRequest
	// 	if err := ctx.ShouldBind(&timeoffRequest); err == nil {
	// 		ctx.JSON(http.StatusOK, timeoffRequest)
	// 	} else {
	// 		ctx.String(http.StatusInternalServerError, err.Error())
	// 	}
	// })

	// apiGroup := router.Group("/api")
	// apiGroup.POST("/timeoff", func(ctx *gin.Context) {
	// 	var timeoffRequest TimeoffRequest
	// 	if err := ctx.ShouldBind(&timeoffRequest); err == nil {
	// 		ctx.JSON(http.StatusOK, timeoffRequest)
	// 	} else {
	// 		ctx.String(http.StatusInternalServerError, err.Error())
	// 	}
	// })

	// router.StaticFile("/", "./public/index.html")

	// router.GET("/table_of_two_cities", func(ctx *gin.Context) {
	// 	f, err := os.Open("./a_table_of_two_cities.csv")
	// 	if err != nil {
	// 		ctx.AbortWithError(http.StatusInternalServerError, err)
	// 	}
	// 	defer f.Close()
	// 	ctx.Stream(stremer(f))
	// 	// data, err := io.ReadAll(f)
	// 	// if err != nil {
	// 	// 	ctx.AbortWithError(http.StatusInternalServerError, err)
	// 	// }
	// 	// ctx.Data(http.StatusOK, "text/plain", data)
	// 	// ctx.File("./a_table_of_two_cities.csv")
	// })

	// router.GET("/great_expactation", func(ctx *gin.Context) {
	// 	f, err := os.Open("./great_expactation.csv")
	// 	if err != nil {
	// 		ctx.AbortWithError(http.StatusInternalServerError, err)
	// 	}
	// 	defer f.Close()
	// 	fi, err := f.Stat()
	// 	if err != nil {
	// 		ctx.AbortWithError(http.StatusInternalServerError, err)
	// 	}
	// 	ctx.DataFromReader(http.StatusOK, fi.Size(), "text/plain", f, map[string]string{
	// 		"Content-Disposition": "attachment;filename=great_expactation.csv",
	// 	})
	// ctx.File("./great_expactation.csv")
	// })

	log.Fatal(router.Run(":3001"))
}

// func stremer(r io.Reader) func(io.Writer) bool {
// 	return func(step io.Writer) bool {
// 		for {
// 			buf := make([]byte, 4*2^10)
// 			if _, err := r.Read(buf); err == nil {
// 				_, err := step.Write(buf)
// 				return err == nil
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// }

var Benchmark gin.HandlerFunc = func(ctx *gin.Context) {
	t := time.Now()
	ctx.Next()
	elapsed := time.Since(t)
	log.Print("Time to process", elapsed)
}

var myErrorLog gin.HandlerFunc = func(ctx *gin.Context) {
	ctx.Next()
	for _, err := range ctx.Errors {
		log.Print(map[string]any{
			"Err":  err.Error(),
			"Type": err.Type,
			"Meta": err.Meta,
		})
	}
}

var myRecoveryFunc gin.RecoveryFunc = func(c *gin.Context, err any) {
	log.Print("Custom recovery function can be used to add fine-grained control over recovery stratagies.", err)
}
