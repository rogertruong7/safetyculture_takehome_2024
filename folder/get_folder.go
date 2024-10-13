package folder

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	res := []Folder{}
	if orgID == uuid.Nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid orgID provided")
		return nil
	}
	if name == "" || strings.Contains(name, ".") {
		fmt.Fprintln(os.Stderr, "Error: Invalid folder name provided")
		return nil
	}

	folders := f.folders
	folderExists := false
	folderOrgId := uuid.Nil
	for _, f := range folders {
		if f.Name == name {
			folderExists = true
			folderOrgId = f.OrgId
			break
		}
	}
	if !folderExists {
		fmt.Fprintln(os.Stderr, "Error: Folder does not exist")
		return nil
	} else if folderOrgId != orgID {
		fmt.Fprintln(os.Stderr, "Error: Folder does not exist in the specified organization")
		return nil
	}

	orgIDFolders := f.GetFoldersByOrgID(orgID)
	pattern := fmt.Sprintf(`%s.`, name)
	for _, f := range orgIDFolders {
		if f.Paths == "" {
			continue
		}
		if strings.Contains(f.Paths, pattern) {
			res = append(res, f)
		}
	}
	return res
}
