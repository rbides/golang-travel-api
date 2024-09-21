package travel

import (
	"database/sql"
	"golang-travel-api/internal/repository"
	models "golang-travel-api/pkg/models/travel"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Panic(err)
	}

	return &Repository{db}
}

func (r *Repository) Get() ([]models.Travel, error) {
	log.Println("Repo - Getting travels")
	rows, err := r.db.Query("SELECT * FROM travels")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []models.Travel
	for rows.Next() {
		t := models.Travel{}
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Destination,
			&t.Price,
			&t.Departure,
			&t.Created_At,
			&t.Updated_At,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Err Not found")
				return nil, repository.ErrNotFound
			}
			return nil, err
		}
		res = append(res, t)
	}
	if err := rows.Err(); err != nil {
		log.Println("Query error: ", err)
		return nil, err
	}
	return res, nil
}

func (r *Repository) GetById(id uuid.UUID) (models.Travel, error) {
	log.Println("Repo - Getting travel: ", id)
	row := r.db.QueryRow("SELECT * FROM travels WHERE id=$1", id)
	var t models.Travel
	if err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Destination,
		&t.Price,
		&t.Departure,
		&t.Created_At,
		&t.Updated_At,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("Err Not found")
			return t, repository.ErrNotFound
		}
		return t, err
	}
	return t, nil
}

// Create a Travel and its corresponding Seats
func (r *Repository) Create(c *gin.Context, t models.Travel, seats []models.Seat) error {
	log.Println("Creating travel with id: ", t)

	// TODO: handle errors

	tx, err := r.db.BeginTx(c, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(
		c,
		"INSERT INTO travels (id, name, destination, price, departure) VALUES ($1, $2, $3, $4, $5);",
		t.ID,
		t.Name,
		t.Destination,
		t.Price,
		t.Departure,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	// User CopyIn for faster bulk inserting seats
	stmt, err := tx.PrepareContext(c, pq.CopyIn("seats", "id", "travel_id", "position"))
	if err != nil {
		log.Println(err)
		return err
	}
	for _, s := range seats {
		_, err = stmt.ExecContext(
			c,
			s.ID,
			s.TravelID,
			s.Position,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	_, err = stmt.ExecContext(c)
	if err != nil {
		log.Println(err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println(err) // create error
		return err
	}

	return nil
}

func (r *Repository) Update() { //PATCH would be better?
	log.Println("Update")
	// create error
}

func (r *Repository) Delete(id uuid.UUID) {
	log.Println("Del")
	_, err := r.db.Exec(
		"DELETE FROM travels WHERE id = $1",
		id,
	)

	if err != nil {
		log.Println(err) // create error
	}
}
