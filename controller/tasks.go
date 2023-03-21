package controller

import (
	// "errors"
	"fmt"
	"main/config"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Home(c *gin.Context) {
	c.String(200, "Welcome")
}

// var errors error

func GetPriceList(c *gin.Context) {
	price_list := []models.PriceList{}
	
	// config.DB.Find(&price_list)
	// Check if returns RecordNotFound error
	// err := db.First(&user, 100).Error
	// errors.Is(err, gorm.ErrRecordNotFound)

	if err:= config.DB.Preload("Project").Find(&price_list).Error;err!= nil {
		c.JSON(http.StatusInternalServerError, "no project and price list")
	}

	c.JSON(200, &price_list)
	// err := config.DB.Preload("Project").Find(&price_list).Error
	// errors.Is(err, gorm.ErrRecordNotFound)
	// config.DB.Raw("SELECT * FROM price_lists WHERE project_id IN (SELECT * FROM project WHERE id=?),"project).Scan(&project)
	// c.String(210,"Hello UserController")
	fmt.Println("endpoint hit: /GetPriceList")
}

func CreatePriceList(c *gin.Context) {
	var price_list models.PriceList
	c.BindJSON(&price_list)
	config.DB.Create(&price_list)
	c.JSON(200, &price_list)
	fmt.Println("endpoint hit: /CreatePriceList")
}

func DeletePriceList(c *gin.Context) {
	var price_list models.PriceList
	config.DB.Where("id = ?", c.Param("id")).Delete(&price_list)
	c.JSON(200, &price_list)
	fmt.Println("endpoint hit: /DeletePriceList")
}

func UpdatePriceList(c *gin.Context) {
	var price_list models.PriceList
	config.DB.Where("id = ?", c.Param("id")).First(&price_list)
	c.BindJSON(&price_list)
	config.DB.Save(&price_list)
	c.JSON(200, &price_list)
	fmt.Println("endpoint hit: /UpdatePriceList")
}

// func GetCommission(c *gin.Context){
// 	commission := []models.Commission{}
// 	config.DB.Find(&commission)
// 	c.JSON(200,&commission)
// 	// c.String(210,"Hello UserController")
// 	fmt.Println("endpoint hit: /GetCommission")
// }

// func CreatePriceList(c *gin.Context){
// 	var price_list models.PriceList
// 	c.BindJSON(&price_list)
// 	config.DB.Create(&price_list)
// 	c.JSON(200,&price_list)
// 	fmt.Println("endpoint hit: /CreatePriceList")

// }

// func DeletePriceList(c *gin.Context){
// 	var user models.PriceList
// 	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
// 	c.JSON(200,&user)
// 	fmt.Println("endpoint hit: /DeletePriceList")

// }

// func UpdatePriceList(c *gin.Context){
// 	var user models.PriceList
// 	config.DB.Where("id = ?",c.Param("id")).First(&user)
// 	c.BindJSON(&user)
// 	config.DB.Save(&user)
// 	c.JSON(200,&user)
// 	fmt.Println("endpoint hit: /UpdatePriceList")

// }

func GetCommission(c *gin.Context) {

	// var artikel models.Artikel
	// Conn.Preload("Komentar").Find(&artikle)
	commission := []models.Commission{}
	// var project []models.Project
	// price_list := models.PriceList{}
	// config.DB.Find(&project)
	config.DB.Preload("Project").Preload("PriceList.Project").Find(&commission)
	// test:=config.DB.Raw("Select * from price_lists, projects where price_lists.id = projects.price_list_id")
	c.JSON(200, &commission)
	// c.String(210,"Hello UserController")
	fmt.Println("endpoint hit: /GetCommission")
}

func CreateCommission(c *gin.Context) {
	var commission models.Commission
	c.BindJSON(&commission)
	config.DB.Create(&commission)
	c.JSON(200, &commission)
	fmt.Println("endpoint hit: /CreateCommission")

}

func DeleteCommission(c *gin.Context) {
	var commission models.Commission
	config.DB.Where("id = ?", c.Param("id")).Delete(&commission)
	c.JSON(200, &commission)
	fmt.Println("endpoint hit: /DeleteCommission")
}

func UpdateCommission(c *gin.Context) {
	var commission models.Commission
	config.DB.Where("id = ?", c.Param("id")).First(&commission)
	c.BindJSON(&commission)
	config.DB.Save(&commission)
	c.JSON(200, &commission)
	fmt.Println("endpoint hit: /UpdateCommission")
}

// func LoginUser(ctx *gin.Context){

// 	var loginService service.LoginService = service.StaticLoginService()
// 	var jwtService service.JWTService = service.JWTAuthService()
// 	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

// 	token := loginController.Login(ctx)
// 	if token != "" {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"token": token,
// 		})
// 	} else {
// 		ctx.JSON(http.StatusUnauthorized, nil)
// 	}
// }

func GetUser(c *gin.Context) {

	// var artikel models.Artikel
	// Conn.Preload("Komentar").Find(&artikle)

	// user:=[]models.User{}
	var user []models.User

	// var project []models.Project
	// price_list := models.PriceList{}
	err := config.DB.Find(&user).Error
	if err != nil {
		c.JSON(400, "eror wkkw")
	}
	// test:=config.DB.Raw("Select * from price_lists, projects where price_lists.id = projects.price_list_id")
	c.JSON(200, &user)
	// c.String(210,"Hello UserController")
	fmt.Println("endpoint hit: /GetUser")
}

// func CreateUser(context *gin.Context) {
// 	var user models.User
// 	if err := context.ShouldBindJSON(&user); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}
// 	if err := user.HashPassword(user.Password); err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}
// 	record := config.DB.Create(&user)
// 	if record.Error != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
// 		context.Abort()
// 		return
// 	}
// 	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
// }

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
	fmt.Println("endpoint hit: /DeleteUser")
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
	fmt.Println("endpoint hit: /CreateUser")

}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
	fmt.Println("endpoint hit: /UpdateUser")
}

func GetProject(c *gin.Context) {

	// var artikel models.Artikel
	// Conn.Preload("Komentar").Find(&artikle)
	project := []models.Project{}
	// var project []models.Project
	// price_list := models.PriceList{}
	config.DB.Find(&project)
	// test:=config.DB.Raw("Select * from price_lists, projects where price_lists.id = projects.price_list_id")
	c.JSON(200, &project)
	// c.String(210,"Hello UserController")
	fmt.Println("endpoint hit: /GetProject")
}

func CreateProject(c *gin.Context) {
	var project models.Project
	c.BindJSON(&project)
	config.DB.Create(&project)
	c.JSON(200, &project)
	fmt.Println("endpoint hit: /CreateProject")

}

func DeleteProject(c *gin.Context) {
	var project models.Project
	config.DB.Where("id = ?", c.Param("id")).Delete(&project)
	c.JSON(200, &project)
	fmt.Println("endpoint hit: /DeleteProject")
}

func UpdateProject(c *gin.Context) {
	var project models.Project
	config.DB.Where("id = ?", c.Param("id")).First(&project)
	c.BindJSON(&project)
	config.DB.Save(&project)
	c.JSON(200, &project)
	fmt.Println("endpoint hit: /UpdateProject")
}
