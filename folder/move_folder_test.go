package folder_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	res := folder.GetAllFolders()

	tests := [...]struct {
		name    string
		source  string
		dst     string
		folders []folder.Folder
		want    []folder.Folder
		wantErr bool
	}{
		{
			name: "Move folder to itself - error",
			source:  "fast-watchmen",
			dst:     "fast-watchmen",
			folders: res,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Source folder does not exist - error",
			source:  "ILoveSafetyCulture",
			dst:     "settling-hobgoblin",
			folders: res,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Destination folder does not exist - error",
			source:  "settling-hobgoblin",
			dst:     "ILoveSafetyCulture",
			folders: res,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Cannot move to a different organization - error",
			source:  "fast-watchmen",
			dst:     "endless-red-hulk",
			folders: res, 
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Cannot move folder to a child of itself - error",
			source:  "helped-blackheart",
			dst:     "concrete-golden-guardian",
			folders: res, 
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid move to parent - no error",
			source:  "sensible-stardust",
			dst:     "stunning-horridus",
			folders: res, 
			want:    extractData("./testOutputs/moveFolderParent.json"),
			wantErr: false,
		},
		{
			name:    "Valid move across unrelated folders - no error",
			source:  "sensible-stardust",
			dst:     "noble-vixen",
			folders: res, 
			want:    extractData("./testOutputs/moveFolderUnrelated.json"),
			wantErr: false,
		},
		{
			name:    "Valid move with similar named folders - no error",
			source:  "csteady-insectc",
			dst:     "steady-insect",
			folders: res, 
			want:    extractData("./testOutputs/moveFolderSimilar.json"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.source, tt.dst)

			if tt.wantErr {
				assert.Error(t, err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, get)
			}
		})
	}
}
