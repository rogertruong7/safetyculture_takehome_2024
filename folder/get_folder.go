package folder

import (
	"fmt"
	"regexp"

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
	// Your code here...
	folders := f.folders
	pattern := fmt.Sprintf(`^%s.`, name)
	res := []Folder{}

	parentFolderPattern, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex: ", err)
		return res
	}

	count := 0

	for _, f := range folders {
		if f.OrgId == orgID && parentFolderPattern.MatchString(f.Paths) {
			fmt.Println("hello " + f.Paths)
			fmt.Println(parentFolderPattern)
			res = append(res, f)
			count++
		}
	}
	fmt.Println(count)
	return res
}
