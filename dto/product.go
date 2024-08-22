package dto

import "time"

type ProductRequest struct {
	Name        string  `json:"name" example:"Digital Camera 4K" valid:"required~Name cannot be empty"`
	Description string  `json:"description" example:"Capture breathtaking moments in stunning 4K resolution." valid:"required~Description cannot be empty"`
	Price       float64 `json:"price" example:"7000000" valid:"required~Price cannot be empty"`
	ImageUrl    string  `json:"imageUrl" example:"https://example.com/digital_camera_4k.jpg" valid:"required~Image Url cannot be empty"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name" example:"Portable Bluetooth Speaker"`
	Description string  `json:"description"  example:"Bring the party anywhere by our Portable Bluetooth Speaker." valid:"required~Description cannot be empty"`
	Price       float64 `json:"price" example:"500000" valid:"required~Price cannot be empty"`
	ImageUrl    string  `json:"imageUrl" example:"https://example.com/portable_speaker.jpg" valid:"required~Image Url cannot be empty"`
}

type ProductCreatedResponse struct {
	Result     string             `json:"result" example:"SUCCESS"`
	StatusCode int                `json:"statusCode" example:"201"`
	Message    string             `json:"message" example:"Product created successfully"`
	Data       CreatedProductData `json:"data"`
}

type ProductsRetrievedResponse struct {
	Result     string                 `json:"result" example:"SUCCESS"`
	StatusCode int                    `json:"statusCode" example:"200"`
	Message    string                 `json:"message" example:"Products retrieved successfully"`
	Data       []RetrievedProductData `json:"data"`
}

type ProductRetrievedResponse struct {
	Result     string               `json:"result" example:"SUCCESS"`
	StatusCode int                  `json:"statusCode" example:"200"`
	Message    string               `json:"message" example:"Product retrieved successfully"`
	Data       RetrievedProductData `json:"data"`
}

type ProductUpdatedResponse struct {
	Result     string             `json:"result" example:"SUCCESS"`
	StatusCode int                `json:"statusCode" example:"200"`
	Message    string             `json:"message" example:"Product updated successfully"`
	Data       UpdatedProductData `json:"data"`
}

type ProductDeletedResponse struct {
	Result     string    `json:"result" example:"SUCCESS"`
	StatusCode int       `json:"statusCode" example:"200"`
	Message    string    `json:"message" example:"Product deleted successfully"`
	DeletedAt  time.Time `json:"deletedAt,omitempty" example:"2024-08-21T19:41:27.1757419+07:00"`
}

type CreatedProductData struct {
	Id          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Digital Camera 4K"`
	Description string    `json:"description" example:"Capture breathtaking moments in stunning 4K resolution."`
	Price       float64   `json:"price" example:"7000000"`
	ImageUrl    string    `json:"imageUrl" example:"https://example.com/digital_camera_4k.jpg"`
	UserId      uint      `json:"userId" example:"1"`
	CreatedAt   time.Time `json:"createdAt,omitempty" example:"2024-08-21T13:48:47.729483+07:00"`
}

type UpdatedProductData struct {
	Id          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Portable Bluetooth Speaker"`
	Description string    `json:"description" example:"Bring the party anywhere by our Portable Bluetooth Speaker."`
	Price       float64   `json:"price" example:"500000"`
	ImageUrl    string    `json:"imageUrl" example:"https://example.com/portable_speaker.jpg"`
	UserId      uint      `json:"userId" example:"1"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" example:"2024-08-21T15:21:08.548216+07:00"`
}

type RetrievedProductData struct {
	Id          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Digital Camera 4K"`
	Description string    `json:"description" example:"Capture breathtaking moments in stunning 4K resolution."`
	Price       float64   `json:"price" example:"7000000"`
	ImageUrl    string    `json:"imageUrl" example:"https://example.com/digital_camera_4k.jpg"`
	UserId      uint      `json:"userId" example:"1"`
	CreatedAt   time.Time `json:"createdAt,omitempty" example:"2024-08-21T13:48:47.729483+07:00"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" example:"2024-08-21T15:21:08.548216+07:00"`
}
