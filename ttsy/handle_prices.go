package main
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func HandlePagesPrices(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//images, err := globalImageStore.FindAll(0)
	//if err != nil {
	//	panic(err)
	//}
	RenderTemplate(w, r, "prices/index", map[string]interface{}{
		"SeoDescription": "Это SEO Description",
		"SeoKeywords": "Это SEO Keywords",
		"Title": "Это Титле",
		"SLastDate": "05.05.2015",
	})
}
