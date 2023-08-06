package commons

import "log"

func Getconnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/jepon?chatset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}