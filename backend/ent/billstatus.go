// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/theuo/app/ent/billstatus"
)

// BillStatus is the model entity for the BillStatus schema.
type BillStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BillStatus holds the value of the "BillStatus" field.
	BillStatus string `json:"BillStatus,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillStatusQuery when eager-loading is set.
	Edges BillStatusEdges `json:"edges"`
}

// BillStatusEdges holds the relations/edges for other nodes in the graph.
type BillStatusEdges struct {
	// Bills holds the value of the bills edge.
	Bills []*Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading.
func (e BillStatusEdges) BillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[0] {
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "bills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillStatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // BillStatus
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillStatus fields.
func (bs *BillStatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(billstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	bs.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field BillStatus", values[0])
	} else if value.Valid {
		bs.BillStatus = value.String
	}
	return nil
}

// QueryBills queries the bills edge of the BillStatus.
func (bs *BillStatus) QueryBills() *BillQuery {
	return (&BillStatusClient{config: bs.config}).QueryBills(bs)
}

// Update returns a builder for updating this BillStatus.
// Note that, you need to call BillStatus.Unwrap() before calling this method, if this BillStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (bs *BillStatus) Update() *BillStatusUpdateOne {
	return (&BillStatusClient{config: bs.config}).UpdateOne(bs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (bs *BillStatus) Unwrap() *BillStatus {
	tx, ok := bs.config.driver.(*txDriver)
	if !ok {
		panic("ent: BillStatus is not a transactional entity")
	}
	bs.config.driver = tx.drv
	return bs
}

// String implements the fmt.Stringer.
func (bs *BillStatus) String() string {
	var builder strings.Builder
	builder.WriteString("BillStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", bs.ID))
	builder.WriteString(", BillStatus=")
	builder.WriteString(bs.BillStatus)
	builder.WriteByte(')')
	return builder.String()
}

// BillStatusSlice is a parsable slice of BillStatus.
type BillStatusSlice []*BillStatus

func (bs BillStatusSlice) config(cfg config) {
	for _i := range bs {
		bs[_i].config = cfg
	}
}