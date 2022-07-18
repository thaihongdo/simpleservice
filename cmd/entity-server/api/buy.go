package api

import (
	"net/http"
	"simpleservice/cmd/entity-server/service"
	"simpleservice/internal/pkg/msg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Buy(c *gin.Context) {
	appG := Gin{C: c}

	id, err := strconv.Atoi(c.Param("wager-id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var reqObj service.BuyReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		reqObj.WagerID = uint(id)
		objRes, err := reqObj.Buy()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}
