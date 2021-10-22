package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Username string

type Password string

func (p *Password) Scan(src interface{}) error {
	*p = Password(fmt.Sprintf("@%v@", src))
	return nil
}

func (p *Password) Value() (driver.Value, error) {
	*p = Password(strings.Trim(string(*p), "@"))
	return p, nil
}

type Passport struct {
	ID        int
	Username  Username // will be field.String
	Password  Password // will be field.Field because type Password set Scan and Value
	LoginTime time.Time
}

type Company struct {
	gorm.Model
	ID          int
	Name        string
	CreateAt    time.Time
	Broken      bool
	MarketValue float64
}

type Struct struct {
	Test string
}
type TestField struct {
	TestInt         int       // gen type: field.Int
	TestInt8        int8      // gen type: field.Int8
	TestInt16       int16     // gen type: field.Int16
	TestInt32       int32     // gen type: field.Int32
	TestInt64       int64     // gen type: field.Int64
	TestUint        uint      // gen type: field.Uint
	TestUint8       uint8     // gen type: field.Uint8
	TestUint16      uint16    // gen type: field.Uint16
	TestUint32      uint32    // gen type: field.Uint32
	TestUint64      uint64    // gen type: field.Uint64
	TestString      string    // gen type: field.String
	TestByte        byte      // gen type: field.Byte
	TestFloat32     float32   // gen type: field.Float32
	TestFloat64     float64   // gen type: field.Float64
	TestBool        time.Time // gen type: field.Time
	TestStringPoint *string   // gen type: field.String  will ignore ptr
	TestIntPoint    *int      // gen type: field.Int	will ignore ptr
	name            string    // unexported member will ignore

	// badCase
	//List []string // unsupported data type slice, will panic
	//TestStruct Struct //unsupported user defined  struct, will panic
}

// Associations

type Customer struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:CustomerRefer"`
	// CreditCards []CreditCard `gorm:"many2many:cus_cards"`
}

type CreditCard struct {
	gorm.Model
	Number        string
	CustomerRefer uint
	BankID        uint
	Bank          Bank
}

type Bank struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Address string
	Scale   int
}
