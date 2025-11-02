package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	ID   int
	Name string
}

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		rows, err := db.Query(c.Context(), "SELECT * FROM products")
		if err != nil {
			return err
		}
		products := []Product{}
		for rows.Next() {
			product := Product{}
			err := rows.Scan(&product.ID, &product.Name)
			if err != nil {
				return err
			}
			products = append(products, product)
		}
		c.Set("Content-Type", "text/html")
		return ProductList(products).Render(c.Context(), c)
	})

	app.Get("/add", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return ProductForm(Product{}).Render(c.Context(), c)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		_, err := db.Exec(c.Context(), "INSERT INTO products (name) VALUES ($1)", c.FormValue("name"))
		if err != nil {
			return err
		}
		return c.Redirect("/")
	})

	app.Get("/:id/edit", func(c *fiber.Ctx) error {
		product := Product{}
		err := db.QueryRow(c.Context(), "SELECT * FROM products WHERE id = $1", c.Params("id")).Scan(&product.ID, &product.Name)
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		return ProductForm(product).Render(c.Context(), c)
	})

	app.Post("/:id/edit", func(c *fiber.Ctx) error {
		_, err := db.Exec(
			c.Context(),
			"UPDATE products SET name = $1 WHERE id = $2",
			c.FormValue("name"),
			c.Params("id"),
		)
		if err != nil {
			return err
		}
		return c.Redirect("/")
	})

	app.Get("/:id/remove", func(c *fiber.Ctx) error {
		dummy := 0
		err := db.QueryRow(c.Context(), "SELECT 1 FROM products WHERE id = $1", c.Params("id")).Scan(&dummy)
		if err != nil {
			return err
		}
		c.Set("Content-Type", "text/html")
		return ProductConfirm().Render(c.Context(), c)
	})

	app.Post("/:id/remove", func(c *fiber.Ctx) error {
		_, err := db.Exec(
			c.Context(),
			"DELETE FROM products WHERE id = $1",
			c.Params("id"),
		)
		if err != nil {
			return err
		}
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3000"))
}
