package structsEntity

func UserEntityToResponseAll(user UserEntity) UserResponseAll {
	return UserResponseAll{
		Id:     user.Id,
		Nama:   user.Nama,
		Alamat: user.Alamat,
		Role:   user.Role,
		Profil: user.Profil,
	}
}

func ProductEntityToResponseAll(product ProductEntity) ProductResponseAll {
	var image []ImageProductResponseAll
	for _, value := range product.Image {
		image = append(image, ImageProductEntityToResponseAll(value))
	}
	var review []ReviewResponseAll
	for _, value := range product.Review {
		review = append(review, ReviewEntityToResponseAll(value))
	}
	return ProductResponseAll{
		Id:      product.Id,
		Nama:    product.Nama,
		Harga:   product.Harga,
		Ratings: product.Ratings,
		Image:   image,
		Review:  review,
	}
}

func ImageProductEntityToResponseAll(imageProduct ImageProductEntity) ImageProductResponseAll {
	return ImageProductResponseAll{
		Link: imageProduct.Link,
	}
}

func TransactionEntityToResponseAll(transaction TransactionEntity) TransactionResponseAll {
	return TransactionResponseAll{
		Id:         transaction.Id,
		Jumlah:     transaction.Jumlah,
		TotalHarga: transaction.TotalHarga,
	}
}

func TransactionFinalEntityToResponseAll(transactionFinal TransactionFinalEntity) TransactionFinalResponseAll {
	return TransactionFinalResponseAll{
		Id:         transactionFinal.Id,
		TotalHarga: transactionFinal.TotalHarga,
		OrderID:    transactionFinal.OrderID,
	}
}

func PaymentEntityToResponseAll(payment PaymentEntity) PaymentResponseAll {
	return PaymentResponseAll{
		Id:     payment.Id,
		Status: payment.Status,
	}
}

func ReviewEntityToResponseAll(review ReviewEntity) ReviewResponseAll {
	return ReviewResponseAll{
		Id:         review.Id,
		TextReview: review.TextReview,
		Rating:     review.Rating,
		ProductID:  review.ProductID,
	}
}

func ImageReviewEntityToResponseAll(imageReview ImageReviewEntity) ImageReviewResponseAll {
	return ImageReviewResponseAll{
		Link: imageReview.Link,
	}
}