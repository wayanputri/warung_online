package structsEntity

import "time"

type UserEntity struct {
	Id           uint          `json:"id" form:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	DeletedAt    time.Time     `json:"deleted_at,omitempty"`
	Nama         string        `json:"nama" form:"nama"`
	Email        string        `json:"email" form:"email" validate:"required,email"`
	Password     string        `json:"password" form:"password"`
	Alamat       string        `json:"alamat" form:"alamat"`
	NoTlp        string        `json:"no_tlp" form:"no_tlp"`
	Gender       string        `json:"gender" form:"gender"`
	Role         string        `json:"role" form:"role"`
	Profil       string        `json:"profil,omitempty" form:"profil"`
	DataUpdate   string        `json:"data_update" form:"data_update"`
	Products     []ProductEntity     `json:"products,omitempty"`
	Transactions []TransactionEntity `json:"transaction,omitempty"`
}

type ProductEntity struct {
	Id           uint           `json:"id" form:"id"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    time.Time      `json:"deleted_at,omitempty"`
	UserID       uint           `json:"user_id" form:"user_id"`
	Nama         string         `json:"nama_barang" form:"nama_barang"`
	Kategori     string         `json:"kategori" form:"kategori"`
	Deskripsi    string         `json:"deskripsi" form:"deskripsi"`
	Stok         int            `json:"stok" form:"stok"`
	Harga        int            `json:"harga" form:"harga"`
	Ratings      float64        `json:"ratings" form:"ratings"`
	Users        UserEntity           `json:"users,omitempty"`
	Image        []ImageProductEntity `json:"image,omitempty"`
	Transactions []TransactionEntity  `json:"transactions,omitempty"`
	Review       []ReviewEntity `json:"reviews,omitempty"`
}

type ImageProductEntity struct {
	Id        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	ProductID uint      `json:"product_id" form:"product_id"`
	Link      string    `json:"link" form:"link"`
	Products  ProductEntity   `json:"products,omitempty"`
}

type TransactionEntity struct {
	Id                uint             `json:"id" form:"id"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
	DeletedAt         time.Time        `json:"deleted_at,omitempty"`
	ProductID         uint             `json:"product_id" from:"product_id"`
	UserID            uint             `json:"user_id" form:"user_id"`
	Jumlah            int              `json:"jumlah" form:"jumlah"`
	TotalHarga        int              `json:"total" form:"total"`
	Products          ProductEntity          `json:"products,omitempty"`
	TransactionKeranjang []TransactionKeranjangEntity `json:"transaction_keranjang,omitempty"`
	Users             UserEntity             `json:"users,omitempty"`
}

type TransactionKeranjangEntity struct{
	Id            		uint      `json:"id" form:"id"`
	CreatedAt     		time.Time `json:"created_at,omitempty"`
	UpdatedAt     		time.Time `json:"updated_at,omitempty"`
	DeletedAt     		time.Time `json:"deleted_at,omitempty"`
	TransactionID 		uint 		`json:"transaction_id" form:"transaction_id"`
	Transaction 		TransactionEntity `json:"transaction,omitempty"`
	TransactionFinalID 	uint 		`json:"transaction_final_id" form:"transaction_final_id"`
	TransactionFinal 	TransactionFinalEntity `json:"transaction_final,omitempty"`
}


type TransactionFinalEntity struct {
	Id            uint      `json:"id" form:"id"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	DeletedAt     time.Time `json:"deleted_at,omitempty"`
	TotalHarga    int       `json:"total_harga" form:"total_harga"`
	OrderID       string    `json:"order_id" form:"order_id"`
	TransactionKeranjang  []TransactionKeranjangEntity `json:"transaction_keranjang,omitempty"`
	Payment       []PaymentEntity `json:"payments,omitempty"`
}

type PaymentEntity struct {
	Id                 uint             `json:"id" form:"id"`
	CreatedAt          time.Time        `json:"created_at,omitempty"`
	UpdatedAt          time.Time        `json:"updated_at,omitempty"`
	DeletedAt          time.Time        `json:"deleted_at,omitempty"`
	TransactionFinalID uint             `json:"transaction_final_id" form:"transaction_final_id"`
	Bank               string           `json:"bank" form:"bank"`
	Status             string           `json:"status" form:"status"`
	VA                 string           `json:" VA" form:"VA"`
	TransactionFinals  TransactionFinalEntity `json:"transaction_final,omitempty"`
	Reviews            []ReviewEntity         `json:"reviews,omitempty"`
}

type ReviewEntity struct {
	Id           uint          `json:"id" form:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	DeletedAt    time.Time     `json:"deleted_at,omitempty"`
	PaymentID    uint          `json:"column:payment_id" form:"payment_id"`
	TextReview   string        `json:"text_review" form:"text_review"`
	Rating       float64       `json:"rating" form:"rating"`
	ProductID     uint         `json:"product_id" form:"product_id"`
	Payments     PaymentEntity       `json:"payments,omitempty"`
	Products     ProductEntity       `json:"products,omitempty"`
	ImageReviews []ImageReviewEntity `json:"image_reviews,omitempty"`
}

type ImageReviewEntity struct {
	Id        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	ReviewID  uint      `json:"review_id" form:"review_id"`
	Link      string    `json:"link" form:"link"`
	Reviews   ReviewEntity    `json:"reviews,omitempty"`
}