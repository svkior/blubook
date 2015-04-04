package main

import "testing"

func TestAuthAvatar(t *testing.T) {

	var authAvatar AuthAvatar
	client := new(client)

	_, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no avlue present")
	}

	testUrl := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}

	url, err := authAvatar.GetAvatarURL(client)

	if err != nil {
		t.Error("AuthAvatar.GetAvaterURL should return correct URL")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURL should return correct URL")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {

	var gravatarAvitar GravatarAvatar

	client := new(client)
	client.userData = map[string]interface{}{"email": "svkior@gmail.com"}
	url, err := gravatarAvitar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvitar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/e9945c78f9c0e996b5a2af041b7b043d" {
		t.Errorf("GravatarAvitar.GetAvatarURL wrongly returned %s", url)
	}
}
