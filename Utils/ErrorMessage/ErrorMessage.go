package emsg

const (
	Success = 10000
	Error   = 20000
	//用户错误
	UsernameExist   = 10001
	UsernameNoExist = 10002
	PasswordWrong   = 10003
	//分类错误
	CategoryExist         = 20001
	CreateCategoryFailed  = 20002
	CategoryNoExist       = 20003
	DeleteCategoryFailed  = 20004
	UpdateCategoryFailed  = 20005
	GetCategoryListFailed = 20006
	//帖子错误
	GetPostsFailed     = 30001
	CreatePostsFailed  = 30002
	PostsNoExist       = 30003
	DeletePostsFailed  = 30004
	UpdatePostsFailed  = 30005
	GetPostsListFailed = 30006
	LikePostsFailed    = 30007
	//token错误
	GenerateAccessTokenFailed  = 40001
	GenerateRefreshTokenFailed = 40002
	TokenErrorMalformed        = 40003
	TokenErrorExpired          = 40004
	TokenErrorNotValidYet      = 40005
	TokenCannotRecognized      = 40006
	TokenInvalid               = 40007
	AccessTokenNoExist         = 40008
	RefreshTokenNoExist        = 40009
)

var ErrorMsg map[int]string

func init() {
	ErrorMsg = make(map[int]string)
	ErrorMsg[Success] = "操作成功"
	ErrorMsg[Error] = "操作失败"
	ErrorMsg[UsernameExist] = "用户名已存在"
	ErrorMsg[UsernameNoExist] = "用户名不存在"
	ErrorMsg[PasswordWrong] = "密码错误"
	ErrorMsg[CategoryExist] = "分类已存在"
	ErrorMsg[CreateCategoryFailed] = "创建分类失败"
	ErrorMsg[CategoryNoExist] = "分类不存在"
	ErrorMsg[DeleteCategoryFailed] = "删除分类失败"
	ErrorMsg[UpdateCategoryFailed] = "更新分类信息失败"
	ErrorMsg[GetCategoryListFailed] = "获取分类列表失败"
	ErrorMsg[GetPostsFailed] = "获取分类信息失败"
	ErrorMsg[CreatePostsFailed] = "创建帖子失败"
	ErrorMsg[PostsNoExist] = "帖子不存在"
	ErrorMsg[DeletePostsFailed] = "删除帖子失败"
	ErrorMsg[UpdatePostsFailed] = "更新帖子信息失败"
	ErrorMsg[GetPostsListFailed] = "获取帖子列表失败"
	ErrorMsg[LikePostsFailed] = "点赞帖子失败"
	ErrorMsg[GenerateAccessTokenFailed] = "生成AccessToken失败"
	ErrorMsg[GenerateRefreshTokenFailed] = "生成RefreshToken失败"
	ErrorMsg[TokenErrorMalformed] = "Token格式错误"
	ErrorMsg[TokenErrorExpired] = "Token已过期"
	ErrorMsg[TokenErrorNotValidYet] = "Token未生效"
	ErrorMsg[TokenCannotRecognized] = "无法辨认该Token"
	ErrorMsg[TokenInvalid] = "非法Token"
	ErrorMsg[AccessTokenNoExist] = "AccessToken不存在"
	ErrorMsg[RefreshTokenNoExist] = "RefreshToken不存在"
}

func GetErrorMsg(emsg int) string {
	return ErrorMsg[emsg]
}
