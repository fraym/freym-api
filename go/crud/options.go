package crud

import "github.com/fraym/freym-api/go/proto/crud/deliverypb"

type GetEntryOptionsBuilder struct {
	Target                   deliverypb.DeploymentTarget
	UseStrongConsistency     bool
	UseStrongConsistencyById string
}

func (b *GetEntryOptionsBuilder) Build() *GetEntryOptions {
	return &GetEntryOptions{
		target:                   b.Target,
		useStrongConsistency:     b.UseStrongConsistency,
		useStrongConsistencyById: b.UseStrongConsistencyById,
	}
}

type GetEntryOptions struct {
	target                   deliverypb.DeploymentTarget
	useStrongConsistency     bool
	useStrongConsistencyById string
}

func (o *GetEntryOptions) Target() deliverypb.DeploymentTarget {
	if o == nil {
		return deliverypb.DeploymentTarget_DEPLOYMENT_TARGET_BLUE
	}

	return o.target
}

func (o *GetEntryOptions) UseStrongConsistency() bool {
	return o.useStrongConsistency
}

func (o *GetEntryOptions) UseStrongConsistencyById() string {
	return o.useStrongConsistencyById
}

type GetSingleEntryOptionsBuilder struct {
	GetEntryOptionsBuilder
	ReturnEmptyDataIfNotFound bool
}

func (b *GetSingleEntryOptionsBuilder) Build() *GetSingleEntryOptions {
	return &GetSingleEntryOptions{
		GetEntryOptions:           *b.GetEntryOptionsBuilder.Build(),
		returnEmptyDataIfNotFound: b.ReturnEmptyDataIfNotFound,
	}
}

type GetSingleEntryOptions struct {
	GetEntryOptions
	returnEmptyDataIfNotFound bool
}

func (o *GetSingleEntryOptions) ReturnEmptyDataIfNotFound() bool {
	if o == nil {
		return false
	}

	return o.returnEmptyDataIfNotFound
}
