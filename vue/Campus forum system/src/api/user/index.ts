// 统一管理 用户相关接口
import request from "@/utils/request";
import type { loginFrom, loginResponseData } from "./type";

//统一管理接口 
enum API {
    LOGIN_URL = '/user/login',
    REGISTER_URL = '/user/register',
    LOGOUT_URL = '/user/logout',
    UserInfo_URL = '/user/info/:id',
    CurrentUser_URL = '/user/current',
    UserProfile_URL = '/user/profile'
}

export const reqLogin = (data:loginFrom) => {
    return request.post<any, loginResponseData>(API.LOGIN_URL, data)
}