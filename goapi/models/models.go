package models

type User struct {
	UserID      int    `json:"userid"`
	Email       string `json:"email" form:"email"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	ConfirmPass string `json:"confirm" form:"confirm"`
}

type Instance struct {
	UserID   int             `json:"userid"`
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
