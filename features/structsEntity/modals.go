package structsEntity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama 		string `gorm:"column:nama;not null"`
	Email 		string `gorm:"column:email;unique;not null"`
	Password 	string `gorm:"column:password;not null"`
	Alamat 		string `gorm:"column:alamat;not null"`
	NoTlp 		string `gorm:"column:no_tlp;not null"`
	Gender 		string `gorm:"type:enum('male','female');default:'female';column:gender;not null"`
	Role 		string `gorm:"type:enum('pedagang','pembeli');default:'pembeli';column:role;not null"`
	Profil 		string `gorm:"column:profil;null"`
	DataUpdate 	string `gorm:"column:data_update"`
	Products 		[]Product 		`gorm:"foreignKey:UserID"`
	Transactions 	[]Transaction 	`gorm:"foreignKey:UserID"`
}

type Product struct{
	gorm.Model
	UserID 		uint 	`gorm:"column:user_id;not null"`
	Nama 		string 	`gorm:"column:nama_barang;not null"`
	Kategori 	string 	`gorm:"column:kategori;not null"`
	Deskripsi 	string 	`gorm:"column:deskripsi"`
	Stok 		int 	`gorm:"column:stok;not null"`
	Harga 		int 	`gorm:"column:harga;not null"`
	Ratings 	float64	`gorm:"column:ratings;default:0"`
	Users		User    `gorm:"foreignKey:UserID"`
	Image		[]ImageProduct `gorm:"foreignKey:ProductID"`
	Transaction []Transaction	`gorm:"foreignKey:ProductID"`
	Reviews 	[]Review `gorm:"foreignKey:ProductID"`
}

type ImageProduct struct{
	gorm.Model
	ProductID uint 		`gorm:"column:product_id;not null"`
	Link	  string 	`gorm:"column:link;not null"`
	Products  Product	`gorm:"foreignKey:ProductID"`
}

type Transaction struct{
	gorm.Model
	ProductID 			uint 	`gorm:"column:product_id;not null"`
	UserID 				uint 	`gorm:"column:user_id;not null"`
	Jumlah 				int 	`gorm:"column:jumlah;not null"`
	TotalHarga 			int 	`gorm:"column:total_harga"`
	Products  			Product	`gorm:"foreignKey:ProductID"`
	TransactionKeranjang []TransactionKeranjang `gorm:"foreignKey:TransactionID"`
	Users				User	`gorm:"foreignKey:UserID"`
	
}
type TransactionKeranjang struct{
	gorm.Model
	TransactionID 		uint 				`gorm:"column:transaction_id"`
	Transaction 		Transaction 		`gorm:"foreignKey:TransactionID"`
	TransactionFinalID 	uint 				`gorm:"column:transaction_final_id;default:1"`
	TransactionFinal 	TransactionFinal 	`gorm:"foreignKey:TransactionFinalID"`
}

type TransactionFinal struct{
	gorm.Model
	TotalHarga 		int 	`gorm:"column:total_harga;not null"`
	OrderID 		string 	`gorm:"column:order_id;not null"`
	TransactionKeranjang    []TransactionKeranjang `gorm:"foreignKey:TransactionFinalID"`
	Payment 		[]Payment `gorm:"foreignKey:TransactionFinalID"`
}

type Payment struct{
	gorm.Model
	TransactionFinalID 	uint 	`gorm:"column:transaction_final_id;not null"`
	Bank 				string 	`gorm:"column:bank;not null"`
	Status 				string 	`gorm:"column:status;not null"`
	VA 					string 	`gorm:"column: VA;not null"`
	TransactionFinals	TransactionFinal `gorm:"foreignKey:TransactionFinalID"` 
	Reviews				[]Review	`gorm:"foreignKey:PaymentID"`	
}

type Review struct{
	gorm.Model
	PaymentID 	uint 	`gorm:"column:payment_id;not null"`
	ProductID   uint	`gorm:"column:product_id;not null"`
	TextReview 	string 	`gorm:"text_review"`
	Rating 		float64 `gorm:"rating;not null"`
	Product		Product `gorm:"foreignKey:ProductID"`
	Payments	Payment `gorm:"foreignKey:PaymentID"`
	ImageReviews []ImageReview `gorm:"foreignKey:ReviewID"`
}

type ImageReview struct{
	gorm.Model
	ReviewID 	uint 	`gorm:"column:review_id;not null"`
	Link 		string 	`gorm:"column:link;not null"`
	Reviews 	Review	`gorm:"foreignKey:ReviewID"`
}