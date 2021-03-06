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

func newBank(db *gorm.DB) bank {
	_bank := bank{}

	_bank.bankDo.UseDB(db)
	_bank.bankDo.UseModel(&model.Bank{})

	tableName := _bank.bankDo.TableName()
	_bank.ALL = field.NewField(tableName, "*")
	_bank.ID = field.NewUint(tableName, "id")
	_bank.Name = field.NewString(tableName, "name")
	_bank.Address = field.NewString(tableName, "address")
	_bank.Scale = field.NewInt(tableName, "scale")

	return _bank
}

type bank struct {
	bankDo bankDo

	ALL     field.Field
	ID      field.Uint
	Name    field.String
	Address field.String
	Scale   field.Int
}

func (b *bank) WithContext(ctx context.Context) *bankDo { return b.bankDo.WithContext(ctx) }

func (b bank) TableName() string { return b.bankDo.TableName() }

func (b bank) clone(db *gorm.DB) bank {
	b.bankDo.ReplaceDB(db)
	return b
}

type bankDo struct{ gen.DO }

func (b bankDo) Debug() *bankDo {
	return b.withDO(b.DO.Debug())
}

func (b bankDo) WithContext(ctx context.Context) *bankDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bankDo) Clauses(conds ...clause.Expression) *bankDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bankDo) Not(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bankDo) Or(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bankDo) Select(conds ...field.Expr) *bankDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bankDo) Where(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bankDo) Order(conds ...field.Expr) *bankDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bankDo) Distinct(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bankDo) Omit(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bankDo) Join(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bankDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bankDo) RightJoin(table schema.Tabler, on ...field.Expr) *bankDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bankDo) Group(cols ...field.Expr) *bankDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bankDo) Having(conds ...gen.Condition) *bankDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bankDo) Limit(limit int) *bankDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bankDo) Offset(offset int) *bankDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bankDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bankDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bankDo) Unscoped() *bankDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bankDo) Create(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bankDo) CreateInBatches(values []*model.Bank, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bankDo) Save(values ...*model.Bank) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bankDo) First() (*model.Bank, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Take() (*model.Bank, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Last() (*model.Bank, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) Find() ([]*model.Bank, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bank), err
}

func (b bankDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Bank, error) {
	result, err := b.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Bank), err
}

func (b bankDo) FindInBatches(result []*model.Bank, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(&result, batchSize, fc)
}

func (b bankDo) Attrs(attrs ...field.AssignExpr) *bankDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bankDo) Assign(attrs ...field.AssignExpr) *bankDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bankDo) Joins(field field.RelationField) *bankDo {
	return b.withDO(b.DO.Joins(field))
}

func (b bankDo) Preload(field field.RelationField) *bankDo {
	return b.withDO(b.DO.Preload(field))
}

func (b bankDo) FirstOrInit() (*model.Bank, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FirstOrCreate() (*model.Bank, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bank), nil
	}
}

func (b bankDo) FindByPage(offset int, limit int) (result []*model.Bank, count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	result, err = b.Offset(offset).Limit(limit).Find()
	return
}

func (b *bankDo) withDO(do gen.Dao) *bankDo {
	b.DO = *do.(*gen.DO)
	return b
}
