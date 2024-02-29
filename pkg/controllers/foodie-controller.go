package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scriptnsam/web-server/pkg/models"
	"github.com/scriptnsam/web-server/pkg/utils"
)

var NewFood models.Food

func CreateFood(w http.ResponseWriter, r *http.Request){
	CreateFood := &models.Food{}
	utils.ParseBody(r, CreateFood)
	f := CreateFood.CreateFood()
	res, _ := json.Marshal(f)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFoods(w http.ResponseWriter, r *http.Request){
	newFood := models.GetFoods()
	res, _ := json.Marshal(newFood)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFoodsById(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	foodId:= vars["id"]
	ID,err := strconv.ParseInt(foodId, 0, 0)
	if err != nil{
		panic("error while parsing")
	}
	foodDetails, _ := models.GetFoodById(ID)
	res, _ := json.Marshal(foodDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFood(w http.ResponseWriter,r *http.Request){
	var updateFood = &models.Food{}
	utils.ParseBody(r, updateFood)
	vars := mux.Vars(r)
	foodId := vars["id"]
	ID,err:=strconv.ParseInt(foodId, 0, 0)
	if err != nil{
		panic("error while parsing")
	}
	foodDetails, db :=models.GetFoodById(ID)
	if updateFood.Name != ""{
		foodDetails.Name = updateFood.Name
	}
	if updateFood.Price != 0{
		foodDetails.Price = updateFood.Price
	}
	if updateFood.Description != ""{
		foodDetails.Description = updateFood.Description
	}
	db.Save(&foodDetails)
	res, _ := json.Marshal(foodDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFood(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	foodId := vars["id"]
	ID,err:=strconv.ParseInt(foodId, 0, 0)
	if err != nil{
		panic("error while parsing")
	}
	models.DeleteFood(ID)
	w.WriteHeader(http.StatusOK)
}