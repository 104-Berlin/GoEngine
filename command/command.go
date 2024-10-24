package command

import "github.com/104-Berlin/GoEngine/object"

type Command interface {
	Exec(object.Object)
	Unexec()
}
