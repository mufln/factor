import React from 'react';
import './index.css';
import Form from './Form'

function Register(){
    return (
        <div class="login-wrapper">
      <div class="logo">
        <svg width="60" height="60" viewBox="0 0 44 44" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M39.6889 2V42H2M39.6889 2L37.3778 8.42857M39.6889 2L42 8.42857" stroke="#01A274" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M11.7297 32.6977L18.8898 23.8517L27.9709 32.6977L32.6861 13.1628M32.6861 13.1628L27.9709 17.9544M32.6861 13.1628L34.4324 19.4287" stroke="#01A274" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          
        <h1>Фактор</h1>
      </div>
      <h2>Регистрация</h2>
      <form class="login">
        <input type="text" class="nickname" id="name" placeholder="Логин" required />
        <input type="password" class="password" id="password" placeholder="Пароль" required />
        <input type="password" class="password" id="rep-password" placeholder="Повторите пароль" required />
        <button id="enter" class="enter">Зарегистрироваться</button>
      </form>
    </div>
    );
}

export default Register;