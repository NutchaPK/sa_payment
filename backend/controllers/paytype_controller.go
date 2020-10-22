package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theuo/app/ent"
	"github.com/theuo/app/ent/paytype"
)

// PayTypeController defines the struct for the paytype controller
type PayTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePayType handles POST requests for adding paytype entities
// @Summary Create paytype
// @Description Create paytype
// @ID create-paytype
// @Accept   json
// @Produce  json
// @Param paytype body ent.PayType true "PayType entity"
// @Success 200 {object} ent.PayType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype [post]
func (ctl *PayTypeController) CreatePayType(c *gin.Context) {
	obj := ent.PayType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paytype binding failed",
		})
		return
	}

	pt, err := ctl.client.PayType.
		Create().
		SetTypeInfo(obj.TypeInfo).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pt)
}

// GetPayType handles GET requests to retrieve a paytype entity
// @Summary Get a paytype entity by ID
// @Description get paytype by ID
// @ID get-paytype
// @Produce  json
// @Param id path int true "PayType ID"
// @Success 200 {object} ent.PayType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype/{id} [get]
func (ctl *PayTypeController) GetPayType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pt, err := ctl.client.PayType.
		Query().
		Where(paytype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pt)
}

// ListPayType handles request to get a list of paytype entities
// @Summary List paytype entities
// @Description list paytype entities
// @ID list-paytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PayType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype [get]
func (ctl *PayTypeController) ListPayType(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	users, err := ctl.client.PayType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// DeletePayType handles DELETE requests to delete a paytype entity
// @Summary Delete a paytype entity by ID
// @Description get paytype by ID
// @ID delete-paytype
// @Produce  json
// @Param id path int true "PayType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype/{id} [delete]
func (ctl *PayTypeController) DeletePayType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PayType.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdatePayType handles PUT requests to update a paytype entity
// @Summary Update a paytype entity by ID
// @Description update paytype by ID
// @ID update-paytype
// @Accept   json
// @Produce  json
// @Param id path int true "PayType ID"
// @Param paytype body ent.PayType true "PayType entity"
// @Success 200 {object} ent.PayType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype/{id} [put]
func (ctl *PayTypeController) UpdatePayType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PayType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paytype binding failed",
		})
		return
	}
	obj.ID = int(id)
	pt, err := ctl.client.PayType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, pt)
}

// NewPayTypeController creates and registers handles for the paytype controller
func NewPayTypeController(router gin.IRouter, client *ent.Client) *PayTypeController {
	ptc := &PayTypeController{
		client: client,
		router: router,
	}
	ptc.register()
	return ptc
}

// InitPayTypeController registers routes to the main engine
func (ctl *PayTypeController) register() {
	paytypes := ctl.router.Group("/paytype")

	paytypes.GET("", ctl.ListPayType)

	// CRUD
	paytypes.POST("", ctl.CreatePayType)
	paytypes.GET(":id", ctl.GetPayType)
	paytypes.PUT(":id", ctl.UpdatePayType)
	paytypes.DELETE(":id", ctl.DeletePayType)
}
