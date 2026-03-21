// описание интерфейса для сохранения и нахождения игры по uuid

package domain

type GameRepository interface {
	Save(game *Game) error
	Get(id string) (*Game, error)
}
