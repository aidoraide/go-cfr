package algo

type History interface {
	TurnToAct() int32
	Infoset() *Infoset
	InfosetKey() string
	TakeAction(int32) History
	IsTerminal() bool
	String() string
}
