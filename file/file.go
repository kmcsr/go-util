
package util_file

import (
	os "os"
	io "io"
)

func IsExist(path string)(bool){
	s, err := os.Stat(path)
	return (s != nil) || (err != nil && os.IsExist(err))
}

func IsNotExist(path string)(bool){
	_, err := os.Stat(path)
	return err != nil && os.IsNotExist(err)
}

func IsFile(path string)(bool){
	s, _ := os.Stat(path)
	return s != nil && !s.IsDir()
}

func IsDir(path string)(bool){
	s, _ := os.Stat(path)
	return s != nil && s.IsDir()
}

func MakeDir(folderPath string, mode_ ...os.FileMode)(err error){
	mode := os.ModePerm
	if len(mode_) > 0 {
		mode = mode_[0]
	}
	if IsNotExist(folderPath){
		if ff := DirPath(folderPath); IsNotExist(ff) {
			err = MakeDir(ff, mode)
			if err != nil { return }
		}
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil { return }
	}
	err = os.Chmod(folderPath, mode)
	if err != nil { return }
	return nil
}

func CopyFile(src string, drt string)(written int64, err error){
	var (
		sfd *os.File
		dfd *os.File
		mode os.FileMode = os.ModePerm
	)
	sfd, err = os.Open(src)
	if err != nil { return }
	defer sfd.Close()

	if info, err := sfd.Stat(); err == nil {
		mode = info.Mode()
	}

	dfd, err = os.OpenFile(name, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, mode)
	if err != nil { return }
	defer dfd.Close()

	written, err = io.Copy(dfd, sfd)
	if err != nil { return }

	err = dfd.Chmod(info.Mode())
	return
}

func CopyDir(src string, drt string)(err error){
	var (
		sfd *os.File
		dirinfo os.FileInfo
		finfos []os.FileInfo
	)
	sfd, err = os.Open(src)
	if err != nil { return }
	defer sfd.Close()

	dirinfo, err = sfd.Stat()
	if err != nil { return }

	err = MakeDir(drt, dirinfo.Mode())
	if err != nil { return }
	defer func(){ if err != nil {
		os.RemoveAll(drt)
	}}()

	finfos, err = sfd.Readdir(-1)
	if err != nil { return }

	sfd.Close()

	for _, info := range finfos {
		sf := JoinPath(src, info.Name())
		df := JoinPath(drt, info.Name())
		if info.IsDir() {
			err = CopyDir(df, df)
		}else{
			_, err = CopyFile(sf, df)
		}
		if err != nil { return }
	}
	return nil
}

type WalkEnity struct{
	root string
	path string
	full string
	parent string
	info os.FileInfo
}

func (e *WalkEnity)Root()(string){
	return e.root
}

func (e *WalkEnity)Path()(string){
	return e.path
}

func (e *WalkEnity)FullPath()(string){
	return e.full
}

func (e *WalkEnity)ParentPath()(string){
	return e.parent
}

func (e *WalkEnity)Name()(string){
	return e.info.Name()
}

func (e *WalkEnity)IsDir()(bool){
	return e.info.IsDir()
}

func (e *WalkEnity)Info()(os.FileInfo){
	return e.info
}

type WalkFunc func(e *WalkEnity, err error)(error)

func walkDir(root string, path string, call WalkFunc)(err error){
	base := JoinPath(root, path)
	list, err := os.ReadDir(base)
	if err != nil { return }
	var rinfo os.FileInfo
	for _, f := range list {
		p, p2 := JoinPath(base, f.Name()), JoinPath(path, f.Name())
		rinfo, err = os.Lstat(p)
		err = call(&WalkEnity{
			root: root,
			path: p2,
			full: p,
			parent: path,
			info: rinfo,
		}, err)
		if err != nil { return }
		if f.IsDir(){
			err = walkDir(root, p2, call)
			if err != nil { return }
		}
	}
	return
}

func Walk(root string, call WalkFunc)(err error){
	rinfo, err := os.Lstat(root)
	if err != nil {
		return
	}
	err = call(&WalkEnity{
		root: root,
		path: "",
		parent: "",
		full: root,
		info: rinfo,
	}, nil)
	if err != nil { return }
	if rinfo.IsDir(){
		return walkDir(root, "", call)
	}
	return nil
}

