package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	commands    = make(map[string]string)
	commandKeys []string
)

func commandsInit() {
	commands = make(map[string]string)

	// Commands
	commands["connect"] = "connect  \n\t for developer playing\n"
	commands["disconnect"] = "disconnect  \n\t for developer playing\n"

	commands["setcollection"] = "setcollection  \n\t for developer playing\n"
	commands["getcollection"] = "getcollection  \n\t for developer playing\n"

	// getmasterdata
	commands["getmasterdata"] = "getmasterdata  \n\t get all master data, i.e. users, vehicles, and the current bookings\n"

	commands["delete"] = "delete  \n\t for developer playing\n"
	commands["deleteall"] = "deleteall  \n\t for developer playing\n"

	commands["addvehicle"] = "addvehicle  \n\t for developer playing\n"
	commands["getvehicles"] = "getvehicles  \n\t for developer playing\n"

	commands["storevehicles"] = "storevehicles  \n\t for developer playing\n"
	commands["clearvehicles"] = "clearvehicles  \n\t for developer playing\n"

	commands["adduser"] = "adduser  \n\t for developer playing\n"
	commands["getuser"] = "getuser  \n\t for developer playing\n"

	commands["getusers"] = "getusers  \n\t for developer playing\n"
	commands["storeusers"] = "storeusers  \n\t for developer playing\n"
	commands["clearusers"] = "clearusers  \n\t for developer playing\n"

	commands["addbooking"] = "addbooking  \n\t for developer playing\n"
	commands["getbookings"] = "getbookings  \n\t for developer playing\n"

	commands["storebookings"] = "storebookings  \n\t for developer playing\n"
	commands["clearbookings"] = "clearbookings  \n\t for developer playing\n"

	commands["querybookings"] = "querybookings  \n\t for developer playing\n"

	commands["additem"] = "additem  \n\t for developer playing\n"
	commands["getitems"] = "getitems  \n\t for developer playing\n"
	commands["deleteallitems"] = "deleteallitems  \n\t for developer playing\n"

	commands["listdocs"] = "listdocs  \n\t for developer playing\n"
	commands["addDocument"] = "addDocument  \n\t for developer playing\n"
	commands["getDocument"] = "getDocument  \n\t for developer playing\n"
	commands["cloneDocument"] = "cloneDocument  \n\t for developer playing\n"

	// Developer
	commands["play"] = "play  \n\t for developer playing\n"

	// Internals
	commands["quit"] = "quit  \n\t close the session and exit\n"
	commands["usage"] = "usage  \n\t shows this usage with all command usages\n"

	// Add your own command to call a function interactively
	// commands["mycmd"] = "mycmd  \n\t does something\n"

	// To store the keys in sorted order
	for commandKey := range commands {
		commandKeys = append(commandKeys, commandKey)
	}
	sort.Strings(commandKeys)
}

// Execute a command specified by the argument string
func executeCommand(commandline string) bool {

	// Trim prefix and split string by white spaces
	commandFields := strings.Fields(commandline)

	// Check for empty string without prefix
	if len(commandFields) > 0 {

		// Switch according to the first word and call appropriate function with the rest as arguments
		switch commandFields[0] {

		case "connect":
			connect(commandFields[1:])
			return true

		case "disconnect":
			disconnect(commandFields[1:])
			return true

		case "setcollection":
			setcollection(commandFields[1:])
			return true

		case "getcollection":
			getcollection(commandFields[1:])
			return true

			// MasterData
		case "getmasterdata":
			GetMasterData(commandFields[1:])
			return true

		case "additem":
			AddItem(commandFields[1:])
			return true
		case "getitems":
			GetItems(commandFields[1:])
			return true
		case "deleteallitems":
			DeleteItems(commandFields[1:])
			return true

		case "addvehicle":
			AddVehicle(commandFields[1:])
			return true
		case "getvehicles":
			GetVehicles(commandFields[1:])
			return true
		case "storevehicles":
			StoreVehicles(commandFields[1:])
			return true
		case "clearvehicles":
			ClearVehicles(commandFields[1:])
			return true

		case "adduser":
			AddUser(commandFields[1:])
			return true
		case "getuser":
			GetUser(commandFields[1:])
			return true
		case "getusers":
			GetUsers(commandFields[1:])
			return true
		case "storeusers":
			StoreUsers(commandFields[1:])
			return true
		case "clearusers":
			ClearUsers(commandFields[1:])
			return true

		case "addbooking":
			AddBooking(commandFields[1:])
			return true
		case "getbookings":
			GetBookings(commandFields[1:])
			return true
		case "storebookings":
			StoreBookings(commandFields[1:])
			return true
		case "clearbookings":
			ClearBookings(commandFields[1:])
			return true

		case "querybookings":
			QueryBookings(commandFields[1:])
			return true

		case "addDocument":
			addDocument(commandFields[1:])
			return true
		case "getDocument":
			getDocument(commandFields[1:])
			return true
		case "cloneDocument":
			cloneDocument(commandFields[1:])
			return true

		case "listdocs":
			listdocs(commandFields[1:])
			return true

		case "delete":
			delete(commandFields[1:])
			return true

		case "deleteall":
			deleteall(commandFields[1:])
			return true

		case "quit":
			quitCmdTool(commandFields[1:])
			return true

		case "play":
			play(commandFields[1:])
			return true

		// Link command and function
		//case "mycmd":
		//	myCmd(commandFields[1:])
		//	return true

		default:
			usage()
			return false
		}
	}
	return false
}

// Display the usage of all available commands
func usage() {
	for _, key := range commandKeys {
		fmt.Printf("%v\n", commands[key])
	}

}

func quitCmdTool(arguments []string) {

	// Get rid of warnings
	_ = arguments

	os.Exit(0)
}

// func myCmd(arguments []string) {}
