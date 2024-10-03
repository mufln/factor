import React from 'react';
import '../../styles/SLogin.scss';
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'
// import { Link } from 'react-router-dom';

class Login extends React.Component{
  state = {
    loginOpen: true
  };
  renderRegister = () => {
    this.setState({loginOpen: false});
  }
  renderLogin = () => {
    this.setState({loginOpen: true});
  }
  render() {
    return (
    <div className="login-wrapper">
      <div className="logo">
        <svg width="60" height="60" viewBox="0 0 44 44" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M39.6889 2V42H2M39.6889 2L37.3778 8.42857M39.6889 2L42 8.42857" stroke="#01A274" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M11.7297 32.6977L18.8898 23.8517L27.9709 32.6977L32.6861 13.1628M32.6861 13.1628L27.9709 17.9544M32.6861 13.1628L34.4324 19.4287" stroke="#01A274" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        <h1>Фактор</h1>
      </div>
      {this.state.loginOpen ? 
        <h2>Авторизация</h2> :
        <h2>Регистрация</h2>}
      {this.state.loginOpen ? 
        <LoginForm></LoginForm> : 
        <RegisterForm></RegisterForm>}
      {this.state.loginOpen ?
        <div className="signin">
          Нет учётной записи? 
          <button className='register-href' onClick={() => this.renderRegister()}>Регистрация</button>
        </div> :
        <div className="signin">
          У Вас уже есть аккаунт? 
          <button className='register-href' onClick={() => this.renderLogin()}>Войти</button>
        </div>
         }
    </div> 
    );
  }
}

export default Login;