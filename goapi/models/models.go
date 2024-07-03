package models

import (
	"time"
)

// used for tables
// users table --> username, email, uid, password
type User struct {
	UserID      int    `json:"userid"                     gorm:"column:userid;primaryKey;autoIncrement"`
	Email       string `json:"email"    form:"email"      gorm:"unique;not null"`
	Username    string `json:"username" form:"username"   gorm:"unique;not null"`
	Password    string `json:"password" form:"password"   gorm:"not null"`
	ConfirmPass string `json:"confirm"  form:"confirm"    gorm:"-"`
}

// configs table --> configid, configname, version
type CustomConfigurations struct {
	ConfigID   int    `json:"configid"                     gorm:"primaryKey;autoIncrement"`
	ConfigName string `json:"configname" form:"configname" gorm:"not null"`
	Version    string `json:"version"    form:"version"    gorm:"not null"`
}

// running instance --> runid, userid, password, configId array
type RunningInstance struct {
	RunID       int `json:"runid"   gorm:"column:runid;primaryKey;autoIncrement"`
	UserID      int `               gorm:"column:userid;foreignKey:users.userid"`
	ConfigArray []int64 `           gorm:"column:configid_arr;type:integer[]"`
}

// toverify --> vid, username, email, password, timestamp
type Verification struct {
	VerificationID int       `json:"vid"                        gorm:"column:vid;primaryKey;autoIncrement"`
	Email          string    `json:"email"    form:"email"      gorm:"unique;not null"`
	Username       string    `json:"username" form:"username"   gorm:"unique;not null"`
	Password       string    `json:"password" form:"password"   gorm:"not null"`
	CreatedAt      time.Time `                                  gorm:"column:created_at;index"`
}

// might not be used for tables
type InstanceDetails struct {
	ID       int             `json:"userid"`
	Password string          `json:"password"`
	Port     int             `json:"port"`
	Deadline int64           `json:"deadline"`
	Config   ContainerConfig `json:"config"`
}

type Credentials struct {
	Email    string `json:"email"    form:"email"`
	Password string `json:"password" form:"password"`
}

type ContainerConfig struct {
	Image      string   `json:"image"`
	Dependency []string `json:"dependency"`
}
