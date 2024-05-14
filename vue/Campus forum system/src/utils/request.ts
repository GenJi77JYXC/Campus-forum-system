// 进行axios二次封装：使用请求与响应拦截器

import axios from "axios";

//第一步：利用axios对象的create方法，去创建axios实例（其他的配置：基础路径，超时时间）
let request = axios.create({
    //基础路径
    baseURL:'http://localhost:8080/api', //基础路径上会携带api
    timeout: 5000 ,//超时的时间的设置
});

// 第二步： request实例添加请求与响应拦截器
request.interceptors.request.use((config) => {
    //config配置对象，headers属性请求头, 经常给服务器端携带公共参数 

    config.headers.common['X-Client'] = 'campus-forum-system'
    config.headers.post['Content-Type'] = 'application/json; charset=utf-8'


    //返回配置对象
    return config
});