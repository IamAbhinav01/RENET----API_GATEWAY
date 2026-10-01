package controllers

import (
	"net/http"
	dtos "renet/DTOs"
	"renet/middlewares"
	"renet/redis"
	"renet/services"
	"renet/utils/formatters"
	"renet/utils/helpers"
	"strconv"
	"strings"
	"time"
)

type AuthController struct {
	AuthService    services.UserService
	SessionManager *redis.SessionManager
}

func NewAuthController(serv services.UserService, session *redis.SessionManager) *AuthController {
	return &AuthController{
		AuthService:    serv,
		SessionManager: session,
	}
}

func setSessionCookie(w http.ResponseWriter, r *http.Request, sessionID string, maxAge int) {
	secure := r.TLS != nil || strings.EqualFold(r.Header.Get("X-Forwarded-Proto"), "https")
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func (cntrl AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	payload, ok := r.Context().Value(helpers.PayloadContextKet).(dtos.SignupRequestDTO)
	if !ok {
		formatters.ErrorResponse(w, http.StatusBadRequest, nil, "Invalid signup request")
		return
	}

	userID, err := cntrl.AuthService.SignUpUser(payload)

	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(err.Error(), "1062") {
			status = http.StatusConflict
		}
		formatters.ErrorResponse(w, status, err, "Error occured while signing the user")
		return
	}

	userSession := &redis.Session{
		Data: map[string]string{
			"user_id": strconv.Itoa(userID),
			"email":   payload.Email,
		},
	}

	err = cntrl.SessionManager.Migrate(r.Context(), userSession)
	if err == nil {
		setSessionCookie(w, r, userSession.Id, 86400)
	}

	formatters.SuccessResponse(w, http.StatusCreated, "User sign-up successfully")
}
func (cntrl AuthController) Login(w http.ResponseWriter, r *http.Request) {
	const loggedInDuration = 1 * time.Second
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		if elapsed < loggedInDuration {
			time.Sleep(loggedInDuration - elapsed)
		}
	}()

	payload := r.Context().Value(helpers.PayloadContextKet).(dtos.SignInRequestDTO)

	userID, err := cntrl.AuthService.SignInUser(payload)

	if err != nil {
		formatters.ErrorResponse(w, http.StatusInternalServerError, err, "Error occured while signing the user")
		return
	}

	user_session := &redis.Session{
		Data: map[string]string{
			"user_id": strconv.Itoa(userID),
			"email":   payload.Email,
		},
	}

	err = cntrl.SessionManager.Migrate(r.Context(), user_session)
	if err == nil {
		setSessionCookie(w, r, user_session.Id, 86400)
	} else {
		formatters.ErrorResponse(w, http.StatusInternalServerError, err, "Failed to create session")
		return
	}
	formatters.SuccessResponse(w, http.StatusCreated, "User sign-In successfully")
}
func (cntrl AuthController) LogOut(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_id")

	if err == nil {
		_ = cntrl.SessionManager.Store().Destroy(r.Context(), cookie.Value)
	}

	setSessionCookie(w, r, "", -1)

	formatters.SuccessResponse(w, http.StatusAccepted, "Logged-OUT successfully")
}

func (cntrl AuthController) CurrentUser(w http.ResponseWriter, r *http.Request) {
	data, ok := middlewares.SessionDataFromContext(r.Context())
	if !ok {
		formatters.ErrorResponse(w, http.StatusInternalServerError, nil, "Session data unavailable")
		return
	}

	formatters.SuccessResponse(w, http.StatusOK, data)
}
