package folder_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)
const DefaultSecondOrgID = "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"
// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondOrgId := uuid.FromStringOrNil(DefaultSecondOrgID)
	fakeOrgId := uuid.FromStringOrNil("doesntexist")
	res := folder.GetAllFolders()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "Get folders for first OrgID",
			orgID: orgID,
			folders: res,
			want: extractData("./testOutputs/getFirstOrgID.json"),
		},
		{
			name: "Fake orgID - no results",
			orgID: fakeOrgId,
			folders: res,
			want: []folder.Folder{},
		},
		{
			name: "Get folders for second OrgID",
			orgID: secondOrgId,
			folders: res,
			want: extractData("./testOutputs/getSecondOrgID.json"),
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.ElementsMatch(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondOrgId := uuid.FromStringOrNil(DefaultSecondOrgID)
	fakeOrgId := uuid.FromStringOrNil("doesntexist")
	res := folder.GetAllFolders()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folderName string
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here

		{
			name: "Invalid orgID - no results",
			orgID: fakeOrgId,
			folderName: "stunning-horridus",
			folders: res,
			want: nil,
		},
		{
			name: "Invalid folder name - no results",
			orgID: orgID,
			folderName: "noble-vixen.",
			folders: res,
			want: nil,
		},
		{
			name: "Invalid folder name 2 folders - no results",
			orgID: orgID,
			folderName: "noble-vixen.nearby-secret",
			folders: res,
			want: nil,
		},
		{
			name: "Empty folder name - no results",
			orgID: orgID,
			folderName: "",
			folders: res,
			want: nil,
		},
		{
			name: "First orgId, creative-scalphunter belongs to different orgid",
			orgID: orgID,
			folderName: "creative-scalphunter",
			folders: res,
			want: nil,
		},
		{
			name: "First orgId, folder doesn't exist at all",
			orgID: orgID,
			folderName: "hellohellohello",
			folders: res,
			want: nil,
		},
		{
			name: "Second orgId, creative-scalphunter does belong",
			orgID: secondOrgId,
			folderName: "creative-scalphunter",
			folders: res,
			want: extractData("./testOutputs/getChildrenWorking.json"),
		},
		{
			name: "Second orgId, topical-micromax not first folder does belong",
			orgID: secondOrgId,
			folderName: "topical-micromax",
			folders: res,
			want: extractData("./testOutputs/getChildrenOfChild.json"),
		},
		{
			name: "Second orgId, file has no children, equal-wonder-woman",
			orgID: secondOrgId,
			folderName: "equal-wonder-woman",
			folders: res,
			want: []folder.Folder{},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.folderName)
			assert.ElementsMatch(t, tt.want, get)
		})
	}
}

func extractData(fileName string) []folder.Folder {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, fileName)

	fmt.Println(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	folders := []folder.Folder{}
	err = json.Unmarshal(jsonByte, &folders)
	if err != nil {
		panic(err)
	}

	return folders
}