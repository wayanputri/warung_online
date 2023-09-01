package structsEntity

func UserRequestToEntity(user UserRequest) UserEntity {
	return UserEntity{
		Nama:     user.Nama,
		Email:    user.Email,
		Alamat:   user.Alamat,
		Password: user.Password,
		NoTlp:    user.NoTlp,
		Gender:   user.Gender,
	}
}
func LoginToUserEntity(login Login) UserEntity {
	return UserEntity{
		Email:    login.Email,
		Password: login.Password,
	}
}

func ProductRequestToEntity(product ProductRequest) ProductEntity {
	return ProductEntity{
		Nama:      product.Nama,
		Kategori:  product.Kategori,
		Deskripsi: product.Deskripsi,
		Stok:      product.Stok,
		Harga:     product.Harga,
	}
}
func ImageProductRequestToEntity(imageProduct ImageProductRequest) ImageProductEntity {
	return ImageProductEntity{
		ProductID: imageProduct.ProductID,
	}
}

func TransactionRequestToEntity(transaction TransactionRequest) TransactionEntity {
	return TransactionEntity{
		ProductID:  transaction.ProductID,
		Jumlah:     transaction.Jumlah,
		TotalHarga: transaction.TotalHarga,
	}
}

func TransactionKeranjangRequestToEntity(transactionKeranjang TransactionKeranjangRequest) TransactionKeranjangEntity {
	return TransactionKeranjangEntity{
		TransactionID:      transactionKeranjang.TransactionID,
		TransactionFinalID: transactionKeranjang.TransactionFinalID,
	}
}

func TransactionFinalRequestToEntity(transactionFinal TransactionFinalRequest) TransactionFinalEntity {
	return TransactionFinalEntity{
		TotalHarga: transactionFinal.TotalHarga,
		OrderID:    transactionFinal.OrderID,
	}
}

func PaymentRequestToEntity(payment PaymentRequest) PaymentEntity {
	return PaymentEntity{
		TransactionFinalID: payment.TransactionFinalID,
		Bank:               payment.Bank,
	}
}

func ReviewRequestToEntity(review ReviewRequest) ReviewEntity {
	return ReviewEntity{
		PaymentID:  review.PaymentID,
		ProductID:  review.ProductID,
		TextReview: review.TextReview,
		Rating:     review.Rating,
	}
}