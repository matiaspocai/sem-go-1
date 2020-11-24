package vinoteca

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/vinos",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/vinos/:id",
		function: getByID(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/postvino/",
		function: postVino(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/deletevino/:id",
		function: deleteVino(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/putvino/:id",
		function: putVino(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"vinos": s.FindAll(),
		})
	}
}

func getByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		c.JSON(http.StatusOK, gin.H{
			"vinos": s.FindByID(i),
		})
	}
}

func postVino(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vino Vino
		err := c.BindJSON(&vino)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Imprimiendo variable: ", vino)

		c.JSON(http.StatusOK, gin.H{
			"vinos": s.PostVino(vino),
		})
	}
}

func putVino(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var vino Vino
		erro := c.BindJSON(&vino)
		if erro != nil {
			log.Fatal(erro)
		}
		fmt.Println("Imprimiendo variable: ", vino)

		c.JSON(http.StatusOK, gin.H{
			"vinos": s.PutVino(i, vino),
		})
	}
}

func deleteVino(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		c.JSON(http.StatusOK, gin.H{
			"vinos": s.DeleteVino(i),
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
