package models

// used for tables
// users table --> username, email, uid, password
type User struct {
	UserID      int    `json:"userid"                     gorm:"primaryKey;autoIncrement"`
	Email       string `json:"email"    form:"email"      gorm:"unique;not null"`
	Username    string `json:"username" form:"username"   gorm:"unique;not null"`
	Password    string `json:"password" form:"password"   gorm:"not null"`
	ConfirmPass string `json:"confirm"  form:"confirm"`
}

// configs table --> configid, configname, version
type CustomConfigurations struct {
	ConfigID   int    `json:"configid"                     gorm:"primaryKey;autoIncrement"`
	ConfigName string `json:"configname" form:"configname" gorm:"not null"`
	Version    string `json:"version"    form:"version"    gorm:"not null"`
}

// running instance --> runid, userid, password, configId array
type RunningInstance struct {
	RunID       int    `json:"runid"                gorm:"primaryKey;autoIncrement"`
	UserID      int    `                            gorm:"foreignKey:users.userid"`
	Password    string `json:"instancepassword"     gorm:"not null"`
	ConfigArray []int
}


// might not be used for tables
type Instance struct {
	ID       int             `json:"userid"`
	Password string          `json:"password"`
	Port     int             `json:"port"`
	Deadline int             `json:"deadline"`
	Config   ContainerConfig `json:"config"`
}

type Credentials struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ContainerConfig struct {
	Image      string   `json:"image"`
	Dependency []string `json:"dependency"`
}
