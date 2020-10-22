package main
 
import (
   "context"
   "log"
 
   "github.com/nutzann/app/controllers"
   _ "github.com/nutzann/app/docs"
   "github.com/nutzann/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	Name   string
	Email  string
	UserID string
}

// Diseases  defines the struct for the disease
type Diseases struct {
	Disease []Disease
}

// Disease  defines the struct for the  diseases
type Disease struct {
	Name string
}

// Levels  defines the struct for the level
type Levels struct {
	Level []Level
}

// Level  defines the struct for the  levels
type Level struct {
	Name string
}

// Statistics defines the struct for the statistic
type Statistics struct {
	Statistic []Statistic
}

// Statistic  defines the struct for the  statistics
type Statistic struct {
	Name string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
  
	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
  
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
  
	v1 := router.Group("/api/v1")
	controllers.NewAreaController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewLevelController(v1, client)
	controllers.NewStatisticController(v1, client)

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"คิม จงอิน", "kim@gmail.com", "D12345"},
			Employee{"โอ เซฮุน", "hunnie@gmail.com", "D54231"},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetName(e.Name).
			SetEmail(e.Email).
			SetUserID(e.UserID).
			Save(context.Background())
	}

	// Set Diseases Data
	diseases := Diseases{
		Disease: []Disease{
			Disease{"HIV"},
			Disease{"Covid-19"},
			Disease{"ไข้หวัดใหญ่"},
		},
	}

	for _, d := range diseases.Disease {
		client.Disease.
			Create().
			SetName(d.Name).
			Save(context.Background())
	}
  
	// Set Levels Data
	levels := Levels{
		Level: []Level{
			Level{"ความเสี่ยงระดับต่ำ"},
			Level{"ความเสี่ยงระดับกลาง"},
			Level{"ความเสี่ยงระดับสูง"},
		},
	}

	for _, l := range levels.Level {
		client.Level.
			Create().
			SetName(l.Name).
			Save(context.Background())
	}
	// Set Statistics Data
	statistics := Statistics{
		Statistic: []Statistic{
			Statistic{"ประชากรติดจำนวนน้อยมาก"},
			Statistic{"ประชากรติดจำนวนน้อย"},
			Statistic{"ประชากรติดจำนวนปานกลาง"},
			Statistic{"ประชากรติดจำนวนมาก"},
			Statistic{"ประชากรติดจำนวนมากที่สุด"},
		},
	}

	for _, st := range statistics.Statistic {
		client.Statistic.
			Create().
			SetName(st.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
 }
 