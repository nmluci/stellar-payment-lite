package service

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/commonkey"
	"github.com/nmluci/stellar-payment-lite/internal/model"
	"github.com/nmluci/stellar-payment-lite/internal/util/ctxutil"
	"github.com/nmluci/stellar-payment-lite/internal/util/timeutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(ctx context.Context, payload *dto.UserRegistrationPayload) (err error) {
	if payload.Username == "" || payload.LegalName == "" || payload.Password == "" {
		return errs.ErrBrokenUserReq
	}

	if exist, err := s.repository.FindUserByUsername(ctx, payload.Username); err != nil {
		s.logger.Error().Err(err).Msgf("an error occurred while validating username existence err: %+v", err)
		return errs.ErrUnknown
	} else if exist != nil {
		return errs.ErrDuplicatedResources
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error().Err(err).Msgf("an error occurred while hashing password err: %+v", err)
		return
	} else {
		payload.Password = string(hashed)
	}

	usr := &model.User{
		RoleID:   commonkey.ROLE_USER,
		Username: payload.Username,
		Password: payload.Password,
	}
	cust := &model.Customer{
		LegalName:  payload.LegalName,
		Address:    payload.Address,
		Phone:      payload.Phone,
		Birthplace: payload.Birthplace,
		Birthdate:  timeutil.ParseDate(payload.Birthdate),
		NIK:        payload.NIK,
		Occupation: payload.Occupation,
		KTPUrl:     payload.KTPUrl,
	}

	err = s.repository.InsertUser(ctx, usr, cust)
	if err != nil {
		s.logger.Error().Err(err).Msgf("an error occurred while inserting user err: %+v", err)
		return
	}

	return
}

func (s *service) GetUserDetailByID(ctx context.Context, params *dto.UserQueryParams) (res *dto.UserResponse, err error) {
	ctxMeta := ctxutil.GetUserCTX(ctx)
	if ctxMeta.RoleID != commonkey.ROLE_ADMIN && ctxMeta.UserID != params.UserID {
		return nil, errs.ErrNoAccess
	}

	usermeta, err := s.repository.FindUserByID(ctx, params.UserID)
	if err != nil {
		s.logger.Error().Err(err).Msg("an error occurred while fetching user meta")
		return
	} else if usermeta == nil {
		s.logger.Error().Msg("failed to find user")
		return nil, errs.ErrNotFound
	}

	res = &dto.UserResponse{
		UserID:     usermeta.UserID,
		CustomerID: usermeta.CustomerID,
		Username:   usermeta.Username,
		RoleID:     usermeta.RoleID,
	}

	return
}

func (s *service) UpdateUserByID(ctx context.Context, params *dto.UserQueryParams, payload *dto.UserPayload) (err error) {
	ctxMeta := ctxutil.GetUserCTX(ctx)
	if ctxMeta.RoleID != commonkey.ROLE_ADMIN && ctxMeta.UserID != params.UserID {
		return errs.ErrNoAccess
	}

	if payload.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed to hash password")
			return err
		}

		payload.Password = string(hashed)
	}

	accModel := &model.User{
		UserID:   params.UserID,
		Password: payload.Password,
	}
	err = s.repository.UpdateUser(ctx, accModel)
	if err != nil {
		s.logger.Error().Err(err).Msg("an error occurred while updating userdata")
		return
	}

	return
}
