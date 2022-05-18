package handler
import (
	"encoding/json"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/service"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	type CommentHandler struct {
		Service   *service.CommentService
		Validator *validator.Validate
		LogInfo *logrus.Logger
		LogError *logrus.Logger
	}
)