export interface loginFrom{
    username: string
    email: string
    password: string
}

export interface loginResponseData{
    Code :    number// 返回值code
	Success :  string     // 是否执行成功
	Message :  string  // 附加信息
}

export interface logoutResponseData{
    Code :    number// 返回值code
	Success :  string     // 是否执行成功
	Message :  string  // 附加信息
}