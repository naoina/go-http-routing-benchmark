// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package benchmark

import (
	"net/http"
	"testing"
)

// http://developer.github.com/v3/
var githubAPI = []route{
	// OAuth Authorizations
	{"GET", "/authorizations"},
	{"GET", "/authorizations/:id"},
	{"POST", "/authorizations"},
	//{"PUT", "/authorizations/clients/:client_id"},
	//{"PATCH", "/authorizations/:id"},
	{"DELETE", "/authorizations/:id"},
	{"GET", "/applications/:client_id/tokens/:access_token"},
	{"DELETE", "/applications/:client_id/tokens"},
	{"DELETE", "/applications/:client_id/tokens/:access_token"},

	// Activity
	{"GET", "/events"},
	{"GET", "/repos/:owner/:repo/events"},
	{"GET", "/networks/:owner/:repo/events"},
	{"GET", "/orgs/:org/events"},
	{"GET", "/users/:user/received_events"},
	{"GET", "/users/:user/received_events/public"},
	{"GET", "/users/:user/events"},
	{"GET", "/users/:user/events/public"},
	{"GET", "/users/:user/events/orgs/:org"},
	{"GET", "/feeds"},
	{"GET", "/notifications"},
	{"GET", "/repos/:owner/:repo/notifications"},
	{"PUT", "/notifications"},
	{"PUT", "/repos/:owner/:repo/notifications"},
	{"GET", "/notifications/threads/:id"},
	//{"PATCH", "/notifications/threads/:id"},
	{"GET", "/notifications/threads/:id/subscription"},
	{"PUT", "/notifications/threads/:id/subscription"},
	{"DELETE", "/notifications/threads/:id/subscription"},
	{"GET", "/repos/:owner/:repo/stargazers"},
	{"GET", "/users/:user/starred"},
	{"GET", "/user/starred"},
	{"GET", "/user/starred/:owner/:repo"},
	{"PUT", "/user/starred/:owner/:repo"},
	{"DELETE", "/user/starred/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/subscribers"},
	{"GET", "/users/:user/subscriptions"},
	{"GET", "/user/subscriptions"},
	{"GET", "/repos/:owner/:repo/subscription"},
	{"PUT", "/repos/:owner/:repo/subscription"},
	{"DELETE", "/repos/:owner/:repo/subscription"},
	{"GET", "/user/subscriptions/:owner/:repo"},
	{"PUT", "/user/subscriptions/:owner/:repo"},
	{"DELETE", "/user/subscriptions/:owner/:repo"},

	// Gists
	{"GET", "/users/:user/gists"},
	{"GET", "/gists"},
	//{"GET", "/gists/public"},
	//{"GET", "/gists/starred"},
	{"GET", "/gists/:id"},
	{"POST", "/gists"},
	//{"PATCH", "/gists/:id"},
	{"PUT", "/gists/:id/star"},
	{"DELETE", "/gists/:id/star"},
	{"GET", "/gists/:id/star"},
	{"POST", "/gists/:id/forks"},
	{"DELETE", "/gists/:id"},

	// Git Data
	{"GET", "/repos/:owner/:repo/git/blobs/:sha"},
	{"POST", "/repos/:owner/:repo/git/blobs"},
	{"GET", "/repos/:owner/:repo/git/commits/:sha"},
	{"POST", "/repos/:owner/:repo/git/commits"},
	//{"GET", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/refs"},
	{"POST", "/repos/:owner/:repo/git/refs"},
	//{"PATCH", "/repos/:owner/:repo/git/refs/*ref"},
	//{"DELETE", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/tags/:sha"},
	{"POST", "/repos/:owner/:repo/git/tags"},
	{"GET", "/repos/:owner/:repo/git/trees/:sha"},
	{"POST", "/repos/:owner/:repo/git/trees"},

	// Issues
	{"GET", "/issues"},
	{"GET", "/user/issues"},
	{"GET", "/orgs/:org/issues"},
	{"GET", "/repos/:owner/:repo/issues"},
	{"GET", "/repos/:owner/:repo/issues/:number"},
	{"POST", "/repos/:owner/:repo/issues"},
	//{"PATCH", "/repos/:owner/:repo/issues/:number"},
	{"GET", "/repos/:owner/:repo/assignees"},
	{"GET", "/repos/:owner/:repo/assignees/:assignee"},
	{"GET", "/repos/:owner/:repo/issues/:number/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments/:id"},
	{"POST", "/repos/:owner/:repo/issues/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
	//{"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
	{"GET", "/repos/:owner/:repo/issues/:number/events"},
	//{"GET", "/repos/:owner/:repo/issues/events"},
	//{"GET", "/repos/:owner/:repo/issues/events/:id"},
	{"GET", "/repos/:owner/:repo/labels"},
	{"GET", "/repos/:owner/:repo/labels/:name"},
	{"POST", "/repos/:owner/:repo/labels"},
	//{"PATCH", "/repos/:owner/:repo/labels/:name"},
	{"DELETE", "/repos/:owner/:repo/labels/:name"},
	{"GET", "/repos/:owner/:repo/issues/:number/labels"},
	{"POST", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name"},
	{"PUT", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones"},
	{"GET", "/repos/:owner/:repo/milestones/:number"},
	{"POST", "/repos/:owner/:repo/milestones"},
	//{"PATCH", "/repos/:owner/:repo/milestones/:number"},
	{"DELETE", "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{"GET", "/emojis"},
	{"GET", "/gitignore/templates"},
	{"GET", "/gitignore/templates/:name"},
	{"POST", "/markdown"},
	{"POST", "/markdown/raw"},
	{"GET", "/meta"},
	{"GET", "/rate_limit"},

	// Organizations
	{"GET", "/users/:user/orgs"},
	{"GET", "/user/orgs"},
	{"GET", "/orgs/:org"},
	//{"PATCH", "/orgs/:org"},
	{"GET", "/orgs/:org/members"},
	{"GET", "/orgs/:org/members/:user"},
	{"DELETE", "/orgs/:org/members/:user"},
	{"GET", "/orgs/:org/public_members"},
	{"GET", "/orgs/:org/public_members/:user"},
	{"PUT", "/orgs/:org/public_members/:user"},
	{"DELETE", "/orgs/:org/public_members/:user"},
	{"GET", "/orgs/:org/teams"},
	{"GET", "/teams/:id"},
	{"POST", "/orgs/:org/teams"},
	//{"PATCH", "/teams/:id"},
	{"DELETE", "/teams/:id"},
	{"GET", "/teams/:id/members"},
	{"GET", "/teams/:id/members/:user"},
	{"PUT", "/teams/:id/members/:user"},
	{"DELETE", "/teams/:id/members/:user"},
	{"GET", "/teams/:id/repos"},
	{"GET", "/teams/:id/repos/:owner/:repo"},
	{"PUT", "/teams/:id/repos/:owner/:repo"},
	{"DELETE", "/teams/:id/repos/:owner/:repo"},
	{"GET", "/user/teams"},

	// Pull Requests
	{"GET", "/repos/:owner/:repo/pulls"},
	{"GET", "/repos/:owner/:repo/pulls/:number"},
	{"POST", "/repos/:owner/:repo/pulls"},
	//{"PATCH", "/repos/:owner/:repo/pulls/:number"},
	{"GET", "/repos/:owner/:repo/pulls/:number/commits"},
	{"GET", "/repos/:owner/:repo/pulls/:number/files"},
	{"GET", "/repos/:owner/:repo/pulls/:number/merge"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/merge"},
	{"GET", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments/:number"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},
	//{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},

	// Repositories
	{"GET", "/user/repos"},
	{"GET", "/users/:user/repos"},
	{"GET", "/orgs/:org/repos"},
	{"GET", "/repositories"},
	{"POST", "/user/repos"},
	{"POST", "/orgs/:org/repos"},
	{"GET", "/repos/:owner/:repo"},
	//{"PATCH", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/contributors"},
	{"GET", "/repos/:owner/:repo/languages"},
	{"GET", "/repos/:owner/:repo/teams"},
	{"GET", "/repos/:owner/:repo/tags"},
	{"GET", "/repos/:owner/:repo/branches"},
	{"GET", "/repos/:owner/:repo/branches/:branch"},
	{"DELETE", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/collaborators"},
	{"GET", "/repos/:owner/:repo/collaborators/:user"},
	{"PUT", "/repos/:owner/:repo/collaborators/:user"},
	{"DELETE", "/repos/:owner/:repo/collaborators/:user"},
	{"GET", "/repos/:owner/:repo/comments"},
	{"GET", "/repos/:owner/:repo/commits/:sha/comments"},
	{"POST", "/repos/:owner/:repo/commits/:sha/comments"},
	{"GET", "/repos/:owner/:repo/comments/:id"},
	//{"PATCH", "/repos/:owner/:repo/comments/:id"},
	{"DELETE", "/repos/:owner/:repo/comments/:id"},
	{"GET", "/repos/:owner/:repo/commits"},
	{"GET", "/repos/:owner/:repo/commits/:sha"},
	{"GET", "/repos/:owner/:repo/readme"},
	//{"GET", "/repos/:owner/:repo/contents/*path"},
	//{"PUT", "/repos/:owner/:repo/contents/*path"},
	//{"DELETE", "/repos/:owner/:repo/contents/*path"},
	//{"GET", "/repos/:owner/:repo/:archive_format/:ref"},
	{"GET", "/repos/:owner/:repo/keys"},
	{"GET", "/repos/:owner/:repo/keys/:id"},
	{"POST", "/repos/:owner/:repo/keys"},
	//{"PATCH", "/repos/:owner/:repo/keys/:id"},
	{"DELETE", "/repos/:owner/:repo/keys/:id"},
	{"GET", "/repos/:owner/:repo/downloads"},
	{"GET", "/repos/:owner/:repo/downloads/:id"},
	{"DELETE", "/repos/:owner/:repo/downloads/:id"},
	{"GET", "/repos/:owner/:repo/forks"},
	{"POST", "/repos/:owner/:repo/forks"},
	{"GET", "/repos/:owner/:repo/hooks"},
	{"GET", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks"},
	//{"PATCH", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks/:id/tests"},
	{"DELETE", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/merges"},
	{"GET", "/repos/:owner/:repo/releases"},
	{"GET", "/repos/:owner/:repo/releases/:id"},
	{"POST", "/repos/:owner/:repo/releases"},
	//{"PATCH", "/repos/:owner/:repo/releases/:id"},
	{"DELETE", "/repos/:owner/:repo/releases/:id"},
	{"GET", "/repos/:owner/:repo/releases/:id/assets"},
	{"GET", "/repos/:owner/:repo/stats/contributors"},
	{"GET", "/repos/:owner/:repo/stats/commit_activity"},
	{"GET", "/repos/:owner/:repo/stats/code_frequency"},
	{"GET", "/repos/:owner/:repo/stats/participation"},
	{"GET", "/repos/:owner/:repo/stats/punch_card"},
	{"GET", "/repos/:owner/:repo/statuses/:ref"},
	{"POST", "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{"GET", "/search/repositories"},
	{"GET", "/search/code"},
	{"GET", "/search/issues"},
	{"GET", "/search/users"},
	{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{"GET", "/legacy/repos/search/:keyword"},
	{"GET", "/legacy/user/search/:keyword"},
	{"GET", "/legacy/user/email/:email"},

	// Users
	{"GET", "/users/:user"},
	{"GET", "/user"},
	//{"PATCH", "/user"},
	{"GET", "/users"},
	{"GET", "/user/emails"},
	{"POST", "/user/emails"},
	{"DELETE", "/user/emails"},
	{"GET", "/users/:user/followers"},
	{"GET", "/user/followers"},
	{"GET", "/users/:user/following"},
	{"GET", "/user/following"},
	{"GET", "/user/following/:user"},
	{"GET", "/users/:user/following/:target_user"},
	{"PUT", "/user/following/:user"},
	{"DELETE", "/user/following/:user"},
	{"GET", "/users/:user/keys"},
	{"GET", "/user/keys"},
	{"GET", "/user/keys/:id"},
	{"POST", "/user/keys"},
	//{"PATCH", "/user/keys/:id"},
	{"DELETE", "/user/keys/:id"},
}

var (
	githubGocraftWeb  http.Handler
	githubGorillaMux  http.Handler
	githubHttpRouter  http.Handler
	githubHttpTreeMux http.Handler
	githubMartini     http.Handler
	githubPat         http.Handler
	githubTigerTonic  http.Handler
	githubTraffic     http.Handler
	githubKocha       http.Handler
)

func init() {
	println("#GithubAPI Routes:", len(githubAPI))

	githubGocraftWeb = loadGocraftWeb(githubAPI)
	githubGorillaMux = loadGorillaMux(githubAPI)
	githubHttpRouter = loadHttpRouter(githubAPI)
	githubHttpTreeMux = loadHttpTreeMux(githubAPI)
	githubMartini = loadMartini(githubAPI)
	githubPat = loadPat(githubAPI)
	githubTigerTonic = loadTigerTonic(githubAPI)
	githubTraffic = loadTraffic(githubAPI)
	githubKocha = loadKocha(githubAPI)
}

// Static
func BenchmarkGocraftWeb_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGorillaMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkHttpRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	req.RequestURI = "/user/repos"
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkMartini_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkTigerTonic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTraffic, req)
}

func BenchmarkKocha_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubKocha, req)
}

// Param
func BenchmarkGocraftWeb_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGorillaMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkHttpRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	req.RequestURI = "/repos/julienschmidt/httprouter/stargazers"
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkMartini_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkTigerTonic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkKocha_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubKocha, req)
}

// All routes
func BenchmarkGocraftWeb_GithubAll(b *testing.B) {
	benchRoutes(b, githubGocraftWeb, githubAPI)
}
func BenchmarkGorillaMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubGorillaMux, githubAPI)
}
func BenchmarkHttpRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpRouter, githubAPI)
}
func BenchmarkHttpTreeMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpTreeMux, githubAPI)
}
func BenchmarkMartini_GithubAll(b *testing.B) {
	benchRoutes(b, githubMartini, githubAPI)
}
func BenchmarkPat_GithubAll(b *testing.B) {
	benchRoutes(b, githubPat, githubAPI)
}
func BenchmarkTigerTonic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTigerTonic, githubAPI)
}
func BenchmarkTraffic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTraffic, githubAPI)
}
func BenchmarkKocha_GithubAll(b *testing.B) {
	benchRoutes(b, githubKocha, githubAPI)
}
