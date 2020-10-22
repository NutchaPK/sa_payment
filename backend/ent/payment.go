// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/theuo/app/ent/bill"
	"github.com/theuo/app/ent/payment"
	"github.com/theuo/app/ent/paymentstatus"
	"github.com/theuo/app/ent/paytype"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AddDatetime holds the value of the "AddDatetime" field.
	AddDatetime time.Time `json:"AddDatetime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges            PaymentEdges `json:"edges"`
	bill_id          *int
	paytype_id       *int
	paymentstatus_id *int
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// Paytype holds the value of the paytype edge.
	Paytype *PayType
	// Paymentstatus holds the value of the paymentstatus edge.
	Paymentstatus *PaymentStatus
	// Bill holds the value of the bill edge.
	Bill *Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PaytypeOrErr returns the Paytype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) PaytypeOrErr() (*PayType, error) {
	if e.loadedTypes[0] {
		if e.Paytype == nil {
			// The edge paytype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: paytype.Label}
		}
		return e.Paytype, nil
	}
	return nil, &NotLoadedError{edge: "paytype"}
}

// PaymentstatusOrErr returns the Paymentstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) PaymentstatusOrErr() (*PaymentStatus, error) {
	if e.loadedTypes[1] {
		if e.Paymentstatus == nil {
			// The edge paymentstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: paymentstatus.Label}
		}
		return e.Paymentstatus, nil
	}
	return nil, &NotLoadedError{edge: "paymentstatus"}
}

// BillOrErr returns the Bill value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) BillOrErr() (*Bill, error) {
	if e.loadedTypes[2] {
		if e.Bill == nil {
			// The edge bill was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bill.Label}
		}
		return e.Bill, nil
	}
	return nil, &NotLoadedError{edge: "bill"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // AddDatetime
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Payment) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // bill_id
		&sql.NullInt64{}, // paytype_id
		&sql.NullInt64{}, // paymentstatus_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(payment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field AddDatetime", values[0])
	} else if value.Valid {
		pa.AddDatetime = value.Time
	}
	values = values[1:]
	if len(values) == len(payment.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field bill_id", value)
		} else if value.Valid {
			pa.bill_id = new(int)
			*pa.bill_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field paytype_id", value)
		} else if value.Valid {
			pa.paytype_id = new(int)
			*pa.paytype_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field paymentstatus_id", value)
		} else if value.Valid {
			pa.paymentstatus_id = new(int)
			*pa.paymentstatus_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPaytype queries the paytype edge of the Payment.
func (pa *Payment) QueryPaytype() *PayTypeQuery {
	return (&PaymentClient{config: pa.config}).QueryPaytype(pa)
}

// QueryPaymentstatus queries the paymentstatus edge of the Payment.
func (pa *Payment) QueryPaymentstatus() *PaymentStatusQuery {
	return (&PaymentClient{config: pa.config}).QueryPaymentstatus(pa)
}

// QueryBill queries the bill edge of the Payment.
func (pa *Payment) QueryBill() *BillQuery {
	return (&PaymentClient{config: pa.config}).QueryBill(pa)
}

// Update returns a builder for updating this Payment.
// Note that, you need to call Payment.Unwrap() before calling this method, if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return (&PaymentClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", AddDatetime=")
	builder.WriteString(pa.AddDatetime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment

func (pa Payments) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
