package routers

import (
	"NULL/blockchain/middleware/cors"
	"NULL/blockchain/pkg/export"
	"NULL/blockchain/pkg/qrcode"
	"NULL/blockchain/pkg/upload"
	"NULL/blockchain/routers/api"
	v1 "NULL/blockchain/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORSMiddleware())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//上传文件
		apiv1.POST("/file/upload", api.UploadFile)
		//文件下载
		apiv1.StaticFS("/file/download", http.Dir(upload.GetImageFullPath()))

		apiv1.POST("/block", v1.CreateBlockchain)
	}
	return r
}
