package swagger

// func TestSwagger(r *chi.Mux) {
// 	d := JSONRoutesDoc(r)
// 	file, err := os.Create("./docs/test.json")
// 	if err != nil {
// 		fmt.Printf("Error opening the file: %v\n", err)
// 		return
// 	}
// 	defer file.Close()
// 	_, err = file.WriteString(d)
// 	if err != nil {
// 		fmt.Printf("Error writing to the file: %v\n", err)
// 		return
// 	}
// 	// fmt.Println(d)
// }

// func BuildDoc(r chi.Routes) (Router, error) {
// 	d := Router{}

// 	// Walk and generate the router docs
// 	d.Route = buildDocRouter(r)
// 	return d, nil
// }

// func buildDocRouter(r chi.Routes) DocRouter {
// 	rts := r
// 	dr := DocRouter{Middlewares: []DocMiddleware{}}
// 	drts := DocRoutes{}
// 	dr.Routes = drts

// 	for _, mw := range rts.Middlewares() {
// 		dmw := DocMiddleware{
// 			FuncInfo: buildFuncInfo(mw),
// 		}
// 		dr.Middlewares = append(dr.Middlewares, dmw)
// 	}

// 	for _, rt := range rts.Routes() {
// 		drt := DocRoute{Pattern: rt.Pattern, Handlers: DocHandlers{}}

// 		if rt.SubRoutes != nil {
// 			subRoutes := rt.SubRoutes
// 			subDrts := buildDocRouter(subRoutes)
// 			drt.Router = &subDrts

// 		} else {
// 			hall := rt.Handlers["*"]
// 			for method, h := range rt.Handlers {
// 				if method != "*" && hall != nil && fmt.Sprintf("%v", hall) == fmt.Sprintf("%v", h) {
// 					continue
// 				}

// 				dh := DocHandler{Method: method, Middlewares: []DocMiddleware{}}

// 				var endpoint http.Handler
// 				chain, _ := h.(*chi.ChainHandler)

// 				if chain != nil {
// 					for _, mw := range chain.Middlewares {
// 						dh.Middlewares = append(dh.Middlewares, DocMiddleware{
// 							FuncInfo: buildFuncInfo(mw),
// 						})
// 					}
// 					endpoint = chain.Endpoint
// 				} else {
// 					endpoint = h
// 				}

// 				dh.FuncInfo = buildFuncInfo(endpoint)

// 				drt.Handlers[method] = dh
// 			}
// 		}

// 		drts[rt.Pattern] = drt
// 	}

// 	return dr
// }

// type Router struct {
// 	Route DocRouter `json:"paths"`
// }

// type DocRouter struct {
// 	Middlewares []DocMiddleware `json:"notUsing"`
// 	Routes      DocRoutes       `json:"routes"`
// }

// type DocMiddleware struct {
// 	FuncInfo
// }

// type DocRoute struct {
// 	Pattern  string      `json:"-"`
// 	Handlers DocHandlers `json:"handlers,omitempty"`
// 	Router   *DocRouter  `json:"path,omitempty"`
// }

// type DocRoutes map[string]DocRoute // Pattern : DocRoute

// type DocHandler struct {
// 	Middlewares []DocMiddleware `json:"middlewares"`
// 	Method      string          `json:"method"`
// 	FuncInfo
// }

// type FuncInfo struct {
// 	Func    string `json:"func"`
// 	Comment string `json:"comment"`
// }

// type DocHandlers map[string]DocHandler // Method : DocHandler

// func buildFuncInfo(i interface{}) FuncInfo {
// 	fi := FuncInfo{}
// 	frame := getCallerFrame(i)

// 	pkgName := getPkgName(frame.File)
// 	funcPath := frame.Func.Name()

// 	idx := strings.Index(funcPath, "/"+pkgName)
// 	if idx > 0 {
// 		fi.Func = funcPath[idx+2+len(pkgName):]
// 	} else {
// 		fi.Func = funcPath
// 	}

// 	fi.Comment = getFuncComment(frame.File, frame.Line)

// 	return fi
// }

// func JSONRoutesDoc(r chi.Routes) string {
// 	doc, _ := BuildDoc(r)
// 	v, err := json.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(v)
// }

// func GetFuncInfo(i interface{}) FuncInfo {
// 	fi := FuncInfo{}
// 	frame := getCallerFrame(i)

// 	pkgName := getPkgName(frame.File)
// 	funcPath := frame.Func.Name()

// 	idx := strings.Index(funcPath, "/"+pkgName)
// 	if idx > 0 {
// 		fi.Func = funcPath[idx+2+len(pkgName):]
// 	} else {
// 		fi.Func = funcPath
// 	}

// 	fi.Comment = getFuncComment(frame.File, frame.Line)

// 	return fi
// }

// func getCallerFrame(i interface{}) *runtime.Frame {
// 	value := reflect.ValueOf(i)
// 	if value.Kind() != reflect.Func {
// 		return nil
// 	}
// 	pc := value.Pointer()
// 	frames := runtime.CallersFrames([]uintptr{pc})
// 	if frames == nil {
// 		return nil
// 	}
// 	frame, _ := frames.Next()
// 	if frame.Entry == 0 {
// 		return nil
// 	}
// 	return &frame
// }

// func getPkgName(file string) string {
// 	fset := token.NewFileSet()
// 	astFile, err := parser.ParseFile(fset, file, nil, parser.PackageClauseOnly)
// 	if err != nil {
// 		return ""
// 	}
// 	if astFile.Name == nil {
// 		return ""
// 	}
// 	return astFile.Name.Name
// }

// func getFuncComment(file string, line int) string {
// 	fset := token.NewFileSet()

// 	astFile, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
// 	if err != nil {
// 		return ""
// 	}

// 	if len(astFile.Comments) == 0 {
// 		return ""
// 	}

// 	for _, cmt := range astFile.Comments {
// 		if fset.Position(cmt.End()).Line+1 == line {
// 			return cmt.Text()
// 		}
// 	}

// 	return ""
// }
