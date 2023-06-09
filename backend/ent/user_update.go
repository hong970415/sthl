// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sthl/ent/imageinfo"
	"sthl/ent/order"
	"sthl/ent/predicate"
	"sthl/ent/product"
	"sthl/ent/siteui"
	"sthl/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetHashedPw sets the "hashed_pw" field.
func (uu *UserUpdate) SetHashedPw(s string) *UserUpdate {
	uu.mutation.SetHashedPw(s)
	return uu
}

// SetEmailVerified sets the "email_verified" field.
func (uu *UserUpdate) SetEmailVerified(b bool) *UserUpdate {
	uu.mutation.SetEmailVerified(b)
	return uu
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetEmailVerified(*b)
	}
	return uu
}

// SetIsArchived sets the "is_archived" field.
func (uu *UserUpdate) SetIsArchived(b bool) *UserUpdate {
	uu.mutation.SetIsArchived(b)
	return uu
}

// SetNillableIsArchived sets the "is_archived" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsArchived(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsArchived(*b)
	}
	return uu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (uu *UserUpdate) AddProductIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddProductIDs(ids...)
	return uu
}

// AddProducts adds the "products" edges to the Product entity.
func (uu *UserUpdate) AddProducts(p ...*Product) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddProductIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (uu *UserUpdate) AddOrderIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddOrderIDs(ids...)
	return uu
}

// AddOrders adds the "orders" edges to the Order entity.
func (uu *UserUpdate) AddOrders(o ...*Order) *UserUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.AddOrderIDs(ids...)
}

// SetSiteuiID sets the "siteui" edge to the Siteui entity by ID.
func (uu *UserUpdate) SetSiteuiID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetSiteuiID(id)
	return uu
}

// SetNillableSiteuiID sets the "siteui" edge to the Siteui entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableSiteuiID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetSiteuiID(*id)
	}
	return uu
}

// SetSiteui sets the "siteui" edge to the Siteui entity.
func (uu *UserUpdate) SetSiteui(s *Siteui) *UserUpdate {
	return uu.SetSiteuiID(s.ID)
}

// AddImagesinfoIDs adds the "imagesinfo" edge to the Imageinfo entity by IDs.
func (uu *UserUpdate) AddImagesinfoIDs(ids ...int) *UserUpdate {
	uu.mutation.AddImagesinfoIDs(ids...)
	return uu
}

// AddImagesinfo adds the "imagesinfo" edges to the Imageinfo entity.
func (uu *UserUpdate) AddImagesinfo(i ...*Imageinfo) *UserUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.AddImagesinfoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (uu *UserUpdate) ClearProducts() *UserUpdate {
	uu.mutation.ClearProducts()
	return uu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (uu *UserUpdate) RemoveProductIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveProductIDs(ids...)
	return uu
}

// RemoveProducts removes "products" edges to Product entities.
func (uu *UserUpdate) RemoveProducts(p ...*Product) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveProductIDs(ids...)
}

// ClearOrders clears all "orders" edges to the Order entity.
func (uu *UserUpdate) ClearOrders() *UserUpdate {
	uu.mutation.ClearOrders()
	return uu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (uu *UserUpdate) RemoveOrderIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveOrderIDs(ids...)
	return uu
}

// RemoveOrders removes "orders" edges to Order entities.
func (uu *UserUpdate) RemoveOrders(o ...*Order) *UserUpdate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uu.RemoveOrderIDs(ids...)
}

// ClearSiteui clears the "siteui" edge to the Siteui entity.
func (uu *UserUpdate) ClearSiteui() *UserUpdate {
	uu.mutation.ClearSiteui()
	return uu
}

// ClearImagesinfo clears all "imagesinfo" edges to the Imageinfo entity.
func (uu *UserUpdate) ClearImagesinfo() *UserUpdate {
	uu.mutation.ClearImagesinfo()
	return uu
}

// RemoveImagesinfoIDs removes the "imagesinfo" edge to Imageinfo entities by IDs.
func (uu *UserUpdate) RemoveImagesinfoIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveImagesinfoIDs(ids...)
	return uu
}

// RemoveImagesinfo removes "imagesinfo" edges to Imageinfo entities.
func (uu *UserUpdate) RemoveImagesinfo(i ...*Imageinfo) *UserUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.RemoveImagesinfoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.HashedPw(); ok {
		if err := user.HashedPwValidator(v); err != nil {
			return &ValidationError{Name: "hashed_pw", err: fmt.Errorf(`ent: validator failed for field "User.hashed_pw": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.HashedPw(); ok {
		_spec.SetField(user.FieldHashedPw, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsArchived(); ok {
		_spec.SetField(user.FieldIsArchived, field.TypeBool, value)
	}
	if uu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !uu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !uu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SiteuiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SiteuiTable,
			Columns: []string{user.SiteuiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: siteui.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SiteuiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SiteuiTable,
			Columns: []string{user.SiteuiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: siteui.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ImagesinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedImagesinfoIDs(); len(nodes) > 0 && !uu.mutation.ImagesinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ImagesinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetHashedPw sets the "hashed_pw" field.
func (uuo *UserUpdateOne) SetHashedPw(s string) *UserUpdateOne {
	uuo.mutation.SetHashedPw(s)
	return uuo
}

// SetEmailVerified sets the "email_verified" field.
func (uuo *UserUpdateOne) SetEmailVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetEmailVerified(b)
	return uuo
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetEmailVerified(*b)
	}
	return uuo
}

// SetIsArchived sets the "is_archived" field.
func (uuo *UserUpdateOne) SetIsArchived(b bool) *UserUpdateOne {
	uuo.mutation.SetIsArchived(b)
	return uuo
}

// SetNillableIsArchived sets the "is_archived" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsArchived(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsArchived(*b)
	}
	return uuo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (uuo *UserUpdateOne) AddProductIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddProductIDs(ids...)
	return uuo
}

// AddProducts adds the "products" edges to the Product entity.
func (uuo *UserUpdateOne) AddProducts(p ...*Product) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddProductIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (uuo *UserUpdateOne) AddOrderIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddOrderIDs(ids...)
	return uuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (uuo *UserUpdateOne) AddOrders(o ...*Order) *UserUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.AddOrderIDs(ids...)
}

// SetSiteuiID sets the "siteui" edge to the Siteui entity by ID.
func (uuo *UserUpdateOne) SetSiteuiID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetSiteuiID(id)
	return uuo
}

// SetNillableSiteuiID sets the "siteui" edge to the Siteui entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSiteuiID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSiteuiID(*id)
	}
	return uuo
}

// SetSiteui sets the "siteui" edge to the Siteui entity.
func (uuo *UserUpdateOne) SetSiteui(s *Siteui) *UserUpdateOne {
	return uuo.SetSiteuiID(s.ID)
}

// AddImagesinfoIDs adds the "imagesinfo" edge to the Imageinfo entity by IDs.
func (uuo *UserUpdateOne) AddImagesinfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddImagesinfoIDs(ids...)
	return uuo
}

// AddImagesinfo adds the "imagesinfo" edges to the Imageinfo entity.
func (uuo *UserUpdateOne) AddImagesinfo(i ...*Imageinfo) *UserUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.AddImagesinfoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (uuo *UserUpdateOne) ClearProducts() *UserUpdateOne {
	uuo.mutation.ClearProducts()
	return uuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (uuo *UserUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveProductIDs(ids...)
	return uuo
}

// RemoveProducts removes "products" edges to Product entities.
func (uuo *UserUpdateOne) RemoveProducts(p ...*Product) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveProductIDs(ids...)
}

// ClearOrders clears all "orders" edges to the Order entity.
func (uuo *UserUpdateOne) ClearOrders() *UserUpdateOne {
	uuo.mutation.ClearOrders()
	return uuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (uuo *UserUpdateOne) RemoveOrderIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveOrderIDs(ids...)
	return uuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (uuo *UserUpdateOne) RemoveOrders(o ...*Order) *UserUpdateOne {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uuo.RemoveOrderIDs(ids...)
}

// ClearSiteui clears the "siteui" edge to the Siteui entity.
func (uuo *UserUpdateOne) ClearSiteui() *UserUpdateOne {
	uuo.mutation.ClearSiteui()
	return uuo
}

// ClearImagesinfo clears all "imagesinfo" edges to the Imageinfo entity.
func (uuo *UserUpdateOne) ClearImagesinfo() *UserUpdateOne {
	uuo.mutation.ClearImagesinfo()
	return uuo
}

// RemoveImagesinfoIDs removes the "imagesinfo" edge to Imageinfo entities by IDs.
func (uuo *UserUpdateOne) RemoveImagesinfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveImagesinfoIDs(ids...)
	return uuo
}

// RemoveImagesinfo removes "imagesinfo" edges to Imageinfo entities.
func (uuo *UserUpdateOne) RemoveImagesinfo(i ...*Imageinfo) *UserUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.RemoveImagesinfoIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.HashedPw(); ok {
		if err := user.HashedPwValidator(v); err != nil {
			return &ValidationError{Name: "hashed_pw", err: fmt.Errorf(`ent: validator failed for field "User.hashed_pw": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.HashedPw(); ok {
		_spec.SetField(user.FieldHashedPw, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsArchived(); ok {
		_spec.SetField(user.FieldIsArchived, field.TypeBool, value)
	}
	if uuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !uuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProductsTable,
			Columns: []string{user.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !uuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SiteuiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SiteuiTable,
			Columns: []string{user.SiteuiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: siteui.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SiteuiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SiteuiTable,
			Columns: []string{user.SiteuiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: siteui.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ImagesinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedImagesinfoIDs(); len(nodes) > 0 && !uuo.mutation.ImagesinfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ImagesinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ImagesinfoTable,
			Columns: []string{user.ImagesinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imageinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
