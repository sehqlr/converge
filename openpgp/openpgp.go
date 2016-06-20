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

package openpgp

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/asteris-llc/converge/util"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// Verify fetches the data, signature, and public key from their respective URLs,
// verifies the data, and returns nil if everything went well.
func Verify(dataURL, sigURL, pubkeyURL *url.URL) error {
	dataRdr, err := util.Retrieve(dataURL)
	if err != nil {
		return err
	}
	defer dataRdr.Close()

	data, err := ioutil.ReadAll(dataRdr)
	if err != nil {
		return err
	}

	fmt.Println("sigpkt")
	sigPkt, err := sigFromURL(sigURL)
	if err != nil {
		return err
	}

	fmt.Println("pubkeypkt")
	pubkeyPkt, err := pubkeyFromURL(pubkeyURL)
	if err != nil {
		return err
	}

	fmt.Println("hash")
	hash := sigPkt.Hash.New()
	if _, err := hash.Write(data); err != nil {
		return err
	}

	return pubkeyPkt.VerifySignature(hash, sigPkt)
}

func pktFromURL(u *url.URL) (packet.Packet, error) {
	fmt.Println("from URL")
	rdr, err := util.Retrieve(u)
	if err != nil {
		return nil, err
	}
	defer rdr.Close()

	// handle armored data
	if block, err := armor.Decode(rdr); err == nil {
		fmt.Println("armored")
		reader := packet.NewReader(block.Body)
		entity, err := openpgp.ReadEntity(reader)
		if err != nil {
			return nil, err
		}
		return entity.PrimaryKey, nil
	}
	content, err := ioutil.ReadAll(rdr)
	fmt.Println(string(content))
	fmt.Println(u.String())
	return packet.Read(rdr)
}

func sigFromURL(sigURL *url.URL) (*packet.Signature, error) {
	pkt, err := pktFromURL(sigURL)
	if err != nil {
		return nil, err
	}

	// is the packet really a PGP signature?
	sigPkt, ok := pkt.(*packet.Signature)
	if !ok {
		return nil, fmt.Errorf("Not a valid signature: %v", sigURL.String())
	}

	return sigPkt, nil
}

func pubkeyFromURL(pubkeyURL *url.URL) (*packet.PublicKey, error) {
	pkt, err := pktFromURL(pubkeyURL)
	if err != nil {
		return nil, err
	}

	// is the packet really a PGP public key?
	pubkeyPkt, ok := pkt.(*packet.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Not a valid public key: %v", pubkeyURL.String())
	}

	return pubkeyPkt, nil
}
