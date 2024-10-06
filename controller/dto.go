package controller

type CreateSongDTO struct {
	Group string `json:"group" validate:"required,min=3,max=100" example:"Muse"`
	Song  string `json:"song" validate:"required,min=3,max=100" example:"Supermassive Black Hole"`
}

type PaginationDTO struct {
	Offset string `json:"offset,omitempty" example:"1"`
	Limit  string `json:"limit,omitempty" example:"1"`
}

type songStringFields struct {
	Group string `json:"group,omitempty" validate:"omitempty,min=3,max=100" example:"Muse"`
	Song  string `json:"song,omitempty" validate:"omitempty,min=3,max=100" example:"Supermassive Black Hole"`
	Text  string `json:"text,omitempty" validate:"omitempty,min=3,max=10000" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link  string `json:"link,omitempty" validate:"omitempty,min=3,max=100" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

type FindSongDTO struct {
	PaginationDTO
	songStringFields

	ReleaseDateBegin string `json:"releaseDateBegin,omitempty" example:"1717188495"`
	ReleaseDateEnd   string `json:"releaseDateEnd,omitempty" example:"1727188495"`
}

type UpdateSongDTO struct {
	IdDTO
	songStringFields

	ReleaseDate string `json:"releaseDate,omitempty" example:"16.07.2006"`
}

type IdDTO struct {
	ID string `json:"id" validate:"required,uuid4" example:"147367f5-93ef-432d-8a97-b06f716f9fad"`
}

type GetTextDTO struct {
	PaginationDTO

	IdDTO
}
