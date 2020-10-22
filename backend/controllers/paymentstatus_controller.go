package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theuo/app/ent"
	"github.com/theuo/app/ent/paymentstatus"
)

// PaymentStatusController defines the struct for the paymentstatus controller
type PaymentStatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePaymentStatus handles POST requests for adding paymentstatus entities
// @Summary Create paymentstatus
// @Description Create paymentstatus
// @ID create-paymentstatus
// @Accept   json
// @Produce  json
// @Param paymentstatus body ent.PaymentStatus true "PaymentStatus entity"
// @Success 200 {object} ent.PaymentStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentstatus [post]
func (ctl *PaymentStatusController) CreatePaymentStatus(c *gin.Context) {
	obj := ent.PaymentStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymentstatus binding failed",
		})
		return
	}

	ps, err := ctl.client.PaymentStatus.
		Create().
		SetPaymentStatus(obj.PaymentStatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ps)
}

// GetPaymentStatus handles GET requests to retrieve a paymentstatus entity
// @Summary Get a paymentstatus entity by ID
// @Description get paymentstatus by ID
// @ID get-paymentstatus
// @Produce  json
// @Param id path int true "PaymentStatus ID"
// @Success 200 {object} ent.PaymentStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentstatus/{id} [get]
func (ctl *PaymentStatusController) GetPaymentStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ps, err := ctl.client.PaymentStatus.
		Query().
		Where(paymentstatus.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ps)
}

// ListPaymentStatus handles request to get a list of paymentstatus entities
// @Summary List paymentstatus entities
// @Description list paymentstatus entities
// @ID list-paymentstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PaymentStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentstatus [get]
func (ctl *PaymentStatusController) ListPaymentStatus(c *gin.Context) {
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

	paymentstatuss, err := ctl.client.PaymentStatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paymentstatuss)
}

// DeletePaymentStatus handles DELETE requests to delete a paymentstatus entity
// @Summary Delete a paymentstatus entity by ID
// @Description get paymentstatus by ID
// @ID delete-paymentstatus
// @Produce  json
// @Param id path int true "PaymentStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentstatus/{id} [delete]
func (ctl *PaymentStatusController) DeletePaymentStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PaymentStatus.
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

// UpdatePaymentStatus handles PUT requests to update a billstats entity
// @Summary Update a paymentstatus entity by ID
// @Description update paymentstatus by ID
// @ID update-paymentstatus
// @Accept   json
// @Produce  json
// @Param id path int true "PaymentStatus ID"
// @Param paymentstatus body ent.PaymentStatus true "PaymentStatus entity"
// @Success 200 {object} ent.PaymentStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentstatus/{id} [put]
func (ctl *PaymentStatusController) UpdatePaymentStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PaymentStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymentstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	ps, err := ctl.client.PaymentStatus.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ps)
}

// NewPaymentStatusController creates and registers handles for the paymentstatus controller
func NewPaymentStatusController(router gin.IRouter, client *ent.Client) *PaymentStatusController {
	psc := &PaymentStatusController{
		client: client,
		router: router,
	}
	psc.register()
	return psc
}

// InitPaymentStatusController registers routes to the main engine
func (ctl *PaymentStatusController) register() {
	paymentstatuss := ctl.router.Group("/paymentstatus")

	paymentstatuss.GET("", ctl.ListPaymentStatus)

	// CRUD
	paymentstatuss.POST("", ctl.CreatePaymentStatus)
	paymentstatuss.GET(":id", ctl.GetPaymentStatus)
	paymentstatuss.PUT(":id", ctl.UpdatePaymentStatus)
	paymentstatuss.DELETE(":id", ctl.DeletePaymentStatus)
}
