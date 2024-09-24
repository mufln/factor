import React, { Component } from 'react';
import './Login.css';

class Form extends Component {
    state = {
        name: '',
        nameError: null,
        password: '',
        passwordError: null
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

    submitHandler = event => {
        event.preventDefault();

        const { name, password } = this.state;

        if (name && password) {
            this.setState({
                name: '',
                nameError: false,
                password: '',
                passwordError: false
            });
            console.log(name, password);
            var xhr = new XMLHttpRequest();
            var url = "http://localhost:5000/login/";
            xhr.open("POST", url, true);
            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.onreadystatechange = function () {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    console.log("Success");
                }
            };
            var data = JSON.stringify({"login": name, "password": password});
            xhr.send(data);

            return;
        }

        this.setState({
            nameError: !name,
            passwordError: !password
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
        {passwordError ? (
                        <div className='error'>! Заполните поле</div>
                    ) : null}
        <button id="enter" className="enter">Войти</button>
      </form>
    );   
}
}

export default Form;