package structsEntity

func UserEntityToModel(user UserEntity) User {
	var products []Product
	for _, product := range user.Products {
		products = append(products, ProductEntityToModel(product))
	}
	var transactions []Transaction
	for _, transaction := range user.Transactions {
		transactions = append(transactions, TransactionEntityToModel(transaction))
	}
	return User{
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

func ProductEntityToModel(product ProductEntity) Product {
	var images []ImageProduct
	for _, image := range product.Image {
		images = append(images, ImageProductEntityToModel(image))
	}
	var transactions []Transaction
	for _, transaction := range product.Transactions {
		transactions = append(transactions, TransactionEntityToModel(transaction))
	}
	return Product{
		UserID:      product.UserID,
		Nama:        product.Nama,
		Kategori:    product.Kategori,
		Deskripsi:   product.Deskripsi,
		Stok:        product.Stok,
		Harga:       product.Harga,
		Ratings:     product.Ratings,
		Users:       UserEntityToModel(product.Users),
		Image:       images,
		Transaction: transactions,
	}
}

func ImageProductEntityToModel(imageProduct ImageProductEntity) ImageProduct {
	return ImageProduct{
		ProductID: imageProduct.ProductID,
		Link:      imageProduct.Link,
		Products:  ProductEntityToModel(imageProduct.Products),
	}
}

func TransactionEntityToModel(transaction TransactionEntity) Transaction {
	var transactions []TransactionKeranjang
	for _, transaction := range transaction.TransactionKeranjang {
		transactions = append(transactions, TransactionKeranjangEntityToModel(transaction))
	}
	return Transaction{
		ProductID:            transaction.ProductID,
		UserID:               transaction.UserID,
		Jumlah:               transaction.Jumlah,
		TotalHarga:           transaction.TotalHarga,
		Products:             ProductEntityToModel(transaction.Products),
		TransactionKeranjang: transactions,
		Users:                UserEntityToModel(transaction.Users),
	}
}

func TransactionKeranjangEntityToModel(transactionKeranjang TransactionKeranjangEntity) TransactionKeranjang {
	return TransactionKeranjang{
		TransactionID:      transactionKeranjang.TransactionID,
		Transaction:        TransactionEntityToModel(transactionKeranjang.Transaction),
		TransactionFinalID: transactionKeranjang.TransactionFinalID,
		TransactionFinal:   TransactionFinalEntityToModel(transactionKeranjang.TransactionFinal),
	}
}

func TransactionFinalEntityToModel(transactionFinal TransactionFinalEntity) TransactionFinal {
	var transactions []TransactionKeranjang
	for _, transaction := range transactionFinal.TransactionKeranjang {
		transactions = append(transactions, TransactionKeranjangEntityToModel(transaction))
	}
	var payments []Payment
	for _, payment := range transactionFinal.Payment {
		payments = append(payments, PaymentEntityToModel(payment))
	}
	return TransactionFinal{
		TotalHarga:           transactionFinal.TotalHarga,
		OrderID:              transactionFinal.OrderID,
		TransactionKeranjang: transactions,
		Payment:              payments,
	}
}

func PaymentEntityToModel(payment PaymentEntity) Payment {
	var reviews []Review
	for _, review := range payment.Reviews {
		reviews = append(reviews, ReviewEntityToModel(review))
	}
	return Payment{
		TransactionFinalID: payment.TransactionFinalID,
		Bank:               payment.Bank,
		Status:             payment.Status,
		VA:                 payment.VA,
		TransactionFinals:  TransactionFinalEntityToModel(payment.TransactionFinals),
		Reviews:            reviews,
	}
}

func ReviewEntityToModel(review ReviewEntity) Review {
	var imageReview []ImageReview
	for _, image := range review.ImageReviews {
		imageReview = append(imageReview, ImageReviewEntityToModel(image))
	}
	return Review{
		PaymentID:    review.PaymentID,
		TextReview:   review.TextReview,
		Rating:       review.Rating,
		Payments:     PaymentEntityToModel(review.Payments),
		ImageReviews: imageReview,
	}
}

func ImageReviewEntityToModel(imageReview ImageReviewEntity) ImageReview {
	return ImageReview{
		ReviewID: imageReview.ReviewID,
		Link:     imageReview.Link,
		Reviews:  ReviewEntityToModel(imageReview.Reviews),
	}
}