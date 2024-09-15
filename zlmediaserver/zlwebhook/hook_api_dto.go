package zlwebhook

type MediaServerID struct {
	MediaServerID        string `json:"mediaServerId"`         // 服务器 ID，通过配置文件设置
	GeneralMediaServerID string `json:"general.mediaServerId"` // 服务器 ID
}

type OnFlowReportRequest struct {
	App           string `json:"app"`           // 流应用名
	Duration      int    `json:"duration"`      // TCP 链接维持时间，单位秒
	Params        string `json:"params"`        // 推流或播放 URL 参数
	Player        bool   `json:"player"`        // true 为播放器，false 为推流器
	Schema        string `json:"schema"`        // 播放或推流的协议，可能是 rtsp、rtmp、http
	Stream        string `json:"stream"`        // 流 ID
	TotalBytes    int    `json:"totalBytes"`    // 耗费上下行流量总和，单位字节
	Vhost         string `json:"vhost"`         // 流虚拟主机
	IP            string `json:"ip"`            // 客户端 IP
	Port          int    `json:"port"`          // 客户端端口号
	ID            string `json:"id"`            // TCP 链接唯一 ID
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OnHttpAccessRequest struct {
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // HTTP 客户端 IP
	IsDir         bool   `json:"is_dir"`        // HTTP 访问路径是文件还是目录
	Params        string `json:"params"`        // HTTP URL 参数
	Path          string `json:"path"`          // 请求访问的文件或目录
	Port          uint16 `json:"port"`          // HTTP 客户端端口号
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type OnHttpAccessResponse struct {
	Code   int    `json:"code"`
	Err    string `json:"err"`
	Path   string `json:"path"`
	Second int    `json:"second"`
}

type OnPlayRequest struct {
	App           string `json:"app"`           // 流应用名
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // 播放器 IP
	Params        string `json:"params"`        // 播放 URL 参数
	Port          uint16 `json:"port"`          // 播放器端口号
	Schema        string `json:"schema"`        // 播放的协议，可能是 RTSP、RTMP、HTTP
	Stream        string `json:"stream"`        // 流 ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type OnPublishRequest struct {
	App           string `json:"app"`           // 流应用名
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // 推流器 IP
	Params        string `json:"params"`        // 推流 URL 参数
	Port          uint16 `json:"port"`          // 推流器端口号
	Schema        string `json:"schema"`        // 推流的协议，可能是 RTSP、RTMP
	Stream        string `json:"stream"`        // 流 ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type OnPublishResponse struct {
	Code           int    `json:"code"`             // 错误代码，0 代表允许推流
	Msg            string `json:"msg"`              // 不允许推流时的错误提示
	EnableHLS      bool   `json:"enable_hls"`       // 是否转换成 HLS-MPEGTS 协议
	EnableHLSFMP4  bool   `json:"enable_hls_fmp4"`  // 是否转换成 HLS-FMP4 协议
	EnableMP4      bool   `json:"enable_mp4"`       // 是否允许 MP4 录制
	EnableRTSP     bool   `json:"enable_rtsp"`      // 是否转 RTSP 协议
	EnableRTMP     bool   `json:"enable_rtmp"`      // 是否转 RTMP/FLV 协议
	EnableTS       bool   `json:"enable_ts"`        // 是否转 HTTP-TS/WS-TS 协议
	EnableFMP4     bool   `json:"enable_fmp4"`      // 是否转 HTTP-FMP4/WS-FMP4 协议
	HlsDemand      bool   `json:"hls_demand"`       // 该协议是否有人观看才生成
	RtspDemand     bool   `json:"rtsp_demand"`      // 该协议是否有人观看才生成
	RtmpDemand     bool   `json:"rtmp_demand"`      // 该协议是否有人观看才生成
	TsDemand       bool   `json:"ts_demand"`        // 该协议是否有人观看才生成
	Fmp4Demand     bool   `json:"fmp4_demand"`      // 该协议是否有人观看才生成
	EnableAudio    bool   `json:"enable_audio"`     // 转协议时是否开启音频
	AddMuteAudio   bool   `json:"add_mute_audio"`   // 转协议时，无音频是否添加静音 AAC 音频
	Mp4SavePath    string `json:"mp4_save_path"`    // MP4 录制文件保存根目录，置空使用默认
	Mp4MaxSecond   int    `json:"mp4_max_second"`   // MP4 录制切片大小，单位秒
	Mp4AsPlayer    bool   `json:"mp4_as_player"`    // MP4 录制是否当作观看者参与播放人数计数
	HlsSavePath    string `json:"hls_save_path"`    // HLS 文件保存根目录，置空使用默认
	ModifyStamp    int    `json:"modify_stamp"`     // 该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)
	ContinuePushMs uint32 `json:"continue_push_ms"` // 断连续推延时，单位毫秒，置空使用配置文件默认值
	AutoClose      bool   `json:"auto_close"`       // 无人观看是否自动关闭流(不触发无人观看 hook)
	StreamReplace  string `json:"stream_replace"`   // 是否修改流 ID，通过此参数可以自定义流 ID(譬如替换 SSRC)
}

type OnRecordMp4Request struct {
	App           string  `json:"app"`           // 录制的流应用名
	FileName      string  `json:"file_name"`     // 文件名
	FilePath      string  `json:"file_path"`     // 文件绝对路径
	FileSize      int     `json:"file_size"`     // 文件大小，单位字节
	Folder        string  `json:"folder"`        // 文件所在目录路径
	StartTime     int     `json:"start_time"`    // 开始录制时间戳
	Stream        string  `json:"stream"`        // 录制的流 ID
	TimeLen       float64 `json:"time_len"`      // 录制时长，单位秒
	URL           string  `json:"url"`           // http/rtsp/rtmp 点播相对 url 路径
	Vhost         string  `json:"vhost"`         // 流虚拟主机
	MediaServerID string  `json:"mediaServerId"` // 服务器 id，通过配置文件设置
}

type OnRtspRealmRequest struct {
	App           string `json:"app"`           // 流应用名
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // rtsp 播放器 ip
	Params        string `json:"params"`        // 播放 rtsp url 参数
	Port          uint16 `json:"port"`          // rtsp 播放器端口号
	Schema        string `json:"schema"`        // rtsp 或 rtsps
	Stream        string `json:"stream"`        // 流 ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器 id，通过配置文件设置
}

type OnRtspRealmResponse struct {
	BaseResponse
	Realm string `json:"realm"` //该 rtsp 流是否需要 rtsp 专有鉴权，空字符串代码不需要鉴权
}

type OnRtspAuthRequest struct {
	App           string `json:"app"`             // 流应用名
	ID            string `json:"id"`              // TCP 链接唯一 ID
	IP            string `json:"ip"`              // rtsp 播放器 ip
	MustNoEncrypt bool   `json:"must_no_encrypt"` // 请求的密码是否必须为明文(base64 鉴权需要明文密码)
	Params        string `json:"params"`          // rtsp url 参数
	Port          uint16 `json:"port"`            // rtsp 播放器端口号
	Realm         string `json:"realm"`           // rtsp 播放鉴权加密 realm
	Schema        string `json:"schema"`          // rtsp 或 rtsps
	Stream        string `json:"stream"`          // 流 ID
	UserName      string `json:"user_name"`       // 播放用户名
	Vhost         string `json:"vhost"`           // 流虚拟主机
	MediaServerID string `json:"mediaServerId"`   // 服务器 id，通过配置文件设置
}

type OnRtspAuthResponse struct {
	BaseResponse
	Encrypted bool   `json:"encrypted"` //用户密码是明文还是摘要
	Passwd    string `json:"passwd"`    //用户密码明文或摘要(md5(username:realm:password))
}

type OnShellLoginRequest struct {
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // telnet 终端 ip
	Passwd        bool   `json:"passwd"`        // telnet 终端登录用户密码
	Port          uint16 `json:"port"`          // telnet 终端端口号
	UserName      string `json:"user_name"`     // telnet 终端登录用户名
	MediaServerID string `json:"mediaServerId"` // 服务器 id，通过配置文件设置
}

type OriginSock struct {
	Identifier string `json:"identifier"` // 唯一标识符
	LocalIP    string `json:"local_ip"`   // 本机 IP
	LocalPort  int    `json:"local_port"` // 本机端口
	PeerIP     string `json:"peer_ip"`    // 对端 IP
	PeerPort   int    `json:"peer_port"`  // 对端端口
}

type Track struct {
	Channels   int    `json:"channels,omitempty"`    // 音频通道数
	CodecID    int    `json:"codec_id"`              // 编码类型 ID
	CodecName  string `json:"codec_id_name"`         // 编码类型名称
	CodecType  int    `json:"codec_type"`            // 编码类型（视频或音频）
	Ready      bool   `json:"ready"`                 // 轨道是否准备就绪
	SampleBit  int    `json:"sample_bit,omitempty"`  // 音频采样位数
	SampleRate int    `json:"sample_rate,omitempty"` // 音频采样率
	FPS        int    `json:"fps,omitempty"`         // 视频帧率
	Height     int    `json:"height,omitempty"`      // 视频高度
	Width      int    `json:"width,omitempty"`       // 视频宽度
}

type OnStreamChangedRequest struct {
	Regist           bool       `json:"regist"`           // 注册标志
	AliveSecond      int        `json:"aliveSecond"`      // 存活时间，单位秒
	App              string     `json:"app"`              // 应用名
	BytesSpeed       int        `json:"bytesSpeed"`       // 数据产生速度，单位 byte/s
	CreateStamp      int64      `json:"createStamp"`      // 创建时间戳，单位秒
	MediaServerID    string     `json:"mediaServerId"`    // 服务器 ID
	OriginSock       OriginSock `json:"originSock"`       // 原始套接字信息
	OriginType       int        `json:"originType"`       // 产生源类型
	OriginTypeStr    string     `json:"originTypeStr"`    // 产生源类型字符串
	OriginUrl        string     `json:"originUrl"`        // 产生源的 URL
	ReaderCount      int        `json:"readerCount"`      // 观看人数
	Schema           string     `json:"schema"`           // 协议
	Stream           string     `json:"stream"`           // 流 ID
	TotalReaderCount int        `json:"totalReaderCount"` // 总观看人数
	Tracks           []Track    `json:"tracks"`           // 轨道信息
	Vhost            string     `json:"vhost"`            // 流虚拟主机
}

type OnStreamNoneReaderRequest struct {
	App           string `json:"app"`           // 流应用名
	Schema        string `json:"schema"`        // rtsp 或 rtmp
	Stream        string `json:"stream"`        // 流 ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type OnStreamNoneReaderResponse struct {
	BaseResponse
	Close bool `json:"close"` //是否关闭推流或拉流
}

type OnStreamNotFoundRequest struct {
	App           string `json:"app"`           // 流应用名
	ID            string `json:"id"`            // TCP 链接唯一 ID
	IP            string `json:"ip"`            // 播放器 IP
	Params        string `json:"params"`        // 播放 URL 参数
	Port          uint16 `json:"port"`          // 播放器端口号
	Schema        string `json:"schema"`        // 播放的协议，可能是 rtsp、rtmp
	Stream        string `json:"stream"`        // 流 ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器 ID，通过配置文件设置
}

type OnServerStartedRequest struct {
	APIAPIDebug                    string `json:"api.apiDebug"`                    // API 调试
	APISecret                      string `json:"api.secret"`                      // API 密钥
	FFmpegBin                      string `json:"ffmpeg.bin"`                      // FFmpeg 路径
	FFmpegCmd                      string `json:"ffmpeg.cmd"`                      // FFmpeg 命令
	FFmpegLog                      string `json:"ffmpeg.log"`                      // FFmpeg 日志路径
	GeneralMediaServerID           string `json:"general.mediaServerId"`           // 服务器 ID
	GeneralAddMuteAudio            string `json:"general.addMuteAudio"`            // 添加静音音频
	GeneralEnableVhost             string `json:"general.enableVhost"`             // 启用虚拟主机
	GeneralFlowThreshold           string `json:"general.flowThreshold"`           // 流量阈值
	GeneralMaxStreamWaitMS         string `json:"general.maxStreamWaitMS"`         // 最大流等待时间
	GeneralPublishToHls            string `json:"general.publishToHls"`            // 发布到 HLS
	GeneralPublishToMP4            string `json:"general.publishToMP4"`            // 发布到 MP4
	GeneralPublishToRtxp           string `json:"general.publishToRtxp"`           // 发布到 RTX
	GeneralResetWhenRePlay         string `json:"general.resetWhenRePlay"`         // 重播时重置
	GeneralStreamNoneReaderDelayMS string `json:"general.streamNoneReaderDelayMS"` // 无观看者流延迟
	GeneralUltraLowDelay           string `json:"general.ultraLowDelay"`           // 超低延迟
	HlsFileBufSize                 string `json:"hls.fileBufSize"`                 // HLS 文件缓冲大小
	HlsFilePath                    string `json:"hls.filePath"`                    // HLS 文件路径
	HlsSegDur                      string `json:"hls.segDur"`                      // HLS 段持续时间
	HlsSegNum                      string `json:"hls.segNum"`                      // HLS 段数量
	HlsSegRetain                   string `json:"hls.segRetain"`                   // HLS 段保留数量
	HookAdminParams                string `json:"hook.admin_params"`               // Hook 管理参数
	HookEnable                     string `json:"hook.enable"`                     // 启用 Hook
	HookOnFlowReport               string `json:"hook.on_flow_report"`             // 流报告 Hook
	HookOnHttpAccess               string `json:"hook.on_http_access"`             // HTTP 访问 Hook
	HookOnPlay                     string `json:"hook.on_play"`                    // 播放 Hook
	HookOnPublish                  string `json:"hook.on_publish"`                 // 发布 Hook
	HookOnRecordMp4                string `json:"hook.on_record_mp4"`              // 记录 MP4 Hook
	HookOnRtspAuth                 string `json:"hook.on_rtsp_auth"`               // RTSP 认证 Hook
	HookOnRtspRealm                string `json:"hook.on_rtsp_realm"`              // RTSP Realm Hook
	HookOnServerStarted            string `json:"hook.on_server_started"`          // 服务器启动 Hook
	HookOnShellLogin               string `json:"hook.on_shell_login"`             // Shell 登录 Hook
	HookOnStreamChanged            string `json:"hook.on_stream_changed"`          // 流改变 Hook
	HookOnStreamNoneReader         string `json:"hook.on_stream_none_reader"`      // 无观看者流 Hook
	HookOnStreamNotFound           string `json:"hook.on_stream_not_found"`        // 流未找到 Hook
	HookTimeoutSec                 string `json:"hook.timeoutSec"`                 // Hook 超时时间
	HttpCharSet                    string `json:"http.charSet"`                    // HTTP 字符集
	HttpKeepAliveSecond            string `json:"http.keepAliveSecond"`            // HTTP 保活时间
	HttpMaxReqCount                string `json:"http.maxReqCount"`                // HTTP 最大请求数量
	HttpMaxReqSize                 string `json:"http.maxReqSize"`                 // HTTP 最大请求大小
	HttpNotFound                   string `json:"http.notFound"`                   // HTTP 404 页面
	HttpPort                       string `json:"http.port"`                       // HTTP 端口
	HttpRootPath                   string `json:"http.rootPath"`                   // HTTP 根路径
	HttpSendBufSize                string `json:"http.sendBufSize"`                // HTTP 发送缓冲大小
	HttpSslport                    string `json:"http.sslport"`                    // HTTPS 端口
	MulticastAddrMax               string `json:"multicast.addrMax"`               // 组播最大地址
	MulticastAddrMin               string `json:"multicast.addrMin"`               // 组播最小地址
	MulticastUdpTTL                string `json:"multicast.udpTTL"`                // 组播 UDP TTL
	RecordAppName                  string `json:"record.appName"`                  // 记录应用名
	RecordFastStart                string `json:"record.fastStart"`                // 快速启动记录
	RecordFileBufSize              string `json:"record.fileBufSize"`              // 记录文件缓冲大小
	RecordFilePath                 string `json:"record.filePath"`                 // 记录文件路径
	RecordFileRepeat               string `json:"record.fileRepeat"`               // 记录文件重复
	RecordFileSecond               string `json:"record.fileSecond"`               // 记录文件秒数
	RecordSampleMS                 string `json:"record.sampleMS"`                 // 记录采样时间
	RtmpHandshakeSecond            string `json:"rtmp.handshakeSecond"`            // RTMP 握手时间
	RtmpKeepAliveSecond            string `json:"rtmp.keepAliveSecond"`            // RTMP 保活时间
	RtmpModifyStamp                string `json:"rtmp.modifyStamp"`                // RTMP 修改时间戳
	RtmpPort                       string `json:"rtmp.port"`                       // RTMP 端口
	RtpAudioMtuSize                string `json:"rtp.audioMtuSize"`                // RTP 音频 MTU 大小
	RtpClearCount                  string `json:"rtp.clearCount"`                  // RTP 清除计数
	RtpCycleMS                     string `json:"rtp.cycleMS"`                     // RTP 周期
	RtpMaxRtpCount                 string `json:"rtp.maxRtpCount"`                 // RTP 最大 RTP 计数
	RtpVideoMtuSize                string `json:"rtp.videoMtuSize"`                // RTP 视频 MTU 大小
	RtspAuthBasic                  string `json:"rtsp.authBasic"`                  // RTSP 基本认证
	RtspDirectProxy                string `json:"rtsp.directProxy"`                // RTSP 直接代理
	RtspHandshakeSecond            string `json:"rtsp.handshakeSecond"`            // RTSP 握手时间
	RtspKeepAliveSecond            string `json:"rtsp.keepAliveSecond"`            // RTSP 保活时间
	RtspModifyStamp                string `json:"rtsp.modifyStamp"`                // RTSP 修改时间戳
	RtspPort                       string `json:"rtsp.port"`                       // RTSP 端口
	RtspSslport                    string `json:"rtsp.sslport"`                    // RTSP SSL 端口
	ShellMaxReqSize                string `json:"shell.maxReqSize"`                // Shell 最大请求大小
	ShellPort                      string `json:"shell.port"`                      // Shell 端口
}

type OnServerKeepaliveData struct {
	Buffer                int `json:"Buffer"`
	BufferLikeString      int `json:"BufferLikeString"`
	BufferList            int `json:"BufferList"`
	BufferRaw             int `json:"BufferRaw"`
	Frame                 int `json:"Frame"`
	FrameImp              int `json:"FrameImp"`
	MediaSource           int `json:"MediaSource"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer"`
	RtmpPacket            int `json:"RtmpPacket"`
	RtpPacket             int `json:"RtpPacket"`
	Socket                int `json:"Socket"`
	TcpClient             int `json:"TcpClient"`
	TcpServer             int `json:"TcpServer"`
	TcpSession            int `json:"TcpSession"`
	UdpServer             int `json:"UdpServer"`
	UdpSession            int `json:"UdpSession"`
}

type OnServerKeepaliveRequest struct {
	Data          OnServerKeepaliveData `json:"data"`
	MediaServerID string                `json:"mediaServerId"`
}

type OnRtpServerTimeoutRequest struct {
	LocalPort     int    `json:"local_port"`    // openRtpServer 输入的参数
	ReUsePort     bool   `json:"re_use_port"`   // openRtpServer 输入的参数
	Ssrc          uint32 `json:"ssrc"`          // openRtpServer 输入的参数
	StreamID      string `json:"stream_id"`     // openRtpServer 输入的参数
	TcpMode       int    `json:"tcp_mode"`      // openRtpServer 输入的参数
	MediaServerID string `json:"mediaServerId"` // 服务器 id,通过配置文件设置
}
