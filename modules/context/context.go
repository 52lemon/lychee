package context

import (
    "fmt"
    "html/template"
    "io"
    "net/http"
    "strings"
    "time"
)

// Context represents context of a request.
type Context struct {
    Session session.Store
    User        *models.User
    IsSigned    bool
}

func (ctx *Context) GetErrMsg() string {
    return ctx.Data["ErrorMsg"].(string)
}


// NotFoundOrServerError use error check function to determine if the error
// is about not found. It responses with 404 status code for not found error,
// or error context description for logging purpose of 500 server error.
func (ctx *Context) NotFoundOrServerError(title string, errck func(error) bool, err error) {
    if errck(err) {
        ctx.Handle(404, title, err)
        return
    }

    ctx.Handle(500, title, err)
}

// Contexter initializes a classic context for a request.
func Contexter() macaron.Handler {
    return func(c *macaron.Context, l i18n.Locale, cache cache.Cache, sess session.Store, f *session.Flash, x csrf.CSRF) {
        ctx := &Context{
            Context: c,
            Session: sess,
        }
        // Get user from session if logined.
        ctx.User, ctx.IsBasicAuth = auth.SignedInUser(ctx.Context, ctx.Session)

        if ctx.User != nil {
            ctx.IsSigned = true
            ctx.Data["IsSigned"] = ctx.IsSigned
            ctx.Data["SignedUser"] = ctx.User
            ctx.Data["SignedUserID"] = ctx.User.ID
            ctx.Data["SignedUserName"] = ctx.User.Name
            ctx.Data["IsAdmin"] = ctx.User.IsAdmin
        } else {
            ctx.Data["SignedUserID"] = 0
            ctx.Data["SignedUserName"] = ""
        }

        // If request sends files, parse them here otherwise the Query() can't be parsed and the CsrfToken will be invalid.
        if ctx.Req.Method == "POST" && strings.Contains(ctx.Req.Header.Get("Content-Type"), "multipart/form-data") {
            if err := ctx.Req.ParseMultipartForm(setting.AttachmentMaxSize << 20); err != nil && !strings.Contains(err.Error(), "EOF") { // 32MB max size
                ctx.Handle(500, "ParseMultipartForm", err)
                return
            }
        }

        ctx.Data["CsrfToken"] = x.GetToken()
        ctx.Data["CsrfTokenHtml"] = template.HTML(`<input type="hidden" name="_csrf" value="` + x.GetToken() + `">`)
        log.Debug("Session ID: %s", sess.ID())
        log.Debug("CSRF Token: %v", ctx.Data["CsrfToken"])

        ctx.Data["ShowRegistrationButton"] = setting.Service.ShowRegistrationButton
        ctx.Data["ShowFooterBranding"] = setting.ShowFooterBranding
        ctx.Data["ShowFooterVersion"] = setting.ShowFooterVersion

        c.Map(ctx)
    }
}
