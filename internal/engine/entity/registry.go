package entity

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type (
	factory            = func() Component
	ComponentsRegistry struct {
		components map[UUID]factory
	}
)

func NewComponentsRegistry() *ComponentsRegistry {
	return &ComponentsRegistry{
		components: map[UUID]factory{},
	}
}

func (r *ComponentsRegistry) RegisterComponent(example Component) {
	if _, exist := r.components[example.Id()]; exist {
		panic(fmt.Sprintf("Component %s (%s) already registered in engine", example.Id(), example.Title()))
	}

	// copy gold component template to heap
	goldComponent := example

	r.components[example.Id()] = func() Component {
		// copy initial state from gold copy
		// and return it
		return goldComponent
	}

	log.Printf("Component '%s' (%s) registered\n", example.Title(), example.Id())
}

func (r *ComponentsRegistry) CreateComponentWithProps(id UUID, props map[string]string) Component {
	factory, ok := r.components[id]
	if !ok {
		panic(fmt.Errorf("component '%s' not registered in engine, and can`t be created", id))
	}

	created := factory()
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused:      true,
		ZeroFields:       true,
		WeaklyTypedInput: true,
		Squash:           false,
		TagName:          "editable",
		Result:           &created,
	})
	if err != nil {
		panic(fmt.Errorf("failed create decoder for decoding '%s' component: %w", id, err))
	}

	err = decoder.Decode(props)
	if err != nil {
		panic(fmt.Errorf("failed decode component '%s' props: %w (props: %#v)", id, err, props))
	}

	return created
}
