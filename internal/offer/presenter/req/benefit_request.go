package req

import "bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data/filedata"

type UpdateBenefitByIDRequest struct {
	Image *filedata.UploadFile
}
