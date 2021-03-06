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

func newCreditCard(db *gorm.DB) creditCard {
	_creditCard := creditCard{}

	_creditCard.creditCardDo.UseDB(db)
	_creditCard.creditCardDo.UseModel(&model.CreditCard{})

	tableName := _creditCard.creditCardDo.TableName()
	_creditCard.ALL = field.NewField(tableName, "*")
	_creditCard.ID = field.NewUint(tableName, "id")
	_creditCard.CreatedAt = field.NewTime(tableName, "created_at")
	_creditCard.UpdatedAt = field.NewTime(tableName, "updated_at")
	_creditCard.DeletedAt = field.NewField(tableName, "deleted_at")
	_creditCard.Number = field.NewString(tableName, "number")
	_creditCard.CustomerRefer = field.NewUint(tableName, "customer_refer")
	_creditCard.BankID = field.NewUint(tableName, "bank_id")
	_creditCard.Bank = creditCardBank{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Bank", "model.Bank"),
	}

	return _creditCard
}

type creditCard struct {
	creditCardDo creditCardDo

	ALL           field.Field
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Number        field.String
	CustomerRefer field.Uint
	BankID        field.Uint
	Bank          creditCardBank
}

func (c *creditCard) WithContext(ctx context.Context) *creditCardDo {
	return c.creditCardDo.WithContext(ctx)
}

func (c creditCard) TableName() string { return c.creditCardDo.TableName() }

func (c creditCard) clone(db *gorm.DB) creditCard {
	c.creditCardDo.ReplaceDB(db)
	return c
}

type creditCardBank struct {
	db *gorm.DB

	field.RelationField
}

func (a creditCardBank) Where(conds ...field.Expr) *creditCardBank {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a creditCardBank) WithContext(ctx context.Context) *creditCardBank {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a creditCardBank) Model(m *model.CreditCard) *creditCardBankTx {
	return &creditCardBankTx{a.db.Model(m).Association(a.Name())}
}

type creditCardBankTx struct{ tx *gorm.Association }

func (a creditCardBankTx) Find() (result *model.Bank, err error) {
	return result, a.tx.Find(&result)
}

func (a creditCardBankTx) Append(values ...*model.Bank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a creditCardBankTx) Replace(values ...*model.Bank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a creditCardBankTx) Delete(values ...*model.Bank) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a creditCardBankTx) Clear() error {
	return a.tx.Clear()
}

func (a creditCardBankTx) Count() int64 {
	return a.tx.Count()
}

type creditCardDo struct{ gen.DO }

func (c creditCardDo) Debug() *creditCardDo {
	return c.withDO(c.DO.Debug())
}

func (c creditCardDo) WithContext(ctx context.Context) *creditCardDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c creditCardDo) Clauses(conds ...clause.Expression) *creditCardDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c creditCardDo) Not(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c creditCardDo) Or(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c creditCardDo) Select(conds ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c creditCardDo) Where(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c creditCardDo) Order(conds ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c creditCardDo) Distinct(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c creditCardDo) Omit(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c creditCardDo) Join(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c creditCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c creditCardDo) RightJoin(table schema.Tabler, on ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c creditCardDo) Group(cols ...field.Expr) *creditCardDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c creditCardDo) Having(conds ...gen.Condition) *creditCardDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c creditCardDo) Limit(limit int) *creditCardDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c creditCardDo) Offset(offset int) *creditCardDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c creditCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *creditCardDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c creditCardDo) Unscoped() *creditCardDo {
	return c.withDO(c.DO.Unscoped())
}

func (c creditCardDo) Create(values ...*model.CreditCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c creditCardDo) CreateInBatches(values []*model.CreditCard, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c creditCardDo) Save(values ...*model.CreditCard) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c creditCardDo) First() (*model.CreditCard, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Take() (*model.CreditCard, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Last() (*model.CreditCard, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) Find() ([]*model.CreditCard, error) {
	result, err := c.DO.Find()
	return result.([]*model.CreditCard), err
}

func (c creditCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.CreditCard, error) {
	result, err := c.DO.FindInBatch(batchSize, fc)
	return result.([]*model.CreditCard), err
}

func (c creditCardDo) FindInBatches(result []*model.CreditCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(&result, batchSize, fc)
}

func (c creditCardDo) Attrs(attrs ...field.AssignExpr) *creditCardDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c creditCardDo) Assign(attrs ...field.AssignExpr) *creditCardDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c creditCardDo) Joins(field field.RelationField) *creditCardDo {
	return c.withDO(c.DO.Joins(field))
}

func (c creditCardDo) Preload(field field.RelationField) *creditCardDo {
	return c.withDO(c.DO.Preload(field))
}

func (c creditCardDo) FirstOrInit() (*model.CreditCard, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) FirstOrCreate() (*model.CreditCard, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreditCard), nil
	}
}

func (c creditCardDo) FindByPage(offset int, limit int) (result []*model.CreditCard, count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	result, err = c.Offset(offset).Limit(limit).Find()
	return
}

func (c *creditCardDo) withDO(do gen.Dao) *creditCardDo {
	c.DO = *do.(*gen.DO)
	return c
}
