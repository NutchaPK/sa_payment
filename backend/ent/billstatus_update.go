// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/theuo/app/ent/bill"
	"github.com/theuo/app/ent/billstatus"
	"github.com/theuo/app/ent/predicate"
)

// BillStatusUpdate is the builder for updating BillStatus entities.
type BillStatusUpdate struct {
	config
	hooks      []Hook
	mutation   *BillStatusMutation
	predicates []predicate.BillStatus
}

// Where adds a new predicate for the builder.
func (bsu *BillStatusUpdate) Where(ps ...predicate.BillStatus) *BillStatusUpdate {
	bsu.predicates = append(bsu.predicates, ps...)
	return bsu
}

// SetBillStatus sets the BillStatus field.
func (bsu *BillStatusUpdate) SetBillStatus(s string) *BillStatusUpdate {
	bsu.mutation.SetBillStatus(s)
	return bsu
}

// AddBillIDs adds the bills edge to Bill by ids.
func (bsu *BillStatusUpdate) AddBillIDs(ids ...int) *BillStatusUpdate {
	bsu.mutation.AddBillIDs(ids...)
	return bsu
}

// AddBills adds the bills edges to Bill.
func (bsu *BillStatusUpdate) AddBills(b ...*Bill) *BillStatusUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bsu.AddBillIDs(ids...)
}

// Mutation returns the BillStatusMutation object of the builder.
func (bsu *BillStatusUpdate) Mutation() *BillStatusMutation {
	return bsu.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (bsu *BillStatusUpdate) RemoveBillIDs(ids ...int) *BillStatusUpdate {
	bsu.mutation.RemoveBillIDs(ids...)
	return bsu
}

// RemoveBills removes bills edges to Bill.
func (bsu *BillStatusUpdate) RemoveBills(b ...*Bill) *BillStatusUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bsu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bsu *BillStatusUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bsu.hooks) == 0 {
		affected, err = bsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bsu.mutation = mutation
			affected, err = bsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bsu.hooks) - 1; i >= 0; i-- {
			mut = bsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bsu *BillStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := bsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bsu *BillStatusUpdate) Exec(ctx context.Context) error {
	_, err := bsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsu *BillStatusUpdate) ExecX(ctx context.Context) {
	if err := bsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsu *BillStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   billstatus.Table,
			Columns: billstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: billstatus.FieldID,
			},
		},
	}
	if ps := bsu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsu.mutation.BillStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: billstatus.FieldBillStatus,
		})
	}
	if nodes := bsu.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billstatus.BillsTable,
			Columns: []string{billstatus.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billstatus.BillsTable,
			Columns: []string{billstatus.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BillStatusUpdateOne is the builder for updating a single BillStatus entity.
type BillStatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *BillStatusMutation
}

// SetBillStatus sets the BillStatus field.
func (bsuo *BillStatusUpdateOne) SetBillStatus(s string) *BillStatusUpdateOne {
	bsuo.mutation.SetBillStatus(s)
	return bsuo
}

// AddBillIDs adds the bills edge to Bill by ids.
func (bsuo *BillStatusUpdateOne) AddBillIDs(ids ...int) *BillStatusUpdateOne {
	bsuo.mutation.AddBillIDs(ids...)
	return bsuo
}

// AddBills adds the bills edges to Bill.
func (bsuo *BillStatusUpdateOne) AddBills(b ...*Bill) *BillStatusUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bsuo.AddBillIDs(ids...)
}

// Mutation returns the BillStatusMutation object of the builder.
func (bsuo *BillStatusUpdateOne) Mutation() *BillStatusMutation {
	return bsuo.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (bsuo *BillStatusUpdateOne) RemoveBillIDs(ids ...int) *BillStatusUpdateOne {
	bsuo.mutation.RemoveBillIDs(ids...)
	return bsuo
}

// RemoveBills removes bills edges to Bill.
func (bsuo *BillStatusUpdateOne) RemoveBills(b ...*Bill) *BillStatusUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bsuo.RemoveBillIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (bsuo *BillStatusUpdateOne) Save(ctx context.Context) (*BillStatus, error) {

	var (
		err  error
		node *BillStatus
	)
	if len(bsuo.hooks) == 0 {
		node, err = bsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bsuo.mutation = mutation
			node, err = bsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bsuo.hooks) - 1; i >= 0; i-- {
			mut = bsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bsuo *BillStatusUpdateOne) SaveX(ctx context.Context) *BillStatus {
	bs, err := bsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// Exec executes the query on the entity.
func (bsuo *BillStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := bsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsuo *BillStatusUpdateOne) ExecX(ctx context.Context) {
	if err := bsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsuo *BillStatusUpdateOne) sqlSave(ctx context.Context) (bs *BillStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   billstatus.Table,
			Columns: billstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: billstatus.FieldID,
			},
		},
	}
	id, ok := bsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BillStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := bsuo.mutation.BillStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: billstatus.FieldBillStatus,
		})
	}
	if nodes := bsuo.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billstatus.BillsTable,
			Columns: []string{billstatus.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bsuo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   billstatus.BillsTable,
			Columns: []string{billstatus.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	bs = &BillStatus{config: bsuo.config}
	_spec.Assign = bs.assignValues
	_spec.ScanValues = bs.scanValues()
	if err = sqlgraph.UpdateNode(ctx, bsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{billstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return bs, nil
}