package v1

import (
	"net/http"
	"vote/models"
	"vote/utils/app"
	"vote/utils/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary 获取投票项目列表
// @Product json
// @Param mobile query string true "手机号码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/votes [get]
func GetVotes(c *gin.Context) {
	appG := app.Gin{c}
	valid := validation.Validation{}

	mobile := c.Query("mobile")
	if !valid.Mobile(mobile, "mobile").Ok {
		appG.Response(http.StatusOK, e.MobileError, nil)
		return
	}

	data := make(map[string]interface{})

	data["list"] = models.GetTuses(mobile)
	//fmt.Println(data)
	appG.Response(http.StatusOK, e.Success, data)
}
