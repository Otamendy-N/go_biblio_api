package catalog

import "go_biblio_api/entities"

type BookCollection struct {
	List []entities.Book
	Seed rune
}

func (bc *BookCollection) SeedData() {
	bc.Seed = 0
}

func (bc *BookCollection) NextId() rune {
	bc.Seed++
	return bc.Seed
}

func (bc *BookCollection) Add(book *entities.Book) {
	book.Id = bc.NextId()
	bc.List = append(bc.List, *book)
}

func (bc BookCollection) GetById(id rune) entities.Book {
	for _, book := range bc.List {
		if book.Id == id {
			return book
		}
	}
	return entities.Book{}
}

func (bc *BookCollection) Update(id rune, book *entities.Book) {
	for i, b := range bc.List {
		if b.Id == id {
			bc.List[i] = *book
			return
		}
	}
}

func (bc *BookCollection) Remove(id rune) {
	for i, b := range bc.List {
		if b.Id == id {
			bc.List = append(bc.List[:i], bc.List[i+1:]...)
			return
		}
	}
}

func (bc BookCollection) GetAll() []entities.Book {
	return bc.List
}
