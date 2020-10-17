package main

import (
	"fmt"
	"os/exec"
	"sync"
)

var (
	gits = []string{
		"bbloom",
		"go-bitswap",
		"go-block-format",
		"go-blockservice",
		"go-cid",
		"go-cidutil",
		"go-datastore",
		"go-ds-badger",
		"go-ds-flatfs",
		"go-ds-leveldb",
		"go-ds-measure",
		"go-filestore",
		"go-fs-lock",
		"go-graphsync",
		"go-ipfs",
		"go-ipfs-blockstore",
		"go-ipfs-chunker",
		"go-ipfs-cmds",
		"go-ipfs-config",
		"go-ipfs-delay",
		"go-ipfs-ds-help",
		"go-ipfs-exchange-interface",
		"go-ipfs-exchange-offline",
		"go-ipfs-files",
		"go-ipfs-pinner",
		"go-ipfs-posinfo",
		"go-ipfs-pq",
		"go-ipfs-provider",
		"go-ipfs-routing",
		"go-ipfs-util",
		"go-ipld-cbor",
		"go-ipld-format",
		"go-ipld-git",
		"go-ipns",
		"go-log",
		"go-merkledag",
		"go-metrics-interface",
		"go-mfs",
		"go-path",
		"go-peertaskqueue",
		"go-unixfs",
		"go-verifcid",
		"interface-go-ipfs-core",
	}
)

func main() {
	path := ""
	var wg sync.WaitGroup
	for _, g := range gits {
		wg.Add(1)
		go func(g string) {
			defer wg.Done()
			path = "git@github.com:ipfs/" + g + ".git"
			fmt.Println(path)
			cmd := exec.Command("git", "clone", path)
			cmd.Dir = "/Users/shuxian/Projects/joincloud/debug-libs/ipfs"
			stdout, err := cmd.CombinedOutput()
			fmt.Printf("start %s-%s\n", stdout, err)
		}(g)
	}

	wg.Wait()
}
