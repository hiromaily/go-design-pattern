package prototype

// TODO: this way seems nice.
// http://blog.ralch.com/tutorial/design-patterns/golang-prototype/

type producter interface {
	clone() producter
	GetName() string
}

type Manager struct {
	product producter
}

func (m *Manager) Register(producter producter) {
	m.product = producter
}

func (m *Manager) Create(name string) producter {
	//TODO: name is not used anywhere
	producter := m.product
	return producter.clone()
}

type Product struct {
	Name string
}

func (p *Product) SetUp() {
	// something takes time...
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) clone() producter {
	return &Product{p.Name}
}
