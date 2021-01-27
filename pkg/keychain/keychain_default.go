// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

// Package keychain implements a native secure password store for each platform.
package keychain

import (
	"github.com/docker/docker-credential-helpers/credentials"
)

// NewMissingKeychain returns a new keychain that always returns an error.
func NewMissingKeychain() *Keychain {
	return newKeychain(&missingHelper{}, "")
}

// missingHelper is a helper which is used when no other helper is available.
// It always returns ErrNoKeychain.
type missingHelper struct{}

func (h *missingHelper) Add(*credentials.Credentials) error {
	return ErrNoKeychain
}

func (h *missingHelper) Delete(string) error {
	return ErrNoKeychain
}

func (h *missingHelper) Get(string) (string, string, error) {
	return "", "", ErrNoKeychain
}

func (h *missingHelper) List() (map[string]string, error) {
	return nil, ErrNoKeychain
}
