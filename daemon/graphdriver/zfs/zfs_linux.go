package zfs

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/graphdriver"
	"golang.org/x/sys/unix"
)

func checkRootdirFs(rootdir string) error {
	var buf unix.Statfs_t
	if err := unix.Statfs(rootdir, &buf); err != nil {
		return fmt.Errorf("Failed to access '%s': %s", rootdir, err)
	}

	if graphdriver.FsMagic(buf.Type) != graphdriver.FsMagicZfs {
		logrus.Debugf("[zfs] no zfs dataset found for rootdir '%s', found '%s' filesystem", rootdir, buf.Type)
		return graphdriver.ErrPrerequisites
	}

	return nil
}

func getMountpoint(id string) string {
	return id
}
