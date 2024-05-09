package main

import(
	"Time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type MongoInstense struct {
	Client *mongo.Client
	Db    *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeOut(context.Background(), 30*time.Second)
   defer cancel()

   err = client.Connect(ctx)
   db := client.Database(dbName)

   if err != nil{
	return err
   }
   mg = MongoInstance{
	Client: client,
	Db: db,
   }
   return nil
}

func main(){
	if err := Connect(); err !=nil{
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}

		cursor, err := mg.Db.collection("employees").Find(c.Context())
        if err != nil{
			return c.Status(500).SendString(err.Error())
		}
		var employees []Employees = Make([]Employees,  0)
	    if err := cursor.All(c.Context(), &employees); error{
			return c.Status(500).SendString(err.Error())
		}
		return c.Json(employees)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employees := new(Employees)

		if err := BodyParser(employee); err != nil{
			return c.Status(400).sendString(err.Error)
		}
		employee.ID = " "

		insertionResult, err := collection.InsertOne(c.Context(), employee)

		if err != nil{
			return c.Status(500).SendString(err, Error())
		}
		filter := bson.D{{key: "_id", value: insertionResult.Inserted}}
		createdRecord := collection.FindOne(c.context(), filter)

		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)
	    
		return c.Status(201).json(createdEmployee)
	})
	app.Put("/employee/:id", func(c *fiber.Ctx) error{
		idParams := c.Params("id")

		employeeID, err := primitive.ObjectIDFromHex(idParam)

		if err !=nil {
			return c.SendStatus(400)
		}
		
		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
	}) 
	query := bson.D{{key: "_id", value: employeeID}}
	update := bson.D{
		{
			key: "$set", 
			Value: bson.D{
				{Key: "name", value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value:employee.Salary},
			},
		},
	}
	err = mg.Db.Collection("employees").FindOneUpdate(c.Contex(), query update) error{

	}
	app.Delete("/employee/:id", func(c *fiber.Ctx) error{
		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"),)

		if err != nil {
			return c.SendStatus(500)
		}
        query := bson.D{{key:"_id", Value: employeeID}}
		result, err := mg.db.Collection("employees").DeleteOne(c.Count(), &query)

		if err != nil {
			return c.SendStatus(500)
		}
		if result.DeletedCount < 1{
			return c.SendStatus(404)
		}
		return c.Status(200).JSON("record deleted")
	} )
	log.Fatal(app.Listen(":3000"))
}
