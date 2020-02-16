package algo

type Game interface {
	NewGame() History
	PlayerSet() []int
}
