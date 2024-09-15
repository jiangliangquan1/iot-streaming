package zlwebhook

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ZlWebHookController struct {
	logger *logrus.Logger
}

func (z *ZlWebHookController) RegisterRoute(group *gin.RouterGroup) {

	g := group.Group("v1/zl_hook")

	g.POST("on_flow_report", z.onFlowReport)
	g.POST("on_http_access", z.onHttpAccess)
	g.POST("on_play", z.onPlay)
	g.POST("on_publish", z.onPublish)
	g.POST("on_record_mp4", z.onRecordMp4)
	g.POST("on_rtsp_realm", z.onRtspRealm)
	g.POST("on_rtsp_auth", z.onRtspAuth)
	g.POST("on_shell_login", z.onShellLogin)
	g.POST("on_stream_changed", z.onStreamChanged)
	g.POST("on_stream_none_reader", z.onStreamNoneReader)
	g.POST("on_stream_not_found", z.onStreamNotFound)
	g.POST("on_server_started", z.onServerStarted)
	g.POST("on_server_keepalive", z.onServerKeepalive)
	g.POST("on_rtp_server_timeout", z.onRtpServerTimeout)

}

// 流量统计事件
func (z *ZlWebHookController) onFlowReport(ctx *gin.Context) {

	var req OnFlowReportRequest
	ctx.BindJSON(&req)

}

// 访问 http 文件服务器上 hls 之外的文件时触发
func (z *ZlWebHookController) onHttpAccess(ctx *gin.Context) {
}

// 播放鉴权事件
func (z *ZlWebHookController) onPlay(ctx *gin.Context) {
}

// rtsp/rtmp/rtp 推流鉴权事件
func (z *ZlWebHookController) onPublish(ctx *gin.Context) {
}

// 录制 mp4 完成后通知事件
func (z *ZlWebHookController) onRecordMp4(ctx *gin.Context) {
}

// 该 rtsp 流是否开启 rtsp 专用方式的鉴权事件，开启后才会触发 on_rtsp_auth 事件
func (z *ZlWebHookController) onRtspRealm(ctx *gin.Context) {
}

// rtsp 专用的鉴权事件，先触发 on_rtsp_realm 事件然后才会触发 on_rtsp_auth 事件
func (z *ZlWebHookController) onRtspAuth(ctx *gin.Context) {
}

// shell 登录鉴权，ZLMediaKit 提供简单的 telnet 调试方式
func (z *ZlWebHookController) onShellLogin(ctx *gin.Context) {
}

// rtsp/rtmp 流注册或注销时触发此事件；此事件对回复不敏感。
func (z *ZlWebHookController) onStreamChanged(ctx *gin.Context) {
}

// 流无人观看时事件，用户可以通过此事件选择是否关闭无人看的流。
// 一个直播流注册上线了，如果一直没人观看也会触发一次无人观看事件，触发时的协议 schema 是随机的，看哪种协议最晚注册(一般为 hls)。
// 后续从有人观看转为无人观看，触发协议 schema 为最后一名观看者使用何种协议。
// 目前 mp4/hls 录制不当做观看人数(mp4 录制可以通过配置文件 mp4_as_player 控制，但是 rtsp/rtmp/rtp 转推算观看人数，也会触发该事件。
func (z *ZlWebHookController) onStreamNoneReader(ctx *gin.Context) {
}

// 流未找到事件，用户可以在此事件触发时，立即去拉流，这样可以实现按需拉流；此事件对回复不敏感。
func (z *ZlWebHookController) onStreamNotFound(ctx *gin.Context) {
}

// 服务器启动事件，可以用于监听服务器崩溃重启；此事件对回复不敏感。
func (z *ZlWebHookController) onServerStarted(ctx *gin.Context) {
}

// 服务器定时上报时间，上报间隔可配置，默认 10s 上报一次
func (z *ZlWebHookController) onServerKeepalive(ctx *gin.Context) {
}

// 调用 openRtpServer 接口，rtp server 长时间未收到数据,执行此 web hook,对回复不敏感
func (z *ZlWebHookController) onRtpServerTimeout(ctx *gin.Context) {
}

func NewZlWebHookController(l *logrus.Logger) *ZlWebHookController {
	return &ZlWebHookController{logger: l}
}
