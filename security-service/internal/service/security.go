package service

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/Dimoonevs/SportsApp/auth-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/auth-service/pkg/data"
	"github.com/Dimoonevs/SportsApp/auth-service/pkg/email"
	"github.com/Dimoonevs/SportsApp/auth-service/pkg/utils"
	proto "github.com/Dimoonevs/SportsApp/auth-service/proto/security"
)

type SecurityService struct {
	proto.UnimplementedSecurityServiceServer
	Repo       repository.Repository
	JwtWrapper utils.JwtWrapper
}

func (s *SecurityService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	requestPayload := data.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Username:    req.Username,
		DateOfBirth: req.DateOfBirth,
		Status:      data.Status(req.Status),
		Email:       req.Email,
		Password:    utils.HashPassword(req.Password),
		TimeZone:    req.TimeZone,
	}

	exist, idUser, err := s.Repo.UsernameAndEmailExist(ctx, requestPayload.Username, requestPayload.Email)
	if err != nil {
		return nil, err
	}
	var userData *data.User
	if idUser > 0 {
		userData, err = s.Repo.GetUserById(ctx, idUser)
		if err != nil {
			return nil, err
		}
	}
	if exist && userData.Active {
		return nil, fmt.Errorf("username or email already exist")
	} else if exist && !userData.Active {
		s.SendConfirmCode(ctx, requestPayload.Email, idUser)
		return nil, errors.New(fmt.Sprintf("User with username %s and email %s already exist, id: %d. But none active", requestPayload.Username, requestPayload.Email, idUser))
	}

	id, err := s.Repo.Register(ctx, &requestPayload)
	if err != nil {
		return nil, err
	}
	err = s.SendConfirmCode(ctx, requestPayload.Email, id)
	if err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{
		Message: fmt.Sprintf("User %s %s registered successfully with username %s, id: %d", requestPayload.FirstName, requestPayload.LastName, requestPayload.Username, id),
		Id:      id,
	}, nil
}

func (s *SecurityService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	exist, idUser, err := s.Repo.UsernameAndEmailExist(ctx, req.Username, req.Email)
	if !exist {
		return nil, errors.New(fmt.Sprintf("username or email does not exist"))
	}
	if err != nil {
		return nil, err
	}
	userData, err := s.Repo.GetUserById(ctx, idUser)
	if err != nil {
		return nil, err
	}
	if exist && !userData.Active {
		s.SendConfirmCode(ctx, userData.Email, idUser)
		return nil, errors.New("User not active. id: " + fmt.Sprint(idUser))
	}
	if !utils.CheckPasswordHash(req.Password, userData.Password) {
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := s.JwtWrapper.GenerateToken(*userData)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{
		Token:    token,
		Message:  fmt.Sprintf("User %s logged in successfully", userData.Username),
		Id:       idUser,
		Username: userData.Username,
	}, nil
}

func (s *SecurityService) Validate(ctx context.Context, req *proto.ValidateRequest) (*proto.Response, error) {
	resp, err := s.JwtWrapper.ValidateToken(req.Token)
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}
	exist, _, err := s.Repo.UsernameAndEmailExist(ctx, resp.Usrname, resp.Email)
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}
	if !exist {
		return nil, fmt.Errorf("username or email does not exist")
	}
	user, err := s.Repo.GetUserById(ctx, resp.Id)
	if err != nil {
		return nil, err
	}
	if !user.Active {
		return nil, fmt.Errorf("user not active")
	}
	return &proto.Response{
		Message: resp.Usrname,
		Id:      resp.Id,
	}, nil
}

func (s *SecurityService) GetCoachTimeZone(ctx context.Context, req *proto.GetCoachTimeZoneRequest) (*proto.Response, error) {
	resp, err := s.Repo.GetCoachTimeZone(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.Response{
		Message: *resp,
	}, nil
}
func (s *SecurityService) ConfirmEmail(ctx context.Context, req *proto.ConfirmEmailRequest) (*proto.Response, error) {
	err := s.Repo.ConfirmEmailByUserId(ctx, req.Code, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Response{
		Message: "Email confirmed successfully",
	}, nil
}

func SendEmailForConfirmEmail(to, code string) error {
	err := email.SendEmail(to, "Secret CODE", "d12k13v15b@gmail.com", "tphy kblx mwww ajsj", "smtp.gmail.com", code)
	if err != nil {
		return err
	}
	return nil
}
func (s *SecurityService) SendConfirmCode(ctx context.Context, emailUser string, id int32) error {
	code, err := email.GenerateCode()
	if err != nil {
		return err
	}

	errs := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err = SendEmailForConfirmEmail(emailUser, code)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		defer wg.Done()
		err = s.Repo.SaveCode(ctx, id, code)
		if err != nil {
			errs <- err
		}
	}()

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SecurityService) ResetPassword(ctx context.Context, req *proto.ResetPasswordRequest) (*proto.Response, error) {
	exist, idUser, err := s.Repo.UsernameAndEmailExist(ctx, "", req.Email)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("username or email does not exist")
	}
	s.SendConfirmCode(ctx, req.Email, idUser)
	return &proto.Response{
		Message: fmt.Sprintf("Code sent to %s", req.Email),
		Id:      idUser,
	}, nil
}

func (s *SecurityService) ConfirmResetPassword(ctx context.Context, req *proto.ConfirmResetPasswordRequest) (*proto.Response, error) {
	req.Password = utils.HashPassword(req.Password)
	err := s.Repo.ConfirmResetPassword(ctx, req.Code, req.Password, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Response{
		Message: "Password changed successfully",
		Id:      req.Id,
	}, nil
}

func (s *SecurityService) GetCoach(ctx context.Context, req *proto.GetCoachRequest) (*proto.GetCoachResponse, error) {
	user, err := s.Repo.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.GetCoachResponse{
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}, nil
}

func (s *SecurityService) UpdateCoach(ctx context.Context, req *proto.CoachData) (*proto.ValidateRequest, error) {
	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		token string
	)
	errc := make(chan error, 2)
	user, err := s.Repo.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		exist, _, err := s.Repo.UsernameAndEmailExist(ctx, req.Username, "")
		if err != nil || (exist && req.Username != req.Username) {
			errc <- err
			return
		}
		if err := s.Repo.UpdateCoach(ctx, req); err != nil {
			errc <- err
			return
		}
		err = s.SendConfirmCode(ctx, user.Email, user.Id)
		if err != nil {
			errc <- err
			return
		}
	}()
	if req.Username != user.Username {
		wg.Add(1)
		go func() {
			defer wg.Done()
			user.Username = req.Username
			mu.Lock()
			defer mu.Unlock()
			token, err = s.JwtWrapper.GenerateToken(*user)
			if err != nil {
				errc <- err
				return
			}
		}()
	}

	wg.Wait()
	close(errc)

	for err := range errc {
		if err != nil {
			return nil, err
		}
	}
	return &proto.ValidateRequest{
		Token: token,
	}, nil
}

func (s *SecurityService) ChangeEmail(ctx context.Context, req *proto.ChangeEmailReq) (*proto.Response, error) {
	err := s.Repo.ChangeEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	s.SendConfirmCode(ctx, req.Email, req.Id)
	return &proto.Response{
		Message: fmt.Sprintf("Code sent to %s", req.Email),
		Id:      req.Id,
	}, nil
}
