package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xyz/model"
	"github.com/xyz/repository"
)

type EmployeeRestService struct {
}

func (employeeRestService *EmployeeRestService) GetAllDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	rep := new(repository.EmployeeRepository)
	res := rep.GetAllEmployee()
	c.JSON(res.StatusCode, res)
}

func (employeeRestService *EmployeeRestService) GetById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var id model.Employee
	c.ShouldBind(&id)
	rep1 := new(repository.EmployeeRepository)
	res1 := rep1.GetById(id)
	c.JSON(res1.StatusCode, res1)
}

func (employeeRestService *EmployeeRestService) AddDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var obj model.Employee
	c.ShouldBind(&obj)
	rep := new(repository.EmployeeRepository)
	res := rep.AddDesignation(obj)
	c.JSON(res.StatusCode, res)
}

func (employeeRestService *EmployeeRestService) UpdateDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var obj1 model.Employee
	c.ShouldBind(&obj1)
	repo := new(repository.EmployeeRepository)
	result2 := repo.UpdateDesignation(obj1)
	c.JSON(result2.StatusCode, result2)
}

func (employeeRestService *EmployeeRestService) DeleteDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var id model.Employee
	c.ShouldBind(&id)
	repo := new(repository.EmployeeRepository)
	result3 := repo.Delete(id)
	c.JSON(result3.StatusCode, result3)


}

func(employeeRestService *EmployeeRestService) GetMaxDeptCode(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")	
	var input model.Employee
	c.ShouldBind(&input)

	rep := new(repository.EmployeeRepository)
	res := rep.MaxDeptCode(input)
	c.JSON(res.StatusCode, res)

}


func (employeeRestService *EmployeeRestService) AddRouters(router *gin.Engine) {
	router.GET("/getalldesignation", employeeRestService.GetAllDesignation)
	router.POST("/getdesigbyid", employeeRestService.GetById )
	router.POST("/adddasignation", employeeRestService.AddDesignation)
	router.PATCH("/updatedesigntion", employeeRestService.UpdateDesignation)
	router.DELETE("/deletedesigntion", employeeRestService.DeleteDesignation)
	router.GET("/maxdesigcode",employeeRestService.GetMaxDeptCode)
}
