// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/asteris-llc/converge/openpgp"
	"github.com/spf13/pflag"
)

// These are both URLs
var pubkey string
var signature string

func addPGPArguments(flags *pflag.FlagSet) {
	flags.StringVar(&signature, "signature", "signature.asc", "URL of (armored or unarmored) openpgp module signature")
	flags.StringVar(&pubkey, "pubkey", "pubkey.asc", "URL of public key of the user who signed the module")
}

// verify the module's PGP signature, failing and exiting if it isn't valid
func verifyModuleSignature(module string) {
	moduleURL, err := url.Parse(module)
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't parse module's URL")
	}

	pubkeyURL, err := url.Parse(pubkey)
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't parse public key's URL")
	}

	signatureURL, err := url.Parse(pubkey)
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't parse signature's URL")
	}

	// if neither were specified, we can't validate
	if pubkey == "" {
		logrus.Fatal("Signer's public key wasn't specified")
	} else if signature == "" {
		logrus.Fatal("Signature wasn't specified")
	}

	if err := openpgp.Verify(moduleURL, signatureURL, pubkeyURL); err != nil {
		logrus.WithError(err).Fatal("Couldn't verify module's PGP signature")
	}
}
