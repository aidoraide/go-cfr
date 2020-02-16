package algo

type Action int
type History interface {
	TurnToAct() int
	Infoset() *Infoset
	InfosetKey() string
	TakeAction(Action) History
	IsTerminal() bool
	Value() []float32 // one value for each player
	String() string
}
