package gof_openfeature

import "github.com/open-feature/go-sdk/openfeature"

// MyFeatureProvider implements the FeatureProvider interface and provides functions for evaluating flags
type MyFeatureProvider struct{}
type Metadata = openfeature.Metadata
type EvaluationContext = openfeature.EvaluationContext
type ResolutionDetail = openfeature.ResolutionDetail
type IntResolutionDetail = openfeature.IntResolutionDetail
type FloatResolutionDetail = openfeature.FloatResolutionDetail
type BoolResolutionDetail = openfeature.BoolResolutionDetail
type StringResolutionDetail = openfeature.StringResolutionDetail
type Hook = openfeature.Hook

// Metadata returns the metadata of the provider
func (e MyFeatureProvider) Metadata() Metadata {
	return Metadata{Name: "MyFeatureProvider"}
}

// BooleanEvaluation returns a boolean flag
func (e MyFeatureProvider) BooleanEvaluation(flag string, defaultValue bool, evalCtx EvaluationContext) BoolResolutionDetail {
	// code to evaluate boolean
	return BoolResolutionDetail{}
}

// StringEvaluation returns a string flag
func (e MyFeatureProvider) StringEvaluation(flag string, defaultValue string, evalCtx EvaluationContext) StringResolutionDetail {
	// code to evaluate string
	return StringResolutionDetail{}
}

// FloatEvaluation returns a float flag
func (e MyFeatureProvider) FloatEvaluation(flag string, defaultValue float64, evalCtx EvaluationContext) FloatResolutionDetail {
	// code to evaluate float
	return FloatResolutionDetail{}
}

// IntEvaluation returns an int flag
func (e MyFeatureProvider) IntEvaluation(flag string, defaultValue int64, evalCtx EvaluationContext) IntResolutionDetail {
	// code to evaluate int
	return IntResolutionDetail{}
}

// ObjectEvaluation returns an object flag
func (e MyFeatureProvider) ObjectEvaluation(flag string, defaultValue interface{}, evalCtx EvaluationContext) ResolutionDetail {
	// code to evaluate object
	return ResolutionDetail{}
}

// Hooks returns hooks
func (e MyFeatureProvider) Hooks() []Hook {
	// code to retrieve hooks
	return []Hook{}
}
