package app_default

const Default_private_help = "这里是帮助信息,如下内容括号中的不要输入:\r\n" +
	"acfur私聊功能表：\r\n" +
	"	1.acfur登录：登录acfur软件\r\n" +
	"	2.acfur清除登录：T出所有已经登录APP的设备\r\n" +
	"	3.acfur绑定(+)本机器人密码：绑定当前机器人\r\n" +
	"高危操作：\r\n" +
	"	1.acfur解绑(+)本机器人密码：解绑当前机器人（仅限号主）\r\n" +
	"	2.acfur修改密码(+)新密码：修改当前机器人密码（仅限号主）\r\n" +
	"\r\n"

const Default_private_help_for_RobotOwner = "机器人主功能（仅拥有人可见）：\r\n" +
	"	1.acfur绑定群：展示本机器人允许使用的群\r\n" +
	"	2.acfur绑定群(+)群号：增加一个机器人允许使用的群\r\n" +
	"	3.acfur解绑群(+)群号：删除一个机器人允许使用的群\r\n" +
	""

const Default_group_help = "这里是群帮助信息：\r\n" +
	"acfur功能表：\r\n" +
	"	1.acfur设定：查看设定\r\n" +
	"	2.acfur刷新：查看更多的刷新选项\r\n" +
	"	3.acfurapp：查看控制软件的下载方式(普通群员也可用)\r\n" +
	"如下内容括号中的不要输入\r\n" +
	"	1.acfur设定(+)功能:设定值：对功能进行设定\r\n" +
	"	2.acfur屏蔽词：查看屏蔽词列表\r\n" +
	"	3.acfurT出词：查看T出词列表\r\n" +
	"	4.acfur撤回词：查看撤回词列表\r\n" +
	"	5.acfur屏蔽：查看屏蔽词/T出词/撤回词添加方法\r\n" +
	"		屏蔽词支持正则表达式等，但请不要使用命令行设定避免出错，强烈建议在APP中设定" +
	"\r\n" +
	"acfur测试项目：\r\n" +
	"	1.acfur测试撤回：测试撤回功能（权限问题普通群员测试有效）\r\n" +
	"	2.acfur测试拼音：测试拼音检测功能\r\n" +
	"	3.acfur测试自动撤回：测试机器人延时撤回功能（机器人撤回自己的发言）\r\n" +
	"群员开放功能：\r\n" +
	"	1.签到：群签到（奖励威望）\r\n" +
	"	2.轮盘：直接发送“轮盘”参与，或“轮盘A10”参与左轮模式；可以使用acfur轮盘来查看详细帮助\r\n"

const Default_str_ban_word = "acfur屏蔽：查看屏蔽词/T出词/撤回词添加方法\r\n" +
	"	1.简单添加屏蔽词（禁言+撤回）：acfur屏蔽1(+)屏蔽词\r\n" +
	"	2.简单添加T出词（禁言+撤回）：acfur屏蔽2(+)屏蔽词\r\n" +
	"	3.简单添加撤回（仅撤回）：acfur屏蔽3(+)屏蔽词\r\n" +
	"	4.删除屏蔽词/T出词/撤回词：acfur屏蔽-屏蔽词\r\n" +
	"	5.高级添加屏蔽词/T出词/撤回词：acfur屏蔽+屏蔽词#处罚，例如acfur屏蔽+触发词#T出撤回\r\n" +
	"	6.高级修改屏蔽词/T出词/撤回词：acfur屏蔽=屏蔽词#处罚，例如acfur屏蔽=触发词#屏蔽撤回\r\n"

const Default_str_login_text = "如果忘记密码或需要修改密码，再次发送登录即可\r\n如果需要自定义密码，" +
	"请输入:\r\nacfur密码+要设定的密码\r\n例如:\"acfur密码123456\"\r\n密码长度允许1-16位,没有复杂度限制，可设定中文密码"

const Default_error_alert = "系统故障，向机器人小组反馈：\r\n542749156\r\n谢谢！"

const Default_app_download_url = "https://pandorabox.tuuz.cc:444/app/gobotq/app-release.apk"

const Default_ban_url = "本群不允许发送链接，请勿发送链接或者分享"

const Deulfat_ban_share = "本群不允许发送链接，请勿发送链接或者分享"

const Default_ban_word = "你已经触发了屏蔽词，现在将对你屏蔽"
const Default_ban_group = "请不要分享群"
const Default_ban_weixin = "请不要分享微信"
const Default_ban_share = "请勿发送分享"
const Default_length_limit = "请勿发送长文本"

const Default_kick_word = "用户触发T出词，和他说再见吧？哦不过他好像已经收不到了hhh"

const Default_retract_word = "用户触发了撤回"

const Default_over_time = "您的服务已经到期，如需增加时间请加入反馈：542749156"

const Default_welcome = "Hi~我是Acfur，如果需要查看功能可以发acfurhelp哦~"

const Default_lunpan_help = "轮盘模式说明：" +
	"\r\n\r\n	1.经典轮盘模式：直接发送“轮盘”即可参与经典轮盘模式，有2分5分10分以及翻倍翻10倍等奖励，也有扣5，随机扣，清零以及减半等惩罚" +
	"\r\n\r\n	2.左轮模式，使用“轮盘[模式字母][数字]" +
	"\r\n	例如“轮盘A10”，来挑战1/6的失败几率，挑战成功，奖励1/6押注威望" +
	"\r\n	同时你可以使用轮盘B10来挑战2/6的败率，获得2/6的奖励" +
	"\r\n	可选模式有ABCDE，左轮模式可以押注部分威望或者全部威望来参与" +
	"\r\n	你可以使用威望查询来查看自己的积分可用情况" +
	"\r\n	“轮盘完整度”代表了左轮轮盘的故障率，完整度越低则卡弹率越高" +
	"\r\n	例如使用轮盘F10来押注左轮满弹的必死模式，例如轮盘完整度只剩15％，那么就有85％的机率出现卡弹的情况从而逃脱" +
	"\r\n	所以完整度越低胜率越高，每次有人玩轮盘的时候，都会消耗轮盘的完整度，参与人数越多，胜率越高"

const Default_daoju = "道具说明：" +
	"\r\n	1.使用“acfur道具”来查看道具帮助" +
	"\r\n	2.使用“道具商城”/“道具商店”来查看可购买的道具" +
	"\r\n	3.使用“道具购买”/“道具兑换”[道具名称]来购买一个道具，例如“道具购买免死金牌”来购买对应的道具" +
	"\r\n	4.使用“道具我的/道具列出/道具背包”来查看你所拥有的全部道具" +
	"\r\n	5.使用“道具清空我的背包”来清空你所拥有的全部道具" +
	"\r\n道具获得后将自动启用/抵消无需二次操作"

const Daoju_notfound = "未找到当前道具，或者道具不可获得，请换一个哦，如需查看列表，请输入“道具商店”来查看可兑换列表"

const Daoju_goumai = "您可以使用使用“道具购买”/“道具兑换”[道具名称]来购买一个道具，例如“道具购买免死金牌”来购买对应的道具"

const Daoju_use_for_ban = "您使用道具抵抗了一次禁言伤害"
const Daoju_use_for_kick = "您使用道具抵抗了一次T出伤害"

const Default_send_daoju = "你可以使用道具赠送[道具名称]给[qq号码]来将你购买的道具送给群友"
