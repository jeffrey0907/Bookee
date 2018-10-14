package user

type User struct {
    Uid     int64   `gorm:"type:int,AUTO_INCREMENT,PRIMARY_KEY"`
    Name    string  `gorm:"type:varchar(20)"`
    OpenId  string  `gorm:"type:varchar(50),UNIQUE_INDEX:user_openid_idx"`
    UnionId string  `gorm:"type:varchar(50)"`
}
