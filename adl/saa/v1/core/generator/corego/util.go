package corego

import (
	"github.com/golangee/architecture/adl/saa/v1/astutil"
	"github.com/golangee/src/ast"
	"strings"
	"unicode"
)

// ModName returns the modules name.
func ModName(n ast.Node) string {
	var mod *ast.Mod
	if ast.ParentAs(n, &mod) {
		return mod.Name
	}

	return ""
}

// PkgName returns the parents package name.
func PkgName(n ast.Node) string {
	if p, ok := n.(*ast.Pkg); ok {
		elems := strings.Split(p.Path, "/")
		if len(elems) == 0 {
			return ""
		}

		return strings.Join(elems[:len(elems)-1], "/")
	}

	var pkg *ast.Pkg
	if ast.ParentAs(n, &pkg) {
		return PkgName(pkg)
	}

	return ""
}

// PkgRelativeName returns the relative path within the given module name.
func PkgRelativeName(n ast.Node) string {
	modName := ModName(n)
	pkgName := PkgName(n)

	return pkgName[len(modName)+1:]
}

func ShortModName(n ast.Node) string {
	name := ModName(n)
	elems := strings.Split(name, "/")
	if len(elems) > 0 {
		return elems[len(elems)-1]
	}

	return name
}

func MkFile(dst *ast.Prj, modName, pkgName, fname string) *ast.File {
	const preamble = "Code generated by golangee/architecture. DO NOT EDIT."

	mod := astutil.MkMod(dst, modName)
	mod.SetLang(ast.LangGo)
	pkg := astutil.MkPkg(mod, pkgName)
	file := astutil.MkFile(pkg, fname)
	file.SetPreamble(preamble)

	return file
}

// MakePublic converts aBc to ABc.
// Special cases:
//  * id becomes ID
func MakePublic(str string) string {
	if len(str) == 0 {
		return str
	}

	switch str {
	case "id":
		return "ID"
	default:
		return string(unicode.ToUpper(rune(str[0]))) + str[1:]
	}
}