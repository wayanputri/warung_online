package structsEntity

func UserModelToEntity(user User) UserEntity {
	var products []ProductEntity
	for _, product := range user.Products {
		products = append(products, ProductModelToEntity(product))
	}
	var transactions []TransactionEntity
	for _, transaction := range user.Transactions {
		transactions = append(transactions, TransactionModelToEntity(transaction))
	}
	return UserEntity{
		Id:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt.Time,
		Nama:         user.Nama,
		Email:        user.Email,
		Password:     user.Password,
		Alamat:       user.Alamat,
		NoTlp:        user.NoTlp,
		Gender:       user.Gender,
		Role:         user.Role,
		Profil:       user.Profil,
		DataUpdate:   user.DataUpdate,
		Products:     products,
		Transactions: transactions,
	}
}

func ProductModelToEntity(product Product) ProductEntity {
	var images []ImageProductEntity
	for _, image := range product.Image {
		images = append(images, ImageProductModelToEntity(image))
	}
	var transactions []TransactionEntity
	for _, transaction := range product.Transaction {
		transactions = append(transactions, TransactionModelToEntity(transaction))
	}
	return ProductEntity{
		Id:           product.ID,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		DeletedAt:    product.DeletedAt.Time,
		UserID:       product.UserID,
		Nama:         product.Nama,
		Kategori:     product.Kategori,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		Harga:        product.Harga,
		Ratings:      product.Ratings,
		Users:        UserModelToEntity(product.Users),
		Image:        images,
		Transactions: transactions,
	}
}

func ImageProductModelToEntity(imageProduct ImageProduct) ImageProductEntity {
	return ImageProductEntity{
		Id:        imageProduct.ID,
		CreatedAt: imageProduct.CreatedAt,
		UpdatedAt: imageProduct.UpdatedAt,
		DeletedAt: imageProduct.DeletedAt.Time,
		ProductID: imageProduct.ProductID,
		Link:      imageProduct.Link,
		Products:  ProductModelToEntity(imageProduct.Products),
	}
}

func TransactionModelToEntity(transaction Transaction) TransactionEntity {
	var keranjang []TransactionKeranjangEntity
	for _, value := range transaction.TransactionKeranjang {
		keranjang = append(keranjang, TransactionKeranjangModelToEntity(value))
	}
	return TransactionEntity{
		Id:                   transaction.ID,
		CreatedAt:            transaction.CreatedAt,
		UpdatedAt:            transaction.UpdatedAt,
		DeletedAt:            transaction.DeletedAt.Time,
		ProductID:            transaction.ProductID,
		UserID:               transaction.UserID,
		Jumlah:               transaction.Jumlah,
		TotalHarga:           transaction.TotalHarga,
		Products:             ProductModelToEntity(transaction.Products),
		TransactionKeranjang: keranjang,
		Users:                UserModelToEntity(transaction.Users),
	}
}

func TransactionKeranjangModelToEntity(transactionKeranjang TransactionKeranjang) TransactionKeranjangEntity {
	return TransactionKeranjangEntity{
		Id:                 transactionKeranjang.ID,
		CreatedAt:          transactionKeranjang.CreatedAt,
		UpdatedAt:          transactionKeranjang.UpdatedAt,
		DeletedAt:          transactionKeranjang.DeletedAt.Time,
		TransactionID:      transactionKeranjang.TransactionID,
		Transaction:        TransactionModelToEntity(transactionKeranjang.Transaction),
		TransactionFinalID: transactionKeranjang.TransactionFinalID,
		TransactionFinal:   TransactionFinalModelToEntity(transactionKeranjang.TransactionFinal),
	}
}

func TransactionFinalModelToEntity(transactionFinal TransactionFinal) TransactionFinalEntity {
	var keranjang []TransactionKeranjangEntity
	for _, transaction := range transactionFinal.TransactionKeranjang {
		keranjang = append(keranjang, TransactionKeranjangModelToEntity(transaction))
	}
	var payments []PaymentEntity
	for _, payment := range transactionFinal.Payment {
		payments = append(payments, PaymentModelToEntity(payment))
	}
	return TransactionFinalEntity{
		Id:                   transactionFinal.ID,
		CreatedAt:            transactionFinal.CreatedAt,
		UpdatedAt:            transactionFinal.UpdatedAt,
		DeletedAt:            transactionFinal.DeletedAt.Time,
		TotalHarga:           transactionFinal.TotalHarga,
		OrderID:              transactionFinal.OrderID,
		TransactionKeranjang: keranjang,
		Payment:              payments,
	}
}

func PaymentModelToEntity(payment Payment) PaymentEntity {
	var reviews []ReviewEntity
	for _, review := range payment.Reviews {
		reviews = append(reviews, ReviewModelToEntity(review))
	}
	return PaymentEntity{
		Id:                 payment.ID,
		CreatedAt:          payment.CreatedAt,
		UpdatedAt:          payment.UpdatedAt,
		DeletedAt:          payment.DeletedAt.Time,
		TransactionFinalID: payment.TransactionFinalID,
		Bank:               payment.Bank,
		Status:             payment.Status,
		VA:                 payment.VA,
		TransactionFinals:  TransactionFinalModelToEntity(payment.TransactionFinals),
		Reviews:            reviews,
	}
}

func ReviewModelToEntity(review Review) ReviewEntity {
	var imageReview []ImageReviewEntity
	for _, image := range review.ImageReviews {
		imageReview = append(imageReview, ImageReviewModelToEntity(image))
	}
	return ReviewEntity{
		Id:           review.ID,
		CreatedAt:    review.CreatedAt,
		UpdatedAt:    review.UpdatedAt,
		DeletedAt:    review.DeletedAt.Time,
		PaymentID:    review.PaymentID,
		TextReview:   review.TextReview,
		Rating:       review.Rating,
		Payments:     PaymentModelToEntity(review.Payments),
		ImageReviews: imageReview,
	}
}

func ImageReviewModelToEntity(imageReview ImageReview) ImageReviewEntity {
	return ImageReviewEntity{
		Id:        imageReview.ID,
		CreatedAt: imageReview.CreatedAt,
		UpdatedAt: imageReview.UpdatedAt,
		DeletedAt: imageReview.DeletedAt.Time,
		ReviewID:  imageReview.ReviewID,
		Link:      imageReview.Link,
		Reviews:   ReviewModelToEntity(imageReview.Reviews),
	}
}