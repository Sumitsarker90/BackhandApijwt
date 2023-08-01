package repository

import (
	"net/http"
	"strings"

	"github.com/xyz/model"
	"github.com/xyz/util"
)

type EmployeeRepository struct {
}

func (designationrepo *EmployeeRepository) GetAllEmployee() model.ResponseDto {
	var output model.ResponseDto
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var obj []model.Employee
	result := db.Order("id").Find(&obj)
	if result.RowsAffected == 0 {
		output.Message = "No country info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutPut struct {
		T_output    []model.Employee `json:"output"`
		OutputCount int           `json:"outputCount"`
	}
	var tOutput tempOutPut
	tOutput.T_output = obj
	tOutput.OutputCount = len(obj)
	output.Message = "List of designations"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK

	return output
}
func (designationrepo *EmployeeRepository) GetById(id model.Employee) model.ResponseDto {
	var output model.ResponseDto
	if id.Id <= 0 {
		output.Message = "Employee can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}

	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	result := db.Raw("select * from public.Employee where code = ?", id.Id).First(&id)
	if result.RowsAffected == 0 {
		output.Message = "No country info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutput struct {
		Output model.Employee `json:"output"`
	}
	var tOutput tempOutput
	tOutput.Output = id
	output.Message = "Designation info details found for given criteria"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK
	return output
}

func (designationrepo *EmployeeRepository) AddDesignation(c model.Employee) model.ResponseDto {
	var output model.ResponseDto
	if c.Id <= 0 {
		output.Message = "Invalid Id"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output

	}
	if c.Name == "" {
		output.Message = "Name can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output

	}
	if c.Designation == "" {
		output.Message = "Dept can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//result := db.Raw("Select * from public.desig where code =?", c.Code).First(&c)
	result := db.Where(&model.Employee{Id: c.Id}).First(&c)
	if result.RowsAffected != 0 {
		output.Message = "Department Code is already exist"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusConflict
		return output
	}
	result1 := db.Raw("Select * from public.desig where lower(designation) =? and lower(sdesignation) = ?", strings.ToLower(c.Name), strings.ToLower(c.Designation)).First(&c)
	if result1.RowsAffected != 0 {
		output.Message = "Name Or Designation is alread exist"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusConflict
		return output
	}
	// result2 := db.Raw("Select * from public.desig where sdesignation =?", c.Sdesignation).First(&c)
	// if result2.RowsAffected !=0{
	// 	output.Message ="Designation is alread exist"
	// 	output.IsSuccess = false
	// 	output.Payload=nil
	// 	output.StatusCode = http.StatusBadRequest
	// 	return output
	// }
	result3 := db.Create(&c)
	if result3.RowsAffected == 0 {
		output.Message = "Designation creation failed"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusInternalServerError
		return output
	}

	type abc struct {
		Output model.Employee `json:"output"`
	}
	var a abc
	a.Output = c
	output.Message = "Designation create succesfully"
	output.IsSuccess = true
	output.Payload = a
	output.StatusCode = http.StatusOK
	return output
}

func (designationrepo *EmployeeRepository) UpdateDesignation(input model.Employee) model.ResponseDto {
	var response model.ResponseDto
	if input.Id <= 0 {
		response.Message = " Code can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	if input.Name == "" {
		response.Message = "Name can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response

	}
	if input.Designation == "" {
		response.Message = "Designation can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var output model.Employee
	result := db.Where(&model.Employee{Id: input.Id}).First(&output)
	if result.RowsAffected == 0 {
		response.Message = "this code doesnot exists"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusNotFound
		return response
	}
	output.Name = input.Name
	output.Designation = input.Designation
	result1 := db.Where(&model.Employee{Id: input.Id}).Updates(&output)
	if result1.RowsAffected == 0 {
		response.Message = "No Employee info found for given criteria"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusInternalServerError
		return response
	}
	response.Message = "Employee info updated successfully"
	response.IsSuccess = true
	response.Payload = output
	response.StatusCode = http.StatusOK

	return response
}

func (designationrepo *EmployeeRepository) Delete(c model.Employee) model.ResponseDto {
	var output model.ResponseDto
	if c.Id <= 0 {
		output.Message = "Invalid code"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	result := db.Where("code = ?", c.Id).Delete(&c)
	if result.RowsAffected == 0 {
		output.Message = "No info found for given criteria"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	output.Message = "Deleted successfully"
	output.IsSuccess = true
	output.Payload = nil
	output.StatusCode = http.StatusOK
	return output
}
func (designationrepo *EmployeeRepository) MaxDeptCode(c model.Employee) model.ResponseDto {
	var output model.ResponseDto
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tx := db.Begin()
	tx.SavePoint("savepoint")
	var outputcode model.Employee
	result := tx.Raw("select max(code)+1 from public.desig").First(&outputcode.Id)
	if result.RowsAffected == 0 {
		tx.RollbackTo("savepoint")
		output.IsSuccess = false
		output.StatusCode = http.StatusNotFound
		output.Message = "Internal Server error!"
		output.Payload = nil
		return output
	}

	tx.Commit()

	output.IsSuccess = true
	output.StatusCode = 200
	output.Message = "Max code for new dept entry"
	output.Payload = outputcode

	return output
}

