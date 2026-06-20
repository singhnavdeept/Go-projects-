package routes
import(
	
	"github.com/gorilla/mux"
	"github.com/singhnavdeept/bookstore-go/pkg/controllers"

)
var RegisterBookstoreRoutes = func( router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.getBook ).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
	
}

func main(){

}