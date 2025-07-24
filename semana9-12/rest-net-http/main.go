package main

import (
	"net/http"
	"regexp"
	"rest-net-http/recipe"
)

var mux = http.NewServeMux() //Nos devuelve una struct ServerMux

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

type homeHandler struct{} //Para ser un Handler, debemos de implementar ServeHttp

type RecipesHandler struct {
}

type RecipeStore interface {
	Add(name string, recipe recipe.Recipe) error
	Get(name string) (recipe.Recipe, error)
	Update(name string, recipe recipe.Recipe) error
	List() (map[string]recipe.Recipe, error)
	Remove(name string) error
}

func (h homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario! que tal estás?"))
}

func (h RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}

}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request)  {}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request)    {}
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

func main() {
	mux.Handle("/", &homeHandler{}) //El mux es el encargado de multiplexar quien es el encargado de manejar la petición. En este caso es tu homeHandler.
	mux.Handle("/recipes", &RecipesHandler{})
	mux.Handle("/recipes/", &RecipesHandler{})

	///Crea y sirve un servidor en ese puerto
	http.ListenAndServe(":8080", mux) //Le podemos pasar el multiplexor ya que implementa ServeHTTP. Pero internamente tiene un map[string]handler (la línea 14)

}
