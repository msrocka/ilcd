package ilcd

type MemStore struct {
	contacts   map[string]*Contact
	flowProps  map[string]*FlowProperty
	flows      map[string]*Flow
	methods    map[string]*Method
	models     map[string]*Model
	processes  map[string]*Process
	sources    map[string]*Source
	unitGroups map[string]*UnitGroup
}

func NewMemStore() *MemStore {
	return &MemStore{
		contacts:   make(map[string]*Contact),
		flowProps:  make(map[string]*FlowProperty),
		flows:      make(map[string]*Flow),
		methods:    make(map[string]*Method),
		models:     make(map[string]*Model),
		processes:  make(map[string]*Process),
		sources:    make(map[string]*Source),
		unitGroups: make(map[string]*UnitGroup),
	}
}

func (store *MemStore) FillFromZip(reader *ZipReader) error {

	// contacts
	if err := reader.EachContact(func(c *Contact) bool {
		store.contacts[c.UUID()] = c
		return true
	}); err != nil {
		return err
	}

	// flow properties
	if err := reader.EachFlowProperty(func(p *FlowProperty) bool {
		store.flowProps[p.UUID()] = p
		return true
	}); err != nil {
		return err
	}

	// flows
	if err := reader.EachFlow(func(f *Flow) bool {
		store.flows[f.UUID()] = f
		return true
	}); err != nil {
		return err
	}

	// methods
	if err := reader.EachMethod(func(m *Method) bool {
		store.methods[m.UUID()] = m
		return true
	}); err != nil {
		return err
	}

	// models
	if err := reader.EachModel(func(m *Model) bool {
		store.models[m.UUID()] = m
		return true
	}); err != nil {
		return err
	}

	// processes
	if err := reader.EachProcess(func(p *Process) bool {
		store.processes[p.UUID()] = p
		return true
	}); err != nil {
		return err
	}

	// sources
	if err := reader.EachSource(func(s *Source) bool {
		store.sources[s.UUID()] = s
		return true
	}); err != nil {
		return err
	}

	// unit groups
	if err := reader.EachUnitGroup(func(u *UnitGroup) bool {
		store.unitGroups[u.UUID()] = u
		return true
	}); err != nil {
		return err
	}

	return nil
}

func (store *MemStore) Contacts() []*Contact {
	contacts := make([]*Contact, 0, len(store.contacts))
	for _, contact := range store.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (store *MemStore) FlowProperties() []*FlowProperty {
	flowProperties := make([]*FlowProperty, 0, len(store.flowProps))
	for _, flowProperty := range store.flowProps {
		flowProperties = append(flowProperties, flowProperty)
	}
	return flowProperties
}

func (store *MemStore) Flows() []*Flow {
	flows := make([]*Flow, 0, len(store.flows))
	for _, flow := range store.flows {
		flows = append(flows, flow)
	}
	return flows
}

func (store *MemStore) Methods() []*Method {
	methods := make([]*Method, 0, len(store.methods))
	for _, method := range store.methods {
		methods = append(methods, method)
	}
	return methods
}

func (store *MemStore) Models() []*Model {
	models := make([]*Model, 0, len(store.models))
	for _, model := range store.models {
		models = append(models, model)
	}
	return models
}

func (store *MemStore) Processes() []*Process {
	processes := make([]*Process, 0, len(store.processes))
	for _, process := range store.processes {
		processes = append(processes, process)
	}
	return processes
}

func (store *MemStore) Sources() []*Source {
	sources := make([]*Source, 0, len(store.sources))
	for _, source := range store.sources {
		sources = append(sources, source)
	}
	return sources
}

func (store *MemStore) UnitGroups() []*UnitGroup {
	unitGroups := make([]*UnitGroup, 0, len(store.unitGroups))
	for _, unitGroup := range store.unitGroups {
		unitGroups = append(unitGroups, unitGroup)
	}
	return unitGroups
}

func (store *MemStore) Contact(id string) *Contact {
	return store.contacts[id]
}

func (store *MemStore) FlowProperty(id string) *FlowProperty {
	return store.flowProps[id]
}

func (store *MemStore) Flow(id string) *Flow {
	return store.flows[id]
}

func (store *MemStore) Method(id string) *Method {
	return store.methods[id]
}

func (store *MemStore) Model(id string) *Model {
	return store.models[id]
}

func (store *MemStore) Process(id string) *Process {
	return store.processes[id]
}

func (store *MemStore) Source(id string) *Source {
	return store.sources[id]
}

func (store *MemStore) UnitGroup(id string) *UnitGroup {
	return store.unitGroups[id]
}

func (store *MemStore) EachContact(fn func(*Contact) bool) {
	for _, contact := range store.contacts {
		if !fn(contact) {
			break
		}
	}
}

func (store *MemStore) EachFlowProperty(fn func(*FlowProperty) bool) {
	for _, flowProperty := range store.flowProps {
		if !fn(flowProperty) {
			break
		}
	}
}

func (store *MemStore) EachFlow(fn func(*Flow) bool) {
	for _, flow := range store.flows {
		if !fn(flow) {
			break
		}
	}
}

func (store *MemStore) EachMethod(fn func(*Method) bool) {
	for _, method := range store.methods {
		if !fn(method) {
			break
		}
	}
}

func (store *MemStore) EachModel(fn func(*Model) bool) {
	for _, model := range store.models {
		if !fn(model) {
			break
		}
	}
}

func (store *MemStore) EachProcess(fn func(*Process) bool) {
	for _, process := range store.processes {
		if !fn(process) {
			break
		}
	}
}

func (store *MemStore) EachSource(fn func(*Source) bool) {
	for _, source := range store.sources {
		if !fn(source) {
			break
		}
	}
}

func (store *MemStore) EachUnitGroup(fn func(*UnitGroup) bool) {
	for _, unitGroup := range store.unitGroups {
		if !fn(unitGroup) {
			break
		}
	}
}
