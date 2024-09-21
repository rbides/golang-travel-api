package travel

import (
	travel "golang-travel-api/internal/repository/postgresql"
	request "golang-travel-api/pkg/models/requests"
	models "golang-travel-api/pkg/models/travel"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	repo *travel.Repository
}

func New(r *travel.Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(c *gin.Context) { // return
	log.Println("Get Service")
	travels, err := s.repo.Get()
	if err != nil {
		log.Println("error", err)
	}
	for _, t := range travels {
		log.Println(t)
	}

}

func (s *Service) GetById(id uuid.UUID) { // return
	log.Println("GetById Service")
	t, err := s.repo.GetById(id)
	if err != nil {
		log.Println("error", err)
	}
	log.Println(t)
}

func (s *Service) Create(c *gin.Context, t *request.TravelRequest) error {
	log.Println("Create Service", time.Now())
	travel := models.Travel{
		ID:          uuid.New(),
		Name:        t.Name,
		Destination: t.Destination,
		Price:       models.Money(t.Price),
		Departure:   t.Departure,
	}
	seats := make([]models.Seat, t.Seats)
	for i := range t.Seats {
		seats[i].ID = uuid.New()
		seats[i].TravelID = travel.ID
		seats[i].Position = i + 1
	}

	return s.repo.Create(c, travel, seats)
}

func (s *Service) Update(c *gin.Context) {
	log.Println("Update Service")
	s.repo.Update()
}

func (s *Service) Delete(id uuid.UUID) {
	log.Println("Del Service")
	s.repo.Delete(id)
}
