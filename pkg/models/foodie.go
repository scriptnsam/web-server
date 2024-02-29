package models

import (
	"github.com/jinzhu/gorm"
	"github.com/scriptnsam/web-server/pkg/config"
)

var db *gorm.DB

type Food struct{
	gorm.Model
	Name string `json:"name"`
	Price float32 `json:"price"`
	Description string `json:"description"`
	Image string `json:"image"`
	CategoryID uint `json:"category_id"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Food{})
}

func (f *Food) CreateFood() *Food{
	db.NewRecord(f)
	db.Create(&f)
	return f
}

func GetFoods()[]Food{
	var Foods []Food
	db.Find(&Foods)
	return Foods
}

func GetFoodById(Id int64) (*Food, *gorm.DB){
	var getFood Food
	db := db.Where("ID=?", Id).Find(&getFood)
	return &getFood, db
}

func UpdateFood(Id int64, food *Food) *gorm.DB{
	var updateFood Food
	db := db.Where("ID=?", Id).Find(&updateFood)
	updateFood.Name = food.Name
	updateFood.Price = food.Price
	updateFood.Description = food.Description
	updateFood.Image = food.Image
	db.Save(&updateFood)
	return db
}

func DeleteFood(Id int64) *gorm.DB{
	var deleteFood Food
	db := db.Where("ID=?", Id).Delete(&deleteFood)
	return db
}

