package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)


func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	
	// Error: Cannot move a folder to a child of itself
	// assume that there cannot be two folders with the same name, because according
	// to example, the dst string is just a folder name, with no path
	
	// By the spec there cannot be two folders with the same name because we do not specify the path of the folder
	if (name == dst) {
		return nil, fmt.Errorf("error: cannot move a folder to itself")
	}
	
	sourceExists := false
	sourcePath := ""
	sourceOrgId := uuid.Nil

	dstExists := false
	dstPath := ""
	dstOrgId := uuid.Nil

	folders := f.folders
	for _, f := range folders {
		if f.Name == name {
			sourceExists = true
			sourceOrgId = f.OrgId
			sourcePath = f.Paths
		} else if f.Name == dst {
			dstExists = true
			dstOrgId = f.OrgId
			dstPath = f.Paths
		}
	}
	if !sourceExists {
		return nil, fmt.Errorf("error: source folder does not exist")
	} else if !dstExists {
		return nil, fmt.Errorf("error: destination folder does not exist")
	} else if sourceOrgId != dstOrgId {
		return nil, fmt.Errorf("error: cannot move a folder to a different organization")
	}

	// Checking if dst is a child of source
	dstIndex := -1
	sourceIndex := -1
	sourcePathParts := strings.Split(sourcePath, ".")
	for i, folder := range sourcePathParts {
		if folder == dst {
			dstIndex = i
		} else if folder == name {
			sourceIndex = i
		}
	}

	if dstIndex != -1 {
		if (sourceIndex < dstIndex) {
			return nil, fmt.Errorf("error: cannot move a folder to a child of itselfn")
		}
	}

	// Changing the path string
	res := []Folder{}
    dstPath = fmt.Sprintf(`%s.`, dstPath)
    for _, f := range folders { 
        if f.OrgId == sourceOrgId && strings.Contains(f.Paths, name) {
			parts := strings.Split(f.Paths, ".")
			nameIndex := -1
			for index, part := range parts {
				if part == name {
					nameIndex = index
				}
			}
			if nameIndex == -1 {
				// We know that source folder does exist from above. This ensures that we are not finding a path that just contains the name of the folder e.g. we want abc, but abcd also contains abc
				continue
			}
			combinedString := strings.Join(parts[nameIndex:], ".")
			fmt.Println(combinedString)
            newFolder := Folder{
				Name: f.Name,
				OrgId: f.OrgId,
				Paths: fmt.Sprintf(`%s%s`, dstPath, combinedString),
			}
            res = append(res, newFolder)
        } else {
            res = append(res, f)
        }
    }

	return res, nil
}
