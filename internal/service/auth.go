package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nmluci/stellar-payment-lite/internal/config"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) FindUserByAccessToken(ctx context.Context, at string) (res *indto.UserRole, err error) {
	conf := config.Get()

	token, err := jwt.Parse(at, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("sign method invalid")
		} else if method != conf.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("sign method invalid")
		}
		return conf.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to parse JWT AT")
		return nil, errs.ErrInvalidCred
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		s.logger.Error().Err(err).Msgf("failed to authenticate JWT AT")
		return nil, err
	}

	if claims["sub"] != "at" {
		s.logger.Error().Err(err).Msgf("wrong JWT sub-claims")
		err = errs.ErrNoAccess
		return
	}

	data := claims["data"].(map[string]interface{})
	userID := int64(data["id"].(float64))

	res, err = s.repository.FindUserRoleByID(ctx, userID)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to fetch userdata")
		return
	}

	return
}

func (s *service) FindUserByRefreshToken(ctx context.Context, rt string) (res *indto.UserRole, err error) {
	conf := config.Get()

	token, err := jwt.Parse(rt, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("sign method invalid")
		} else if method != conf.JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("sign method invalid")
		}
		return conf.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to parse JWT AT")
		return nil, errs.ErrInvalidCred
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		s.logger.Error().Err(err).Msgf("failed to authenticate JWT AT")
		return nil, err
	}

	if claims["sub"] != "rt" {
		s.logger.Error().Err(err).Msgf("wrong JWT sub-claims")
		err = errs.ErrNoAccess
		return
	}

	data := claims["data"].(map[string]interface{})
	userID := int64(data["id"].(float64))

	res, err = s.repository.FindUserRoleByID(ctx, userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to fetch userdata")
		return
	}

	return
}

func (s *service) AuthRefreshToken(ctx context.Context, payload *dto.AuthRefreshTokenPayload) (res *dto.AuthResponse, err error) {
	userMeta, err := s.FindUserByRefreshToken(ctx, payload.RT)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to validate refresh token")
		err = errs.ErrNoAccess
		return
	}

	userData := &indto.UserDetail{
		UserID:    userMeta.UserID,
		Username:  userMeta.Username,
		LegalName: userMeta.Name,
		RoleID:    userMeta.RoleID,
	}

	newAT, err := s.newAccessToken(userData)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to generate access token")
		return
	}

	newRT, err := s.newRefreshToken(userData)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to generate refresh token")
		return
	}

	res = &dto.AuthResponse{
		UserID:       userMeta.UserID,
		RoleID:       userMeta.RoleID,
		Name:         userData.LegalName,
		Username:     userData.Username,
		AccessToken:  newAT,
		RefreshToken: newRT,
	}

	return
}

func (s *service) AuthLogin(ctx context.Context, payload *dto.AuthLoginPayload) (res *dto.AuthResponse, err error) {
	user, err := s.repository.FindUserByUsername(ctx, payload.Username)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to fetch userdata")
		return
	} else if user == nil {
		return nil, errs.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to validate user credentials")
		return nil, errs.ErrInvalidCred
	}

	var at, rt string
	at, err = s.newAccessToken(user)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to generate access token")
		return
	}

	rt, err = s.newRefreshToken(user)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to generate refresh token")
		return
	}

	res = &dto.AuthResponse{
		UserID:       user.UserID,
		RoleID:       user.RoleID,
		Name:         user.LegalName,
		Username:     user.Username,
		AccessToken:  at,
		RefreshToken: rt,
	}

	return
}

func (s *service) newAccessToken(user *indto.UserDetail) (signed string, err error) {
	conf := config.Get()

	claims := s.newATUserClaim(user.UserID, user.Username, conf.JWT_AT_EXPIRATION)
	accessToken := jwt.NewWithClaims(conf.JWT_SIGNING_METHOD, claims)
	signed, err = accessToken.SignedString(conf.JWT_SIGNATURE_KEY)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to signed new access token")
		return "", err
	}

	return
}

func (s *service) newRefreshToken(user *indto.UserDetail) (signed string, err error) {
	conf := config.Get()

	claims := s.newRTUserClaim(user.UserID, user.Username, conf.JWT_RT_EXPIRATION)
	refreshToken := jwt.NewWithClaims(conf.JWT_SIGNING_METHOD, claims)
	signed, err = refreshToken.SignedString(conf.JWT_SIGNATURE_KEY)
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to signed new access token")
		return "", err
	}

	return
}

func (s *service) newATUserClaim(id int64, username string, exp time.Duration) *jwt.MapClaims {
	return &jwt.MapClaims{
		"iss": config.Get().JWT_ISSUER,
		"exp": time.Now().Add(exp).Unix(),
		"data": map[string]interface{}{
			"id":       id,
			"username": username,
		},
		"sub": "at",
	}
}

func (s *service) newRTUserClaim(id int64, username string, exp time.Duration) *jwt.MapClaims {
	return &jwt.MapClaims{
		"iss": config.Get().JWT_ISSUER,
		"exp": time.Now().Add(exp).Unix(),
		"data": map[string]interface{}{
			"id":       id,
			"username": username,
		},
		"sub": "rt",
	}
}
