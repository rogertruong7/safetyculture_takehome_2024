package folder

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	
	// Error: Cannot move a folder to a child of itself
	// Error: Cannot move a folder to itself
	// Error: Cannot move a folder to a different organization
	// Error: Source folder does not exist
	// Error: Destination folder does not exist
	folders := f.folders
	sourceExists := false
	dstExists := false

	sourceOrgId := uuid.Nil
	dstOrgId := uuid.Nil
	for _, f := range folders {
		if f.Name == name {
			folderExists = true
			folderOrgId = f.OrgId
			break
		}
	}
	if !sourceExists {
		return nil, fmt.Errorf("Error: Source folder does not exist")
	} else if !dstExists {
		return nil, fmt.Errorf("Error: Destination folder does not exist")
	}
	
	return []Folder{}, nil
}
