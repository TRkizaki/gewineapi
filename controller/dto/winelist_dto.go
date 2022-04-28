package dto

type WinelistResponse struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
}

type WinelistRequest struct {
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
}

type WinelistsResponse struct {
	Winelists []WinelistResponse `json:"winelists"`
}
