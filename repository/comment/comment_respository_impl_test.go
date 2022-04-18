package comment

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestcommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_db.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Aolia@gmail.com",
		Comment: "belajar golang itu mudah ",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_db.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_db.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
