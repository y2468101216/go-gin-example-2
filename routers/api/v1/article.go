package v1

import (
	"net/http"
	"strconv"

	"gogin/example/models"
	"gogin/example/pkg/e"
	"gogin/example/pkg/setting"
	"gogin/example/pkg/util"

	"github.com/gin-gonic/gin"
)

//获取单个文章
func GetArticle(c *gin.Context) {

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

	code := e.ERROR_NOT_EXIST_ARTICLE
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

//获取多个文章
func GetArticles(c *gin.Context) {
	var query models.GetArticlesForm
	code := e.INVALID_PARAMS
    if err := c.ShouldBindQuery(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

	maps := make(map[string]interface{})
	if c.Query("state") != "" {
        maps["state"] = query.State
    }

	if c.Query("tag_id") != "" {
        maps["tag_id"] = query.TagID
    }

	data := make(map[string]interface{})

	data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

//新增文章
func AddArticle(c *gin.Context) {
	var query models.AddArticleForm
	code := e.INVALID_PARAMS
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

	models.AddArticle(query)

	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]interface{}),
    })
}

//修改文章
func EditArticle(c *gin.Context) {
	var query models.EditArticleForm
	code := e.INVALID_PARAMS
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
	}

	data := make(map[string]interface{})
	if query.TagID != 0 {
		data["tag_id"] = query.TagID
	}

	if query.Title != "" {
		data["title"] = query.Title
	}

	if query.Desc != "" {
		data["desc"] = query.Desc
	}

	if query.Content != "" {
		data["content"] = query.Content
	}

	if query.ModifiedBy != "" {
		data["modified_by"] = query.ModifiedBy
	}

	models.EditArticle(id, data)

	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	code := e.INVALID_PARAMS
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
	}

	models.DeleteArticle(id)

	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}