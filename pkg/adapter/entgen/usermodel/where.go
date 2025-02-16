// Code generated by ent, DO NOT EDIT.

package usermodel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUpdatedAt, v))
}

// CognitoID applies equality check predicate on the "cognito_id" field. It's identical to CognitoIDEQ.
func CognitoID(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldCognitoID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldName, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUserName, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldBirthDate, v))
}

// SelfIntroduction applies equality check predicate on the "self_introduction" field. It's identical to SelfIntroductionEQ.
func SelfIntroduction(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldSelfIntroduction, v))
}

// ProfileImage applies equality check predicate on the "profile_image" field. It's identical to ProfileImageEQ.
func ProfileImage(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldProfileImage, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldNotNull(FieldUpdatedAt))
}

// CognitoIDEQ applies the EQ predicate on the "cognito_id" field.
func CognitoIDEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldCognitoID, v))
}

// CognitoIDNEQ applies the NEQ predicate on the "cognito_id" field.
func CognitoIDNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldCognitoID, v))
}

// CognitoIDIn applies the In predicate on the "cognito_id" field.
func CognitoIDIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldCognitoID, vs...))
}

// CognitoIDNotIn applies the NotIn predicate on the "cognito_id" field.
func CognitoIDNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldCognitoID, vs...))
}

// CognitoIDGT applies the GT predicate on the "cognito_id" field.
func CognitoIDGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldCognitoID, v))
}

// CognitoIDGTE applies the GTE predicate on the "cognito_id" field.
func CognitoIDGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldCognitoID, v))
}

// CognitoIDLT applies the LT predicate on the "cognito_id" field.
func CognitoIDLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldCognitoID, v))
}

// CognitoIDLTE applies the LTE predicate on the "cognito_id" field.
func CognitoIDLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldCognitoID, v))
}

// CognitoIDContains applies the Contains predicate on the "cognito_id" field.
func CognitoIDContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldCognitoID, v))
}

// CognitoIDHasPrefix applies the HasPrefix predicate on the "cognito_id" field.
func CognitoIDHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldCognitoID, v))
}

// CognitoIDHasSuffix applies the HasSuffix predicate on the "cognito_id" field.
func CognitoIDHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldCognitoID, v))
}

// CognitoIDEqualFold applies the EqualFold predicate on the "cognito_id" field.
func CognitoIDEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldCognitoID, v))
}

// CognitoIDContainsFold applies the ContainsFold predicate on the "cognito_id" field.
func CognitoIDContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldCognitoID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldName, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldUserName, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldBirthDate, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldGender, vs...))
}

// GenderIsNil applies the IsNil predicate on the "gender" field.
func GenderIsNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldIsNull(FieldGender))
}

// GenderNotNil applies the NotNil predicate on the "gender" field.
func GenderNotNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldNotNull(FieldGender))
}

// SelfIntroductionEQ applies the EQ predicate on the "self_introduction" field.
func SelfIntroductionEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldSelfIntroduction, v))
}

// SelfIntroductionNEQ applies the NEQ predicate on the "self_introduction" field.
func SelfIntroductionNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldSelfIntroduction, v))
}

// SelfIntroductionIn applies the In predicate on the "self_introduction" field.
func SelfIntroductionIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldSelfIntroduction, vs...))
}

// SelfIntroductionNotIn applies the NotIn predicate on the "self_introduction" field.
func SelfIntroductionNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldSelfIntroduction, vs...))
}

// SelfIntroductionGT applies the GT predicate on the "self_introduction" field.
func SelfIntroductionGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldSelfIntroduction, v))
}

// SelfIntroductionGTE applies the GTE predicate on the "self_introduction" field.
func SelfIntroductionGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldSelfIntroduction, v))
}

// SelfIntroductionLT applies the LT predicate on the "self_introduction" field.
func SelfIntroductionLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldSelfIntroduction, v))
}

// SelfIntroductionLTE applies the LTE predicate on the "self_introduction" field.
func SelfIntroductionLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldSelfIntroduction, v))
}

// SelfIntroductionContains applies the Contains predicate on the "self_introduction" field.
func SelfIntroductionContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldSelfIntroduction, v))
}

// SelfIntroductionHasPrefix applies the HasPrefix predicate on the "self_introduction" field.
func SelfIntroductionHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldSelfIntroduction, v))
}

// SelfIntroductionHasSuffix applies the HasSuffix predicate on the "self_introduction" field.
func SelfIntroductionHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldSelfIntroduction, v))
}

// SelfIntroductionIsNil applies the IsNil predicate on the "self_introduction" field.
func SelfIntroductionIsNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldIsNull(FieldSelfIntroduction))
}

// SelfIntroductionNotNil applies the NotNil predicate on the "self_introduction" field.
func SelfIntroductionNotNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldNotNull(FieldSelfIntroduction))
}

// SelfIntroductionEqualFold applies the EqualFold predicate on the "self_introduction" field.
func SelfIntroductionEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldSelfIntroduction, v))
}

// SelfIntroductionContainsFold applies the ContainsFold predicate on the "self_introduction" field.
func SelfIntroductionContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldSelfIntroduction, v))
}

// ProfileImageEQ applies the EQ predicate on the "profile_image" field.
func ProfileImageEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldProfileImage, v))
}

// ProfileImageNEQ applies the NEQ predicate on the "profile_image" field.
func ProfileImageNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldProfileImage, v))
}

// ProfileImageIn applies the In predicate on the "profile_image" field.
func ProfileImageIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldProfileImage, vs...))
}

// ProfileImageNotIn applies the NotIn predicate on the "profile_image" field.
func ProfileImageNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldProfileImage, vs...))
}

// ProfileImageGT applies the GT predicate on the "profile_image" field.
func ProfileImageGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldProfileImage, v))
}

// ProfileImageGTE applies the GTE predicate on the "profile_image" field.
func ProfileImageGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldProfileImage, v))
}

// ProfileImageLT applies the LT predicate on the "profile_image" field.
func ProfileImageLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldProfileImage, v))
}

// ProfileImageLTE applies the LTE predicate on the "profile_image" field.
func ProfileImageLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldProfileImage, v))
}

// ProfileImageContains applies the Contains predicate on the "profile_image" field.
func ProfileImageContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldProfileImage, v))
}

// ProfileImageHasPrefix applies the HasPrefix predicate on the "profile_image" field.
func ProfileImageHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldProfileImage, v))
}

// ProfileImageHasSuffix applies the HasSuffix predicate on the "profile_image" field.
func ProfileImageHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldProfileImage, v))
}

// ProfileImageIsNil applies the IsNil predicate on the "profile_image" field.
func ProfileImageIsNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldIsNull(FieldProfileImage))
}

// ProfileImageNotNil applies the NotNil predicate on the "profile_image" field.
func ProfileImageNotNil() predicate.UserModel {
	return predicate.UserModel(sql.FieldNotNull(FieldProfileImage))
}

// ProfileImageEqualFold applies the EqualFold predicate on the "profile_image" field.
func ProfileImageEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldProfileImage, v))
}

// ProfileImageContainsFold applies the ContainsFold predicate on the "profile_image" field.
func ProfileImageContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldProfileImage, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldEmail, v))
}

// HasFollowers applies the HasEdge predicate on the "followers" edge.
func HasFollowers() predicate.UserModel {
	return predicate.UserModel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FollowersTable, FollowersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFollowersWith applies the HasEdge predicate on the "followers" edge with a given conditions (other predicates).
func HasFollowersWith(preds ...predicate.FollowsModel) predicate.UserModel {
	return predicate.UserModel(func(s *sql.Selector) {
		step := newFollowersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFollowees applies the HasEdge predicate on the "followees" edge.
func HasFollowees() predicate.UserModel {
	return predicate.UserModel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FolloweesTable, FolloweesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFolloweesWith applies the HasEdge predicate on the "followees" edge with a given conditions (other predicates).
func HasFolloweesWith(preds ...predicate.FollowsModel) predicate.UserModel {
	return predicate.UserModel(func(s *sql.Selector) {
		step := newFolloweesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.NotPredicates(p))
}
