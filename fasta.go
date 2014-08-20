package main

type Fasta interface {
	SetId(string)
	GetId() string

	SetDescription(string)
	GetDescription() string

	SetSequence(string)
	GetSequence() string
}

type FastaImp struct {
	id          string
	description string
	sequence    string
}

func (fa *FastaImp) SetId(id string) {
	fa.id = id
}

func (fa *FastaImp) GetId() string {
	return fa.id
}

func (fa *FastaImp) SetDescription(description string) {
	fa.description = description
}

func (fa *FastaImp) GetDescription() string {
	return fa.description
}

func (fa *FastaImp) SetSequence(sequence string) {
	fa.sequence = sequence
}

func (fa *FastaImp) GetSequence() string {
	return fa.sequence
}
