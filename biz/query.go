package biz

import (
	"context"
	"log"

	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/hints"

	"gendemo/dal"
	"gendemo/dal/model"
	"gendemo/dal/query"
)

var q = query.Use(dal.DB.Debug())

func Create(ctx context.Context) {
	var err error
	ud := q.User.WithContext(ctx)

	userData := &model.User{ID: 1, Name: "modi"}
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`id`) VALUES ('2021-09-13 20:05:51.389','2021-09-13 20:05:51.389',NULL,'modi',0,'',1)
	err = ud.Create(userData)
	catchErr("create one item fail", err)

	userDataArray := []*model.User{{ID: 2, Name: "A"}, {ID: 3, Name: "B"}}
	err = ud.CreateInBatches(userDataArray, 2)
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`id`) VALUES ('2021-09-13 20:05:51.403','2021-09-13 20:05:51.403',NULL,'A',0,'',2),('2021-09-13 20:05:51.403','2021-09-13 20:05:51.403',NULL,'B',0,'',3)
	catchErr("create in batches fail", err)

	userData.Name = "new name"
	err = ud.Save(userData)
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`id`) VALUES ('2021-09-13 20:05:51.389','2021-09-13 20:05:51.409',NULL,'new name',0,'',1) ON DUPLICATE KEY UPDATE `updated_at`=VALUES(`updated_at`),`deleted_at`=VALUES(`deleted_at`),`name`=VALUES(`name`),`age`=VALUES(`age`),`role`=VALUES(`role`)
	catchErr("save user fail", err)
}

func Delete(ctx context.Context) {
	var err error
	u, ud := q.User, q.User.WithContext(ctx)

	_, err = ud.Where(u.ID.Eq(1)).Delete()
	// UPDATE `users` SET `deleted_at`='2021-09-13 20:05:51.418' WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	catchErr("delete item(id=1) fail", err)

	_, err = ud.Where(u.ID.In(2, 3)).Delete()
	// UPDATE `users` SET `deleted_at`='2021-09-13 20:05:51.428' WHERE `users`.`id` IN (2,3) AND `users`.`deleted_at` IS NULL
	catchErr("delete items fail", err)

	_, err = ud.Where(u.ID.Gt(100)).Unscoped().Delete()
	// DELETE FROM `users` WHERE `users`.`id` > 100
	catchErr("unscoped delete item fail", err)
}

func Query(ctx context.Context) {
	var err error
	var user *model.User
	var users []*model.User

	u, ud := q.User, q.User.WithContext(ctx)

	/*--------------Basic query-------------*/
	user, err = ud.Take()
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1
	catchErr("take 1 item fail", err)
	log.Printf("query 1 item: %+v\n", user)

	user, err = ud.Where(u.ID.Gt(100), u.Name.Like("%T%")).Take()
	// SELECT * FROM `users` WHERE `users`.`id` > 100 AND `users`.`name` LIKE '%T%' AND `users`.`deleted_at` IS NULL LIMIT 1
	catchErr("query with conditions fail", err)
	log.Printf("query conditions got: %+v\n", user)

	user, err = ud.Where(ud.Columns(u.ID).In(ud.Select(u.ID.Min()))).First()
	// SELECT * FROM `users` WHERE `users`.`id` IN (SELECT MIN(`users`.`id`) FROM `users` WHERE `users`.`deleted_at` IS NULL) AND `users`.`deleted_at` IS NULL
	// ORDER BY `users`.`id` LIMIT 1
	catchErr("subquery 1 fail", err)
	log.Printf("subquery 1 got item: %+v\n", user)

	user, err = ud.Where(ud.Columns(u.ID).Eq(ud.Select(u.ID.Max()))).First()
	// SELECT * FROM `users` WHERE `users`.`id` = (SELECT MAX(`users`.`id`) FROM `users` WHERE `users`.`deleted_at` IS NULL) AND `users`.`deleted_at` IS NULL
	// ORDER BY `users`.`id` LIMIT 1
	catchErr("subquery 2 fail", err)
	log.Printf("subquery 2 got item: %+v\n", user)

	users, err = ud.Distinct(u.Name).Find()
	// SELECT DISTINCT `users`.`name` FROM `users` WHERE `users`.`deleted_at` IS NULL
	// users, err = u.Select(u.Name).Distinct().Find()
	// users, err = u.Distinct(u.ID, u.Name).Find()
	// users, err = u.Distinct(u.ID, u.Name.As("n")).Find()
	catchErr("select distinct fail", err)
	log.Printf("select distinct got: %d\n", len(users))

	/*--------------Diy query-------------*/
	user, err = ud.FindByNameAndAge("tom", 29)
	// SELECT * FROM `users` WHERE name='tom' and age=29 AND `users`.`deleted_at` IS NULL
	catchErr("FindByNameAndAge fail", err)
	log.Printf("FindByNameAndAge: %+v\n", user)

	users, err = ud.FindBySimpleName()
	// select id,name,age from users where age>18
	catchErr("FindBySimpleName fail", err)
	log.Printf("FindBySimpleName: (%d)%+v\n", len(users), users)

	user, err = ud.FindByIDOrName(false, 0, "tom", "user")
	// select id,name,age from users where age>18
	catchErr("FindByIDOrName fail", err)
	log.Printf("FindByIDOrName: %+v\n", user)

	/*--------------Advanced query-------------*/
	users, err = ud.Clauses(hints.New("MAX_EXECUTION_TIME(10000)")).Find()
	// SELECT /*+ MAX_EXECUTION_TIME(10000) */ * FROM `users` WHERE `users`.`deleted_at` IS NULL
	catchErr("Find with hints 1 fail", err)
	log.Printf("find with hints 2: (%d)%+v\n", len(users), users)

	users, err = ud.Clauses(hints.New("idx_user_name")).Find()
	// SELECT /*+ idx_user_name */ * FROM `users` WHERE `users`.`deleted_at` IS NULL
	catchErr("Find with hints 2 fail", err)
	log.Printf("find with hints 2: (%d)%+v\n", len(users), users)

	users, err = ud.Clauses(hints.New("hint")).Select(u.ID, u.Name).Where(u.ID.IsNotNull(), u.Age.Gt(18)).Find()
	// SELECT `users`.`id`,`users`.`name` FROM `users` WHERE `users`.`id` IS NOT NULL AND `users`.`age` > 18 AND `users`.`deleted_at` IS NULL
	catchErr("Find with hints 3 fail", err)
	log.Printf("find with hints 3: (%d)%+v\n", len(users), users)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).FirstOrInit()
	catchErr("FirstOrInit fail", err)
	log.Printf("FirstOrInit got: %+v\n", user)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).Attrs(u.Name.Value("modi")).FirstOrInit()
	catchErr("FirstOrInit fail", err)
	log.Printf("FirstOrInit got: %+v\n", user)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).Attrs(u.Name.Value("modi")).Assign(u.Age.Value(17)).FirstOrInit()
	catchErr("FirstOrInit fail", err)
	log.Printf("FirstOrInit got: %+v\n", user)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).FirstOrCreate()
	catchErr("FirstOrCreate fail", err)
	log.Printf("FirstOrCreate got: %+v\n", user)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).Attrs(u.Name.Value("modi")).FirstOrCreate()
	catchErr("FirstOrCreate fail", err)
	log.Printf("FirstOrCreate got: %+v\n", user)

	user, err = ud.Select(u.ID, u.Name).Where(u.ID.Eq(1)).Attrs(u.Name.Value("modi")).Assign(u.Age.Value(17)).FirstOrCreate()
	// UPDATE `users` SET `age`=17 WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	catchErr("FirstOrCreate fail", err)
	log.Printf("FirstOrCreate got: %+v\n", user)
}

func Update(ctx context.Context) {
	var err error

	u, ud := q.User, q.User.WithContext(ctx)

	user, err := ud.First()
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	catchErr("First fail", err)

	user.Name = "save test"
	err = ud.Save(user)
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`id`) VALUES ('2021-09-13 20:12:18.655','2021-09-13 20:12:18.655',NULL,'save test',190,'',4) ON DUPLICATE KEY UPDATE `updated_at`=VALUES(`updated_at`),`deleted_at`=VALUES(`deleted_at`),`name`=VALUES(`name`),`age`=VALUES(`age`),`role`=VALUES(`role`)
	catchErr("Save fail", err)

	_, err = ud.Where(u.ID.Eq(user.ID)).Update(u.Name, "update test")
	// UPDATE `users` SET `name`='update test',`updated_at`='2021-09-13 20:12:18.664' WHERE `users`.`id` = 4 AND `users`.`deleted_at` IS NULL
	catchErr("Update fail", err)

	_, err = ud.Where(u.ID.Eq(user.ID)).Updates(model.User{Name: "updates test"})
	// UPDATE `users` SET `updated_at`='2021-09-28 20:14:41.139',`name`='updates test' WHERE `users`.`id` = 4 AND `users`.`deleted_at` IS NULL
	catchErr("Updates fail", err)

	_, err = ud.Where(u.ID.Eq(user.ID)).UpdateSimple(u.Age.Add(1), u.CreatedAt.Null(), u.Name.Value("modi"), u.UpdatedAt.Zero())
	// UPDATE `users` SET `age`=`users`.`age`+1,`created_at`=NULL,`name`='modi',`updated_at`='0000-00-00 00:00:00'
	// WHERE `users`.`id` = 4 AND `users`.`deleted_at` IS NULL
	catchErr("UpdateSimple fail", err)
}

func Association(ctx context.Context) {
	c, cd := q.Customer, q.Customer.WithContext(ctx)

	err := cd.Save(&model.Customer{
		Model: gorm.Model{ID: 1},
		CreditCards: []model.CreditCard{
			{
				Model:  gorm.Model{ID: 1},
				Number: "123",
				Bank: model.Bank{
					ID:   1,
					Name: "zhaoshang",
				},
			},
			{
				Model:  gorm.Model{ID: 2},
				Number: "456",
				Bank: model.Bank{
					ID:   2,
					Name: "pufa",
				},
			},
		}})
	catchErr("create with association fail", err)

	// query without associations
	customer, err := cd.Where(c.ID.Eq(1)).Take()
	catchErr("Take customer without association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// preload CreditCards
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).Preload(c.CreditCards).Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// preload with conditions
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).Preload(c.CreditCards.On(q.CreditCard.ID.Gt(1))).Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// nested preload with conditions
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).
		Preload(c.CreditCards.On(q.CreditCard.ID.Gt(1))).
		Preload(c.CreditCards.Bank).
		Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// more compalicate preload conditions
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).Preload(
		c.CreditCards.Clauses(hints.New("new hints")).On(q.CreditCard.ID.Gt(0)).Order(q.CreditCard.ID.Desc(), q.CreditCard.Number),
	).Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// long path preload
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).Preload(c.CreditCards.Bank).Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	// preload all fields（1 level）
	customer, err = c.WithContext(ctx).Where(c.ID.Eq(1)).Preload(field.Associations).Take()
	catchErr("Take customer with association fail", err)
	log.Printf("Take customer got: %+v\n", customer)

	count := c.CreditCards.Model(&model.Customer{Model: gorm.Model{ID: 1}}).Count()
	catchErr("Count customer's CreditCards fail", err)
	log.Printf("Count cards got: %+v\n", count)

	cards, err := c.CreditCards.Model(&model.Customer{Model: gorm.Model{ID: 1}}).Find()
	catchErr("Find all customer's cards fail", err)
	log.Printf("Find cards got: %+v\n", cards)

	cards, err = c.CreditCards.Where(q.CreditCard.ID.Gte(2)).Model(&model.Customer{Model: gorm.Model{ID: 1}}).Find()
	catchErr("Find Cards with conditions", err)
	log.Printf("Find cards got: %+v\n", cards)

	err = c.CreditCards.Model(&model.Customer{Model: gorm.Model{ID: 1}}).Append(
		&model.CreditCard{
			Model:  gorm.Model{ID: 3},
			Number: "789", Bank: model.Bank{ID: 3, Name: "ustr"},
		})
	catchErr("Append CreditCards fail", err)
	log.Printf("Append cards got: %+v\n", cards)

	err = c.CreditCards.Model(&model.Customer{Model: gorm.Model{ID: 1}}).Delete(
		&model.CreditCard{Model: gorm.Model{ID: 2}})
	catchErr("Delete CreditCards fail", err)
	log.Printf("Delete cards got: %+v\n", cards)

	err = c.CreditCards.Model(&model.Customer{Model: gorm.Model{ID: 1}}).Clear()
	catchErr("Clear all CreditCards fail", err)
	log.Printf("clear cards")
}

func Transaction(ctx context.Context) {
	err := q.Transaction(func(tx *query.Query) error {
		err := tx.User.WithContext(ctx).Create(&model.User{Name: "modi"})
		if err != nil {
			return err
		}

		err = tx.Transaction(func(tx2 *query.Query) error {
			_, err := tx2.User.WithContext(ctx).Where(tx.User.ID.Eq(1)).Delete()
			return err
		})
		return err
	})
	if err != nil {
		panic(err)
	}
	_ = createWithTx(ctx, nil)
}

func createWithTx(ctx context.Context, customers ...*model.Customer) (err error) {
	tx := q.Begin()
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	// do something here
	err = tx.Customer.WithContext(ctx).Create(customers...)
	if err != nil {
		return
	}
	return tx.Commit()
}
