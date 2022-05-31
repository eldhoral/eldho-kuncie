package service

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/repository"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{
		offerRepo: repo,
	}
}

type service struct {
	offerRepo repository.Repository
}

//Loan limit
func (s service) GetLoanLimitByID(id int64) (int, *modelLanding.LoanLimit, error) {
	repo, err := s.offerRepo.GetLoanLimitByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Loan limit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) GetLoanLimit() (int, *modelLanding.LoanLimit, error) {
	repo, err := s.offerRepo.GetLoanLimit()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Loan limit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateLoanLimit(limit string) (int, *modelLanding.LoanLimit, error) {
	limitLoan, _ := strconv.Atoi(limit)
	model := &modelLanding.LoanLimit{
		Limit: int64(limitLoan),
	}
	repo, err := s.offerRepo.CreateLoanLimit(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateLoanLimit(limit string) (int, error) {
	_, err := s.offerRepo.UpdateLoanLimit(limit)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Loan limit ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteLoanLimit() (int, error) {
	status, err := s.offerRepo.DeleteLoanLimit()
	if err == sql.ErrNoRows {
		return status, errors.New("Loan limit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

//Benefit
func (s service) GetBenefitByID(id int64) (int, *modelLanding.Benefit, error) {
	repo, err := s.offerRepo.GetBenefitByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListBenefit() (int, []modelLanding.Benefit, error) {
	repo, err := s.offerRepo.ListBenefit()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateBenefit(title string, description string, path string) (int, *modelLanding.Benefit, error) {
	model := &modelLanding.Benefit{
		Title:       title,
		Description: description,
		Image:       path,
	}
	repo, err := s.offerRepo.CreateBenefit(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateBenefitByID(id int64, params data.Params, path string) (int, error) {
	_, err := s.offerRepo.UpdateBenefitByID(id, params, path)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteBenefitByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteBenefitByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Benefit ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Loan method
func (s service) GetLoanMethodByID(id int64) (int, *modelLanding.LoanMethod, error) {
	repo, err := s.offerRepo.GetLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Loan method not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListLoanMethod() (int, []modelLanding.LoanMethod, error) {
	repo, err := s.offerRepo.ListLoanMethod()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Loan method not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateLoanMethod(title string, description string) (int, *modelLanding.LoanMethod, error) {
	model := &modelLanding.LoanMethod{
		Title:       title,
		Description: description,
	}
	repo, err := s.offerRepo.CreateLoanMethod(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateLoanMethodByID(id int64, title string, description string) (int, error) {
	_, err := s.offerRepo.UpdateLoanMethodByID(id, title, description)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Loan method ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteLoanMethodByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Loan method ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Tnc
func (s service) GetTncByID(id int64) (int, *modelTnc.Tnc, error) {
	repo, err := s.offerRepo.GetTncByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Tnc not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListTnc() (int, []modelTnc.Tnc, error) {
	repo, err := s.offerRepo.ListTnc()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Tnc not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateTnc(title string) (int, *modelTnc.Tnc, error) {
	model := &modelTnc.Tnc{
		Title: title,
	}
	repo, err := s.offerRepo.CreateTnc(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateTncByID(id int64, params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Tnc ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) UpdateTncMobile(params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncMobile(params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Tnc Row not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteTncByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteTncByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Tnc ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Tnc title
func (s service) GetTncTitleByID(id int64) (int, *modelTnc.TncTitle, error) {
	repo, err := s.offerRepo.GetTncTitleByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Tnc Title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListTncTitle() (int, []modelTnc.TncTitle, error) {
	repo, err := s.offerRepo.ListTncTitle()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Tnc Title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateTncTitle(idTnc int64, title string) (int, *modelTnc.TncTitle, error) {
	model := &modelTnc.TncTitle{
		IDTnc: idTnc,
		Title: title,
	}
	repo, err := s.offerRepo.CreateTncTitle(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateTncTitleByID(id int64, params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncTitleByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID Tnc Title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteTncTitleByID(id int64) (httpStatus int, err error) {
	status, err := s.offerRepo.DeleteTncTitleByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Tnc title ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Tnc subtitle
func (s service) GetTncSubtitleByID(id int64) (int, *modelTnc.TncSubtitle, error) {
	repo, err := s.offerRepo.GetTncSubtitleByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Tnc Subtitle not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListTncSubtitle() (int, []modelTnc.TncSubtitle, error) {
	repo, err := s.offerRepo.ListTncSubtitle()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Tnc Subtitle not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateTncSubtitle(idTncTitle int64, subtitle string) (int, *modelTnc.TncSubtitle, error) {
	model := &modelTnc.TncSubtitle{
		IDTncTitle: idTncTitle,
		Subtitle:   &subtitle,
	}
	repo, err := s.offerRepo.CreateTncSubtitle(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateTncSubtitleByID(id int64, params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncSubtitleByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Tnc subtitle ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteTncSubtitleByID(id int64) (httpStatus int, err error) {
	status, err := s.offerRepo.DeleteTncSubtitleByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Tnc subtitle ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Tnc explain
func (s service) GetTncExplainByID(id int64) (int, *modelTnc.TncExplain, error) {
	repo, err := s.offerRepo.GetTncExplainByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Tnc Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListTncExplain() (int, []modelTnc.TncExplain, error) {
	repo, err := s.offerRepo.ListTncExplain()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Tnc Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateTncExplain(idTnc int64, idTncTitle int64, idTncSubtitle *int64, description string) (int, *modelTnc.TncExplain, error) {
	if idTncSubtitle == nil || *idTncSubtitle == 0 {
		idTncSubtitle = nil
	}
	model := &modelTnc.TncExplain{
		IDTnc:         idTnc,
		IDTncTitle:    idTncTitle,
		IDTncSubtitle: idTncSubtitle,
		Description:   description,
	}
	repo, err := s.offerRepo.CreateTncExplain(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateTncExplainByID(id int64, params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncExplainByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Tnc explain ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteTncExplainByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteTncExplainByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Tnc explain ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Paylater Offer Page - Landing Page
func (s service) GetLandingPage() (int, *resp.LandingPage, error) {
	httpStatus, repoLoanLimit, err := s.GetLoanLimit()
	if err != nil {
		return httpStatus, nil, err
	}
	respLoanLimit := resp.LoanLimit{
		ID:    repoLoanLimit.ID,
		Limit: repoLoanLimit.Limit,
	}

	respLandingPage := resp.LandingPage{
		LoanLimit: &respLoanLimit,
	}

	httpStatus, repoBenefit, err := s.ListBenefit()
	if err != nil {
		return httpStatus, nil, err
	}
	for _, dataBenefit := range repoBenefit {
		respBenefit := resp.Benefit{
			ID:          dataBenefit.ID,
			Title:       dataBenefit.Title,
			Description: dataBenefit.Description,
			Image:       dataBenefit.Image,
		}
		respLandingPage.Benefit = append(respLandingPage.Benefit, respBenefit)
	}

	httpStatus, repoLoanMethod, err := s.ListLoanMethod()
	if err != nil {
		return httpStatus, nil, err
	}
	for _, dataLoanMethod := range repoLoanMethod {
		respLoanMethod := resp.LoanMethod{
			ID:          dataLoanMethod.ID,
			Title:       dataLoanMethod.Title,
			Description: dataLoanMethod.Description,
		}
		respLandingPage.LoanMethod = append(respLandingPage.LoanMethod, respLoanMethod)
	}

	return http.StatusOK, &respLandingPage, nil
}

//Paylater Offer Page - Tnc Page
func (s service) GetTncPage() (int, []*resp.TncPage, error) {
	httpStatus, repoListTnc, err := s.ListTnc()
	if err != nil {
		return httpStatus, nil, err
	}
	responses := make([]*resp.TncPage, 0)
	for _, dataListTnc := range repoListTnc {
		respPage := resp.TncPage{}
		respPage.MainTitle = dataListTnc.Title

		repoListTncTitle, err := s.offerRepo.ListTncTitleByID(dataListTnc.ID)
		if err != nil {
			return http.StatusNotFound, nil, err
		}
		for _, dataListTncTitle := range repoListTncTitle {
			respTncTitle := resp.TncTitle{}
			respTncTitle.Title = dataListTncTitle.Title

			repoListTncSubtitle, err := s.offerRepo.ListTncSubtitleByID(dataListTncTitle.ID)
			if err != nil {
				return http.StatusNotFound, nil, err
			}

			if repoListTncSubtitle != nil {
				for _, dataListTncSubtitle := range repoListTncSubtitle {
					respTncSubtitle := resp.TncSubtitle{}
					respTncSubtitle.Subtitle = dataListTncSubtitle.Subtitle

					repoListTncExplain, err := s.offerRepo.ListTncExplainByIDWithSubtitle(dataListTnc.ID,
						dataListTncTitle.ID,
						dataListTncSubtitle.ID)
					if err != nil {
						return http.StatusNotFound, nil, err
					}

					respTncExplain := resp.TncExplain{}
					for _, dataListTncExplain := range repoListTncExplain {
						respTncExplain.Explain = dataListTncExplain.Description

						respTncSubtitle.Explain = append(respTncSubtitle.Explain, respTncExplain)
					}
					respTncTitle.Subtitle = append(respTncTitle.Subtitle, respTncSubtitle)

				}
			}

			if len(repoListTncSubtitle) == 0 {
				respTncSubtitle := resp.TncSubtitle{}
				repoListTncExplain, err := s.offerRepo.ListTncExplainByID(dataListTnc.ID, dataListTncTitle.ID)
				if err == sql.ErrNoRows {
					return http.StatusNotFound, nil, errors.New("Any ID Tnc explain not found")
				}
				if err != nil {
					return http.StatusInternalServerError, nil, err
				}

				for _, dataListTncExplain := range repoListTncExplain {
					respTncExplain := resp.TncExplain{}
					respTncExplain.Explain = dataListTncExplain.Description

					respTncSubtitle.Explain = append(respTncSubtitle.Explain, respTncExplain)
				}

				respTncTitle.Subtitle = append(respTncTitle.Subtitle, respTncSubtitle)
			}
			respPage.Title = append(respPage.Title, respTncTitle)

		}
		responses = append(responses, &respPage)

	}
	return http.StatusOK, responses, nil
}

func (s service) GetTncPageMobile() (int, *modelTnc.TncMobile, error) {
	repo, err := s.offerRepo.GetTncMobile()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Tnc Mobile not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
