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
// @Success 200 {array} models.Tuse "{"code":200,"data":{},"msg":"ok"}"
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

	//data["list"] = models.GetTuses(mobile)
	data["list"] = []models.Tuse{
		{
			TuseHouse_id:        13765,
			Tuse_id:             29,
			Thouse_id:           158,
			TuseHouse_sumAmount: "56.56",
			TuseHouse_ZJfentan:  "56.56",
			TuseHouse_XJfentan:  "0.00",
			TuseHouse_voteValue: 0,
			Tuse_content:        "公共设施维修",
			Tuse_hezhunAmount:   "46829.00",
			Tuse_fentanHouse:    "姓名158",
		},
		{
			TuseHouse_id:        16335,
			Tuse_id:             34,
			Thouse_id:           158,
			TuseHouse_sumAmount: "49.92",
			TuseHouse_ZJfentan:  "49.92",
			TuseHouse_XJfentan:  "0.00",
			TuseHouse_voteValue: 0,
			Tuse_content:        "公共设施维修",
			Tuse_hezhunAmount:   "41326.60",
			Tuse_fentanHouse:    "姓名158",
		},
	}
	//fmt.Println(data)
	appG.Response(http.StatusOK, e.Success, data)
}
