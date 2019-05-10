const api = {
  Login: '/v1/login',
  Logout: '/v1/login/exit',
  // 根据验证码id获得图像
  LoginCaptcha: '/v1/login/captcha',
  // 获得当前用户
  // LoginCurrentUser: '/v1/current/user',
  LoginCurrentUser: '/v1/mock/current/user',
  ForgePassword: '/auth/forge-password',
  Register: '/auth/register',
  // 获得验证码id
  twoStepCode: '/v1/login/captchaid',
  SendSms: '/account/sms',
  SendSmsErr: '/account/sms_err',
  // 获得当前用户
  // UserInfo: '/v1/current/user'
  UserInfo: '/v1/mock/current/user'
}
export default api
