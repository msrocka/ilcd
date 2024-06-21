package ilcd

type StoreReader interface {
	EachSource(fn func(*Source) bool) error
	EachContact(fn func(*Contact) bool) error
	EachUnitGroup(fn func(*UnitGroup) bool) error
	EachFlowProperty(fn func(*FlowProperty) bool) error
	EachFlow(fn func(*Flow) bool) error
	EachProcess(fn func(*Process) bool) error
	EachModel(fn func(*Model) bool) error
	EachMethod(fn func(*Method) bool) error

	ReadSource(uuid string) (*Source, error)
	ReadContact(uuid string) (*Contact, error)
	ReadUnitGroup(uuid string) (*UnitGroup, error)
	ReadFlowProperty(uuid string) (*FlowProperty, error)
	ReadFlow(uuid string) (*Flow, error)
	ReadProcess(uuid string) (*Process, error)
	ReadModel(uuid string) (*Model, error)
	ReadMethod(uuid string) (*Method, error)
}
