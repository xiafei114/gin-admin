const api = {
  Login: '/v1/login',
  Logout: '/v1/login/exit',
  // 根据验证码id获得图像
  LoginCaptcha: '/v1/login/captcha',
  ForgePassword: '/auth/forge-password',
  Register: '/auth/register',
  // 获得验证码id
  twoStepCode: '/v1/login/captchaid',
  SendSms: '/account/sms',
  SendSmsErr: '/account/sms_err',
  // get my info
  UserInfo: '/user/info'
}
export default api
