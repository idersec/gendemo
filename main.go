package main

import (
	"context"

	"gendemo/biz"
)

func main() {
	// dal.init()
	ctx := context.Background()

	biz.Create(ctx)
	biz.Delete(ctx)
	biz.Query(ctx)
	biz.Update(ctx)
	biz.Association(ctx)
	biz.Transaction(ctx)

}
