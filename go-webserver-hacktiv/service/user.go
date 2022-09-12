package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// pertama bikin stuct user dan userservice
// kedua bikin func NewUserService untuk dilempar ke main.go
// ketiga bikin func Register untuk menambahkan user ke db
// keempat bikin func GetUser untuk mengambil semua user dari db
// kelima bikin func GetUserById untuk mengambil user berdasarkan id
// keenam bikin interface UserIface untuk kumpulan semua fungsi yang ada di UserService
type User struct {
	Id   int    `json:"id" form:"id"`
	Nama string `json:"name" form:"name"`
}

type UserService struct {
	db []*User
}

type UserIface interface {
	RegisterAPI(c *gin.Context)
	GetUserAPI(c *gin.Context)
	UpdateUserAPI(c *gin.Context)
	DeleteUserAPI(c *gin.Context)
	Register(u *User) string
	GetUser() []*User
	GetUserById(id int) *User
	UpdateUser(id int, nama string) string
	DeleteUser(id int) string
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

func (u *UserService) UpdateUser(id int, nama string) string {
	for _, v := range u.db {
		if v.Id == id {
			v.Nama = nama
			return nama + " berhasil diupdate"
		}
	}
	return "user tidak ditemukan"
}

func (u *UserService) DeleteUser(id int) string {
	for i, v := range u.db {
		if v.Id == id {
			fmt.Printf("%# v", v)
			fmt.Printf("%# v", u.db[:i])
			fmt.Printf("%# v", u.db[:i+1])
			u.db = append(u.db[:i], u.db[i+1:]...)
			return "user berhasil dihapus"
		}
	}
	return "user tidak ditemukans"
}

func (u *UserService) RegisterAPI(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Printf("%# v", user)

	if c.Request.Method == "POST" {
		user := u.Register(
			&User{
				Id:   user.Id,
				Nama: user.Nama,
			},
		)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": user,
		})
		return
	}
}

func (u *UserService) GetUserAPI(c *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Content-Type", "application/json")

	if c.Request.Method == "GET" {
		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			result := u.GetUser()

			if result == nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "data kosong",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": result,
				})
				return
			}

		} else {
			result := u.GetUserById(id)
			if result == nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "user not found",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": result,
				})
			}

			return
		}
	}
}

func (u *UserService) UpdateUserAPI(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if c.Request.Method == "PUT" {
		user := u.UpdateUser(
			id,
			user.Nama,
		)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": user,
		})
		return
	}
}

func (u *UserService) DeleteUserAPI(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))
	if c.Request.Method == "DELETE" {
		user := u.DeleteUser(
			id,
		)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": user,
		})
		return
	}
}
