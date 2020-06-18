package object

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	o1 := Create("Object 1")
	o2 := Create("Object 2")

	if o1.Uuid() == o2.Uuid() {
		t.Fail()
	}

	o1c := o1

	if o1 != o1c {
		t.Fail()
	}

	if o1 == o2 {
		t.Fail()
	}

	o1.SetName("Renamed")

	if "Renamed" != o1c.Name() {
		t.Fail()
	}
}

func TestIsSet(t *testing.T) {
	var o Object
	if o.IsSet() {
		t.Fail()
	}
	o = Create("Obj")
	if !o.IsSet() {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var o Object
	fmt.Println(o.String())
	o = Create("Obj")
	fmt.Println(o.String())
}

func createObjTree(count int) Object {
	objs := make([]Object, count)
	for i := range objs {
		objs[i] = Create(fmt.Sprintf("Obj-%d", i))
	}
	for i := 1; i < len(objs); i++ {
		err := objs[(i+1)/2-1].AddChild(objs[i])
		if err != nil {
			panic(err)
		}
	}
	return objs[0]
}

func TestAddChild(t *testing.T) {
	var p, c Object
	if p.AddChild(c) == nil {
		t.Fail()
	}

	p = Create("Parent")
	if p.AddChild(c) == nil {
		t.Fail()
	}

	p.object = nil
	c = Create("Child")
	if p.AddChild(c) == nil {
		t.Fail()
	}

	p = Create("Parent")
	if p.AddChild(c) != nil {
		t.Fail()
	}

	if c.AddChild(p) == nil {
		t.Fail()
	}

	np := Create("New Parent")
	if np.AddChild(c) != nil || c.parent != np || len(p.children) != 0 {
		t.Fail()
	}

	if p.AddChild(p) == nil {
		t.Fail()
	}

	objs := [128]Object{}
	for i := range objs {
		objs[i] = Create(fmt.Sprintf("Obj-%d", i))
	}
	for i := 1; i < len(objs); i++ {
		err := objs[(i+1)/2-1].AddChild(objs[i])
		if err != nil {
			t.Fail()
		}
	}
	if objs[3].AddChild(objs[5]) != nil {
		t.Fail()
	}
}

func TestTraverse(t *testing.T) {
	c := 0
	createObjTree(100).Traverse(func(o Object) {
		c++
	})
	if c != 100 {
		t.Fail()
	}
	var no Object
	c2 := 0
	no.Traverse(func(o Object) {
		c2++
	})
	if c2 != 0 {
		t.Fail()
	}
}
