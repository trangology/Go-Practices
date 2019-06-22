package actions

import (
	"errors"
	"homework08/elevennote/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// NotesResource is the resource for the Note model
type NotesResource struct {
	buffalo.Resource
}

// List gets all Notes. This function is mapped to the path
// GET /notes
func (v NotesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Note
	// note := &models.Note{}
	notes := &models.Notes{}

	q := tx.PaginateFromParams(c.Params())

	if err := q.All(notes); err != nil {
		return err
	}

	// if c.Request().Method == "GET" {
	// 	note := &models.Note{}
	// 	query := c.Request().FormValue("q")

	// 	tags := strings.Split(note.Tag, ",")
	// 	for _, t := range tags {
	// 		if query == t {
	// 			return c.Render(200, r.HTML('show.html'))
	// 		}
	// 	}
	// }

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, notes))
}

// Show gets the data for one Note. This function is mapped to
// the path GET /notes/{note_id}
func (v NotesResource) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	note := &models.Note{}

	if err := tx.Find(note, c.Param("note_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, note))
}

func (v NotesResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Note{}))
}

func (v NotesResource) Create(c buffalo.Context) error {
	note := &models.Note{}

	if err := c.Bind(note); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(note)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, note))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "note.created.success"))
	// and redirect to the notes index page
	return c.Render(201, r.Auto(c, note))
}

func (v NotesResource) Edit(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	note := &models.Note{}

	if err := tx.Find(note, c.Param("note_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, note))
}

func (v NotesResource) Update(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	note := &models.Note{}

	if err := tx.Find(note, c.Param("note_id")); err != nil {
		return c.Error(404, err)
	}

	if err := c.Bind(note); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(note)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, note))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "note.updated.success"))
	// and redirect to the notes index page
	return c.Render(200, r.Auto(c, note))
}

func (v NotesResource) Destroy(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	note := &models.Note{}

	if err := tx.Find(note, c.Param("note_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(note); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "note.destroyed.success"))
	// Redirect to the notes index page
	return c.Render(200, r.Auto(c, note))
}
