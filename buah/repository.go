package buah

import (
	"ProjectBuahIn/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Buah, error)
	FindByIDuser(ID int) (models.User, error)
	FindByNama(nama string) (Buah, error)
	FindByID(ID int) (Buah, error)
	Create(buah Buah) (Buah, error)
	Update(buah Buah) (Buah, error)
	Delete(buah Buah) (Buah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Buah, error) {
	var buahs []Buah

	err := r.db.Find(&buahs).Error

	return buahs, err
}

func (r *repository) FindByID(ID int) (Buah, error) {
	var buah Buah

	err := r.db.Find(&buah, ID).Error

	return buah, err
}

func (r *repository) FindByIDuser(ID int) (models.User, error) {
	var user models.User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *repository) Create(buah Buah) (Buah, error) {
	err := r.db.Create(&buah).Error

	return buah, err
}

func (r *repository) Update(buah Buah) (Buah, error) {
	err := r.db.Save(&buah).Error
	return buah, err
}

func (r *repository) Delete(buah Buah) (Buah, error) {
	err := r.db.Delete(&buah).Error
	return buah, err
}

func (r *repository) FindByNama(nama string) (Buah, error) {
	var buah Buah

	err := r.db.Find(&buah, nama).Error

	return buah, err
}

func DB() *gorm.DB {

	var err error
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	db.AutoMigrate(&models.User{}, &Buah{}, &Order{}, &Cart{})
	return db

}

type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

// NewOrderRepository --> returns new order repository
func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: DB(),
	}
}

func (db *orderRepository) OrderProduct(userID int, productID int, quantity int) error {
	return db.connection.Create(&Order{
		ProductID: uint(productID),
		UserID:    uint(userID),
		Quantity:  quantity,
	}).Error

}
