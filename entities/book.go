package entities

const (
	LightNovel = "Light Novel"
	Comic      = "Comic"
	Manga      = "Manga"
	Manhua     = "Manhua"
)

type Book struct {
	Id          rune   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Vols        int    `json:"vols"`
	Type        string `json:"type"`
	CoverUrl    string `json:"coverUrl"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}