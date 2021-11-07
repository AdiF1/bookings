package render
import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/justinas/nosurf"
)

// FuncMap is the type of the map defining the mapping from names to functions.
var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer sets the config for the template package
func NewRenderer (a *config.AppConfig) {
	app = a
}

// AddDefaultData allows app-wide info to be added to templates before rendering
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	return td
}

// Template renders templates using html/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {

	var tc map[string]*template.Template
	if app.UseCache {
		// production, so get the template cache from app config
		tc = app.TemplateCache
	} else {
		// dev mode, sp rebuild template cache on every render
		tc, _ = CreateTemplateCache()
	}

	// get the templateCache from app config
	//tc = app.TemplateCache
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("can't get template from cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buff, td)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}
	return nil
}
// CreateTemplateCache creates a template cache as a map.
// It's called in main.go
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))

	// the only error would be from a malformed pattern (ErrBadPattern)
	if err != nil {
		return myCache, err
	}
	// range through the templates
	for _, page := range pages {

		// Base returns the last element of path. Trailing path separators are removed before 
		// extracting the last element. If the path is empty, Base returns ".". If the path 
		// consists entirely of separators, Base returns a single separator.
		
		name := filepath.Base(page)

		// New allocates a new HTML template with the given name.

		// Funcs adds the elements of the argument map to the template's function map. 
		// It must be called before the template is parsed. It panics if a value in the map 
		// is not a function with appropriate return type. However, it is legal to 
		// overwrite elements of the map. The return value is the template, so calls can be chained.

		// ParseFiles parses the named files and associates the resulting templates with t. If 
		// an error occurs, parsing stops and the returned template is nil; otherwise it is t. 
		// There must be at least one file.

		// create a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts , err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache ,err
		}
		// MAYBE wrap the layout around each page before adding to ts
		if len(layouts) > 0 {
			// ParseGlob parses the template definitions in the files identified by the pattern 
			// and associates the resulting templates with t. The files are matched according to 
			// the semantics of filepath.Match, and the pattern must match at least one file. 
			// ParseGlob is equivalent to calling t.ParseFiles with the list of files matched by the pattern.
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}


/*_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache", err)
	}*/

	/*parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}*/