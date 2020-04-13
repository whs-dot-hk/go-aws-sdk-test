package template

import (
	cf "github.com/awslabs/goformation/v4/cloudformation"
)

type OpenvpnTemplate struct {
	KeyName string
}

func NewOpenvpnTemplate() *OpenvpnTemplate {
	return &OpenvpnTemplate{}
}

func (o *OpenvpnTemplate) GetCfTemplate() *cf.Template {
	return getCfTemplate(o)
}
