package structsEntity

import "github.com/midtrans/midtrans-go/coreapi"

func UserEntityToResponse(user UserEntity) UserResponse {
	var products []ProductResponse
	for _, product := range user.Products {
		products = append(products, ProductEntityToResponse(product))
	}
	var transactions []TransactionResponse
	for _, transaction := range user.Transactions {
		transactions = append(transactions, TransactionEntityToResponse(transaction))
	}
	return UserResponse{
		Id:           user.Id,
		CreatedAt:    user.CreatedAt,
		Nama:         user.Nama,
		Email:        user.Email,
		Alamat:       user.Alamat,
		NoTlp:        user.NoTlp,
		Gender:       user.Gender,
		Role:         user.Role,
		Profil:       user.Profil,
		Products:     products,
		Transactions: transactions,
	}
}

func ProductEntityToResponse(product ProductEntity) ProductResponse {
	var images []ImageProductResponse
	for _, image := range product.Image {
		images = append(images, ImageProductEntityToResponse(image))
	}
	var transactions []TransactionResponse
	for _, transaction := range product.Transactions {
		transactions = append(transactions, TransactionEntityToResponse(transaction))
	}
	return ProductResponse{
		Id:           product.Id,
		CreatedAt:    product.CreatedAt,
		UserID:       product.UserID,
		Nama:         product.Nama,
		Kategori:     product.Kategori,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		Harga:        product.Harga,
		Ratings:      product.Ratings,
		Users:        UserEntityToResponse(product.Users),
		Image:        images,
		Transactions: transactions,
	}
}

func ImageProductEntityToResponse(imageProduct ImageProductEntity) ImageProductResponse {
	return ImageProductResponse{
		CreatedAt: imageProduct.CreatedAt,
		ProductID: imageProduct.ProductID,
		Link:      imageProduct.Link,
		Products:  ProductEntityToResponse(imageProduct.Products),
	}
}

func TransactionEntityToResponse(transaction TransactionEntity) TransactionResponse {
	var keranjang []TransactionKeranjangResponse
	for _, value := range transaction.TransactionKeranjang {
		keranjang = append(keranjang, TransactionKeranjangEntityToResponse(value))
	}
	return TransactionResponse{
		Id:                   transaction.Id,
		CreatedAt:            transaction.CreatedAt,
		ProductID:            transaction.ProductID,
		UserID:               transaction.UserID,
		Jumlah:               transaction.Jumlah,
		TotalHarga:           transaction.TotalHarga,
		Products:             ProductEntityToResponse(transaction.Products),
		TransactionKeranjang: keranjang,
		Users:                UserEntityToResponse(transaction.Users),
	}
}

func TransactionKeranjangEntityToResponse(transactionKeranjang TransactionKeranjangEntity) TransactionKeranjangResponse {
	return TransactionKeranjangResponse{
		Id:                 transactionKeranjang.Id,
		CreatedAt:          transactionKeranjang.CreatedAt,
		TransactionID:      transactionKeranjang.TransactionID,
		Transaction:        TransactionEntityToResponse(transactionKeranjang.Transaction),
		TransactionFinalID: transactionKeranjang.TransactionFinalID,
		TransactionFinal:   TransactionFinalEntityToResponse(transactionKeranjang.TransactionFinal),
	}
}

func TransactionFinalEntityToResponse(transactionFinal TransactionFinalEntity) TransactionFinalResponse {
	var transactions []TransactionKeranjangResponse
	for _, transaction := range transactionFinal.TransactionKeranjang {
		transactions = append(transactions, TransactionKeranjangEntityToResponse(transaction))
	}
	var payments []PaymentResponse
	for _, payment := range transactionFinal.Payment {
		payments = append(payments, PaymentEntityToResponse(payment))
	}
	return TransactionFinalResponse{
		Id:                   transactionFinal.Id,
		CreatedAt:            transactionFinal.CreatedAt,
		TotalHarga:           transactionFinal.TotalHarga,
		OrderID:              transactionFinal.OrderID,
		TransactionKeranjang: transactions,
		Payment:              payments,
	}
}

func PaymentEntityToResponse(payment PaymentEntity) PaymentResponse {
	var reviews []ReviewResponse
	for _, review := range payment.Reviews {
		reviews = append(reviews, ReviewEntityToResponse(review))
	}
	return PaymentResponse{
		Id:                 payment.Id,
		CreatedAt:          payment.CreatedAt,
		TransactionFinalID: payment.TransactionFinalID,
		Bank:               payment.Bank,
		Status:             payment.Status,
		VA:                 payment.VA,
		TransactionFinals:  TransactionFinalEntityToResponse(payment.TransactionFinals),
		Reviews:            reviews,
	}
}

func ReviewEntityToResponse(review ReviewEntity) ReviewResponse {
	var imageReview []ImageReviewResponse
	for _, image := range review.ImageReviews {
		imageReview = append(imageReview, ImageReviewEntityToResponse(image))
	}
	return ReviewResponse{
		Id:           review.Id,
		CreatedAt:    review.CreatedAt,
		PaymentID:    review.PaymentID,
		TextReview:   review.TextReview,
		Rating:       review.Rating,
		Payments:     PaymentEntityToResponse(review.Payments),
		ImageReviews: imageReview,
	}
}

func ImageReviewEntityToResponse(imageReview ImageReviewEntity) ImageReviewResponse {
	return ImageReviewResponse{
		Id:        imageReview.Id,
		CreatedAt: imageReview.CreatedAt,
		ReviewID:  imageReview.ReviewID,
		Link:      imageReview.Link,
		Reviews:   ReviewEntityToResponse(imageReview.Reviews),
	}
}

func ResponseMitrans(paymentModel PaymentEntity, transactionFinalId uint, dataResponse *coreapi.ChargeResponse) PaymentEntity {
	paymentModel.TransactionFinalID = transactionFinalId
	paymentModel.Bank = dataResponse.VaNumbers[0].Bank
	paymentModel.VA = dataResponse.VaNumbers[0].VANumber
	paymentModel.Status = dataResponse.TransactionStatus
	return paymentModel
}
