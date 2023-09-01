package structsEntity

import "time"

type UserResponse struct {
	Id           uint          `json:"id,omitempty" form:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	Nama         string        `json:"nama,omitempty" form:"nama"`
	Email        string        `json:"email,omitempty" form:"email" validate:"required,email"`
	Alamat       string        `json:"alamat,omitempty" form:"alamat"`
	NoTlp        string        `json:"no_tlp,omitempty" form:"no_tlp"`
	Gender       string        `json:"gender,omitempty" form:"gender"`
	Role         string        `json:"role,omitempty" form:"role"`
	Profil       string        `json:"profil,omitempty" form:"profil"`
	Products     []ProductResponse     `json:"products,omitempty"`
	Transactions []TransactionResponse `json:"transaction,omitempty"`
}

type ProductResponse struct {
	Id           uint           `json:"id,omitempty" form:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UserID       uint           `json:"user_id,omitempty" form:"user_id"`
	Nama         string         `json:"nama_barang,omitempty" form:"nama_barang"`
	Kategori     string         `json:"kategori,omitempty" form:"kategori"`
	Deskripsi    string         `json:"deskripsi,omitempty" form:"deskripsi"`
	Stok         int            `json:"stok,omitempty" form:"stok"`
	Harga        int            `json:"harga,omitempty" form:"harga"`
	Ratings      float64        `json:"ratings,omitempty" form:"ratings"`
	Users        UserResponse           `json:"users,omitempty"`
	Image        []ImageProductResponse `json:"image,omitempty"`
	Transactions []TransactionResponse  `json:"transactions,omitempty"`
	Review       []ReviewResponse       `json:"reviews,omitempty"`
}

type ImageProductResponse struct {
	Id        uint      `json:"id,omitempty" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ProductID uint      `json:"product_id,omitempty" form:"product_id"`
	Link      string    `json:"link,omitempty" form:"link"`
	Products  ProductResponse   `json:"products,omitempty"`
}

type TransactionResponse struct {
	Id                uint             `json:"id,omitempty" form:"id"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	ProductID         uint             `json:"product_id,omitempty" from:"product_id"`
	UserID            uint             `json:"user_id,omitempty" form:"user_id"`
	Jumlah            int              `json:"jumlah,omitempty" form:"jumlah"`
	TotalHarga        int              `json:"total,omitempty" form:"total"`
	Products          ProductResponse          `json:"products,omitempty"`
	TransactionKeranjang []TransactionKeranjangResponse `json:"transaction_keranjang,omitempty"`
	Users             UserResponse             `json:"users,omitempty"`
}

type TransactionKeranjangResponse struct{
	Id            		uint      `json:"id" form:"id"`
	CreatedAt         time.Time   `json:"created_at,omitempty"`
	TransactionID 		uint 		`json:"transaction_id" form:"transaction_id"`
	Transaction 		TransactionResponse `json:"transaction,omitempty"`
	TransactionFinalID 	uint 		`json:"transaction_final_id" form:"transaction_final_id"`
	TransactionFinal 	TransactionFinalResponse `json:"transaction_final,omitempty"`
}


type TransactionFinalResponse struct {
	Id            uint      `json:"id,omitempty" form:"id"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	TotalHarga    int       `json:"total_harga,omitempty" form:"total_harga"`
	OrderID       string    `json:"order_id,omitempty" form:"order_id"`
	TransactionKeranjang  []TransactionKeranjangResponse `json:"transaction_keranjang,omitempty"`
	Payment       []PaymentResponse `json:"payments,omitempty"`
}

type PaymentResponse struct {
	Id                 uint             `json:"id,omitempty" form:"id"`
	CreatedAt          time.Time        `json:"created_at,omitempty"`
	TransactionFinalID uint             `json:"transaction_final_id,omitempty" form:"transaction_final_id"`
	Bank               string           `json:"bank,omitempty" form:"bank"`
	Status             string           `json:"status,omitempty" form:"status"`
	VA                 string           `json:"VA,omitempty" form:"VA"`
	TransactionFinals  TransactionFinalResponse `json:"transaction_final,omitempty"`
	Reviews            []ReviewResponse         `json:"reviews,omitempty"`
}

type ReviewResponse struct {
	Id           uint          `json:"id,omitempty" form:"id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	PaymentID    uint          `json:"payment_id,omitempty" form:"payment_id"`
	ProductID    uint			`json:"product_id" form:"product_id"`
	TextReview   string        `json:"text_review,omitempty" form:"text_review"`
	Rating       float64       `json:"rating,omitempty" form:"rating"`
	Payments     PaymentResponse       `json:"payments,omitempty"`
	Products     ProductResponse `json:"products,omitempty"`
	ImageReviews []ImageReviewResponse `json:"image_reviews,omitempty"`
}

type ImageReviewResponse struct {
	Id        uint      `json:"id,omitempty" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ReviewID  uint      `json:"review_id,omitempty" form:"review_id"`
	Link      string    `json:"link,omitempty" form:"link"`
	Reviews   ReviewResponse    `json:"reviews,omitempty"`
}

type UserResponseAll struct {
	Id           uint          `json:"id,omitempty" form:"id"`
	Nama         string        `json:"nama,omitempty" form:"nama"`
	Alamat       string        `json:"alamat,omitempty" form:"alamat"`
	Role         string        `json:"role,omitempty" form:"role"`
	Profil       string        `json:"profil,omitempty" form:"profil"`
}

type ProductResponseAll struct {
	Id           uint           `json:"id,omitempty" form:"id"`
	Nama         string         `json:"nama_barang,omitempty" form:"nama_barang"`
	Harga        int            `json:"harga,omitempty" form:"harga"`
	Ratings      float64        `json:"ratings,omitempty" form:"ratings"`
	Image        []ImageProductResponseAll `json:"image,omitempty"`
	Review       []ReviewResponseAll `json:"review,omitempty"`
}


type ImageProductResponseAll struct {
	Id        uint      `json:"id,omitempty" form:"id"`
	Link      string    `json:"link,omitempty" form:"link"`
	
}

type TransactionResponseAll struct {
	Id                uint             `json:"id,omitempty" form:"id"`
	Jumlah            int              `json:"jumlah,omitempty" form:"jumlah"`
	TotalHarga        int              `json:"total,omitempty" form:"total"`
}

type TransactionKeranjangResponseAll struct{
	Id            		uint      `json:"id" form:"id"`
	TransactionID 		uint 		`json:"transaction_id" form:"transaction_id"`
	Transaction 		TransactionResponseAll `json:"transaction,omitempty"`
	TransactionFinalID 	uint 		`json:"transaction_final_id" form:"transaction_final_id"`
	TransactionFinal 	TransactionFinalResponseAll `json:"transaction_final,omitempty"`
}

type TransactionFinalResponseAll struct {
	Id            uint      `json:"id,omitempty" form:"id"`
	TotalHarga    int       `json:"total_harga,omitempty" form:"total_harga"`
	OrderID       string    `json:"order_id,omitempty" form:"order_id"`
}

type PaymentResponseAll struct {
	Id                 uint             `json:"id,omitempty" form:"id"`
	Status             string           `json:"status,omitempty" form:"status"`
}

type ReviewResponseAll struct {
	Id           uint          `json:"id,omitempty" form:"id"`
	TextReview   string        `json:"text_review,omitempty" form:"text_review"`
	Rating       float64       `json:"rating,omitempty" form:"rating"`
	ProductID 	 uint 			`json:"product_id" form:"product_id"`
}

type ImageReviewResponseAll struct {
	Id        uint      `json:"id,omitempty" form:"id"`
	Link      string    `json:"link,omitempty" form:"link"`
}