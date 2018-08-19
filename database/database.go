package database

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "github.com/AlyHKafoury/oaas/userdefinedmodels"
)

var db sql.DB

func init (){
  db, err := sql.Open("mysql", "root:@/oaas?charset=utf8")
  fmt.Println("%v+ %v+",db,err)
}

func CreateModel(userModel *userdefinedmodels.UserModel){
  fmt.Println("%v+", userModel)
}
