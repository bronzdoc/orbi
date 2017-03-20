package skeleton

import (
	"github.com/bronzdoc/skeletor/skeleton/actions"
	"github.com/bronzdoc/skeletor/template"
)

type skeleton struct {
	template template.Template
}

func New(template template.Template) *skeleton {
	return &skeleton{
		template: template,
	}
}

func (s *skeleton) Create() {
	actions.Create(s.template)
}
