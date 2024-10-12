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
const DefaultWrongOrgID = "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"
// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondOrgId := uuid.FromStringOrNil(DefaultWrongOrgID)
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
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondOrgId := uuid.FromStringOrNil(DefaultWrongOrgID)
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
		// invalid org id
		// invalid path
		// creative-scalphunter, wrong orgid
		// close-layla-miller wrong orgid
		// noble-vixen, doesnt contain just noble-vixen
		// nearby-secret
		// stunning-horridus
		{
			name: "Invalid orgID - no results",
			orgID: fakeOrgId,
			folderName: "stunning-horridus",
			folders: res,
			want: []folder.Folder{},
		},
		{
			name: "Get folders for first OrgID",
			orgID: orgID,
			folders: res,
			want: extractData("./testOutputs/getFirstOrgID.json"),
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
			get := f.GetAllChildFolders(tt.orgID, tt.folderName)
			assert.Equal(t, tt.want, get)
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