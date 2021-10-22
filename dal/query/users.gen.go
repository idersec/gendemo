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

func newUser(db *gorm.DB) user {
	_user := user{}

	_user.userDo.UseDB(db)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewField(tableName, "*")
	_user.ID = field.NewInt32(tableName, "id")
	_user.Name = field.NewString(tableName, "name")
	_user.Age = field.NewInt32(tableName, "age")
	_user.Password = field.NewString(tableName, "password")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Title = field.NewString(tableName, "title")
	_user.IsAdult = field.NewInt32(tableName, "is_adult")
	_user.UpdatedTime = field.NewTime(tableName, "updated_time")
	_user.Pass = field.NewString(tableName, "pass")
	_user.Role = field.NewString(tableName, "role")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.Address = field.NewString(tableName, "address")

	return _user
}

type user struct {
	userDo userDo

	ALL         field.Field
	ID          field.Int32
	Name        field.String
	Age         field.Int32
	Password    field.String
	DeletedAt   field.Field
	Title       field.String
	IsAdult     field.Int32
	UpdatedTime field.Time
	Pass        field.String
	Role        field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	Address     field.String
}

func (u *user) WithContext(ctx context.Context) *userDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

//Where("name=@name and age=@age")
func (u userDo) FindByNameAndAge(name string, age int) (result *model.User, err error) {
	params := map[string]interface{}{
		"name": name,
		"age":  age,
	}

	var generateSQL string
	generateSQL += "name=@name and age=@age"

	executeSQL := u.UnderlyingDB().Where(generateSQL, params).Take(&result)
	err = executeSQL.Error
	return
}

//sql(select id,name,age from users where age>18)
func (u userDo) FindBySimpleName() (result []*model.User, err error) {
	var generateSQL string
	generateSQL += "select id,name,age from users where age>18"

	executeSQL := u.UnderlyingDB().Raw(generateSQL).Find(&result)
	err = executeSQL.Error
	return
}

//sql(select id,name,age from @@table where age>18
//{{if cond1}}and id=@id {{end}}
//{{if name == ""}}and @@col=@name{{end}})
func (u userDo) FindByIDOrName(cond1 bool, id int, col string, name string) (result *model.User, err error) {
	params := map[string]interface{}{
		"id":   id,
		"name": name,
	}

	var generateSQL string
	generateSQL += "select id,name,age from users where age>18 "

	ifCond0 := make([]helper.Cond, 0, 100)
	ifCond0 = append(ifCond0, helper.Cond{cond1, "and id=@id"})
	generateSQL += helper.IfClause(ifCond0)

	ifCond1 := make([]helper.Cond, 0, 100)
	ifCond1 = append(ifCond1, helper.Cond{name == "", "and " + u.Quote(col) + "=@name"})
	generateSQL += helper.IfClause(ifCond1)

	executeSQL := u.UnderlyingDB().Raw(generateSQL, params).Take(&result)
	err = executeSQL.Error
	return
}

//sql(select * from users)
func (u userDo) FindAll() (result []map[string]interface{}, err error) {
	var generateSQL string
	generateSQL += "select * from users"

	executeSQL := u.UnderlyingDB().Raw(generateSQL).Find(&result)
	err = executeSQL.Error
	return
}

//sql(select * from users limit 1)
func (u userDo) FindOne() (result map[string]interface{}) {
	var generateSQL string
	generateSQL += "select * from users limit 1"

	result = make(map[string]interface{})
	_ = u.UnderlyingDB().Raw(generateSQL).Take(result)
	return
}

//sql(select address from users limit 1)
func (u userDo) FindAddress() (result *model.User, err error) {
	var generateSQL string
	generateSQL += "select address from users limit 1"

	executeSQL := u.UnderlyingDB().Raw(generateSQL).Take(&result)
	err = executeSQL.Error
	return
}

//where(id=@id)
func (u userDo) FindByID(id int) (result *model.User, err error) {
	params := map[string]interface{}{
		"id": id,
	}

	var generateSQL string
	generateSQL += "id=@id"

	executeSQL := u.UnderlyingDB().Where(generateSQL, params).Take(&result)
	err = executeSQL.Error
	return
}

//select * from users where age>18
func (u userDo) FindAdult() (result []*model.User, err error) {
	var generateSQL string
	generateSQL += "select * from users where age>18"

	executeSQL := u.UnderlyingDB().Raw(generateSQL).Find(&result)
	err = executeSQL.Error
	return
}

//select * from @@table
//	{{where}}
//		{{if role=="user"}}
//			id=@id
//		{{else if role=="admin"}}
//			role="user" or rule="normal-admin"
//		{{else}}
//			role="user" or role="normal-admin" or role="admin"
//		{{end}}
//	{{end}}
func (u userDo) FindByRole(role string, id int) {
	params := map[string]interface{}{
		"id": id,
	}

	var generateSQL string
	generateSQL += "select * from users"

	whereCond0 := make([]string, 0, 100)
	ifCond0 := make([]helper.Cond, 0, 100)
	ifCond0 = append(ifCond0, helper.Cond{role == "user", " id=@id"})
	ifCond0 = append(ifCond0, helper.Cond{!(role == "user") && role == "admin", " role=\"user\" or rule=\"normal-admin\" "})
	ifCond0 = append(ifCond0, helper.Cond{!(role == "user" || role == "admin"), " role=\"user\" or role=\"normal-admin\" or role=\"admin\" "})
	whereCond0 = append(whereCond0, helper.IfClause(ifCond0))
	generateSQL += helper.WhereClause(whereCond0)

	_ = u.UnderlyingDB().Exec(generateSQL, params)
	return
}

//update users
//	{{set}}
//		update_time=now(),
//		{{if name != ""}}
//			name=@name
//		{{end}}
//	{{end}}
//where id=@id
func (u userDo) UpdateUserName(name string, id int) (err error) {
	params := map[string]interface{}{
		"name": name,
		"id":   id,
	}

	var generateSQL string
	generateSQL += "update users "

	setCond0 := make([]string, 0, 100)
	setCond0 = append(setCond0, " update_time=now(), ")
	ifCond0 := make([]helper.Cond, 0, 100)
	ifCond0 = append(ifCond0, helper.Cond{name != "", " name=@name"})
	setCond0 = append(setCond0, helper.IfClause(ifCond0))
	generateSQL += helper.SetClause(setCond0)

	generateSQL += " where id=@id"

	executeSQL := u.UnderlyingDB().Exec(generateSQL, params)
	err = executeSQL.Error
	return
}

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) ([]*model.User, error) {
	result, err := u.DO.FindInBatch(batchSize, fc)
	return result.([]*model.User), err
}

func (u userDo) FindInBatches(result []*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(&result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(field field.RelationField) *userDo {
	return u.withDO(u.DO.Joins(field))
}

func (u userDo) Preload(field field.RelationField) *userDo {
	return u.withDO(u.DO.Preload(field))
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	return
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}