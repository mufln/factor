import React, { Component } from 'react';

class Teams extends Component {
    
    render() {
    return(
        <div className='content'>
            <div class="teams-top">
                <h2>Команды</h2>
                <button class="green-button" id="new-team">Создать команду</button>
            </div>
            <ul id="teams" class="teams"> 
            </ul>    
        </div>
    );   
}
}

export default Teams;