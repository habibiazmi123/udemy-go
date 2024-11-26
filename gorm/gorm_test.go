package belajar_golang_gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:teuingatuh@tcp(localhost:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(name) values(?)", "Azmi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(name) values(?)", "Azmi Cumi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(name) values(?)", "Azmi X").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(name) values(?)", "Azmi Y").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   int
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, sample.Id)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       1,
		Password: "rahasia",
		Name: Name{
			FirstName:  "Muhamad",
			MiddleName: "Habibi",
			LastName:   "Azmi",
		},
		Information: "Ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, 1, int(response.RowsAffected))
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID: int64(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Password: "rahasia",
		})
	}
	result := db.CreateInBatches(&users, 1000)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: 10, Password: "rahasia", Name: Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 2, user.ID)

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 10, user.ID)
}

func TestQuerySingleInlineCondition(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", 5)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}

func TestQueryAllObject(t *testing.T) {
	var users []User
	result := db.Find(&users, "id in ?", []int{3, 4, 5})
	assert.Nil(t, result.Error)
	assert.Equal(t, 3, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 9, len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").
		Or("password = ?", "rahasia").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 9, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User
	result := db.Not("first_name like ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 1, len(users))
}

func TestSelect(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").
		Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 9, len(users))
}

func TestStructConditions(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
		},
	}
	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

func TestMapConditions(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
	}
	var users []User
	result := db.Where(mapCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, len(users))
}

func TestLimitOrder(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
	assert.Equal(t, 7, users[0].ID)
}

type UserResponse struct {
	ID        int
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	result := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 9, len(users))
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", 2)
	assert.Nil(t, result.Error)

	user.Name.FirstName = "John wick"
	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestSelectedColumns(t *testing.T) {
	result := db.Model(&User{}).Where("id = ?", 1).Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "oke sih",
	})
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", 1).Update("password", "teuingatuh")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", 1).Updates(User{
		Name: Name{
			FirstName:  "Muhamad",
			MiddleName: "Habibi",
			LastName:   "Azmi",
		},
	})
	assert.Nil(t, result.Error)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserID: 1,
			Action: "Test Action",
		}

		result := db.Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestUpdateOrCreate(t *testing.T) {
	userLog := UserLog{
		UserID: 1,
		Action: "Test Action",
	}
	result := db.Save(&userLog)
	assert.Nil(t, result.Error)

	userLog.UserID = 2
	result = db.Save(&userLog)
	assert.Nil(t, result.Error)
}

func TestSaveOrUpdateNonIncrement(t *testing.T) {
	user := User{
		ID: 99,
		Name: Name{
			FirstName: "Asep",
		},
	}
	result := db.Save(&user)
	assert.Nil(t, result.Error)

	user.Name.FirstName = "Asep Updated"
	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: 88,
		Name: Name{
			FirstName: "Ujang",
		},
	}

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)

	assert.Nil(t, result.Error)
}

func TestDelete(t *testing.T) {
	var user User
	result := db.First(&user, "id = ?", 99)
	assert.Nil(t, result.Error)
	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	result = db.Delete(&User{}, "id = ?", 88)
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", 77).Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserID:      1,
		Title:       "Todo 12",
		Description: "Isi Todo 13",
	}

	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo Todo
	result := db.Unscoped().First(&todo, "id = ?", 4)
	assert.Nil(t, result.Error)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "id = ?", 10).Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Anwar"
		return tx.Save(&user).Error
	})

	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		UserId:  2,
		Balance: 100000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").First(&user).Error
	assert.Nil(t, err)

	assert.Equal(t, 2, user.ID)
	assert.Equal(t, 4, user.Wallet.ID)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var users []User
	err := db.Model(&User{}).Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)

	assert.Equal(t, 9, len(users))
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       11,
		Password: "rahasia",
		Name: Name{
			FirstName: "testing",
		},
		Wallet: Wallet{
			ID:      11,
			UserId:  11,
			Balance: 1000000,
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       13,
		Password: "rahasia",
		Name: Name{
			FirstName: "User 12",
		},
		Wallet: Wallet{
			ID:      13,
			UserId:  13,
			Balance: 1000000,
		},
	}
	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       14,
		Password: "rahasia",
		Name:     Name{FirstName: "User 14"},
		Wallet: Wallet{
			ID:      12,
			UserId:  14,
			Balance: 10000,
		},
		Addresses: []Address{
			{
				UserId:  14,
				Address: "Jl antapani",
			},
			{
				UserId:  14,
				Address: "Jl antapani 1",
			},
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var usersPreload []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&usersPreload).Error
	assert.Nil(t, err)
}

func TestBelongsToAddress(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(addresses))

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(addresses))
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateMany2Many(t *testing.T) {
	product := Product{
		Name:  "Product 1",
		Price: 10000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    2,
		"product_id": 1,
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    3,
		"product_id": 1,
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").Take(&product, "id = ?", 1).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	err := db.Preload("LikeProducts").Take(&user, "id = ?", 2).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationMode(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", 1).Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestAssociationAdd(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", 4).Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", 1).Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := db.First(&user, "id = ?", 2).Error
		assert.Nil(t, err)

		wallet := Wallet{
			UserId:  2,
			Balance: 2000,
		}

		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", 4).Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", 1).Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", 1).Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 1000).Take(&user, "id = ?", 14).Error
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestNestedPreloading(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Take(&wallet, "id = ?", 12).Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).First(&user, "id = ?", 2).Error
	assert.Nil(t, err)

	fmt.Println(user)
	fmt.Println(user.Wallet)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

func TestJoinQueryCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 9999).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 9999).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 1000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(3), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select(
		"sum(balance) as total_balance",
		"min(balance) as min_balance",
		"max(balance) as max_balance",
		"avg(balance) as avg_balance",
	).Take(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(1012000), result.TotalBalance)
	assert.Equal(t, int64(2000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(337333.3333), result.AvgBalance)

	fmt.Println(result)
}

func TestGroupByHaving(t *testing.T) {
	var result []AggregationResult
	err := db.Model(&Wallet{}).Select(
		"sum(balance) as total_balance",
		"min(balance) as min_balance",
		"max(balance) as max_balance",
		"avg(balance) as avg_balance",
	).
		Joins("User").
		Group("User.id").
		Having("sum(balance) > ?", 1000000).
		Find(&result).
		Error

	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestUserHook(t *testing.T) {
	user := User{
		ID:       15,
		Password: "secret",
		Name: Name{
			FirstName: "Mantap",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)

	assert.NotNil(t, user.Name.LastName)
	assert.NotEqual(t, "", user.ID)
}
