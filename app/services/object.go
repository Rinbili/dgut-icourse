package services

import (
	"context"
	"dgut-icourse/app/utils"
	"dgut-icourse/ent"
	"dgut-icourse/ent/object"
	"fmt"
	"github.com/google/uuid"
)

// CreateObject
// @Description: 创建对象
// @param ctx
// @param obj
// @return *ent.Object
// @return error
func CreateObject(ctx context.Context, obj ent.Object) (object *ent.Object, err error) {
	if err = WithTx(ctx, Client, func(tx *ent.Tx) error {
		var err error
		if object, err = tx.Object.Create().
			SetName(obj.Name).
			SetDesc(obj.Desc).
			SetType(obj.Type).
			Save(ctx); err != nil {
			return err
		}
		switch obj.Type {
		case 0:
			if _, err = tx.Person.Create().
				SetObject(object).Save(ctx); err != nil {
				return err
			}
			break
		case 1:
			if _, err = tx.Organization.Create().
				SetAddress(obj.Edges.Organization.Address).
				SetType(obj.Edges.Organization.Type).
				SetObject(object).Save(ctx); err != nil {
				return err
			}
			break
		case 2:
			if _, err = tx.Course.Create().
				SetCourseID(obj.Edges.Course.CourseID).
				SetTeacherID(obj.Edges.Course.Edges.Teacher.ID).
				SetOrganizationID(obj.Edges.Course.Edges.Organization.ID).
				SetObject(object).Save(ctx); err != nil {
				return err
			}
			break
		case 3:
			if _, err = tx.Other.Create().
				SetObject(object).Save(ctx); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid object type: %d", obj.Type)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return
}

// GetObjectByUUID
// @Description: 通过UUID获取对象
// @param ctx
// @param uuidString
// @return *ent.Object
// @return error
func GetObjectByUUID(ctx context.Context, uuidString string) (obj *ent.Object, err error) {
	if oid, err := uuid.Parse(uuidString); err != nil {
		return nil, err
	} else {
		// TODO: 优化查询
		return Client.Object.Query().
			Where(object.ID(oid)).
			WithCourse(func(q *ent.CourseQuery) {
				q.WithOrganization(func(query *ent.ObjectQuery) {
					query.WithOrganization()
				})
				q.WithTeacher(func(q *ent.ObjectQuery) {
					q.WithPerson()
				})
			}).
			WithOrganization().
			WithPerson().
			WithOther().
			WithTeachCourses().
			WithSetupCourses().
			WithComments().
			Only(ctx)
	}
}

// GetObjects
// @Description: 获取对象列表
// @param ctx
// @param paging
// @return []*ent.Object
// @return error
func GetObjects(ctx context.Context, paging utils.Paging) (objs []*ent.Object, err error) {
	return Client.Object.Query().
		Offset(paging.Offset()).
		Order(ent.Desc(object.FieldUpdatedAt)).
		WithPerson().
		WithOrganization().
		WithCourse(func(q *ent.CourseQuery) {
			q.WithOrganization(func(query *ent.ObjectQuery) {
				query.WithOrganization()
			})
			q.WithTeacher(func(q *ent.ObjectQuery) {
				q.WithPerson()
			})
		}).
		WithOther().
		WithComments().
		Limit(paging.Limit()).
		All(ctx)
}

// GetObjectsByType
// @Description: 通过类型获取对象列表
// @param ctx
// @param objType
// @return []*ent.Object
// @return error
func GetObjectsByType(ctx context.Context, objType int) (objs []*ent.Object, err error) {
	return Client.Object.Query().Where(object.Type(int8(objType))).
		WithCourse(func(q *ent.CourseQuery) {
			q.WithOrganization(func(query *ent.ObjectQuery) {
				query.WithOrganization()
			})
			q.WithTeacher(func(q *ent.ObjectQuery) {
				q.WithPerson()
			})
		}).
		All(ctx)
}
