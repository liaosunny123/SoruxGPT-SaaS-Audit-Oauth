package config

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT             = 8080
	ForbiddenWords   = []string{} // 禁止词
	OAIKEY           = ""         // OAIKEY
	OAIKEYLOG        = ""         // OAIKEYLOG 隐藏
	MODERATION       = ""
	Oauth            = ""
	PlusModels       = garray.NewStrArrayFrom([]string{"gpt-4", "gpt-4o", "gpt-4-browsing", "gpt-4-plugins", "gpt-4-mobile", "gpt-4-code-interpreter", "gpt-4-dalle", "gpt-4-gizmo", "gpt-4-magic-create", "o1-preview", "o1-mini", "video_gen", "o1", "o1-pro"})
	AdminApiKey      = ""
	SyncAdmitPostUrl = ""
)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	oaikey := g.Cfg().MustGetWithEnv(ctx, "OAIKEY").String()
	// oaikey 不为空
	if oaikey != "" {
		OAIKEY = oaikey
		// 日志隐藏 oaikey，有 * 代表有值
		OAIKEYLOG = "******"
	}
	g.Log().Info(ctx, "OAIKEY:", OAIKEYLOG)
	moderation := g.Cfg().MustGetWithEnv(ctx, "MODERATION").String()
	if moderation != "" {
		MODERATION = moderation
	}
	g.Log().Info(ctx, "MODERATION:", MODERATION)

	oauthUrl := g.Cfg().MustGetWithEnv(ctx, "OAUTH").String()
	if oauthUrl != "" {
		Oauth = oauthUrl
	}
	g.Log().Info(ctx, "OAUTH:", Oauth)

	syncAdmitPostUrl := g.Cfg().MustGetWithEnv(ctx, "SYNC_ADMIT_POST_URL").String()
	if syncAdmitPostUrl != "" {
		SyncAdmitPostUrl = syncAdmitPostUrl
	}
	g.Log().Info(ctx, "SYNC_ADMIT_POST_URL:", SyncAdmitPostUrl)
}
