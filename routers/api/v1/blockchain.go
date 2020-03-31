package v1

import (
	"NULL/blockchain/models"
	"NULL/blockchain/pkg/app"
	"NULL/blockchain/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// Message takes incoming JSON payload for writing heart rate
type BussinessForm struct {
	BS string `json:"bs"`
}

// Blockchain is a series of validated Blocks
var Blockchain []models.Block

var mutex = &sync.Mutex{}

func CreateBlockchain(c *gin.Context) {
	var (
		appG      = app.Gin{C: c}
		prevBlock = models.Block{}
		form      BussinessForm
		bs        string
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	bs = form.BS

	mutex.Lock()
	prevBlock = Blockchain[len(Blockchain)-1]
	newBlock := models.GenerateBlock(prevBlock, bs)

	if models.IsBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
	}
	mutex.Unlock()
	appG.Response(http.StatusOK, e.SUCCESS, Blockchain)
}
