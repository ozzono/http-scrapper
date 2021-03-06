// Copyright 2015 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package scraper implements a simple HTTP automation utility.


Motivation

The web today is widelly used by several companies,
but not all web sites are implemented equal.
Besides the variety of technologies that can be used,
some apps simply does not offer a clean, REST or SOAP like
interface of interaction.

Scraper is a simple tool that allows developers to interact with
legacy websites and webservices, that are not yet in the new age.
It is a statefull HTTP client, meaning that it uses a Cookie Jar
implementation, allowing you to automate tasks that would usually
be done manually.


Getting Started

Scraper is designed to allow you to embedd it in your own automation procedures.
One primary use case is to allow you to submit an authentication request
and have the session data reused in the next requests to the same target.
Scraper acomplished this with the use of a Cookie Jar.

Scraper uses GoQuery internally, and exposes some fun entry points
to extract data from the resulting response.
Here is a sample snippet that logins into a service,
and updates the user preferences by resubmitting a form.

	b := scraper.New().BaseURL("http://website")
	page, err := b.POST("/login", url.Values{ "username": {"myuser"}, "password": {"not-a-secret"}})
	if err != nil {
		// If the request is not in the 2xx range, it is an error
		log.Fatal(err)
	}
	for _, f := range page.Forms() {
		if f.ID == "preferences" {
			f.Fields["mailings"] = []string{""}
			if _, err := b.POST("/account/", f.Fields) {
				log.Fatal(err)
			}
		}
	}

Scraper does not provide any JavaScript inmplementation, nor a Rendering engine.
It is just a headless, statefull HTTP client.
*/
package scraper
