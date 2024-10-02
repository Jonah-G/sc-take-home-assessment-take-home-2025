package main

import (
	"fmt"
	"github.com/Jonah-G/sc-take-home-assessment-take-home-2025/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	orgID = uuid.FromStringOrNil("61850505-a112-4035-8f38-0a9879811fb0")

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
	name := "lucky-defenders"
	children, err := folderDriver.GetAllChildFolders(orgID, name)
	fmt.Printf("\n Child folders for orgID: %s and name: %s \n", orgID, name)

	if err != nil {
		fmt.Println(err)
	} else {
		// folder.PrettyPrint(children)
		fmt.Println(children)
	}
}
