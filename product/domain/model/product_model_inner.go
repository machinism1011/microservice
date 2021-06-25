package model

type ProductImage struct {
	ID				int64	`gorm:"primary_key;not_null;auto_increment" json:"id"`
	ImageName		string	`json:"image_name"`
	ImageCode		string	`gorm:"unique_index;not_null" json:"image_code"`
	ImageUrl		string	`json:"image_url"`
	ImageProductID	int64	`json:"image_product_id"`
}

type ProductSize struct {
	ID				int64	`gorm:"primary_key;not_null;auto_increment" json:"id"`
	SizeName		string	`json:"size_name"`
	SizeCode		string	`gorm:"unique_index;not_null" json:"size_code"`
	SizeProductID	int64	`json:"size_product_id"`
}

type ProductSeo struct {
	ID				int64	`gorm:"primary_key;not_null;auto_increment" json:"id"`
	SeoTitle		string	`json:"seo_title"`
	SeoKeyWords		string	`json:"seo_keywords"`
	SeoDescription	string	`json:"seo_description"`
	SeoCode			string	`json:"seo_code"`
	SeoProductID	int64	`json:"seo_product_id"`
}
