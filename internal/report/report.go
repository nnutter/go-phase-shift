package report

import "fmt"

func MutatesPointerParameter(name string) string {
	return fmt.Sprintf("//constable:nonmutating function mutates pointer parameter %s", name)
}

func MutatesReceiver(name string) string {
	return fmt.Sprintf("//constable:nonmutating method mutates receiver %s", name)
}

func MutatesParameter(name string) string {
	return fmt.Sprintf("//constable:nonmutating function mutates parameter %s", name)
}

func DeletesFromMapParameter(name string) string {
	return fmt.Sprintf("//constable:nonmutating function deletes from map parameter %s", name)
}
