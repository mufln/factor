import React, { Component } from 'react';
import './Register.css';

class Form extends Component {
    state = {
        name: '',
        nameError: null,
        password: '',
        passwordError: null,
        repeatpassword: '',
        repeatpasswordError: null,
    };

    nameChangeHandler = event => {
        const name = event.target.value;
        this.setState({
            name,
            nameError: !name
        });
    };

    passwordChangeHandler = event => {
        const password = event.target.value;
        this.setState({
            password,
            passwordError: !password
        });
    };

    repeatPasswordChangeHandler = event => {
        const repeatpassword = event.target.value;
        this.setState({
            repeatpassword,
            repeatPasswordError: !repeatpassword
        });
    };

    submitHandler = event => {
        event.preventDefault();

        const { name, password, repeatpassword } = this.state;

        if (name && password) {
            this.setState({
                name: '',
                nameError: false,
                password: '',
                passwordError: false,
                repeatpassword: '',
                repeatpasswordError: false
            });
            console.log(name, password, repeatpassword);
            fetch("http://127.0.0.1:5000/login/",{credentials:'include',method:'POST',body:`{"login":"User", "password":"12345"}`})
            return;
        }

        this.setState({
            nameError: !name,
            passwordError: !password,
            repeatPasswordError: !repeatpassword
        });
    };

    render() {
        const { name, nameError, password, passwordError } = this.state;
    return(
        <form className="login" onSubmit={this.submitHandler}>
        <input 
            value={name}
            onChange={this.nameChangeHandler}
            type="text" 
            className="nickname" 
            id="name" 
            placeholder="Логин" 
            required
        />
        {nameError ? (
                        <div className='error'>! Заполните поле</div>
                    ) : null}
        <input 
            value={password}
            onChange={this.passwordChangeHandler}
            type="password" 
            className="password" 
            id="password" 
            placeholder="Пароль" 
            required
        />
        <input 
           
            onChange={this.repeatPasswordChangeHandler}
            type="password" 
            className="repeat-password" 
            id="repeat-password" 
            placeholder="Повторите пароль" 
            required
        />
        {passwordError ? (
                        <div className='error'>! Заполните поле</div>
                    ) : null}
        <button id="enter" className="enter">Войти</button>
      </form>
    );   
}
}

export default Form;