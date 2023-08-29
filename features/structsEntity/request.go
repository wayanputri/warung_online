package structsEntity

type UserRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Alamat   string `json:"alamat" form:"alamat"`
	Password string `json:"password" form:"password"`
	NoTlp    string `json:"no_tlp" form:"no_tlp"`
	Gender   string `json:"gender" form:"gender"`
}
type Login struct {
	Email    string `json:"email" form:"email" validate:"required, email"`
	Password string `json:"password" form:"password"`
}
type ProductRequest struct {
	Nama      string `json:"nama_barang" form:"nama_barang"`
	Kategori  string `json:"kategori" form:"kategori"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Stok      int    `json:"stok" form:"stok"`
	Harga     int    `json:"harga" form:"harga"`
}
type ImageProductRequest struct {
	ProductID uint `json:"product_id" from:"product_id"`
}

type TransactionRequest struct {
	ProductID  uint `json:"product_id" from:"product_id"`
	Jumlah     int  `json:"jumlah" form:"jumlah"`
	TotalHarga int  `json:"total" form:"total"`
}

type TransactionKeranjangRequest struct {
	TransactionID      uint `json:"transaction_id" form:"transaction_id"`
	TransactionFinalID uint `json:"transaction_final_id" form:"transaction_final_id"`
}

type TransactionFinalRequest struct {
	TotalHarga int    `json:"total_harga" form:"total_harga"`
	OrderID    string `json:"order_id" form:"order_id"`
}

type PaymentRequest struct {
	TransactionFinalID uint   `json:"transaction_final_id" form:"transaction_final_id"`
	Bank               string `json:"bank" form:"bank"`
}

type ReviewRequest struct {
	PaymentID  uint    `json:"column:payment_id" form:"payment_id"`
	TextReview string  `json:"text_review" form:"text_review"`
	Rating     float64 `json:"rating" form:"rating"`
}
