package consts


const (
	// Success 业务响应状态码
	Success              = 200  //成功（获取成功、操作成功、创建成功、更新新成功、调用成功、响应成功）
	ServiceInternalError = -1   //服务内部错误
	GainFail             = 1001 //获取信息失败
	SubmitFail           = 1002 //提交失败
	CreationFail         = 1003 //创建失败
	UpdateFail           = 1004 //更新失败
	CallFail             = 1005 //调用失败
	ResponseFail         = 1006 //响应失败
	RequestFail          = 1007 //请求失败
	ParameterError       = 1008 //参数异常
	OverTime             = 1009 //请求超时
	UnknownErr           = 1010 //未知错误

	// ErrorsConfigInitFail 配置文件
	ErrorsConfigInitFail      string = "初始化配置文件发生错误"
	ErrorsConfigYamlNotExists string = "配置文件不存在"

	//CurdStatusOkCode CURD常用状态码
	CurdStatusOkCode   int    = 0
	CurdStatusOkMsg    string = "Success"
	CurdSelectFailCode int    = 2000 //查询
	CurdSelectFailMsg  string = "查询失败"
	CurdUpdateFailCode int    = 2001 //更新
	CurdUpdateFailMsg  string = "更新失败"
	CurdWriteFailCode  int    = 2002 // 写入
	CurdWriteFailMsg   string = "写入失败"

	// ValidatorPrefix 表单验证
	ValidatorPrefix              string = "Form_Validator_" //表单验证前缀
	ValidatorParamsCheckFailCode int    = -1
	ValidatorParamsCheckFailMsg  string = "参数校验失败"
	ValidatorParamsToJSONFail    string = "验证器参数 json 反序列化失败"

	// DefaultMysqlPool 数据库连接池配置
	DefaultMysqlPool = "default"
	ProMysqlPool     = "pro"

	//角色
	RoleSys      = "system_control"
	RoleAudit    = "audit_control"
	RoleSecurity = "security_control"

	ParameterFormatError = "参数格式错误"
)
