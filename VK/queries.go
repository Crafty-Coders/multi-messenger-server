package VK

import (
	"multi-messenger-server/config"
	"multi-messenger-server/tools"
)

func GetVKAuthData(code string) (map[string]interface{}, error) {
	urlTemplate := `https://oauth.vk.com/access_token?client_id={{.client_id}}&client_secret={{.client_secret}}&redirect_uri={{.redirect_uri}}&signin&code={{.code}}&v=5.131`
	data := map[string]interface{}{
		"client_id":     config.VkConfig.ClientId,
		"client_secret": config.VkConfig.ClientSecret,
		"redirect_uri":  config.AppConfig.FrontUrl + "/",
		"code":          code,
	}

	body, err := tools.CreateGETQueryFromTemplate(urlTemplate, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getDataFromVkResponse(response map[string]interface{}) map[string]interface{} {
	return response["response"].([]interface{})[0].(map[string]interface{})
}

func GetVkUser(access_token string, user_id int) (map[string]interface{}, error) {

	urlTemplate := `https://api.vk.com/method/users.get?user_ids={{.user_id}}&fields=photo_400,has_mobile,home_town,contacts,mobile_phone&access_token={{.access_token}}&v=5.131`
	data := map[string]interface{}{
		"access_token": access_token,
		"user_id":      user_id,
	}

	body, err := tools.CreateGETQueryFromTemplate(urlTemplate, data)
	if err != nil {
		return nil, err
	}
	body = getDataFromVkResponse(body)
	return body, nil
}
