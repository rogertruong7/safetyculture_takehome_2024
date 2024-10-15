package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)


func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// By the spec there cannot be two folders with the same name because we do not specify the path of the folder
	if (name == dst) {
		return nil, fmt.Errorf("error: cannot move a folder to itself")
	}
	
	sourceExists := false
	sourceOrgId := uuid.Nil
	dstExists := false
	dstPath := ""
	dstOrgId := uuid.Nil
	folders := f.folders
	for _, f := range folders {
		if f.Name == name {
			sourceExists = true
			sourceOrgId = f.OrgId
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
	dstPathParts := strings.Split(dstPath, ".")
	for i, folder := range dstPathParts {
		if folder == dst {
			dstIndex = i
		} else if folder == name {
			sourceIndex = i
		}
	}

	if sourceIndex != -1 {
		if (sourceIndex < dstIndex) {
			return nil, fmt.Errorf("error: cannot move a folder to a child of itself")
		}
	}

	// Changing the path string
	res := []Folder{}
    dstPath = fmt.Sprintf(`%s.`, dstPath)
    for _, f := range folders { 
        if f.OrgId == sourceOrgId && strings.Contains(f.Paths, name) {
			// this ensures that we are not finding a path that just contains the name of the folder
			parts := strings.Split(f.Paths, ".")
			nameIndex := -1
			for index, part := range parts {
				if part == name {
					nameIndex = index
				}
			}
			if nameIndex == -1 {
				res = append(res, f)
				continue
			}
			combinedString := strings.Join(parts[nameIndex:], ".")
            res = append(res, Folder{
					Name: f.Name,
					OrgId: f.OrgId,
					Paths: fmt.Sprintf(`%s%s`, dstPath, combinedString),
				})
        } else {
            res = append(res, f)
        }
    }

	return res, nil
}
