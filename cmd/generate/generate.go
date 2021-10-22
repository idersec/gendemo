package main

import (
	"gendemo/dal/method"

	"gorm.io/gen"

	"gendemo/dal"
	"gendemo/dal/model"
)

// unexported struct will be ignored.
type test struct {
	id  int
	Xxx string
	Ttt int
}

func main() {
	// dal.init()

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		// Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(dal.DB)

	userTpl := g.GenerateModel("users")
	personTpl := g.GenerateModelAs("people", "Person")

	g.ApplyBasic(model.Customer{}, model.CreditCard{}, model.Bank{}) // Associations
	g.ApplyBasic(personTpl, userTpl)
	g.ApplyInterface(func(method.Method) {}, userTpl, &model.Company{}, test{}) // struct test will be ignored

	g.ApplyBasic(model.Company{}, model.TestField{})
	g.ApplyBasic(model.Passport{}, personTpl,
		g.GenerateModelAs("address", "Addr",
			gen.FieldIgnore("deleted_at"),
			gen.FieldNewTag("id", `newTag:"tag info"`),
		),
	)
	g.ApplyInterface(func(method.Method, model.TestField) {}, &model.Company{}, userTpl, test{})
	g.ApplyInterface(func(method method.UserMethod) {}, userTpl)
	g.Execute()
}
