// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/predicate"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/usermodel"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUserModel = "UserModel"
)

// UserModelMutation represents an operation that mutates the UserModel nodes in the graph.
type UserModelMutation struct {
	config
	op                Op
	typ               string
	id                *int64
	created_at        *time.Time
	updated_at        *time.Time
	cognito_id        *string
	name              *string
	user_name         *string
	birth_date        *time.Time
	gender            *usermodel.Gender
	self_introduction *string
	profile_image     *string
	email             *string
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*UserModel, error)
	predicates        []predicate.UserModel
}

var _ ent.Mutation = (*UserModelMutation)(nil)

// usermodelOption allows management of the mutation configuration using functional options.
type usermodelOption func(*UserModelMutation)

// newUserModelMutation creates new mutation for the UserModel entity.
func newUserModelMutation(c config, op Op, opts ...usermodelOption) *UserModelMutation {
	m := &UserModelMutation{
		config:        c,
		op:            op,
		typ:           TypeUserModel,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserModelID sets the ID field of the mutation.
func withUserModelID(id int64) usermodelOption {
	return func(m *UserModelMutation) {
		var (
			err   error
			once  sync.Once
			value *UserModel
		)
		m.oldValue = func(ctx context.Context) (*UserModel, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserModel.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserModel sets the old UserModel of the mutation.
func withUserModel(node *UserModel) usermodelOption {
	return func(m *UserModelMutation) {
		m.oldValue = func(context.Context) (*UserModel, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserModelMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserModelMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("entgen: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of UserModel entities.
func (m *UserModelMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserModelMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserModelMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserModel.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserModelMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserModelMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ClearCreatedAt clears the value of the "created_at" field.
func (m *UserModelMutation) ClearCreatedAt() {
	m.created_at = nil
	m.clearedFields[usermodel.FieldCreatedAt] = struct{}{}
}

// CreatedAtCleared returns if the "created_at" field was cleared in this mutation.
func (m *UserModelMutation) CreatedAtCleared() bool {
	_, ok := m.clearedFields[usermodel.FieldCreatedAt]
	return ok
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserModelMutation) ResetCreatedAt() {
	m.created_at = nil
	delete(m.clearedFields, usermodel.FieldCreatedAt)
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserModelMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserModelMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *UserModelMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[usermodel.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *UserModelMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[usermodel.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserModelMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, usermodel.FieldUpdatedAt)
}

// SetCognitoID sets the "cognito_id" field.
func (m *UserModelMutation) SetCognitoID(s string) {
	m.cognito_id = &s
}

// CognitoID returns the value of the "cognito_id" field in the mutation.
func (m *UserModelMutation) CognitoID() (r string, exists bool) {
	v := m.cognito_id
	if v == nil {
		return
	}
	return *v, true
}

// OldCognitoID returns the old "cognito_id" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldCognitoID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCognitoID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCognitoID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCognitoID: %w", err)
	}
	return oldValue.CognitoID, nil
}

// ResetCognitoID resets all changes to the "cognito_id" field.
func (m *UserModelMutation) ResetCognitoID() {
	m.cognito_id = nil
}

// SetName sets the "name" field.
func (m *UserModelMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserModelMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserModelMutation) ResetName() {
	m.name = nil
}

// SetUserName sets the "user_name" field.
func (m *UserModelMutation) SetUserName(s string) {
	m.user_name = &s
}

// UserName returns the value of the "user_name" field in the mutation.
func (m *UserModelMutation) UserName() (r string, exists bool) {
	v := m.user_name
	if v == nil {
		return
	}
	return *v, true
}

// OldUserName returns the old "user_name" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldUserName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserName: %w", err)
	}
	return oldValue.UserName, nil
}

// ResetUserName resets all changes to the "user_name" field.
func (m *UserModelMutation) ResetUserName() {
	m.user_name = nil
}

// SetBirthDate sets the "birth_date" field.
func (m *UserModelMutation) SetBirthDate(t time.Time) {
	m.birth_date = &t
}

// BirthDate returns the value of the "birth_date" field in the mutation.
func (m *UserModelMutation) BirthDate() (r time.Time, exists bool) {
	v := m.birth_date
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthDate returns the old "birth_date" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldBirthDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBirthDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBirthDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthDate: %w", err)
	}
	return oldValue.BirthDate, nil
}

// ResetBirthDate resets all changes to the "birth_date" field.
func (m *UserModelMutation) ResetBirthDate() {
	m.birth_date = nil
}

// SetGender sets the "gender" field.
func (m *UserModelMutation) SetGender(u usermodel.Gender) {
	m.gender = &u
}

// Gender returns the value of the "gender" field in the mutation.
func (m *UserModelMutation) Gender() (r usermodel.Gender, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldGender(ctx context.Context) (v usermodel.Gender, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// ClearGender clears the value of the "gender" field.
func (m *UserModelMutation) ClearGender() {
	m.gender = nil
	m.clearedFields[usermodel.FieldGender] = struct{}{}
}

// GenderCleared returns if the "gender" field was cleared in this mutation.
func (m *UserModelMutation) GenderCleared() bool {
	_, ok := m.clearedFields[usermodel.FieldGender]
	return ok
}

// ResetGender resets all changes to the "gender" field.
func (m *UserModelMutation) ResetGender() {
	m.gender = nil
	delete(m.clearedFields, usermodel.FieldGender)
}

// SetSelfIntroduction sets the "self_introduction" field.
func (m *UserModelMutation) SetSelfIntroduction(s string) {
	m.self_introduction = &s
}

// SelfIntroduction returns the value of the "self_introduction" field in the mutation.
func (m *UserModelMutation) SelfIntroduction() (r string, exists bool) {
	v := m.self_introduction
	if v == nil {
		return
	}
	return *v, true
}

// OldSelfIntroduction returns the old "self_introduction" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldSelfIntroduction(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSelfIntroduction is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSelfIntroduction requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSelfIntroduction: %w", err)
	}
	return oldValue.SelfIntroduction, nil
}

// ClearSelfIntroduction clears the value of the "self_introduction" field.
func (m *UserModelMutation) ClearSelfIntroduction() {
	m.self_introduction = nil
	m.clearedFields[usermodel.FieldSelfIntroduction] = struct{}{}
}

// SelfIntroductionCleared returns if the "self_introduction" field was cleared in this mutation.
func (m *UserModelMutation) SelfIntroductionCleared() bool {
	_, ok := m.clearedFields[usermodel.FieldSelfIntroduction]
	return ok
}

// ResetSelfIntroduction resets all changes to the "self_introduction" field.
func (m *UserModelMutation) ResetSelfIntroduction() {
	m.self_introduction = nil
	delete(m.clearedFields, usermodel.FieldSelfIntroduction)
}

// SetProfileImage sets the "profile_image" field.
func (m *UserModelMutation) SetProfileImage(s string) {
	m.profile_image = &s
}

// ProfileImage returns the value of the "profile_image" field in the mutation.
func (m *UserModelMutation) ProfileImage() (r string, exists bool) {
	v := m.profile_image
	if v == nil {
		return
	}
	return *v, true
}

// OldProfileImage returns the old "profile_image" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldProfileImage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProfileImage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProfileImage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProfileImage: %w", err)
	}
	return oldValue.ProfileImage, nil
}

// ClearProfileImage clears the value of the "profile_image" field.
func (m *UserModelMutation) ClearProfileImage() {
	m.profile_image = nil
	m.clearedFields[usermodel.FieldProfileImage] = struct{}{}
}

// ProfileImageCleared returns if the "profile_image" field was cleared in this mutation.
func (m *UserModelMutation) ProfileImageCleared() bool {
	_, ok := m.clearedFields[usermodel.FieldProfileImage]
	return ok
}

// ResetProfileImage resets all changes to the "profile_image" field.
func (m *UserModelMutation) ResetProfileImage() {
	m.profile_image = nil
	delete(m.clearedFields, usermodel.FieldProfileImage)
}

// SetEmail sets the "email" field.
func (m *UserModelMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserModelMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the UserModel entity.
// If the UserModel object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserModelMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserModelMutation) ResetEmail() {
	m.email = nil
}

// Where appends a list predicates to the UserModelMutation builder.
func (m *UserModelMutation) Where(ps ...predicate.UserModel) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserModelMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserModelMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.UserModel, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserModelMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserModelMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (UserModel).
func (m *UserModelMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserModelMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.created_at != nil {
		fields = append(fields, usermodel.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, usermodel.FieldUpdatedAt)
	}
	if m.cognito_id != nil {
		fields = append(fields, usermodel.FieldCognitoID)
	}
	if m.name != nil {
		fields = append(fields, usermodel.FieldName)
	}
	if m.user_name != nil {
		fields = append(fields, usermodel.FieldUserName)
	}
	if m.birth_date != nil {
		fields = append(fields, usermodel.FieldBirthDate)
	}
	if m.gender != nil {
		fields = append(fields, usermodel.FieldGender)
	}
	if m.self_introduction != nil {
		fields = append(fields, usermodel.FieldSelfIntroduction)
	}
	if m.profile_image != nil {
		fields = append(fields, usermodel.FieldProfileImage)
	}
	if m.email != nil {
		fields = append(fields, usermodel.FieldEmail)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserModelMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case usermodel.FieldCreatedAt:
		return m.CreatedAt()
	case usermodel.FieldUpdatedAt:
		return m.UpdatedAt()
	case usermodel.FieldCognitoID:
		return m.CognitoID()
	case usermodel.FieldName:
		return m.Name()
	case usermodel.FieldUserName:
		return m.UserName()
	case usermodel.FieldBirthDate:
		return m.BirthDate()
	case usermodel.FieldGender:
		return m.Gender()
	case usermodel.FieldSelfIntroduction:
		return m.SelfIntroduction()
	case usermodel.FieldProfileImage:
		return m.ProfileImage()
	case usermodel.FieldEmail:
		return m.Email()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserModelMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case usermodel.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case usermodel.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case usermodel.FieldCognitoID:
		return m.OldCognitoID(ctx)
	case usermodel.FieldName:
		return m.OldName(ctx)
	case usermodel.FieldUserName:
		return m.OldUserName(ctx)
	case usermodel.FieldBirthDate:
		return m.OldBirthDate(ctx)
	case usermodel.FieldGender:
		return m.OldGender(ctx)
	case usermodel.FieldSelfIntroduction:
		return m.OldSelfIntroduction(ctx)
	case usermodel.FieldProfileImage:
		return m.OldProfileImage(ctx)
	case usermodel.FieldEmail:
		return m.OldEmail(ctx)
	}
	return nil, fmt.Errorf("unknown UserModel field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserModelMutation) SetField(name string, value ent.Value) error {
	switch name {
	case usermodel.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case usermodel.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case usermodel.FieldCognitoID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCognitoID(v)
		return nil
	case usermodel.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case usermodel.FieldUserName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserName(v)
		return nil
	case usermodel.FieldBirthDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthDate(v)
		return nil
	case usermodel.FieldGender:
		v, ok := value.(usermodel.Gender)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case usermodel.FieldSelfIntroduction:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSelfIntroduction(v)
		return nil
	case usermodel.FieldProfileImage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProfileImage(v)
		return nil
	case usermodel.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	}
	return fmt.Errorf("unknown UserModel field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserModelMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserModelMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserModelMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserModel numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserModelMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(usermodel.FieldCreatedAt) {
		fields = append(fields, usermodel.FieldCreatedAt)
	}
	if m.FieldCleared(usermodel.FieldUpdatedAt) {
		fields = append(fields, usermodel.FieldUpdatedAt)
	}
	if m.FieldCleared(usermodel.FieldGender) {
		fields = append(fields, usermodel.FieldGender)
	}
	if m.FieldCleared(usermodel.FieldSelfIntroduction) {
		fields = append(fields, usermodel.FieldSelfIntroduction)
	}
	if m.FieldCleared(usermodel.FieldProfileImage) {
		fields = append(fields, usermodel.FieldProfileImage)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserModelMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserModelMutation) ClearField(name string) error {
	switch name {
	case usermodel.FieldCreatedAt:
		m.ClearCreatedAt()
		return nil
	case usermodel.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	case usermodel.FieldGender:
		m.ClearGender()
		return nil
	case usermodel.FieldSelfIntroduction:
		m.ClearSelfIntroduction()
		return nil
	case usermodel.FieldProfileImage:
		m.ClearProfileImage()
		return nil
	}
	return fmt.Errorf("unknown UserModel nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserModelMutation) ResetField(name string) error {
	switch name {
	case usermodel.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case usermodel.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case usermodel.FieldCognitoID:
		m.ResetCognitoID()
		return nil
	case usermodel.FieldName:
		m.ResetName()
		return nil
	case usermodel.FieldUserName:
		m.ResetUserName()
		return nil
	case usermodel.FieldBirthDate:
		m.ResetBirthDate()
		return nil
	case usermodel.FieldGender:
		m.ResetGender()
		return nil
	case usermodel.FieldSelfIntroduction:
		m.ResetSelfIntroduction()
		return nil
	case usermodel.FieldProfileImage:
		m.ResetProfileImage()
		return nil
	case usermodel.FieldEmail:
		m.ResetEmail()
		return nil
	}
	return fmt.Errorf("unknown UserModel field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserModelMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserModelMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserModelMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserModelMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserModelMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserModelMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserModelMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown UserModel unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserModelMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown UserModel edge %s", name)
}