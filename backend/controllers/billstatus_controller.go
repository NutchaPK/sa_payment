package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theuo/app/ent"
	"github.com/theuo/app/ent/billstatus"
)

// BillStatusController defines the struct for the billstatus controller
type BillStatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBillStatus handles POST requests for adding billstatus entities
// @Summary Create billstatus
// @Description Create billstatus
// @ID create-billstatus
// @Accept   json
// @Produce  json
// @Param billstatus body ent.BillStatus true "BillStatus entity"
// @Success 200 {object} ent.BillStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billstatus [post]
func (ctl *BillStatusController) CreateBillStatus(c *gin.Context) {
	obj := ent.BillStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "billstatus binding failed",
		})
		return
	}

	bs, err := ctl.client.BillStatus.
		Create().
		SetBillStatus(obj.BillStatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bs)
}

// GetBillStatus handles GET requests to retrieve a billstatus entity
// @Summary Get a billstatus entity by ID
// @Description get billstatus by ID
// @ID get-billstatus
// @Produce  json
// @Param id path int true "BillStatus ID"
// @Success 200 {object} ent.BillStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billstatus/{id} [get]
func (ctl *BillStatusController) GetBillStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bs, err := ctl.client.BillStatus.
		Query().
		Where(billstatus.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bs)
}

// ListBillStatus handles request to get a list of billstatus entities
// @Summary List billstatus entities
// @Description list billstatus entities
// @ID list-billstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.BillStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billstatus [get]
func (ctl *BillStatusController) ListBillStatus(c *gin.Context) {
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

	billstatuss, err := ctl.client.BillStatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, billstatuss)
}

// DeleteBillStatus handles DELETE requests to delete a billstatus entity
// @Summary Delete a billstatus entity by ID
// @Description get billstatus by ID
// @ID delete-billstatus
// @Produce  json
// @Param id path int true "BillStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billstatus/{id} [delete]
func (ctl *BillStatusController) DeleteBillStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.BillStatus.
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

// UpdateBillStatus handles PUT requests to update a billstats entity
// @Summary Update a billstatus entity by ID
// @Description update billstatus by ID
// @ID update-billstatus
// @Accept   json
// @Produce  json
// @Param id path int true "BillStatus ID"
// @Param billstatus body ent.BillStatus true "BillStatus entity"
// @Success 200 {object} ent.BillStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /billstatus/{id} [put]
func (ctl *BillStatusController) UpdateBillStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.BillStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "billstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	bs, err := ctl.client.BillStatus.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, bs)
}

// NewBillStatusController creates and registers handles for the billstatus controller
func NewBillStatusController(router gin.IRouter, client *ent.Client) *BillStatusController {
	bsc := &BillStatusController{
		client: client,
		router: router,
	}
	bsc.register()
	return bsc
}

// InitBillStatusController registers routes to the main engine
func (ctl *BillStatusController) register() {
	billstatuss := ctl.router.Group("/billstatus")

	billstatuss.GET("", ctl.ListBillStatus)

	// CRUD
	billstatuss.POST("", ctl.CreateBillStatus)
	billstatuss.GET(":id", ctl.GetBillStatus)
	billstatuss.PUT(":id", ctl.UpdateBillStatus)
	billstatuss.DELETE(":id", ctl.DeleteBillStatus)
}
