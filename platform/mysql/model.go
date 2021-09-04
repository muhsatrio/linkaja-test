package mysql

type Customer struct {
	CustomerNumber uint `gorm:"primarykey"`
	Name           string
}

type Account struct {
	AccountNumber  uint `gorm:"primarykey"`
	Balance        int
	CustomerNumber uint
	Customer       Customer `gorm:"foreignKey:CustomerNumber"`
}
