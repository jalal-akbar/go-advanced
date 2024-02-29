package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("App", "Simple App")

	commandAdd             = app.Command("add", "add user")
	commandAddFlagOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
	commandAddArgUser      = commandAdd.Arg("user", "username").Required().String()

	commandUpdate           = app.Command("update", "update user")
	commandUpdateArgOldUser = commandUpdate.Arg("old", "old user").Required().String()
	commandUpdateArgNewUser = commandUpdate.Arg("new", "new user").Required().String()

	commandDelete             = app.Command("delete", "delete user")
	commandDeleteFlagOverride = commandDelete.Flag("force", "force deletion").Short('f').Bool()
	commandDeleteArgUser      = commandDelete.Arg("user", "username").Required().String()
)

func main() {
	// commandAdd.Action(func(pc *kingpin.ParseContext) error {
	// 	user := *commandAddArgUser
	// 	override := *commandAddFlagOverride
	// 	fmt.Printf("adding user: %s, override: %t\n", user, override)
	// 	return nil
	// })
	// commandUpdate.Action(func(pc *kingpin.ParseContext) error {
	// 	newUser := *commandUpdateArgNewUser
	// 	oldUser := *commandUpdateArgOldUser
	// 	fmt.Printf("updating user from %s = %s\n", oldUser, newUser)
	// 	return nil
	// })
	// commandDelete.Action(func(pc *kingpin.ParseContext) error {
	// 	user := *commandDeleteArgUser
	// 	force := *commandDeleteFlagOverride
	// 	fmt.Printf("deletiing user: %s, force: %t\n", user, force)
	// 	return nil
	// })

	// kingpin.MustParse(app.Parse(os.Args[1:]))

	// or

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch cmd {
	case commandAdd.FullCommand():
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("adding user: %s, override: %t\n", user, override)
	case commandUpdate.FullCommand():
		newUser := *commandUpdateArgNewUser
		oldUser := *commandUpdateArgOldUser
		fmt.Printf("updating user from %s = %s\n", oldUser, newUser)
	case commandDelete.FullCommand():
		user := *commandDeleteArgUser
		force := *commandDeleteFlagOverride
		fmt.Printf("deleting user: %s, force: %t\n", user, force)
	}

}
