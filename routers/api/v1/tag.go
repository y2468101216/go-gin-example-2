package v1

import (
	"fmt"
	"gogin/example/models"
	"gogin/example/pkg/e"
	"gogin/example/pkg/logging"
	"gogin/example/pkg/setting"
	"gogin/example/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
    var query models.GetTagsForm
    if err := c.ShouldBindQuery(&query); err != nil {
        logging.Info(err.Error())
        code := e.INVALID_PARAMS
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if query.Name != "" {
        maps["name"] = query.Name
    }

    if c.Query("state") != "" {
        maps["state"] = query.State
    }

    fmt.Print(c.Query("state"))

    code := e.SUCCESS

    data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
    data["total"] = models.GetTagTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

//新增文章标签
func AddTag(c *gin.Context) {
    fmt.Println("controller")
    var json models.AddTagFrom
    if err := c.ShouldBindJSON(&json); err != nil {
        code := e.INVALID_PARAMS
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

    code := e.INVALID_PARAMS
    if ! models.ExistTagByName(json.Name) {
        code = e.SUCCESS
        models.AddTag(json.Name, json.State, json.CreatedBy)
    } else {
        code = e.ERROR_EXIST_TAG
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

//修改文章标签
func EditTag(c *gin.Context) {
    var json models.EditTagForm
    if err := c.ShouldBindJSON(&json); err != nil {
        code := e.INVALID_PARAMS
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    
    if err != nil {
        code := e.INVALID_PARAMS
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

    code := e.INVALID_PARAMS
    if models.ExistTagById(id) {
        code = e.SUCCESS
        data := make(map[string]interface{})
        data["modified_by"] = json.ModifiedBy
        if json.Name != "" {
            data["name"] = json.Name
        }


        models.EditTag(id, data)
    } else {
        code = e.ERROR_NOT_EXIST_TAG
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

//删除文章标签
func DeleteTag(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    
    if err != nil {
        code := e.INVALID_PARAMS
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

    models.DeleteTag(id)

    code := e.SUCCESS
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}