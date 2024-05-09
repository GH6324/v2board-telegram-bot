///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package v2_telegram_checkin

import (
	"fmt"
	"time"
	"v2board-telegram-bot/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *V2TelegramCheckin {
	return new(V2TelegramCheckin)
}

func NewQueryBuilder() *v2TelegramCheckinQueryBuilder {
	return new(v2TelegramCheckinQueryBuilder)
}

func (t *V2TelegramCheckin) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type v2TelegramCheckinQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *v2TelegramCheckinQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	if qb.limit > 0 {
		ret = ret.Limit(qb.limit).Offset(qb.offset)
	}

	return ret
}

func (qb *v2TelegramCheckinQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&V2TelegramCheckin{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *v2TelegramCheckinQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&V2TelegramCheckin{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *v2TelegramCheckinQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&V2TelegramCheckin{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *v2TelegramCheckinQueryBuilder) First(db *gorm.DB) (*V2TelegramCheckin, error) {
	ret := &V2TelegramCheckin{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *v2TelegramCheckinQueryBuilder) QueryOne(db *gorm.DB) (*V2TelegramCheckin, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *v2TelegramCheckinQueryBuilder) QueryAll(db *gorm.DB) ([]*V2TelegramCheckin, error) {
	var ret []*V2TelegramCheckin
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *v2TelegramCheckinQueryBuilder) Limit(limit int) *v2TelegramCheckinQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) Offset(offset int) *v2TelegramCheckinQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereId(p mysql.Predicate, value int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereIdIn(value []int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereIdNotIn(value []int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderById(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserTgId(p mysql.Predicate, value int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_tg_id", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserTgIdIn(value []int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_tg_id", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserTgIdNotIn(value []int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_tg_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderByUserTgId(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_tg_id "+order)
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserId(p mysql.Predicate, value int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserIdIn(value []int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUserIdNotIn(value []int32) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderByUserId(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereAward(p mysql.Predicate, value int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereAwardIn(value []int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereAwardNotIn(value []int64) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "award", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderByAward(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "award "+order)
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereCreatedAtIn(value []time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderByCreatedAt(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUpdatedAtIn(value []time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *v2TelegramCheckinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *v2TelegramCheckinQueryBuilder) OrderByUpdatedAt(asc bool) *v2TelegramCheckinQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}