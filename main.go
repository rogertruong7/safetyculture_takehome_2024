package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	
	// folder.PrettyPrint(res)
	
	// childrenFolder := folderDriver.GetAllChildFolders(orgID, "fast-watchmen");
	// folder.PrettyPrint(childrenFolder)

	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	movefolder, err := folderDriver.MoveFolder("fast-watchmen", "pure-blastaar");
	fmt.Print(err)
	folder.PrettyPrint(movefolder)
}
