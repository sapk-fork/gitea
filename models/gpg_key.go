// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"

	"code.gitea.io/gitea/modules/log"
	"github.com/go-xorm/xorm"
)

// GPGKey represents a GPG key.
type GPGKey struct {
	ID                int64     `xorm:"pk autoincr"`
	OwnerID           int64     `xorm:"INDEX NOT NULL"`
	KeyID             string    `xorm:"INDEX TEXT NOT NULL"`
	PrimaryKeyID      string    `xorm:"TEXT"`
	Content           string    `xorm:"TEXT NOT NULL"`
	Created           time.Time `xorm:"-"`
	CreatedUnix       int64
	Expired           time.Time `xorm:"-"`
	ExpiredUnix       int64
	Added             time.Time `xorm:"-"`
	AddedUnix         int64
	SubsKey           []*GPGKey `xorm:"-"`
	Emails            []*EmailAddress
	CanSign           bool
	CanEncryptComms   bool
	CanEncryptStorage bool
	CanCertify        bool
}

// BeforeInsert will be invoked by XORM before inserting a record
func (key *GPGKey) BeforeInsert() {
	key.AddedUnix = time.Now().Unix()
	key.ExpiredUnix = key.Expired.Unix()
	key.CreatedUnix = key.Created.Unix()
}

// AfterInsert will be invoked by XORM after inserting a record
func (key *GPGKey) AfterInsert() {
	log.Debug("AfterInsert Subkeys: %v", key.SubsKey)
	sess := x.NewSession()
	defer sessionRelease(sess)
	sess.Begin()
	for _, subkey := range key.SubsKey {
		if err := addGPGKey(sess, subkey); err != nil {
			log.Warn("Failed to add subKey: [err:%v, subkey:%v]", err, subkey)
		}
	}
	sess.Commit()
}

// AfterSet is invoked from XORM after setting the value of a field of this object.
func (key *GPGKey) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "key_id":
		x.Where("primary_key_id=?", key.KeyID).Find(&key.SubsKey)
	case "added_unix":
		key.Added = time.Unix(key.AddedUnix, 0).Local()
	case "expired_unix":
		key.Expired = time.Unix(key.ExpiredUnix, 0).Local()
	case "created_unix":
		key.Created = time.Unix(key.CreatedUnix, 0).Local()
	}
}

// ListGPGKeys returns a list of public keys belongs to given user.
func ListGPGKeys(uid int64) ([]*GPGKey, error) {
	keys := make([]*GPGKey, 0, 5)
	return keys, x.Where("owner_id=? AND primary_key_id=''", uid).Find(&keys)
}

// GetGPGKeyByID returns public key by given ID.
func GetGPGKeyByID(keyID int64) (*GPGKey, error) {
	key := new(GPGKey)
	has, err := x.Id(keyID).Get(key)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrGPGKeyNotExist{keyID}
	}
	return key, nil
}

// checkArmoredGPGKeyString checks if the given key string is a valid GPG armored key.
// The function returns the actual public keyon success
func checkArmoredGPGKeyString(content string) (*openpgp.Entity, error) {
	list, err := openpgp.ReadArmoredKeyRing(strings.NewReader(content))
	log.Trace("GPG Err : %v", err)
	log.Trace("GPG List : %v", list)
	if err != nil {
		return nil, err
	}
	return list[0], nil
}

func addGPGKey(e Engine, key *GPGKey) (err error) {
	// Save GPG key.
	if _, err = e.Insert(key); err != nil {
		return err
	}
	return nil
}

// AddGPGKey adds new public key to database.
func AddGPGKey(ownerID int64, content string) (*GPGKey, error) {
	log.Trace(content)
	ekey, err := checkArmoredGPGKeyString(content)
	if err != nil {
		return nil, err
	}

	// Key ID cannot be duplicated.
	has, err := x.Where("key_id=?", ekey.PrimaryKey.KeyIdString()).
		Get(new(GPGKey))
	if err != nil {
		return nil, err
	} else if has {
		return nil, ErrGPGKeyIDAlreadyUsed{ekey.PrimaryKey.KeyIdString()}
	}

	//Get DB session
	sess := x.NewSession()
	defer sessionRelease(sess)
	if err = sess.Begin(); err != nil {
		return nil, err
	}

	key, err := parseGPGKey(ownerID, ekey)
	if err != nil {
		return nil, err
	}

	if err = addGPGKey(sess, key); err != nil {
		return nil, fmt.Errorf("addKey: %v", err)
	}

	return key, sess.Commit()
}

func parseSubGPGKey(ownerID int64, primaryID string, pubkey *packet.PublicKey) *GPGKey {
	content := new(bytes.Buffer)
	b64 := base64.NewEncoder(base64.StdEncoding, content)
	if err := pubkey.Serialize(b64); err != nil {
		log.Warn("Failed to serialize public key content: %v", pubkey.Fingerprint)
	}
	return &GPGKey{
		OwnerID:           ownerID,
		KeyID:             pubkey.KeyIdString(),
		PrimaryKeyID:      primaryID,
		Content:           content.String(),
		Created:           pubkey.CreationTime,
		CanSign:           pubkey.CanSign(),
		CanEncryptComms:   pubkey.PubKeyAlgo.CanEncrypt(),
		CanEncryptStorage: pubkey.PubKeyAlgo.CanEncrypt(),
		CanCertify:        pubkey.PubKeyAlgo.CanSign(),
	}
}
func parseGPGKey(ownerID int64, e *openpgp.Entity) (*GPGKey, error) {
	pubkey := e.PrimaryKey
	content := new(bytes.Buffer)
	b64 := base64.NewEncoder(base64.StdEncoding, content)
	if err := pubkey.Serialize(b64); err != nil {
		log.Warn("Failed to serialize public key content: %v", pubkey.Fingerprint)
	}
	subkeys := make([]*GPGKey, len(e.Subkeys))
	for i, k := range e.Subkeys {
		subkeys[i] = parseSubGPGKey(ownerID, pubkey.KeyIdString(), k.PublicKey)
	}

	//Check email
	userEmails, err := GetEmailAddresses(ownerID)
	if err != nil {
		return nil, err
	}
	emails := make([]*EmailAddress, len(e.Identities))
	n := 0
	for _, ident := range e.Identities {
		for _, e := range userEmails {
			if e.Email == ident.UserId.Email && e.IsActivated {
				emails[n] = e
				break
			}
		}
		if emails[n] == nil {
			return nil, fmt.Errorf("Failed to found email or is not confirmed : %s", ident.UserId.Email)
		}
		n++
	}

	log.Debug("Subkeys: %v", subkeys)
	return &GPGKey{
		OwnerID:           ownerID,
		KeyID:             pubkey.KeyIdString(),
		PrimaryKeyID:      "",
		Content:           content.String(),
		Created:           pubkey.CreationTime,
		Expired:           time.Time{},
		Emails:            emails,
		SubsKey:           subkeys,
		CanSign:           pubkey.CanSign(),
		CanEncryptComms:   pubkey.PubKeyAlgo.CanEncrypt(),
		CanEncryptStorage: pubkey.PubKeyAlgo.CanEncrypt(),
		CanCertify:        pubkey.PubKeyAlgo.CanSign(),
	}, nil
}

// deleteGPGKey does the actual key deletion
func deleteGPGKey(e *xorm.Session, keyIDs ...int64) error {

	log.Debug("deleteGPGKey: %v", keyIDs)
	if len(keyIDs) == 0 {
		return nil
	}

	_, err := e.In("id", keyIDs).Delete(new(GPGKey))
	return err
}

// DeleteGPGKey deletes GPG key information in database.
func DeleteGPGKey(doer *User, id int64) (err error) {
	key, err := GetGPGKeyByID(id)
	if err != nil {
		if IsErrGPGKeyNotExist(err) {
			return nil
		}
		return fmt.Errorf("GetPublicKeyByID: %v", err)
	}

	// Check if user has access to delete this key.
	if !doer.IsAdmin && doer.ID != key.OwnerID {
		return ErrGPGKeyAccessDenied{doer.ID, key.ID}
	}

	sess := x.NewSession()
	defer sessionRelease(sess)
	if err = sess.Begin(); err != nil {
		return err
	}

	//Add subkeys to remove
	subkeys := make([]*GPGKey, 0, 5)
	x.Where("primary_key_id=?", key.KeyID).Find(&subkeys)
	ids := make([]int64, len(subkeys)+1)
	for i, sk := range subkeys {
		ids[i] = sk.ID
	}

	//Add primary key to remove at last
	ids[len(subkeys)] = id

	if err = deleteGPGKey(sess, ids...); err != nil {
		return err
	}

	if err = sess.Commit(); err != nil {
		return err
	}

	return nil
}
