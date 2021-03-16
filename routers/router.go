package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_blog/docs"
	"go_blog/middleware/jwt"
	"go_blog/pkg/setting"
	"go_blog/routers/api"
	v1 "go_blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 取得標籤列表
		apiv1.GET("/tags", v1.GetTags)
		// 建立標籤
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定標籤
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 刪除指定標籤
		apiv1.DELETE("/tags/:id", v1.DeleteTag)


		// 取得文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 取得指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 建立文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 刪除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
