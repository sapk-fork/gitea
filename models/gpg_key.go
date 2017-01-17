// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/openpgp"

	"github.com/go-xorm/xorm"
	"github.com/gogits/gogs/modules/log"
)

// GPGKey represents a GPG key.
type GPGKey struct {
	ID      int64  `xorm:"pk autoincr"`
	OwnerID int64  `xorm:"INDEX NOT NULL"`
	KeyID   string `xorm:"INDEX TEXT NOT NULL"`
	//PrimaryKeyID string    `xorm:"TEXT"`
	Content     string    `xorm:"TEXT NOT NULL"`
	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Added       time.Time `xorm:"-"`
	AddedUnix   int64
}

// BeforeInsert will be invoked by XORM before inserting a record
func (key *GPGKey) BeforeInsert() {
	key.AddedUnix = time.Now().Unix()
	key.CreatedUnix = key.Created.Unix()
}

// AfterSet is invoked from XORM after setting the value of a field of this object.
func (key *GPGKey) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "added_unix":
		key.Added = time.Unix(key.CreatedUnix, 0).Local()
	case "created_unix":
		key.Created = time.Unix(key.CreatedUnix, 0).Local()
	}
}

// ListGPGKeys returns a list of public keys belongs to given user.
func ListGPGKeys(uid int64) ([]*GPGKey, error) {
	keys := make([]*GPGKey, 0, 5)
	return keys, x.Where("owner_id=?", uid).Find(&keys)
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
func checkArmoredGPGKeyString(content string) (*openpgp.Entity, error) { //TODO test
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

	key := &GPGKey{
		OwnerID: ownerID,
		KeyID:   ekey.PrimaryKey.KeyIdString(),
		//PrimaryKeyID: "",
		Content: content,
		Created: ekey.PrimaryKey.CreationTime,
	}
	/*
		for k, v := range ekey.Subkeys {
			eskey := v.PublicKey
			skey := &GPGKey{
				OwnerID:      ownerID,
				KeyID:        eskey.KeyIdString(),
				PrimaryKeyID: key.KeyID,
				Content:      eskey.PublicKey,
				CreatedUnix:  eskey.CreationTime,
			}
		}
		//TODO add recursively subkeys ?
	*/

	if err = addGPGKey(sess, key); err != nil {
		return nil, fmt.Errorf("addKey: %v", err)
	}

	return key, sess.Commit()
}

// DeleteGPGKey deletes GPG key information in database.
func DeleteGPGKey(doer *User, id int64) (err error) {
	//TODO Implement
	return nil
}

/*  TODO
// CheckCommitWithSign checks if author's signature of commit is corresponsind to a user.
func CheckCommitWithSign(c *git.Commit) *User {
	u, err := GetUserByEmail(c.Author.Email)
	if err != nil {
		return nil
	}
	ks, err := ListPublicGPGKeys(u.ID)
	if err != nil {
		return nil
	}
	return u
}

// CheckCommitsWithSign checks if author's signature of commits are corresponding to users.
func CheckCommitsWithSign(oldCommits *list.List) *list.List {
	var (
		u          *User
		emails     = map[string]*User{}
		newCommits = list.New()
		e          = oldCommits.Front()
	)
	for e != nil {
		c := e.Value.(*git.Commit)

		if v, ok := emails[c.Author.Email]; !ok {
			u, _ = GetUserByEmail(c.Author.Email)
			emails[c.Author.Email] = u
		} else {
			u = v
		}

		newCommits.PushBack(UserCommit{
			User:   u,
			Commit: c,
		})
		e = e.Next()
	}
	return newCommits
}
*/
