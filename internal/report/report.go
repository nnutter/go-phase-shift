package report

import "fmt"

func MutatesPointerParameter(name string) string {
	return fmt.Sprintf("//phase:nonmutating function mutates pointer parameter %s", name)
}

func MutatesReceiver(name string) string {
	return fmt.Sprintf("//phase:nonmutating method mutates receiver %s", name)
}

func MutatesParameter(name string) string {
	return fmt.Sprintf("//phase:nonmutating function mutates parameter %s", name)
}

func DeletesFromMapParameter(name string) string {
	return fmt.Sprintf("//phase:nonmutating function deletes from map parameter %s", name)
}
