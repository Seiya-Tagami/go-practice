package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"techtrain-go-practice/model"
	"techtrain-go-practice/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedUrl, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Println(err)
			return
		}

		var (
			prevID int = 0
			size   int = 3
		)

		if len(parsedUrl["prev_id"]) != 0 {
			prevID, _ = strconv.Atoi(parsedUrl["prev_id"][0])
		}

		if len(parsedUrl["size"]) != 0 {
			size, _ = strconv.Atoi(parsedUrl["size"][0])
		}

		req := &model.ReadTODORequest{
			PrevID: int64(prevID),
			Size:   int64(size),
		}

		res, err := h.Read(r.Context(), req)
		if err != nil {
			log.Println(err)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			return
		}
	}

	if r.Method == "POST" {
		req := &model.CreateTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			return
		}

		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := h.Create(r.Context(), req)
		if err != nil {
			log.Println(err)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			return
		}
	}

	if r.Method == "PUT" {
		req := &model.UpdateTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			return
		}

		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := h.Update(r.Context(), req)
		if err != nil {
			log.Println(err)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			return
		}
	}

	if r.Method == "DELETE" {
		req := &model.DeleteTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			return
		}

		if len(req.IDs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := h.Delete(r.Context(), req)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}

	return &model.CreateTODOResponse{
		TODO: *todo,
	}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	todos, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
	if err != nil {
		return nil, err
	}
	return &model.ReadTODOResponse{
		TODOs: todos,
	}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}
	return &model.UpdateTODOResponse{
		TODO: *todo,
	}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	err := h.svc.DeleteTODO(ctx, req.IDs)
	if err != nil {
		return nil, err
	}
	return &model.DeleteTODOResponse{}, nil
}
