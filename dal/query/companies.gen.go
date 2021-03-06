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
	"gorm.io/gen/helper"
)

func newCompany(db *gorm.DB) company {
	_company := company{}

	_company.companyDo.UseDB(db)
	_company.companyDo.UseModel(&model.Company{})

	tableName := _company.companyDo.TableName()
	_company.ALL = field.NewField(tableName, "*")
	_company.ID = field.NewInt(tableName, "id")
	_company.CreatedAt = field.NewTime(tableName, "created_at")
	_company.UpdatedAt = field.NewTime(tableName, "updated_at")
	_company.DeletedAt = field.NewField(tableName, "deleted_at")
	_company.Name = field.NewString(tableName, "name")
	_company.CreateAt = field.NewTime(tableName, "create_at")
	_company.Broken = field.NewBool(tableName, "broken")
	_company.MarketValue = field.NewFloat64(tableName, "market_value")

	return _company
}

type company struct {
	companyDo companyDo

	ALL         field.Field
	ID          field.Int
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Name        field.String
	CreateAt    field.Time
	Broken      field.Bool
	MarketValue field.Float64
}

func (c *company) WithContext(ctx context.Context) *companyDo { return c.companyDo.WithContext(ctx) }

func (c company) TableName() string { return c.companyDo.TableName() }

func (c company) clone(db *gorm.DB) company {
	c.companyDo.ReplaceDB(db)
	return c
}

type companyDo struct{ gen.DO }

//Where("name=@name and age=@age")
func (c companyDo) FindByNameAndAge(name string, age int) (result *model.Company, err error) {
	params := map[string]interface{}{
		"name": name,
		"age":  age,
	}

	var generateSQL string
	generateSQL += "name=@name and age=@age"

	executeSQL := c.UnderlyingDB().Where(generateSQL, params).Take(&result)
	err = executeSQL.Error
	return
}

//sql(select id,name,age from users where age>18)
func (c companyDo) FindBySimpleName() (result []*model.Company, err error) {
	var generateSQL string
	generateSQL += "select id,name,age from users where age>18"

	executeSQL := c.UnderlyingDB().Raw(generateSQL).Find(&result)
	err = executeSQL.Error
	return
}

//sql(select id,name,age from @@table where age>18
//{{if cond1}}and id=@id {{end}}
//{{if name == ""}}and @@col=@name{{end}})
func (c companyDo) FindByIDOrName(cond1 bool, id int, col string, name string) (result *model.Company, err error) {
	params := map[string]interface{}{
		"id":   id,
		"name": name,
	}

	var generateSQL string
	generateSQL += "select id,name,age from companies where age>18 "

	ifCond0 := make([]helper.Cond, 0, 100)
	ifCond0 = append(ifCond0, helper.Cond{cond1, "and id=@id"})
	generateSQL += helper.IfClause(ifCond0)

	ifCond1 := make([]helper.Cond, 0, 100)
	ifCond1 = append(ifCond1, helper.Cond{name == "", "and " + c.Quote(col) + "=@name"})
	generateSQL += helper.IfClause(ifCond1)

	executeSQL := c.UnderlyingDB().Raw(generateSQL, params).Take(&result)
	err = executeSQL.Error
	return
}

//sql(select * from users)
func (c companyDo) FindAll() (result []map[string]interface{}, err error) {
	var generateSQL string
	generateSQL += "select * from users"

	executeSQL := c.UnderlyingDB().Raw(generateSQL).Find(&result)
	err = executeSQL.Error
	return
}

//sql(select * from users limit 1)
func (c companyDo) FindOne() (result map[string]interface{}) {
	var generateSQL string
	generateSQL += "select * from users limit 1"

	result = make(map[string]interface{})
	_ = c.UnderlyingDB().Raw(generateSQL).Take(result)
	return
}

//sql(select address from users limit 1)
func (c companyDo) FindAddress() (result *model.Company, err error) {
	var generateSQL string
	generateSQL += "select address from users limit 1"

	executeSQL := c.UnderlyingDB().Raw(generateSQL).Take(&result)
	err = executeSQL.Error
	return
}

func (c companyDo) Debug() *companyDo {
	return c.withDO(c.DO.Debug())
}

func (c companyDo) WithContext(ctx context.Context) *companyDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c companyDo) Clauses(conds ...clause.Expression) *companyDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c companyDo) Not(conds ...gen.Condition) *companyDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c companyDo) Or(conds ...gen.Condition) *companyDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c companyDo) Select(conds ...field.Expr) *companyDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c companyDo) Where(conds ...gen.Condition) *companyDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c companyDo) Order(conds ...field.Expr) *companyDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c companyDo) Distinct(cols ...field.Expr) *companyDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c companyDo) Omit(cols ...field.Expr) *companyDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c companyDo) Join(table schema.Tabler, on ...field.Expr) *companyDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c companyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *companyDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c companyDo) RightJoin(table schema.Tabler, on ...field.Expr) *companyDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c companyDo) Group(cols ...field.Expr) *companyDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c companyDo) Having(conds ...gen.Condition) *companyDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c companyDo) Limit(limit int) *companyDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c companyDo) Offset(offset int) *companyDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c companyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *companyDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c companyDo) Unscoped() *companyDo {
	return c.withDO(c.DO.Unscoped())
}

func (c companyDo) Create(values ...*model.Company) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c companyDo) CreateInBatches(values []*model.Company, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c companyDo) Save(values ...*model.Company) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c companyDo) First() (*model.Company, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Take() (*model.Company, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Last() (*model.Company, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) Find() ([]*model.Company, error) {
	result, err := c.DO.Find()
	return result.([]*model.Company), err
}

func (c companyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.Company, error) {
	result, err := c.DO.FindInBatch(batchSize, fc)
	return result.([]*model.Company), err
}

func (c companyDo) FindInBatches(result []*model.Company, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(&result, batchSize, fc)
}

func (c companyDo) Attrs(attrs ...field.AssignExpr) *companyDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c companyDo) Assign(attrs ...field.AssignExpr) *companyDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c companyDo) Joins(field field.RelationField) *companyDo {
	return c.withDO(c.DO.Joins(field))
}

func (c companyDo) Preload(field field.RelationField) *companyDo {
	return c.withDO(c.DO.Preload(field))
}

func (c companyDo) FirstOrInit() (*model.Company, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) FirstOrCreate() (*model.Company, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Company), nil
	}
}

func (c companyDo) FindByPage(offset int, limit int) (result []*model.Company, count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	result, err = c.Offset(offset).Limit(limit).Find()
	return
}

func (c *companyDo) withDO(do gen.Dao) *companyDo {
	c.DO = *do.(*gen.DO)
	return c
}
