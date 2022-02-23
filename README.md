### 1. Entity Analysis ####
- Menganalisa kebutuhan entity dari business proses

### 2. ERD Analysis ###
- Menganalisa Relational antar Entrity
- Mandatory : suatu entity diharuskan memiliki keterkaitan dengan entity lain
- Optional : Suatu entity tidak diharuskan memiliki hubungan dengan entity lain
- dapat menggunakan tools : https://erdplus.com/
- Nama entity menggunakan plural dari kata benda

		- User -<<create>>- Campaign (Mandatory <-> Optional) (One to Many)
		- Campaign -<<has>>- Campain Image (Mandatory <-> Mandatory) (One to Many)
		- User -<<has>>- transaction (Mandatory <-> Optional) (One to Many)
		- Campaign -<<create>>- transaction (Mandatory <-> Optional) (One to Many)
		
### Instal postgreeSql ###

- Instalation :
https://www.niagahoster.co.id/blog/cara-install-postgresql-di-ubuntu-18-04/?amp&gclid=CjwKCAiA6seQBhAfEiwAvPqu15g5e40OBMaCl1sIcnUCUzire_GxlajAmnpXifEMima3OZID7jLNvhoCbPMQAvD_BwE

- Check host & port
	sudo netstat -plunt |grep postgres
	
- Add auto increment with navivate
	https://www.youtube.com/watch?v=dwQYjg4gl58
	
- add foreign key
	https://www.youtube.com/watch?v=LjR0X_T2JBo

## 3.Field Analysis ##
	Users :
		- id : int
		- name : varchar
		- occupation : varchar
		- email : varchar
		- password_hash : varchar
		- avatar_file_name : varchar
		- role : varchar
		- token : varchar
		- created_at : datetime
		- updated_at : datetime
		
	Campaigns :
		- id : int
		- user_id : int
		- name : varchar
		- short_description : varchar
		- description : text
		- goal_amount : int
		- current_amaount : int
		- perks : text
		- becker_count : int
		- slug : varchar ? slug digunakan untuk mengganti nomor id di url saat mengakses web, manfaatnya agar SEO friendly
		- created_at : datetime
		- updated_at : datetime
	
	Campaign Images:
		- id : int
		- campaign_id : int
		- file_name : varchar
		- is_primary : boolean (tinyint)
		- created_at : datetime
		- updated_at : datetime		
	
	Transactions :
		- id : int
		- campaign_id : int
		- user_id : int
		- amount : int
		- status : varchar ? sebuah status pembayaran (belum dibayar, lunas, dll)
		- code : varchar ? sebuah kode transaksi
		- created_at : datetime
		- updated_at : datetime
		
#### 4. Setup Backend API ###
- Menggunakan framework Gin (web framework)
	- instalation : https://github.com/gin-gonic/gin#installation
		
- Menggunakan framework gorm (db access framework)
	- instalation : https://gorm.io/docs/index.html#Install
	- untuk instalasi driver, harus disesuaikan dengan DBMS yang digunakan
		
- Membuat koneksi ke database
	- instruction : https://gorm.io/docs/connecting_to_the_database.html
	
### 5. Entity Maping ####
- Membuat struct sesuai dengan entity / table pada database
	Nama dari struct harus merupakan nama tunggal dari table
	Contoh : struct = user -> table = users
	
- Test Retrieving data from db data 

	func TestRetrieveUserTableFromDB(t *testing.T) {

		// get database connection
		dsn := "host=localhost user=rizal password=3748 dbname=db_startup_bwa port=5432 sslmode=disable TimeZone=UTC"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		// pastikan error == nil
		assert.Nil(t, err)

		// buat object array dari entity struct
		var users []user.User

		// pastikan object kosong sebelum query ke db
		assert.Equal(t, 0, len(users))

		// query ke db
		if assert.NotNil(t, db) {
			db.Find(&users)
		}

		// pastikan object user tidak nil lagi
		assert.NotEqual(t, 0, &users)

		for _, user := range users {
			fmt.Println(user.Name)
		}

	}

## 6. Implement Repository layer for Users

- Create file repository.go on user package

- Create interface (contract) for repository struct

	type Repository interface {
		Save(user User) (User, error)
	}

- Create implement struct with db *gorm.DB field

	type repository struct {
		db *gorm.DB
	}

- Create method for instance the object

	func NewRepository(db *gorm.DB) *repository {
		return &repository{db}
	}

- Create contract method

	func (r *repository) Save(user User) (User, error) {
		err := r.db.Create(&user).Error

		if err != nil {
			return user, err
		}

		return user, nil
	}

## 7. Implement Service layer for Users

- Create file service.go on user package

- Create interface (contract) for repository struct

	type Service interface {
		RegisterUser(input RegisterUserInput) (User, error)
	}

- Create implement struct with db *gorm.DB field

	type service struct {
		repository Repository
	}

- Create method for instance the object

	func NewService(repository Repository) *service {
		return &service{repository}
	}

- Create contract method

	Inserted password must be Hashed before passing to repository

	func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
		user := User{}
		user.Name = input.Name
		user.Email = input.Email
		user.Occupation = input.Occupation

		passHashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

		if err != nil {
			return user, err
		}

		user.PasswordHash = string(passHashed)

		user.Role = "user"

		newUser, err := s.repository.Save(user)

		if err != nil {
			return newUser, err
		}

		return newUser, nil
	}


### 8. Create inputUser struct
	
Input struct digunakan untuk menampung data yang diinputkan oleh user atau untuk mapping data dari sebuah request

- Create input.go on user package

	type RegisterUserInput struct {
		Name       string
		Occupation string
		Email      string
		Password   string
	}

## 9. Implement handler a.k.a controller for User

handler digunakan untuk mapping API request

- Create user.go on handler package

- Create struct

	type UserHandler struct {
		userService user.Service
	}

- Create method for instance the object

	func NewUserHandler(userService user.Service) *UserHandler {
		return &UserHandler{userService}
	}

- Create handler method

	func (h *UserHandler) RegisterUser(c *gin.Context) {
		// tangkap input dari user
		// map input ke struct RegisterUserInput
		// passing struct diatas kedalam service

		input := user.RegisterUserInput{}

		err := c.ShouldBindJSON(&input)

		if err != nil {
			log.Fatal(err.Error)
			c.JSON(http.StatusBadRequest, input)
		}

		user, registErr := h.userService.RegisterUser(input)

		if registErr != nil {
			log.Fatal(registErr.Error)
			c.JSON(http.StatusBadRequest, input)
		}

		c.JSON(http.StatusOK, user)
	}
## Test in main.go

Create each needed component and run the program

	func main() {
		db, err := utils.GetDb()

		if err != nil {
			log.Fatal(err.Error)
		}

		userRepository := user.NewRepository(db)

		userService := user.NewService(userRepository)

		userHandler := handler.NewUserHandler(userService)

		router := gin.Default()

		// API versioning
		api := router.Group("/api/v1")

		api.POST("/nasabah", userHandler.RegisterUser)

		router.Run()
	}

### Testing ###
- Sebuah file golang akan secara otomatis terdeteksi sebagai Testing file jika ada _test.go

	contoh : main_test.go
	
- Menggunakan library testify
	- instalation : https://github.com/stretchr/testify#installation
	- instruction : https://github.com/stretchr/testify#assert-package
	

