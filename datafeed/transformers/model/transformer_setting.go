package model

// TransformerSetting provides model to specify transformer settings
type TransformerSetting struct {
	Transformer string
	Parameters  map[string]interface{}
}
