package bank

type AccountType int

const (
	Checking AccountType = iota
	Savings
)

func (t AccountType) String() string {
	switch t {
	case Checking:
		return "Conta Corrente"
	case Savings:
		return "Conta Poupança"
	default:
		return "Conta desconhecida"
	}
}
