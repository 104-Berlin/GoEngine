package object

import (
	"errors"
	"fmt"
)

type ObjId uint

var currObjId ObjId = 0

type Object struct {
	*object
}

type object struct {
	name     string
	uuid     ObjId
	children map[ObjId]Object
	parent   Object
}

func (o Object) Name() string {
	return o.name
}

func (o Object) SetName(name string) {
	o.name = name
}

func (o Object) Uuid() ObjId {
	return o.uuid
}

func Create(name string) Object {
	defer func() { currObjId++ }()

	return Object{
		object: &object{
			name: name,
			uuid: currObjId,
		},
	}
}

func (o Object) String() string {
	if o.object == nil {
		return "Obj: nil"
	}
	return fmt.Sprintf("Obj: %s (ID: %d)", o.name, o.object.uuid)
}

func (o Object) IsSet() bool {
	return o.object != nil
}

func (o Object) AddChild(c Object) error {
	if !o.IsSet() {
		return errors.New("Object not set!")
	}
	if !c.IsSet() {
		return errors.New("Child not set!")
	}

	if o == c {
		return errors.New("Both Objects are the same!")
	}

	for oi := o; oi.IsSet(); oi = oi.parent {
		if oi == c {
			return errors.New("New child is a parent of Object!")
		}
	}

	if c.parent.object != nil {
		delete(c.parent.children, c.uuid)
	}

	if o.children == nil {
		o.children = make(map[ObjId]Object)
	}

	c.parent = o
	o.children[c.uuid] = c
	return nil
}

func (o Object) Traverse(f func(Object)) {
	if o.object == nil {
		return
	}
	f(o)
	if o.children != nil {
		for _, c := range o.children {
			c.Traverse(f)
		}
	}
}
