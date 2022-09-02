package service

// pertama bikin stuct user dan userservice
// kedua bikin func NewUserService untuk dilempar ke main.go
// ketiga bikin func Register untuk menambahkan user ke db
// keempat bikin func GetUser untuk mengambil semua user dari db
// kelima bikin func GetUserById untuk mengambil user berdasarkan id
// keenam bikin interface UserIface untuk kumpulan semua fungsi yang ada di UserService
type User struct {
	Id   int
	Nama string
}

type UserService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
	GetUserById(id int) *User
}

func NewUserService(db []*User) UserIface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(usr *User) string {
	u.db = append(u.db, usr)
	// fmt.Println(usr)
	// fmt.Println(u)
	// fmt.Println(u.db)
	return usr.Nama + " berhasil terdaftar"
}

func (u *UserService) GetUser() []*User {
	return u.db
}

func (u *UserService) GetUserById(id int) *User {
	for _, v := range u.db {
		if v.Id == id {
			return v
		}
	}
	return nil
}
