package main

import (
	"os"
	"path/filepath"
	"strconv"
	"io/ioutil"
	"path"
	"github.com/softleader/captain-kube/tmpl"
)

const (
	readme        = "README.md"
	readmeContent = `# Table Of Contents
{{- range $key, $courses := . }}

## {{ $key }}

| Season | Course |
|---|---|
{{- range $index, $course := $courses }}
| {{ $course.Season }} | [{{ $course.Course }}](../{{ $course.Year }}/{{ $course.Season }}/{{ $course.Course }}) |
{{- end }}
{{- end }}
`
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root := filepath.Dir(wd)

	var courses []Course

	// collect course dir
	walkDir(root, 0, 2, func(path string) {
		season := filepath.Dir(path)
		year := filepath.Dir(season)
		y, e := strconv.Atoi(filepath.Base(year))
		if e == nil {
			c := Course{
				Year:   y,
				Season: filepath.Base(season),
				Course: filepath.Base(path),
			}
			courses = append(courses, c)
		}
	})

	// Reversing
	//for i := len(courses)/2 - 1; i >= 0; i-- {
	//	opp := len(courses) - 1 - i
	//	courses[i], courses[opp] = courses[opp], courses[i]
	//}

	// group by year
	groupByYear := make(map[int][]Course)
	for _, c := range courses {
		groupByYear[c.Year] = append(groupByYear[c.Year], c)
	}

	err = tmpl.CompileTo(readmeContent, groupByYear, filepath.Join(wd, readme))
	if err != nil {
		panic(err)
	}

}

// 在第 depth 層的時候依序針對當下的 dir 執行 walkFn
func walkDir(dirpath string, currentDepth int, depth int, walkFn func(path string)) {
	if currentDepth > depth {
		return
	}
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			p := path.Join(dirpath, file.Name())
			if currentDepth == depth {
				walkFn(p)
			}
			walkDir(p, currentDepth+1, depth, walkFn)
			continue
		} else {
			continue
		}
	}
}

type Course struct {
	Year   int
	Season string
	Course string
}
