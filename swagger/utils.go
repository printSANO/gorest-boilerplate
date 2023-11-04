package swagger

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/printSANO/gorest-boilerplate/config"
)

func GenerateDoc(r *chi.Mux, host string, info Info) {
	pathsInfo, _ := generateEnpointMap(r)
	st := SwaggerTemplate{
		Version:     "2.0",
		SwaggerInfo: info,
		Host:        host,
		BasePath:    config.BasePath,
		Paths:       pathsInfo,
	}
	v, err := json.MarshalIndent(st, "", "	")
	if err != nil {
		panic(err)
	}
	s := string(v)
	log.Println("Generated swagger.json")
	file, err := os.Create("./docs/swaggertest.json")
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(s)
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n", err)
		return
	}
}

type SwaggerTemplate struct {
	Version     string    `json:"swagger"`
	SwaggerInfo Info      `json:"info"`
	Host        string    `json:"host"`
	BasePath    string    `json:"basePath"`
	Paths       MapMethod `json:"paths"`
}

type Info struct {
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Contact     Contacts `json:"contact"`
	License     Licenses `json:"license,omitempty"`
	Version     string   `json:"version"`
}

type Contacts struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type Licenses struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

func generateEnpointMap(r *chi.Mux) (MapMethod, error) {
	p := buildSwaggerPaths(r)

	return p, nil
}

func buildRoutes(r chi.Routes, m MapMethod, pattern string, basePath string) MapMethod {
	for _, rt := range r.Routes() {
		var currentPattern string
		if rt.Pattern != "/" {
			currentPattern = pattern + rt.Pattern
		} else {
			currentPattern = pattern
		}
		temp := strings.Replace(currentPattern, "/*", "", 2)
		if temp != basePath {
			routeMethods := make(MethodInfos)
			for method, h := range rt.Handlers {
				desTemp := strings.Replace(buildFuncInfo(h), "\n", "", 1)
				methodInfoTemp := MethodInfo{Description: desTemp}
				routeMethods[strings.ToLower(method)] = methodInfoTemp
			}
			m[strings.Replace(temp, basePath, "", 1)] = routeMethods
		}

		if rt.SubRoutes != nil {
			buildRoutes(rt.SubRoutes, m, currentPattern, basePath)
		}
	}
	return m
}

func buildFuncInfo(i interface{}) string {
	frame := getCallerFrame(i)
	comment := getFuncComment(frame.File, frame.Line)
	return comment
}

func getCallerFrame(i interface{}) *runtime.Frame {
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Func {
		return nil
	}
	pc := value.Pointer()
	frames := runtime.CallersFrames([]uintptr{pc})
	if frames == nil {
		return nil
	}
	frame, _ := frames.Next()
	if frame.Entry == 0 {
		return nil
	}
	return &frame
}

func getPkgName(file string) string {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, file, nil, parser.PackageClauseOnly)
	if err != nil {
		return ""
	}
	if astFile.Name == nil {
		return ""
	}
	return astFile.Name.Name
}

func getFuncComment(file string, line int) string {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return ""
	}

	if len(astFile.Comments) == 0 {
		return ""
	}

	for _, cmt := range astFile.Comments {
		if fset.Position(cmt.End()).Line+1 == line {
			return cmt.Text()
		}
	}

	return ""
}

func buildSwaggerPaths(r *chi.Mux) MapMethod {
	rts := r
	mm := MapMethod{}
	paths := buildRoutes(rts, mm, "", config.BasePath)

	return paths
}

type Path struct {
	Paths MapMethod `json:"paths"` // "paths"
}

type MapMethod map[string]MethodInfos // /movies : get

type MethodInfos map[string]MethodInfo // get : {description: ... etc}

type MethodInfo struct {
	Description string      `json:"description,omitempty"`
	Consumes    []string    `json:"consumes,omitempty"`
	Produces    []string    `json:"produces,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	Responses   Response    `json:"responses,omitempty"`
}

type Parameter struct {
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Name        string    `json:"name,omitempty"`
	In          string    `json:"in,omitempty"`
	Required    string    `json:"required,omitempty"`
	Schema      ResSchema `json:"schema,omitempty"`
}

type Response map[string]Parameter

type ResSchema struct {
	Type                 string `json:"type,omitempty"`
	Items                Item   `json:"items,omitempty"`
	AdditionalProperties bool   `json:"additionalProperties,omitempty"`
}

type Item struct {
	Ref string `json:"$ref,omitempty"`
}
