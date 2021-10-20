package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kaindy7633/go-programming-tour-book/blog-service/docs"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/global"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/routers/api"
	v1 "github.com/kaindy7633/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	// 新增auth相关路由
	r.POST("/auth", api.GetAuth)

	// 上传文件服务
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 提供静态资源访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
