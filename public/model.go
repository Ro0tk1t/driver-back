package public

type User struct {
	ID    int64  `json:"id" xorm:"'id' pk autoincr"`
	Name  string `json:"name" xorm:"'name' varchar(100)"`
	Email string `json:"email" xorm:"'email' varchar(80) unique" validate:"required,email"`
	Phone string `json:"phone" xorm:"'phone' varchar(20)"`
	// TODO: encrypt password
	Password string `json:"password" validate:"required"`
	Role     int64  `json:"role" xorm:"'role' int"`
	Created  int64  `json:"created" xorm:"'created' int"`
}

func (u *User) TableName() string {
	return "users"
}

type Avatar struct {
	ID     int64  `json:"id" xorm:"'id' pk autoincr"`
	URL    string `json:"url" xorm:"'url' varchar(255)"`
	UserID int64  `json:"user_id" xorm:"'user_id'"`
}

type File struct {
	ID   int64  `json:"id" xorm:"'id' pk autoincr"`
	UID  int64  `json:"uid" xorm:"'uid' int"`
	Path string `json:"path" xorm:"'path' text not null default('/')"`
	Name string `json:"name" xorm:"'name' varchar(255)"`
	Size int64  `json:"size" xorm:"'size' int"`
	Type string `json:"type" xorm:"'type' varchar(255)"`
	Hash string `json:"hash" xorm:"'hash' varchar(65)"`
	Time int64  `json:"time" xorm:"'time' int"`
}

type ShareInfo struct {
	Link     string `json:"link" xorm:"'link' varchar(255)"`
	UserName string `json:"username" xorm:"'username' varchar(255)"`
	Fids     string `json:"fid" xorm:"'fids' text"`
	Uid      int64  `json:"uid" xorm:"'uid' int"`
	Created  int64  `json:"created" xorm:"'created' int"`
	Expired  int64  `json:"expired" xorm:"'expired' int"`
}
