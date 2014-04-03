package participant

import (
	"fmt"
	"net/http"
	"net/url"
	"socialapi/models"
	"socialapi/workers/api/modules/helpers"

	"github.com/jinzhu/gorm"
)

func List(u *url.URL, h http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	channelId, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		fmt.Println(err)
		return helpers.NewBadRequestResponse(err)
	}

	req := models.NewChannelParticipant()
	req.ChannelId = channelId
	participants, err := req.List()
	if err != nil {
		if err == gorm.RecordNotFound {
			return helpers.NewNotFoundResponse()
		}
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(participants)
}

func Add(u *url.URL, h http.Header, req *models.ChannelParticipant) (int, http.Header, interface{}, error) {
	channelId, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	accountId, err := helpers.GetURIInt64(u, "accountId")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	req.AccountId = accountId
	req.ChannelId = channelId
	req.StatusConstant = models.ChannelParticipant_STATUS_ACTIVE

	if err := req.Create(); err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(req)
}

func Delete(u *url.URL, h http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	channelId, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	accountId, err := helpers.GetURIInt64(u, "accountId")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	req := models.NewChannelParticipant()
	req.AccountId = accountId
	req.ChannelId = channelId

	if err := req.Delete(); err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	// yes it is deleted but not removed completely from our system
	return helpers.NewDeletedResponse()
}
