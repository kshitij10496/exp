package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
)

func main() {
	// Connect with DB.
	dsn := "host=localhost port=5432 dbname=test user=test password=secret sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect with DB: %v", err)
	}
	defer db.Close()

	fmt.Println("connected to DB")

	// if err := insertItem(db, 6, "attrs=pointer to empty value", &Attrs{}); err != nil {
	// 	log.Fatalf("failed to insert item: %v", err)
	// }

	for i := 1; i < 7; i++ {
		if err := readItem(db, i); err != nil {
			log.Fatalf("failed to read item %d: %v", i, err)
		}
	}
}

type Item struct {
	ID    int                `db:"id"`
	Name  string             `db:"name"`
	Attrs types.NullJSONText `db:"attrs"`
}

// // func (i Item) String() string {
// // 	return fmt.Sprintf("ID: %d\nName:%s\nAttrs:%s", i.ID, i.Name, i.Attrs)
// // }

// The Attrs struct represents the data in the JSON/JSONB column. We can use
// struct tags to control how each field is encoded.
type Attrs struct {
	Name        string   `json:"name,omitempty"`
	Ingredients []string `json:"ingredients,omitempty"`
	Organic     bool     `json:"organic,omitempty"`
	Dimensions  struct {
		Weight float64 `json:"weight,omitempty"`
	} `json:"dimensions,omitempty"`
}

// // Make the Attrs struct implement the driver.Valuer interface. This method
// // simply returns the JSON-encoded representation of the struct.
// func (a Attrs) Value() (driver.Value, error) {
// 	return json.Marshal(a)
// }

// // Make the Attrs struct implement the sql.Scanner interface. This method
// // simply decodes a JSON-encoded value into the struct fields.
// func (a *Attrs) Scan(value interface{}) error {
// 	b, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}

// 	return json.Unmarshal(b, &a)
// }

func insertItem(db *sqlx.DB, id int, name string, attrs *Attrs) error {
	// Initialize a new Attrs struct and add some values.

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err := db.Exec("INSERT INTO items (id, name, attrs) VALUES($1, $2, $3)", id, name, attrs)
	if err != nil {
		return fmt.Errorf("inserting item: %w", err)
	}
	return nil
}

func readItem(db *sqlx.DB, id int) error {
	var res Item
	if err := db.QueryRowx("SELECT id, name, attrs FROM items WHERE id=$1", id).StructScan(&res); err != nil {
		return fmt.Errorf("error reading item: %w", err)
	}

	fmt.Printf("%+v\n", res)

	if res.Attrs.Valid {
		var attrs Attrs
		if err := res.Attrs.Unmarshal(&attrs); err != nil {
			return fmt.Errorf("failed to unmarshal attrs: %w", err)
		}
		fmt.Println(attrs)
	}
	return nil
}
