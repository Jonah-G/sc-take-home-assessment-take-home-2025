package folder

import "strings"
import "errors"
// import "fmt"

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders
	var srcFolder, dstFolder *Folder

	// make a copy for the output (no state persistance)
	newFolders := make([]Folder, len(folders))
	copy(newFolders, folders)

	// find src & dst folders
	for i := range folders {
		if folders[i].Name == name {
			srcFolder = &f.folders[i]
		}
		if folders[i].Name == dst {
			dstFolder = &f.folders[i]
		}

		if srcFolder != nil && dstFolder != nil {
			break
		}
	}

	// error handling
	if srcFolder == nil {
		return nil, errors.New("Error: Source folder does not exist")
	}
	if dstFolder == nil {
		return nil, errors.New("Error: Destination folder does not exist")
	}
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths+".") {
		return nil, errors.New("Error: Cannot move folder to a child of itself")
	}
	if srcFolder == dstFolder {
		return nil, errors.New("Error: Cannot move a folder to itself")
	}
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("Error: Cannot move a folder to a different organization")
	}

	newPathPrefix := dstFolder.Paths + "." + srcFolder.Name
	oldPathPrefix := srcFolder.Paths
	

	for i := range newFolders {
		if newFolders[i].Paths == srcFolder.Paths {
			// update  source folder path
			newFolders[i].Paths = newPathPrefix
		} else if strings.HasPrefix(newFolders[i].Paths, oldPathPrefix+".") {
			// update  child folders paths
			newFolders[i].Paths = strings.Replace(newFolders[i].Paths, oldPathPrefix, newPathPrefix, 1)
		}
	}

	return newFolders, nil
}
