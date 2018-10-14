package repository

import (
    "Bookee/domain/user"
    "testing"
)

func TestUser(t *testing.T) {
    if !DB().HasTable(&user.User{}) {
        DB().CreateTable(&user.User{})
    }
    usr := user.User{OpenId:"Open_jeff"}
    db := DB().Begin()
    if !db.NewRecord(usr) {
        t.Log("create user failed")
    }
    db.Commit()

    var usr1 user.User

    t.Log(`---------------------------`)
    t.Log(&usr1)
    db1 := DB().New()
    db1.First(&usr1)
    db1.Commit()
    t.Log(&usr1)
}
