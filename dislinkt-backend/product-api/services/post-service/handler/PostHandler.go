package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/service"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"



	type PostHandler struct {
		PostService *service.PostService
		LogInfo *logrus.Logger
		LogError *logrus.Logger
	}

	func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		var postDto dto.PostDto
		err := json.NewDecoder(r.Body).Decode(&postDTO)
	
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "PostHandler",
				"action":   "CPOST530",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast json to PostDTO!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}


)
