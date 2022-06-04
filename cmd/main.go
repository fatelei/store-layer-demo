package main

import (
	"flag"
	"github.com/fatelei/store-layer-demo/pkg/fs"
)

func main() {
	var mountpoint string
	flag.StringVar(&mountpoint, "mountpoint", "", "mountpoint")
	flag.Parse()
	if len(mountpoint) == 0 {
		return
	}

	store := &fs.LayerStore{}
	err := fs.NewLayerStore(mountpoint, store)
	if err != nil {
		panic(err)
	}

	defer func() {
		store.Conn.Close()
	}()
}
