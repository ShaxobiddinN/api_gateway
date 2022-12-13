package handlers

import (
	"net/http"
	"strconv"

	"blogpost/api_gateway/models"
	"blogpost/api_gateway/protogen/blogpost"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "article body"
// @Success     201     {object} models.JSONResponse{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponce
// @Router      /v1/article [post]
func (h handler) CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	//ToDo - validation should be here

	article, err := h.grpcClients.Article.CreateArticle(c.Request.Context(), &blogpost.CreateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		AuthorId: body.AuthorID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | Created",
		Data:    article,
	})
}

// GetArticleById godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponce
// @Router      /v1/article/{id} [get]
func (h handler) GetArticleById(c *gin.Context) {

	idStr := c.Param("id")

	article, err := h.grpcClients.Article.GetArticleById(c.Request.Context(), &blogpost.GetArticleByIdRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})

}

// getArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int     false "0"
// @Param       limit  query    int     false "0"
// @Param       search query    string  false "smth"
// @Success     200    {object} models.JSONResponse{data=[]models.Article}
// @Router      /v1/article [get]
func (h handler) GetArticleList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClients.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    articleList,
	})
}

// UpdateArticle...
// UpdateArticle godoc
// @Summary     Update article
// @Description update a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.UpdateArticleModel true "article body"
// @Success     200     {object} models.JSONResponse{data=[]models.Article}
// @Failure     400     {object} models.JSONErrorResponce
// @Router      /v1/article [put]
func (h handler) UpdateArticle(c *gin.Context) {

	var body models.UpdateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}

	article, err := h.grpcClients.Article.UpdateArticle(c.Request.Context(), &blogpost.UpdateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		Id: body.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    article,
	})
}

// DeleteArticle...
// DeleteArticle godoc
// @Summary     delete article by id
// @Description delete an article by id
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponce
// @Router      /v1/article/{id} [delete]
func (h handler) DeleteArticle(c *gin.Context) {

	idStr := c.Param("id")

	article, err := h.grpcClients.Article.DeleteArticle(c.Request.Context(), &blogpost.DeleteArticleRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    article,
	})
}
