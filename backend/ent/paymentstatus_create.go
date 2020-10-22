// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/theuo/app/ent/payment"
	"github.com/theuo/app/ent/paymentstatus"
)

// PaymentStatusCreate is the builder for creating a PaymentStatus entity.
type PaymentStatusCreate struct {
	config
	mutation *PaymentStatusMutation
	hooks    []Hook
}

// SetPaymentStatus sets the PaymentStatus field.
func (psc *PaymentStatusCreate) SetPaymentStatus(s string) *PaymentStatusCreate {
	psc.mutation.SetPaymentStatus(s)
	return psc
}

// AddPaymentIDs adds the payments edge to Payment by ids.
func (psc *PaymentStatusCreate) AddPaymentIDs(ids ...int) *PaymentStatusCreate {
	psc.mutation.AddPaymentIDs(ids...)
	return psc
}

// AddPayments adds the payments edges to Payment.
func (psc *PaymentStatusCreate) AddPayments(p ...*Payment) *PaymentStatusCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddPaymentIDs(ids...)
}

// Mutation returns the PaymentStatusMutation object of the builder.
func (psc *PaymentStatusCreate) Mutation() *PaymentStatusMutation {
	return psc.mutation
}

// Save creates the PaymentStatus in the database.
func (psc *PaymentStatusCreate) Save(ctx context.Context) (*PaymentStatus, error) {
	if _, ok := psc.mutation.PaymentStatus(); !ok {
		return nil, &ValidationError{Name: "PaymentStatus", err: errors.New("ent: missing required field \"PaymentStatus\"")}
	}
	if v, ok := psc.mutation.PaymentStatus(); ok {
		if err := paymentstatus.PaymentStatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "PaymentStatus", err: fmt.Errorf("ent: validator failed for field \"PaymentStatus\": %w", err)}
		}
	}
	var (
		err  error
		node *PaymentStatus
	)
	if len(psc.hooks) == 0 {
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psc.mutation = mutation
			node, err = psc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PaymentStatusCreate) SaveX(ctx context.Context) *PaymentStatus {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (psc *PaymentStatusCreate) sqlSave(ctx context.Context) (*PaymentStatus, error) {
	ps, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ps.ID = int(id)
	return ps, nil
}

func (psc *PaymentStatusCreate) createSpec() (*PaymentStatus, *sqlgraph.CreateSpec) {
	var (
		ps    = &PaymentStatus{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paymentstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymentstatus.FieldID,
			},
		}
	)
	if value, ok := psc.mutation.PaymentStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paymentstatus.FieldPaymentStatus,
		})
		ps.PaymentStatus = value
	}
	if nodes := psc.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymentstatus.PaymentsTable,
			Columns: []string{paymentstatus.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ps, _spec
}
