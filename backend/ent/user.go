// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sthl/ent/siteui"
	"sthl/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Email holds the value of the "email" field.
	Email string `json:"email"`
	// HashedPw holds the value of the "hashed_pw" field.
	HashedPw string `json:"-"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"emailVerified"`
	// IsArchived holds the value of the "is_archived" field.
	IsArchived bool `json:"isArchived"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"-"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// Siteui holds the value of the siteui edge.
	Siteui *Siteui `json:"siteui,omitempty"`
	// Imagesinfo holds the value of the imagesinfo edge.
	Imagesinfo []*Imageinfo `json:"imagesinfo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// SiteuiOrErr returns the Siteui value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SiteuiOrErr() (*Siteui, error) {
	if e.loadedTypes[2] {
		if e.Siteui == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: siteui.Label}
		}
		return e.Siteui, nil
	}
	return nil, &NotLoadedError{edge: "siteui"}
}

// ImagesinfoOrErr returns the Imagesinfo value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ImagesinfoOrErr() ([]*Imageinfo, error) {
	if e.loadedTypes[3] {
		return e.Imagesinfo, nil
	}
	return nil, &NotLoadedError{edge: "imagesinfo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailVerified, user.FieldIsArchived:
			values[i] = new(sql.NullBool)
		case user.FieldEmail, user.FieldHashedPw:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldHashedPw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_pw", values[i])
			} else if value.Valid {
				u.HashedPw = value.String
			}
		case user.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				u.EmailVerified = value.Bool
			}
		case user.FieldIsArchived:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_archived", values[i])
			} else if value.Valid {
				u.IsArchived = value.Bool
			}
		}
	}
	return nil
}

// QueryProducts queries the "products" edge of the User entity.
func (u *User) QueryProducts() *ProductQuery {
	return NewUserClient(u.config).QueryProducts(u)
}

// QueryOrders queries the "orders" edge of the User entity.
func (u *User) QueryOrders() *OrderQuery {
	return NewUserClient(u.config).QueryOrders(u)
}

// QuerySiteui queries the "siteui" edge of the User entity.
func (u *User) QuerySiteui() *SiteuiQuery {
	return NewUserClient(u.config).QuerySiteui(u)
}

// QueryImagesinfo queries the "imagesinfo" edge of the User entity.
func (u *User) QueryImagesinfo() *ImageinfoQuery {
	return NewUserClient(u.config).QueryImagesinfo(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("hashed_pw=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email_verified=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailVerified))
	builder.WriteString(", ")
	builder.WriteString("is_archived=")
	builder.WriteString(fmt.Sprintf("%v", u.IsArchived))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User