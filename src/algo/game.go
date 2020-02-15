package algo

type Game interface {
	NActions() int32
	NewGame() History
	Value(History) []float32 // one value for each player
	PlayerSet() []int32
}
