package method

import (
	"gorm.io/gen"
)

type X string

type Method interface {
	// Where("name=@name and age=@age")
	FindByNameAndAge(name string, age int) (gen.T, error)
	//sql(select id,name,age from users where age>18)
	FindBySimpleName() ([]gen.T, error)

	//sql(select id,name,age from @@table where age>18
	//{{if cond1}}and id=@id {{end}}
	//{{if name == ""}}and @@col=@name{{end}})
	FindByIDOrName(cond1 bool, id int, col, name string) (gen.T, error)

	//sql(select * from users)
	FindAll() ([]gen.M, error)

	//sql(select * from users limit 1)
	FindOne() gen.M

	//sql(select address from users limit 1)
	FindAddress() (gen.T, error)
}

type Methods interface {
	// Where("name=@name and age=@age")
	SimpleFindByNameAndAge(name string, age int) (gen.T, error)
	//simple query

	// select * from users where id=@id
	FindUserToMap(id int) (gen.M, error)
	// return to map[string]interface{}

	//sql(insert into @@table (name,age) values (@name,@age) )
	InsertValue(age int, name string) error
	// execute only return error

	/*
		select * from @@table
		  {{where}}
				  id>0
		       {{if cond}}id=@id {{end}}
		       {{if key!="" && value != ""}} or @@key=@value{{end}}
		   {{end}}
	*/
	FindByIDOrCustom(cond bool, id int, key, value string) ([]gen.T, error)
	// where clause

	//update @@table
	//	{{set}}
	//		update_time=now(),
	//		{{if name != ""}}
	//			name=@name
	//		{{end}}
	//	{{end}}
	//	{{where}}
	//		id=@id
	//	{{end}}
	UpdateName(name string, id int) error
	// Set clause if clause blend

}

// only used to User
type UserMethod interface {
	//where(id=@id)
	FindByID(id int) (gen.T, error)

	//select * from users where age>18
	FindAdult() ([]gen.T, error)

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
	FindByRole(role string, id int)

	//update users
	//	{{set}}
	//		update_time=now(),
	//		{{if name != ""}}
	//			name=@name
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserName(name string, id int) error

	//Insert()

}
