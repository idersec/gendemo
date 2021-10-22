// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"gendemo/dal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newPassport(db *gorm.DB) passport {
	_passport := passport{}

	_passport.passportDo.UseDB(db)
	_passport.passportDo.UseModel(&model.Passport{})

	tableName := _passport.passportDo.TableName()
	_passport.ALL = field.NewField(tableName, "*")
	_passport.ID = field.NewInt(tableName, "id")
	_passport.Username = field.NewString(tableName, "username")
	_passport.Password = field.NewField(tableName, "password")
	_passport.LoginTime = field.NewTime(tableName, "login_time")

	return _passport
}

type passport struct {
	passportDo passportDo

	ALL       field.Field
	ID        field.Int
	Username  field.String
	Password  field.Field
	LoginTime field.Time
}

func (p *passport) WithContext(ctx context.Context) *passportDo { return p.passportDo.WithContext(ctx) }

func (p passport) TableName() string { return p.passportDo.TableName() }

func (p passport) clone(db *gorm.DB) passport {
	p.passportDo.ReplaceDB(db)
	return p
}

type passportDo struct{ gen.DO }

func (p passportDo) Debug() *passportDo {
	return p.withDO(p.DO.Debug())
}

func (p passportDo) WithContext(ctx context.Context) *passportDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passportDo) Clauses(conds ...clause.Expression) *passportDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passportDo) Not(conds ...gen.Condition) *passportDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passportDo) Or(conds ...gen.Condition) *passportDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passportDo) Select(conds ...field.Expr) *passportDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passportDo) Where(conds ...gen.Condition) *passportDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passportDo) Order(conds ...field.Expr) *passportDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passportDo) Distinct(cols ...field.Expr) *passportDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passportDo) Omit(cols ...field.Expr) *passportDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passportDo) Join(table schema.Tabler, on ...field.Expr) *passportDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *passportDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passportDo) RightJoin(table schema.Tabler, on ...field.Expr) *passportDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passportDo) Group(cols ...field.Expr) *passportDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passportDo) Having(conds ...gen.Condition) *passportDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passportDo) Limit(limit int) *passportDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passportDo) Offset(offset int) *passportDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *passportDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passportDo) Unscoped() *passportDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passportDo) Create(values ...*model.Passport) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passportDo) CreateInBatches(values []*model.Passport, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passportDo) Save(values ...*model.Passport) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passportDo) First() (*model.Passport, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passport), nil
	}
}

func (p passportDo) Take() (*model.Passport, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passport), nil
	}
}

func (p passportDo) Last() (*model.Passport, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passport), nil
	}
}

func (p passportDo) Find() ([]*model.Passport, error) {
	result, err := p.DO.Find()
	return result.([]*model.Passport), err
}

func (p passportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Passport, error) {
	result, err := p.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Passport), err
}

func (p passportDo) FindInBatches(result []*model.Passport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(&result, batchSize, fc)
}

func (p passportDo) Attrs(attrs ...field.AssignExpr) *passportDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passportDo) Assign(attrs ...field.AssignExpr) *passportDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passportDo) Joins(field field.RelationField) *passportDo {
	return p.withDO(p.DO.Joins(field))
}

func (p passportDo) Preload(field field.RelationField) *passportDo {
	return p.withDO(p.DO.Preload(field))
}

func (p passportDo) FirstOrInit() (*model.Passport, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passport), nil
	}
}

func (p passportDo) FirstOrCreate() (*model.Passport, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Passport), nil
	}
}

func (p passportDo) FindByPage(offset int, limit int) (result []*model.Passport, count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	result, err = p.Offset(offset).Limit(limit).Find()
	return
}

func (p *passportDo) withDO(do gen.Dao) *passportDo {
	p.DO = *do.(*gen.DO)
	return p
}
