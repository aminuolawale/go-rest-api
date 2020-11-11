package  repos
import (
	"../entity"
)

type PostRepo interface {
	Save(post *entity.Post)(*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
