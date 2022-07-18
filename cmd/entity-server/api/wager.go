package api

import (
	"net/http"

	"simpleservice/cmd/entity-server/service"
	"simpleservice/configs"
	"simpleservice/internal/pkg/ginutil"
	"simpleservice/internal/pkg/msg"

	"github.com/gin-gonic/gin"
)

func AddWager(c *gin.Context) {
	appG := Gin{C: c}

	var reqObj service.WagerReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		objRes, err := reqObj.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func GetWager(c *gin.Context) {
	appG := Gin{C: c}
	offset, limit := ginutil.GetPage(c, configs.GetConfig().DefaultPageNum, configs.GetConfig().DefaultPageLimit)

	service := service.WagerReq{
		PageNum:  offset,
		PageSize: limit,
	}
	list, err := service.GetWager()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = list
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), data, nil)
}
