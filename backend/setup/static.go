package setup

//
//import (
//  "github.com/labstack/echo/v4"
//  "net/http"
//  "net/url"
//  "os"
//  "path"
//  "path/filepath"
//  "strings"
//)
//
//func StaticWithConfig() echo.MiddlewareFunc {
//
//  return func(next echo.HandlerFunc) echo.HandlerFunc {
//    return func(c echo.Context) (err error) {
//      p := c.Request().URL.Path
//      if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
//        p = c.Param("*")
//      }
//      p, err = url.PathUnescape(p)
//      if err != nil {
//        return
//      }
//      name := filepath.Join(config.Root, path.Clean("/"+p)) // "/"+ for security
//
//      fi, err := os.Stat(name)
//      if err != nil {
//        if os.IsNotExist(err) {
//          if err = next(c); err != nil {
//            if he, ok := err.(*echo.HTTPError); ok {
//              if config.HTML5 && he.Code == http.StatusNotFound {
//                return c.File(filepath.Join(config.Root, config.Index))
//              }
//            }
//            return
//          }
//        }
//        return
//      }
//
//      if fi.IsDir() {
//        index := filepath.Join(name, config.Index)
//        fi, err = os.Stat(index)
//
//        if err != nil {
//          if config.Browse {
//            return listDir(t, name, c.Response())
//          }
//          if os.IsNotExist(err) {
//            return next(c)
//          }
//          return
//        }
//
//        return c.File(index)
//      }
//
//      return c.File(name)
//    }
//  }
//}
