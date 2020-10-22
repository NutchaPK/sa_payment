package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/theuo/app/ent"
	"github.com/theuo/app/ent/bill"
	"github.com/theuo/app/ent/payment"
	"github.com/theuo/app/ent/paymentstatus"
	"github.com/theuo/app/ent/paytype"
)

// PaymentController defines the struct for the payment controller
type PaymentController struct {
	client *ent.Client
	router gin.IRouter
}

type Payment struct {
	AddDatetime   string
	PayType       int
	PaymentStatus int
	Bill          int
}

// CreatePayment handles POST requests for adding payment entities
// @Summary Create payment
// @Description Create payment
// @ID create-payment
// @Accept   json
// @Produce  json
// @Param payment body Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment [post]
func (ctl *PaymentController) CreatePayment(c *gin.Context) {
	obj := Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment binding failed",
		})
		return
	}
	b, err := ctl.client.Bill.
		Query().
		Where(bill.IDEQ(int(obj.Bill))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}
	ps, err := ctl.client.PaymentStatus.
		Query().
		Where(paymentstatus.IDEQ(int(obj.PaymentStatus))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "paymentstatus not found",
		})
		return
	}
	pt, err := ctl.client.PayType.
		Query().
		Where(paytype.IDEQ(int(obj.PayType))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "paytype not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.AddDatetime)
	pm, err := ctl.client.Payment.
		Create().
		SetAddDatetime(times).
		SetBill(b).
		SetPaytype(pt).
		SetPaymentstatus(ps).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pm)
}

// GetPayment handles GET requests to retrieve a payment entity
// @Summary Get a payment entity by ID
// @Description get payment by ID
// @ID get-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment/{id} [get]
func (ctl *PaymentController) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Payment.
		Query().
		WithBill().
		WithPaymentstatus().
		WithPaytype().
		Where(payment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListPayment handles request to get a list of payment entities
// @Summary List payment entities
// @Description list payment entities
// @ID list-payment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment [get]
func (ctl *PaymentController) ListPayment(c *gin.Context) {
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

	users, err := ctl.client.Payment.
		Query().
		WithPaytype().
		WithPaymentstatus().
		WithBill().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// DeletePayment handles DELETE requests to delete a payment entity
// @Summary Delete a payment entity by ID
// @Description get payment by ID
// @ID delete-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment/{id} [delete]
func (ctl *PaymentController) DeletePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Payment.
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

// UpdatePayment handles PUT requests to update a Payment entity
// @Summary Update a payment entity by ID
// @Description update payment by ID
// @ID update-payment
// @Accept   json
// @Produce  json
// @Param id path int true "Payment ID"
// @Param payment body ent.Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment/{id} [put]
func (ctl *PaymentController) UpdatePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Payment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewPaymentController creates and registers handles for the payment controller
func NewPaymentController(router gin.IRouter, client *ent.Client) *PaymentController {
	uc := &PaymentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPaymentController registers routes to the main engine
func (ctl *PaymentController) register() {
	payments := ctl.router.Group("/payment")

	payments.GET("", ctl.ListPayment)

	// CRUD
	payments.POST("", ctl.CreatePayment)
	payments.GET(":id", ctl.GetPayment)
	payments.PUT(":id", ctl.UpdatePayment)
	payments.DELETE(":id", ctl.DeletePayment)
}
