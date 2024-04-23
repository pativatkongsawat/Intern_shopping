package guestcontroller

import (
	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	return e.JSON(200, "Home Page")
}

// func TestGetUsers(ctx echo.Context) error {
// 	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}
// 	pagination := &helper.Pagination{Row: 5}
// 	filter := &helper.UserFilter{}

// 	// *ANCHOR - ดึงค่าจาก QueryParams
// 	err := echo.QueryParamsBinder(ctx).
// 		Int("row", &pagination.Row).
// 		Int("page", &pagination.Page).
// 		String("sort", &pagination.Sort).
// 		String("firstname", &filter.Firstname).
// 		String("lastname", &filter.Lastname).
// 		String("email", &filter.Email).
// 		String("add", &filter.Address).
// 		BindError()
// 	if err != nil {
// 		return ctx.JSON(400, map[string]interface{}{"massage": "Error query param"})
// 	}

// 	users, err := userModelHelper.SelectAll(pagination, filter)
// 	if err != nil {
// 		return ctx.JSON(500, map[string]interface{}{"Error": err.Error()})
// 	}
// 	return ctx.JSON(200, map[string]interface{}{"data": users, "pagination": pagination, "message": "success"})
// }

// func TestCreateUsers(ctx echo.Context) error {
// 	userModelHelper := users.DatabaseRequest{DB: database.DBMYSQL}
// 	now := time.Now()

// 	// ANCHOR -  - ดึงข้อมูลจาก Body มาใส่ตัวแปร
// 	// data := []users.Users{}
// 	mockdata := []users.Users{
// 		{
// 			ID:           "e5c72af8-965d-4c78-9c17-908bbf63754c",
// 			Firstname:    "Alice",
// 			Lastname:     "Walker",
// 			Address:      "123 Main St, Springfield, USA",
// 			Email:        "alice.walker@example.com",
// 			Password:     "$2a$10$XziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9C",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 1,
// 		},
// 		{
// 			ID:           "fc31f1b4-8c2c-4b02-b3d4-d2e97a54b3e1",
// 			Firstname:    "Brian",
// 			Lastname:     "Smith",
// 			Address:      "456 Elm St, Springfield, USA",
// 			Email:        "brian.smith@example.com",
// 			Password:     "$2a$10$YziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9D",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 0,
// 		},
// 		{
// 			ID:           "2a3b3b94-9d0a-4f62-a56b-7f0c23fbc659",
// 			Firstname:    "Catherine",
// 			Lastname:     "Johnson",
// 			Address:      "789 Oak St, Springfield, USA",
// 			Email:        "catherine.johnson@example.com",
// 			Password:     "$2a$10$YxiGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9E",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 1,
// 		},
// 		{
// 			ID:           "f3d6d78e-2d2c-44b8-92ef-f7beef45678a",
// 			Firstname:    "David",
// 			Lastname:     "Brown",
// 			Address:      "101 Birch St, Springfield, USA",
// 			Email:        "david.brown@example.com",
// 			Password:     "$2a$10$ZziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9F",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 0,
// 		},
// 		{
// 			ID:           "1f56e8de-9d0a-4c62-a56b-7f0c23fbc567",
// 			Firstname:    "Emma",
// 			Lastname:     "Williams",
// 			Address:      "123 Maple St, Springfield, USA",
// 			Email:        "emma.williams@example.com",
// 			Password:     "$2a$10$AziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9G",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 1,
// 		},
// 		{
// 			ID:           "1f56e8de-9d0a-4c62-a56b-7f0c23fbc568",
// 			Firstname:    "Frank",
// 			Lastname:     "Jones",
// 			Address:      "456 Pine St, Springfield, USA",
// 			Email:        "frank.jones@example.com",
// 			Password:     "$2a$10$BziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9H",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 0,
// 		},
// 		{
// 			ID:           "2f56e8de-9d0a-4c62-a56b-7f0c23fbc569",
// 			Firstname:    "Grace",
// 			Lastname:     "Taylor",
// 			Address:      "789 Spruce St, Springfield, USA",
// 			Email:        "grace.taylor@example.com",
// 			Password:     "$2a$10$CziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9I",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 1,
// 		},
// 		{
// 			ID:           "3f56e8de-9d0a-4c62-a56b-7f0c23fbc570",
// 			Firstname:    "Henry",
// 			Lastname:     "Garcia",
// 			Address:      "101 Willow St, Springfield, USA",
// 			Email:        "henry.garcia@example.com",
// 			Password:     "$2a$10$DziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9J",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 0,
// 		},
// 		{
// 			ID:           "4f56e8de-9d0a-4c62-a56b-7f0c23fbc571",
// 			Firstname:    "Isabella",
// 			Lastname:     "Martinez",
// 			Address:      "456 Cedar St, Springfield, USA",
// 			Email:        "isabella.martinez@example.com",
// 			Password:     "$2a$10$EziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9K",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 1,
// 		},
// 		{
// 			ID:           "5f56e8de-9d0a-4c62-a56b-7f0c23fbc572",
// 			Firstname:    "James",
// 			Lastname:     "Rodriguez",
// 			Address:      "789 Oakwood St, Springfield, USA",
// 			Email:        "james.rodriguez@example.com",
// 			Password:     "$2a$10$FziGFw2NIQ9Cu6g1E3fKb.NYJCv7RzO0J6qfMFrt2lGDn8ixV1e9L",
// 			CreatedAt:    &now,
// 			UpdatedAt:    now,
// 			DeletedAt:    nil,
// 			PermissionID: 0,
// 		},
// 	}
// 	err := ctx.Bind(&mockdata)
// 	if err != nil {
// 		return ctx.JSON(400, map[string]interface{}{"message": "Invalid request body"})
// 	}
// 	users := []*users.Users{}
// 	for _, user := range mockdata {
// 		user.ID = helper.GenerateUUID()
// 		user.CreatedAt = &now
// 		user.UpdatedAt = now
// 		users = append(users, &user)
// 	}
// 	errors := userModelHelper.InsertArray(users)
// 	if errors != nil {
// 		return ctx.JSON(400, map[string]interface{}{"message": "Invalid insert body"})
// 	}

// 	return ctx.JSON(200, map[string]interface{}{"message": "User created successfully"})
// }
