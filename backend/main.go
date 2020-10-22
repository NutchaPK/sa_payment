package main
 
import (
	"context"
	"fmt"
	
   "log"
   "github.com/theuo/app/controllers"
   _ "github.com/theuo/app/docs"
   "github.com/theuo/app/ent"
   "github.com/theuo/app/ent/user"
   "github.com/theuo/app/ent/billstatus"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}
type User struct {
	Email  string
	Name string
}
type Bills struct {
	Bill []Bill
}
type Bill struct {
	Amount  int
	Resident int
	BillStatus int
}
type BillStatuss struct {
	BillStatus []BillStatus
}
type BillStatus struct {
	BStatus  string
}
type PaymentStatuss struct {
	PaymentStatus []PaymentStatus
}
type PaymentStatus struct {
	PStatus  string
}
type PayTypes struct {
	PayType []PayType
}
type PayType struct {
	TypeInfo string
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

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewBillController(v1, client)
	controllers.NewBillStatusController(v1, client)
	controllers.NewPaymentController(v1, client)
	controllers.NewPayTypeController(v1, client)
	controllers.NewPaymentStatusController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"Nutchaporn@gmail.com", "Nutchaporn Klinrod"},
			User{"B6102845@g.sut.ac.th", "Name Surname"},
		},
	}
	for _, u := range users.User {
		
		client.User.
			Create().
			SetEmail(u.Email).
			SetName(u.Name).
			Save(context.Background())
	}
	// Set BillStatuiss Data
	billstatuss := BillStatuss{
		BillStatus: []BillStatus{
			BillStatus{"Unpaid"},
			BillStatus{"Paid"},
		},
	}

	for _, bs := range billstatuss.BillStatus {
		
		client.BillStatus.
			Create().
			SetBillStatus(bs.BStatus).
			Save(context.Background())
	}

	// Set Bills Data
	bills := Bills{
		Bill: []Bill{
			Bill{4790, 1,1},
			Bill{4500, 1,2},
			Bill{3900, 2,1},
		},
	}

	for _, b := range bills.Bill {
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(b.Resident))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
			bs, err := client.BillStatus.
				Query().
				Where(billstatus.IDEQ(int(b.BillStatus))).
				Only(context.Background())
	
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		client.Bill.
			Create().
			SetAmount(b.Amount).
			SetResident(u).
			SetBillstatus(bs).
			Save(context.Background())
	}
	// Set PaymentStatuss Data
	paymentstatuss := PaymentStatuss{
		PaymentStatus: []PaymentStatus{
			PaymentStatus{"On-time"},
			PaymentStatus{"Overdue"},
		},
	}

	for _, ps := range paymentstatuss.PaymentStatus{
		
		client.PaymentStatus.
			Create().
			SetPaymentStatus(ps.PStatus).
			Save(context.Background())
	}

	// Set PayTypes Data
	paytypes := PayTypes{
		PayType: []PayType{
			PayType{"Online Banking"},
			PayType{"Counter Service"},
		},
	}

	for _, pt := range paytypes.PayType {
		
		client.PayType.
			Create().
			SetTypeInfo(pt.TypeInfo).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
 }
  
 