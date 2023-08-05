package commons

import "log"

func Getconnection() *gormDB {
	db, err := gorm.Open("mysql", "root:@/jepon?chatset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}