package service

import (
	"database/sql"
	"errors"
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
func (s service) GetLoanLimitByID(id int64) (*modelLanding.LoanLimit, error) {
	return s.offerRepo.GetLoanLimitByID(id)
}
func (s service) GetLoanLimit() (*modelLanding.LoanLimit, error) {
	return s.offerRepo.GetLoanLimit()
}
func (s service) CreateLoanLimit(limit string) (*modelLanding.LoanLimit, error) {
	limitLoan, _ := strconv.Atoi(limit)
	model := &modelLanding.LoanLimit{
		Limit: float64(limitLoan),
	}
	repo, err := s.offerRepo.CreateLoanLimit(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateLoanLimit(limit string) error {
	_, err := s.offerRepo.UpdateLoanLimit(limit)
	if err == sql.ErrNoRows {
		return errors.New("Loan limit ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteLoanLimit() error {
	err := s.offerRepo.DeleteLoanLimit()
	if err == sql.ErrNoRows {
		return errors.New("Loan limit not found")
	}
	if err != nil {
		return err
	}

	return nil
}

//Benefit
func (s service) GetBenefitByID(id int64) (*modelLanding.Benefit, error) {
	return s.offerRepo.GetBenefitByID(id)
}
func (s service) ListBenefit() ([]modelLanding.Benefit, error) {
	return s.offerRepo.ListBenefit()
}
func (s service) CreateBenefit(title string, description string, path string) (*modelLanding.Benefit, error) {
	model := &modelLanding.Benefit{
		Title:       title,
		Description: description,
		Image:       path,
	}
	repo, err := s.offerRepo.CreateBenefit(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateBenefitByID(id int64, params data.Params, path string) error {
	_, err := s.offerRepo.UpdateBenefitByID(id, params, path)
	if err == sql.ErrNoRows {
		return errors.New("Benefit ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteBenefitByID(id int64) error {
	err := s.offerRepo.DeleteBenefitByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Benefit ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Loan method
func (s service) GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error) {
	return s.offerRepo.GetLoanMethodByID(id)
}
func (s service) ListLoanMethod() ([]modelLanding.LoanMethod, error) {
	return s.offerRepo.ListLoanMethod()
}
func (s service) CreateLoanMethod(title string, description string) (*modelLanding.LoanMethod, error) {
	model := &modelLanding.LoanMethod{
		Title:       title,
		Description: description,
	}
	repo, err := s.offerRepo.CreateLoanMethod(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateLoanMethodByID(id int64, title string, description string) error {
	_, err := s.offerRepo.UpdateLoanMethodByID(id, title, description)
	if err == sql.ErrNoRows {
		return errors.New("Loan method ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteLoanMethodByID(id int64) error {
	err := s.offerRepo.DeleteLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Loan method ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Tnc
func (s service) GetTncByID(id int64) (*modelTnc.Tnc, error) {
	return s.offerRepo.GetTncByID(id)
}
func (s service) ListTnc() ([]modelTnc.Tnc, error) {
	return s.offerRepo.ListTnc()
}
func (s service) CreateTnc(title string) (*modelTnc.Tnc, error) {
	model := &modelTnc.Tnc{
		Title: title,
	}
	repo, err := s.offerRepo.CreateTnc(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateTncByID(id int64, params data.Params) error {
	_, err := s.offerRepo.UpdateTncByID(id, params)
	if err == sql.ErrNoRows {
		return errors.New("Tnc ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteTncByID(id int64) error {
	err := s.offerRepo.DeleteTncByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Tnc ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Tnc title
func (s service) GetTncTitleByID(id int64) (*modelTnc.TncTitle, error) {
	return s.offerRepo.GetTncTitleByID(id)
}
func (s service) ListTncTitle() ([]modelTnc.TncTitle, error) {
	return s.offerRepo.ListTncTitle()
}
func (s service) CreateTncTitle(idTnc int64, title string) (*modelTnc.TncTitle, error) {
	model := &modelTnc.TncTitle{
		IDTnc: idTnc,
		Title: title,
	}
	repo, err := s.offerRepo.CreateTncTitle(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateTncTitleByID(id int64, params data.Params) error {
	_, err := s.offerRepo.UpdateTncTitleByID(id, params)
	if err == sql.ErrNoRows {
		return errors.New("Tnc title ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteTncTitleByID(id int64) error {
	err := s.offerRepo.DeleteTncTitleByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Tnc title ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Tnc subtitle
func (s service) GetTncSubtitleByID(id int64) (*modelTnc.TncSubtitle, error) {
	return s.offerRepo.GetTncSubtitleByID(id)
}
func (s service) ListTncSubtitle() ([]modelTnc.TncSubtitle, error) {
	return s.offerRepo.ListTncSubtitle()
}
func (s service) CreateTncSubtitle(idTncTitle int64, subtitle string) (*modelTnc.TncSubtitle, error) {
	model := &modelTnc.TncSubtitle{
		IDTncTitle: idTncTitle,
		Subtitle:   &subtitle,
	}
	repo, err := s.offerRepo.CreateTncSubtitle(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateTncSubtitleByID(id int64, params data.Params) error {
	_, err := s.offerRepo.UpdateTncSubtitleByID(id, params)
	if err == sql.ErrNoRows {
		return errors.New("Tnc subtitle ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteTncSubtitleByID(id int64) error {
	err := s.offerRepo.DeleteTncSubtitleByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Tnc subtitle ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Tnc explain
func (s service) GetTncExplainByID(id int64) (*modelTnc.TncExplain, error) {
	return s.offerRepo.GetTncExplainByID(id)
}
func (s service) ListTncExplain() ([]modelTnc.TncExplain, error) {
	return s.offerRepo.ListTncExplain()
}
func (s service) CreateTncExplain(idTnc int64, idTncTitle int64, idTncSubtitle *int64, description string) (*modelTnc.TncExplain, error) {
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
		return nil, err
	}

	return repo, nil
}
func (s service) UpdateTncExplainByID(id int64, params data.Params) error {
	_, err := s.offerRepo.UpdateTncExplainByID(id, params)
	if err == sql.ErrNoRows {
		return errors.New("Tnc explain ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}
func (s service) DeleteTncExplainByID(id int64) error {
	err := s.offerRepo.DeleteTncExplainByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Tnc explain ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Paylater Offer Page - Landing Page
func (s service) GetLandingPage() (*resp.LandingPage, error) {
	repoLoanLimit, err := s.offerRepo.GetLoanLimit()
	if err != nil {
		return nil, err
	}
	respLoanLimit := resp.LoanLimit{
		ID:    repoLoanLimit.ID,
		Limit: repoLoanLimit.Limit,
	}

	respLandingPage := resp.LandingPage{
		LoanLimit: &respLoanLimit,
	}

	repoBenefit, err := s.offerRepo.ListBenefit()
	if err != nil {
		return nil, err
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

	repoLoanMethod, err := s.offerRepo.ListLoanMethod()
	if err != nil {
		return nil, err
	}
	for _, dataLoanMethod := range repoLoanMethod {
		respLoanMethod := resp.LoanMethod{
			ID:          dataLoanMethod.ID,
			Title:       dataLoanMethod.Title,
			Description: dataLoanMethod.Description,
		}
		respLandingPage.LoanMethod = append(respLandingPage.LoanMethod, respLoanMethod)
	}

	return &respLandingPage, nil
}

//Paylater Offer Page - Tnc Page
func (s service) GetTncPage() ([]*resp.TncPage, error) {
	repoListTnc, err := s.offerRepo.ListTnc()
	if err != nil {
		return nil, err
	}
	responses := make([]*resp.TncPage, 0)
	for _, dataListTnc := range repoListTnc {
		respPage := resp.TncPage{}
		respPage.MainTitle = dataListTnc.Title

		repoListTncTitle, err := s.offerRepo.ListTncTitleByID(dataListTnc.ID)
		if err != nil {
			return nil, err
		}
		for _, dataListTncTitle := range repoListTncTitle {
			respTncTitle := resp.TncTitle{}
			respTncTitle.Title = dataListTncTitle.Title

			repoListTncSubtitle, _ := s.offerRepo.ListTncSubtitleByID(dataListTncTitle.ID)

			if repoListTncSubtitle != nil {
				for _, dataListTncSubtitle := range repoListTncSubtitle {
					respTncSubtitle := resp.TncSubtitle{}
					respTncSubtitle.Subtitle = dataListTncSubtitle.Subtitle

					repoListTncExplain, err := s.offerRepo.ListTncExplainByIDWithSubtitle(dataListTnc.ID,
						dataListTncTitle.ID,
						dataListTncSubtitle.ID)
					if err != nil {
						return nil, err
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
				if err != nil {
					return nil, err
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
	return responses, nil
}
