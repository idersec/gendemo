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

func newCustomer(db *gorm.DB) customer {
	_customer := customer{}

	_customer.customerDo.UseDB(db)
	_customer.customerDo.UseModel(&model.Customer{})

	tableName := _customer.customerDo.TableName()
	_customer.ALL = field.NewField(tableName, "*")
	_customer.ID = field.NewUint(tableName, "id")
	_customer.CreatedAt = field.NewTime(tableName, "created_at")
	_customer.UpdatedAt = field.NewTime(tableName, "updated_at")
	_customer.DeletedAt = field.NewField(tableName, "deleted_at")
	_customer.CreditCards = customerCreditCards{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CreditCards", "model.CreditCard"),
		Bank: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("CreditCards.Bank", "model.Bank"),
		},
	}

	return _customer
}

type customer struct {
	customerDo customerDo

	ALL         field.Field
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	CreditCards customerCreditCards
}

func (c *customer) WithContext(ctx context.Context) *customerDo { return c.customerDo.WithContext(ctx) }

func (c customer) TableName() string { return c.customerDo.TableName() }

func (c customer) clone(db *gorm.DB) customer {
	c.customerDo.ReplaceDB(db)
	return c
}

type customerCreditCards struct {
	db *gorm.DB

	field.RelationField

	Bank struct {
		field.RelationField
	}
}

func (a customerCreditCards) Where(conds ...field.Expr) *customerCreditCards {
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

func (a customerCreditCards) WithContext(ctx context.Context) *customerCreditCards {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a customerCreditCards) Model(m *model.Customer) *customerCreditCardsTx {
	return &customerCreditCardsTx{a.db.Model(m).Association(a.Name())}
}

type customerCreditCardsTx struct{ tx *gorm.Association }

func (a customerCreditCardsTx) Find() (result *model.CreditCard, err error) {
	return result, a.tx.Find(&result)
}

func (a customerCreditCardsTx) Append(values ...*model.CreditCard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a customerCreditCardsTx) Replace(values ...*model.CreditCard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a customerCreditCardsTx) Delete(values ...*model.CreditCard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a customerCreditCardsTx) Clear() error {
	return a.tx.Clear()
}

func (a customerCreditCardsTx) Count() int64 {
	return a.tx.Count()
}

type customerDo struct{ gen.DO }

func (c customerDo) Debug() *customerDo {
	return c.withDO(c.DO.Debug())
}

func (c customerDo) WithContext(ctx context.Context) *customerDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c customerDo) Clauses(conds ...clause.Expression) *customerDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c customerDo) Not(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c customerDo) Or(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c customerDo) Select(conds ...field.Expr) *customerDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c customerDo) Where(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c customerDo) Order(conds ...field.Expr) *customerDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c customerDo) Distinct(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c customerDo) Omit(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c customerDo) Join(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c customerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c customerDo) RightJoin(table schema.Tabler, on ...field.Expr) *customerDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c customerDo) Group(cols ...field.Expr) *customerDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c customerDo) Having(conds ...gen.Condition) *customerDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c customerDo) Limit(limit int) *customerDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c customerDo) Offset(offset int) *customerDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c customerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *customerDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c customerDo) Unscoped() *customerDo {
	return c.withDO(c.DO.Unscoped())
}

func (c customerDo) Create(values ...*model.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c customerDo) CreateInBatches(values []*model.Customer, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c customerDo) Save(values ...*model.Customer) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c customerDo) First() (*model.Customer, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Customer), nil
	}
}

func (c customerDo) Take() (*model.Customer, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Customer), nil
	}
}

func (c customerDo) Last() (*model.Customer, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Customer), nil
	}
}

func (c customerDo) Find() ([]*model.Customer, error) {
	result, err := c.DO.Find()
	return result.([]*model.Customer), err
}

func (c customerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Customer, error) {
	result, err := c.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Customer), err
}

func (c customerDo) FindInBatches(result []*model.Customer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(&result, batchSize, fc)
}

func (c customerDo) Attrs(attrs ...field.AssignExpr) *customerDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c customerDo) Assign(attrs ...field.AssignExpr) *customerDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c customerDo) Joins(field field.RelationField) *customerDo {
	return c.withDO(c.DO.Joins(field))
}

func (c customerDo) Preload(field field.RelationField) *customerDo {
	return c.withDO(c.DO.Preload(field))
}

func (c customerDo) FirstOrInit() (*model.Customer, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Customer), nil
	}
}

func (c customerDo) FirstOrCreate() (*model.Customer, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Customer), nil
	}
}

func (c customerDo) FindByPage(offset int, limit int) (result []*model.Customer, count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	result, err = c.Offset(offset).Limit(limit).Find()
	return
}

func (c *customerDo) withDO(do gen.Dao) *customerDo {
	c.DO = *do.(*gen.DO)
	return c
}
