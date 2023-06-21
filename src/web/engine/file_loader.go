package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func (e Engine) fileLoader() (error) {
	component := dirPatches("component", [][]string{
		dirPatch("button", "page_back.html","page_next.html"),
		dirPatch("table", "table.html","tboby.html","thead.html","table_content.html"),
		dirPatch("tag", "active.html","expired.html","inactive.html","permanent.html"),
	})
	filepathes := []string{
		fmt.Sprintf("%s/pages/common.html", e.path),
	}

	for _, c := range component {
		filepathes = append(filepathes, fmt.Sprintf("%s/%s", e.path, c))
	}

	m, err := openFiles(filepathes...)
	if err != nil {
		return err
	}

	e.Htmls = m
	return nil
}

func openFiles(pathes ...string) (map[string]os.File, error){
	var returns map[string]os.File
	eg := new(errgroup.Group)
	
	for _, p := range pathes {
		p := p
		eg.Go(func() error {
			if filepath.Ext(p) == ".html" {
				file, err := os.Open(p)
				if err != nil {
					return err
				}
				defer file.Close()

				name := strings.Replace(filepath.Base(p), filepath.Ext(p), "", -1)
				returns[name] = *file
				return nil
			} else {
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	} else {
		return returns, nil
	}
}

func dirPatch(dires string, pathes ...string) []string {
	var returns []string

	for _, p := range pathes {
		returns = append(returns, fmt.Sprintf("%s/%s", dires, p))
	}

	return returns
}

func dirPatches(dires string, pathes [][]string) []string {
	var returns []string

	for _, p1 := range pathes {
		for _, p2 := range p1 {
			returns = append(returns, fmt.Sprintf("%s/%s", dires, p2))
		}
	}
	return returns
}