// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/openpgp"

	"github.com/Unknwon/com"

	api "code.gitea.io/sdk/gitea"

	"code.gitea.io/git"
	"code.gitea.io/gitea/models"
)

// ToEmail convert models.EmailAddress to api.Email
func ToEmail(email *models.EmailAddress) *api.Email {
	return &api.Email{
		Email:    email.Email,
		Verified: email.IsActivated,
		Primary:  email.IsPrimary,
	}
}

// ToBranch convert a commit and branch to an api.Branch
func ToBranch(b *models.Branch, c *git.Commit) *api.Branch {
	return &api.Branch{
		Name:   b.Name,
		Commit: ToCommit(c),
	}
}

// ToCommit convert a commit to api.PayloadCommit
func ToCommit(c *git.Commit) *api.PayloadCommit {
	authorUsername := ""
	author, err := models.GetUserByEmail(c.Author.Email)
	if err == nil {
		authorUsername = author.Name
	}
	committerUsername := ""
	committer, err := models.GetUserByEmail(c.Committer.Email)
	if err == nil {
		committerUsername = committer.Name
	}
	return &api.PayloadCommit{
		ID:      c.ID.String(),
		Message: c.Message(),
		URL:     "Not implemented",
		Author: &api.PayloadUser{
			Name:     c.Author.Name,
			Email:    c.Author.Email,
			UserName: authorUsername,
		},
		Committer: &api.PayloadUser{
			Name:     c.Committer.Name,
			Email:    c.Committer.Email,
			UserName: committerUsername,
		},
		Timestamp: c.Author.When,
	}
}

// ToPublicKey convert models.PublicKey to api.PublicKey
func ToPublicKey(apiLink string, key *models.PublicKey) *api.PublicKey {
	return &api.PublicKey{
		ID:      key.ID,
		Key:     key.Content,
		URL:     apiLink + com.ToStr(key.ID),
		Title:   key.Name,
		Created: key.Created,
	}
}

// ToGPGKey converts models.PublicGPGKey to api.GPGKey
func ToGPGKey(key *models.GPGKey) *api.GPGKey {
	keyList, _ := openpgp.ReadArmoredKeyRing(strings.NewReader(key.Content))
	pkey := keyList[0].PrimaryKey
	//Generate subkeys array
	subkeys := make([]*api.GPGKey, len(keyList[0].Subkeys))
	for id, k := range keyList[0].Subkeys {
		subkeys[id] = &api.GPGKey{
			ID:           key.ID, // OR int64(id) ?
			PrimaryKeyID: key.KeyID,
			KeyID:        k.PublicKey.KeyIdString(),
			//PublicKey:         key.Content, //TODO replace with pkey.PublicKey.Serialize
			Created: k.PublicKey.CreationTime,
			Expires: time.Time{}, //TODO expire keyList[0].PrimaryKey.PublicKey.(packet.PublicKeyV3).DaysToExpire //TODO expire
			//Emails:            emails,
			//SubsKey:           subkeys,
			CanSign:           k.PublicKey.CanSign(),
			CanEncryptComms:   k.PublicKey.PubKeyAlgo.CanEncrypt(),
			CanEncryptStorage: k.PublicKey.PubKeyAlgo.CanEncrypt(),
			CanCertify:        k.PublicKey.PubKeyAlgo.CanSign(),
		}
	}
	//Generate emails array
	emails := make([]*api.GPGKeyEmail, len(keyList[0].Identities))
	id := 0
	var validIDNName = regexp.MustCompile("^.+ <([A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,64})>$") //"Full Name (comment) <email@example.com>"
	//for name, identity := range keyList[0].Identities {
	for name := range keyList[0].Identities {
		match := validIDNName.FindAllStringSubmatch(name, -1)
		emails[id] = &api.GPGKeyEmail{
			Email:    match[0][len(match[0])-1],
			Verified: false,
		}
		id++
	}
	return &api.GPGKey{
		ID:                key.ID,
		PrimaryKeyID:      "",
		KeyID:             key.KeyID,
		PublicKey:         key.Content, //TODO replace with pkey.PublicKey.Serialize
		Created:           key.Created,
		Expires:           time.Time{}, //TODO expire keyList[0].PrimaryKey.PublicKey.(packet.PublicKeyV3).DaysToExpire //TODO expire
		Emails:            emails,
		SubsKey:           subkeys,
		CanSign:           pkey.CanSign(),
		CanEncryptComms:   pkey.PubKeyAlgo.CanEncrypt(),
		CanEncryptStorage: pkey.PubKeyAlgo.CanEncrypt(),
		CanCertify:        pkey.PubKeyAlgo.CanSign(),
	}
}

// ToHook convert models.Webhook to api.Hook
func ToHook(repoLink string, w *models.Webhook) *api.Hook {
	config := map[string]string{
		"url":          w.URL,
		"content_type": w.ContentType.Name(),
	}
	if w.HookTaskType == models.SLACK {
		s := w.GetSlackHook()
		config["channel"] = s.Channel
		config["username"] = s.Username
		config["icon_url"] = s.IconURL
		config["color"] = s.Color
	}

	return &api.Hook{
		ID:      w.ID,
		Type:    w.HookTaskType.Name(),
		URL:     fmt.Sprintf("%s/settings/hooks/%d", repoLink, w.ID),
		Active:  w.IsActive,
		Config:  config,
		Events:  w.EventsArray(),
		Updated: w.Updated,
		Created: w.Created,
	}
}

// ToDeployKey convert models.DeployKey to api.DeployKey
func ToDeployKey(apiLink string, key *models.DeployKey) *api.DeployKey {
	return &api.DeployKey{
		ID:       key.ID,
		Key:      key.Content,
		URL:      apiLink + com.ToStr(key.ID),
		Title:    key.Name,
		Created:  key.Created,
		ReadOnly: true, // All deploy keys are read-only.
	}
}

// ToOrganization convert models.User to api.Organization
func ToOrganization(org *models.User) *api.Organization {
	return &api.Organization{
		ID:          org.ID,
		AvatarURL:   org.AvatarLink(),
		UserName:    org.Name,
		FullName:    org.FullName,
		Description: org.Description,
		Website:     org.Website,
		Location:    org.Location,
	}
}

// ToTeam convert models.Team to api.Team
func ToTeam(team *models.Team) *api.Team {
	return &api.Team{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		Permission:  team.Authorize.String(),
	}
}
