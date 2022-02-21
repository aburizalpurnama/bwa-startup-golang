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



### Testing ###
- Menggunakan library testify
	- instalation : https://github.com/stretchr/testify#installation
	- instruction : https://github.com/stretchr/testify#assert-package
	

