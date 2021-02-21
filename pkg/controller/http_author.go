package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/command"
	"github.com/maestre3d/newton/internal/query"
	"github.com/maestre3d/newton/pkg/httputil"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// AuthorHTTP aggregate.Author HTTP endpoints
type AuthorHTTP struct {
	app *application.Author
}

// NewAuthorHTTP creates an aggregate.Author HTTP controller
func NewAuthorHTTP(app *application.Author) httputil.ControllersFx {
	return httputil.ControllersFx{
		Controller: &AuthorHTTP{app: app},
	}
}

// Route maps aggregate.Author exposed use cases using the given mux.Router
func (h AuthorHTTP) Route(r *mux.Router) {
	r.Path("/authors").Methods(http.MethodPost).HandlerFunc(h.create)
	r.Path("/authors").Methods(http.MethodGet).HandlerFunc(h.list)
	r.Path("/authors/{id}").Methods(http.MethodGet).HandlerFunc(h.get)
	r.Path("/authors/{id}").Methods(http.MethodPatch, http.MethodPut).HandlerFunc(h.update)
	r.Path("/authors/{id}").Methods(http.MethodDelete).HandlerFunc(h.delete)
	r.Path("/authors/{id}/state").Methods(http.MethodDelete, http.MethodPatch, http.MethodPatch).
		HandlerFunc(h.changeState)
	r.Path("/authors/{id}/image").Methods(http.MethodPatch, http.MethodPut).HandlerFunc(h.uploadPicture)
}

func (h AuthorHTTP) get(w http.ResponseWriter, r *http.Request) {
	a, err := query.GetAuthorByIDHandle(h.app, r.Context(), query.GetAuthorByID{
		ID: mux.Vars(r)["id"],
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, a)
}

func (h AuthorHTTP) create(w http.ResponseWriter, r *http.Request) {
	id, err := gonanoid.New(16)
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	err = command.CreateAuthorHandle(h.app, r.Context(), command.CreateAuthor{
		ID:          id,
		DisplayName: r.PostFormValue("display_name"),
		CreateBy:    r.PostFormValue("create_by"),
		Image:       r.PostFormValue("image"),
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusCreated, id)
}

func (h AuthorHTTP) list(w http.ResponseWriter, r *http.Request) {
	criteria, err := httputil.UnmarshalCriteriaJSON(r)
	if err != nil {
		criteria = httputil.UnmarshalCriteria(r)
	}

	as, err := query.ListAuthorsHandle(h.app, r.Context(), query.ListAuthors{
		Criteria: criteria,
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, as)
}

func (h AuthorHTTP) update(w http.ResponseWriter, r *http.Request) {
	err := command.UpdateAuthorHandle(h.app, r.Context(), command.UpdateAuthor{
		ID:          mux.Vars(r)["id"],
		DisplayName: r.PostFormValue("display_name"),
		CreateBy:    r.PostFormValue("create_by"),
		Image:       r.PostFormValue("image"),
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, nil)
}

func (h AuthorHTTP) changeState(w http.ResponseWriter, r *http.Request) {
	state := true
	if r.Method == http.MethodDelete {
		state = false
	}
	err := command.ChangeAuthorStateHandle(h.app, r.Context(), command.ChangeAuthorState{
		ID:    mux.Vars(r)["id"],
		State: state,
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, nil)
}

func (h AuthorHTTP) delete(w http.ResponseWriter, r *http.Request) {
	err := command.DeleteAuthorHandle(h.app, r.Context(), command.DeleteAuthor{
		ID: mux.Vars(r)["id"],
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, nil)
}

func (h AuthorHTTP) uploadPicture(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Print(err)
			return
		}
	}()
	err = command.UploadAuthorPictureHandle(h.app, r.Context(), command.UploadAuthorPicture{
		ID:       mux.Vars(r)["id"],
		Filename: header.Filename,
		Size:     header.Size,
		Image:    file,
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}

	httputil.RespondJSON(w, http.StatusOK, struct {
		Filename string `json:"filename"`
		Size     int64  `json:"size"`
	}{
		Filename: mux.Vars(r)["id"],
		Size:     header.Size,
	})
}
