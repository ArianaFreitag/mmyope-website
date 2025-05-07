package models

import (
	"database/sql"
)

var DB *sql.DB

// Book - Data Model To store book data in
// id, media_id, title, author, license, book_num, page_count, dimensions_in,dimensions_mm, weight, paper, ink, squiggly_num, icon_path, img_paths, col_text, col_img_path
type Book struct {
	Id            int    `json:"id"`
	Media_Id      string `json:"media_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	License       string `json:"license"`
	Book_Num      string `json:"book_num"`
	Page_Count    int    `json:"page_count"`
	Dimensions_In string `json:"dimensions_in"`
	Dimensions_Mm string `json:"dimensions_mm"`
	Weight        string `json:"weight"`
	Paper         string `json:"paper"`
	Ink           string `json:"ink"`
	Squiggly_Num  string `json:"squiggly_num"`
	Icon_Path     string `json:"icon_path"`
	Img_Paths     string `json:"img_paths"`
	Col_Text      string `json:"col_text"`
	Col_Img_Path  string `json:"col_img_path"`
	Created_At    string `json:"created_at"`
}

type Media struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Year            int    `json:"year"`
	Infosheet_Path  string `json:"infosheet_path"`
	Thumbnail_Path  string `json:"thumbnail_path"`
	Contributor_Ids string `json:"contributor_ids"`
	Description     string `json:"description"`
	Quantity        int    `json:"quanity"`
	Spine_Num       string `json:"spine_num"`
	Publisher       string `json:"publisher"`
	Date_Released   string `json:"date_released"`
	Created_At      string `json:"created_at"`
}

// make sure its returning in sorted order
func GetBooks() ([]Media, error) {
	rows, err := DB.Query("SELECT * FROM mmyope.physical_media WHERE type = 'book'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Media

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var bk Media
		if err := rows.Scan(&bk.Id, &bk.Name, &bk.Type, &bk.Year, &bk.Infosheet_Path, &bk.Thumbnail_Path, &bk.Contributor_Ids, &bk.Description, &bk.Quantity, &bk.Spine_Num, &bk.Publisher, &bk.Date_Released, &bk.Created_At); err != nil {
			return bks, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return bks, err
	}
	return bks, nil
}

// make this where its handling single row vs looping
func GetBook(id string) (*Book, error) {
	rows, err := DB.Query("SELECT * FROM mmyope.books WHERE media_id=$1", id)
	var bk Book
	if err != nil {
		bk := new(Book)
		return bk, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		// id, media_id, title, author, license, book_num, page_count, dimensions_in,dimensions_mm, weight, paper, ink, squiggly_num, icon_path, img_paths, col_text, col_img_path
		if err := rows.Scan(&bk.Id, &bk.Media_Id, &bk.Title, &bk.Author, &bk.License, &bk.Book_Num, &bk.Page_Count, &bk.Dimensions_In, &bk.Dimensions_Mm, &bk.Weight, &bk.Paper, &bk.Ink, &bk.Squiggly_Num, &bk.Icon_Path, &bk.Img_Paths, &bk.Col_Text, &bk.Col_Img_Path, &bk.Created_At); err != nil {
			return &bk, err
		}

	}

	if err = rows.Err(); err != nil {
		return &bk, err
	}
	return &bk, nil
}
