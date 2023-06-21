package engine

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/shiro8613/litebans-go/src/utils"
)

func VarrMapper(r io.Reader, m Mapper) (string, error){

	reg := regexp.MustCompile("{([a-z._]{1,})}")

	var outstring string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		bytes := scanner.Bytes()

		if reg.Match(bytes) {
			sebyteses := reg.FindAllSubmatch(bytes, -1)
			for _, sebytes := range sebyteses {
				if len(sebytes) > 0 && len(sebytes)> 1 {
					b1 := sebytes[0]
					b2 := sebytes[1]
					
					utils.ContainSwitch(string(b2), map[string]func(){
						"var.": func() {
							vs := m.Var
							if vs != nil {
								vari := strings.Replace(string(b2), "var.", "", -1)
								if vari != "" {
									bytes = utils.Replacer(bytes, b1, vs[vari])
								} else {
									bytes = utils.Replacer(bytes, b1, "")
								}
							}
						},
						"component.": func() {
							vs := m.Component
							if vs != nil {
								vari := strings.Replace(string(b2), "component.", "", -1)
								if vari != "" {
									bytes = utils.Replacer(bytes, b1, vs[vari])
								} else {
									bytes = utils.Replacer(bytes, b1, "")
								}
							}
						},
						"content": func() {
							vs := m.Content
							if vs != nil {
								vari := vs["content"]
								if vari != "" {
									bytes = utils.Replacer(bytes, b1, vari)
								} else {
									bytes = make([]byte, 0)
								}
							} else {
								bytes = utils.Replacer(bytes, b1, "")
							}
						},
					})					
				}

			}
		}
				
		outstring += fmt.Sprintln(string(bytes))
	}


	return outstring, nil

}

