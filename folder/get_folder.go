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
		fmt.Fprintln(os.Stderr, "Invalid orgID provided")
		return res
	}
	if name == "" || strings.Contains(name, ".") {
		fmt.Fprintln(os.Stderr, "Invalid folder name provided")
		return res 
	}
	folders := f.folders
	pattern := fmt.Sprintf(`%s.`, name)

	for _, f := range folders {
		if f.Paths == "" {
			fmt.Fprintln(os.Stderr, "Skipping folder with empty path")
			continue
		}
		if f.OrgId == orgID && strings.Contains(f.Paths, pattern) {
			res = append(res, f)
		}
	}
	return res
}
