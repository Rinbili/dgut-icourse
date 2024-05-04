package services

import (
	"context"
	"dgut-icourse/ent"
	"dgut-icourse/ent/comment"
	"dgut-icourse/ent/object"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/google/uuid"
)

// CreateComment
// @Description: 创建评论
// @param ctx: context.Context
// @param source: ent.Comment
// @return *ent.Comment: 评论信息
// @return error: err
func CreateComment(ctx context.Context, source ent.Comment) (comment *ent.Comment, err error) {
	if user := ctx.Value("user").(*ent.User); user == nil {
		return nil, fmt.Errorf("invalid user")
	} else {
		if err = WithTx(ctx, Client, func(tx *ent.Tx) error {
			var err error
			if comment, err = tx.Comment.Create().
				SetScore(source.Score).
				SetContent(source.Content).
				SetObjectID(source.Edges.Object.ID).
				SetAuthorID(user.ID).
				Save(ctx); err != nil {
				return err
			}
			obj, _ := tx.Object.Get(ctx, source.Edges.Object.ID)
			switch obj.Type {
			// TODO: Add other cases
			case 2:
				if _, err = tx.CourseComment.Create().
					SetDifficulty(source.Edges.CourseComment.Difficulty).
					SetMark(source.Edges.CourseComment.Mark).
					SetQuality(source.Edges.CourseComment.Quality).
					SetWorkload(source.Edges.CourseComment.Workload).
					SetComment(comment).
					Save(ctx); err != nil {
					return err
				}
				break
			default:
				return fmt.Errorf("invalid object type: %d", obj.Type)
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}
	return
}

// GetCommentByUUID
// @Description: 通过CID获取评论信息
// @param ctx: context.Context
// @param cidStr: string
// @return *ent.Comment: 评论信息
// @return error: err
func GetCommentByUUID(ctx context.Context, cidStr string) (c *ent.Comment, err error) {
	if id, err := uuid.Parse(cidStr); err != nil {
		return nil, err
	} else {
		return Client.Comment.Query().
			Where(comment.ID(id)).
			WithAuthor().
			WithObject().
			WithCourseComment().
			WithParent().
			WithChildren().
			Only(ctx)
	}
}

func GetMyComments(ctx context.Context) (comments []*ent.Comment, err error) {
	u := ctx.Value("user").(*ent.User)
	if u == nil {
		return nil, fmt.Errorf("invalid user")
	}
	return u.QueryComments().WithParent().WithObject().All(ctx)
}

// GetCommentsByObjectUUID
// @Description: 通过OID获取评论列表
// @param ctx: context.Context
// @param oidStr: string
// @return []*ent.Comment: 评论列表
// @return error: err
func GetCommentsByObjectUUID(ctx context.Context, oidStr string) (comments []*ent.Comment, err error) {
	if oid, err := uuid.Parse(oidStr); err != nil {
		return nil, err
	} else {
		if comments, err = Client.Comment.Query().
			Where(comment.HasObjectWith(object.ID(oid))).
			WithAuthor().
			WithLikedUsers().
			WithCourseComment().
			All(ctx); err != nil {
			return nil, err
		}
		return comments, nil
	}
}

// GetCommentsByTime
// @Description: 通过时间获取评论列表
// @param ctx: context.Context
// @param offset: int
// @param limit: int
func GetCommentsByTime(ctx context.Context, offset int, limit int) (comments []*ent.Comment, err error) {
	return Client.Comment.Query().
		Where(comment.HasObject()).
		Order(comment.ByCreatedAt(sql.OrderDesc())).
		Offset(offset).
		WithAuthor().
		WithLikedUsers().
		WithCourseComment().
		WithObject().
		Limit(limit).
		All(ctx)
}

func UpdateComment(ctx context.Context, source ent.Comment) (c *ent.Comment, err error) {
	if user := ctx.Value("user").(*ent.User); user == nil {
		return nil, fmt.Errorf("invalid user")
	} else {
		if err = WithTx(ctx, Client, func(tx *ent.Tx) (err error) {
			c, err = tx.Comment.Query().
				Where(comment.ID(source.ID)).
				WithAuthor().
				WithObject().
				WithCourseComment().
				Only(ctx)
			if err != nil {
				return err
			}
			if c.Edges.Author.ID != user.ID {
				return fmt.Errorf("invalid user")
			} else {
				switch c.Edges.Object.Type {
				case 2:
					uo := tx.CourseComment.UpdateOne(c.Edges.CourseComment)
					if source.Edges.CourseComment != nil {
						if source.Edges.CourseComment.Difficulty != 0 {
							uo.SetDifficulty(source.Edges.CourseComment.Difficulty)
						}
						if source.Edges.CourseComment.Mark != 0 {
							uo.SetMark(source.Edges.CourseComment.Mark)
						}
						if source.Edges.CourseComment.Quality != 0 {
							uo.SetQuality(source.Edges.CourseComment.Quality)
						}
						if source.Edges.CourseComment.Workload != 0 {
							uo.SetWorkload(source.Edges.CourseComment.Workload)
						}
					}
					_, err = uo.Save(ctx)
					break
				default:
					return fmt.Errorf("invalid object type: %d", c.Edges.Object.Type)
				}
				c, err = tx.Comment.UpdateOne(c).
					SetContent(source.Content).
					SetScore(source.Score).
					Save(ctx)
				return err
			}
		}); err != nil {
			return nil, err
		}
	}
	return
}

// LikeComment
// @Description: 点赞评论
// @param ctx: context.Context
// @param cidStr: string
// @return error: err
func LikeComment(ctx context.Context, cidStr string) (err error) {
	user := ctx.Value("user").(*ent.User)
	if user == nil {
		return fmt.Errorf("invalid user")
	}
	if id, err := uuid.Parse(cidStr); err != nil {
		return err
	} else {
		return WithTx(ctx, Client, func(tx *ent.Tx) error {
			if _, err := tx.Comment.Update().
				Where(comment.ID(id)).
				AddLikedUsers(user).
				Save(ctx); err != nil {
				return err
			}
			return nil
		})
	}
}

// UnlikeComment
// @Description: 取消点赞评论
// @param ctx: context.Context
// @param cidStr: string
// @return error: err
func UnlikeComment(ctx context.Context, cidStr string) (err error) {
	user := ctx.Value("user").(*ent.User)
	if user == nil {
		return fmt.Errorf("invalid user")
	}
	if id, err := uuid.Parse(cidStr); err != nil {
		return err
	} else {
		return WithTx(ctx, Client, func(tx *ent.Tx) error {
			if _, err := tx.Comment.Update().
				Where(comment.ID(id)).
				RemoveLikedUsers(user).
				Save(ctx); err != nil {
				return err
			}
			return nil
		})
	}
}
