package travel

import (
	travel "golang-travel-api/internal/service"
	request "golang-travel-api/pkg/models/requests"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *travel.Service
}

func New(s *travel.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Get(c *gin.Context) {
	log.Println("Get Handler")
	h.svc.Get(c)
}

func (h *Handler) GetById(c *gin.Context) {
	log.Println("GetById Handler")
	id, err := uuid.Parse(c.Param("travelID"))
	if err != nil {
		log.Println("Invalid ID", err)
		return
	}
	h.svc.GetById(id)
}

func (h *Handler) Create(c *gin.Context) {
	log.Println("Create Handler")
	t := request.TravelRequest{}

	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println("Error json", err)
		c.JSON(http.StatusBadRequest, "Invalid input data")
		return
	}
	// Checks if departure date is > Now
	if b := t.Departure.After(time.Now()); !b {
		log.Println("Invalid departure date", t.Departure) // create a custom error for this
		c.JSON(http.StatusBadRequest, "Invalid departure date")
		return
	}

	if err := h.svc.Create(c, &t); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed creating Travel.")
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) Update(c *gin.Context) {
	log.Println("Update Handler")
	h.svc.Update(c)
}

func (h *Handler) Delete(c *gin.Context) {
	log.Println("Del Handler")
	id, err := uuid.Parse(c.Param("travelID"))
	if err != nil {
		log.Println("Invalid ID", err)
		return
	}
	h.svc.Delete(id)
}
