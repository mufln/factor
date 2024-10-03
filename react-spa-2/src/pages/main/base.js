import React from 'react';
import '../../styles/SLogin.scss';
// import { Link } from 'react-router-dom';

class Base extends React.Component{
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
        <div className="base">
            <div class="sidebar">
            <h1>Фактор</h1>
            <ul>
                <li id="main-page-button"><i class="home">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M11.25 27.5V15H18.75V27.5M3.75 11.25L15 2.5L26.25 11.25V25C26.25 25.663 25.9866 26.2989 25.5178 26.7678C25.0489 27.2366 24.413 27.5 23.75 27.5H6.25C5.58696 27.5 4.95107 27.2366 4.48223 26.7678C4.01339 26.2989 3.75 25.663 3.75 25V11.25Z" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>
                    </i>Главная страница</li>
                <li id="profile-button"><i class="user">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M25 26.25V23.75C25 22.4239 24.4732 21.1521 23.5355 20.2145C22.5979 19.2768 21.3261 18.75 20 18.75H10C8.67392 18.75 7.40215 19.2768 6.46447 20.2145C5.52678 21.1521 5 22.4239 5 23.75V26.25M20 8.75C20 11.5114 17.7614 13.75 15 13.75C12.2386 13.75 10 11.5114 10 8.75C10 5.98858 12.2386 3.75 15 3.75C17.7614 3.75 20 5.98858 20 8.75Z" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>                    
                </i>Профиль</li>
                <li id="chat-button"><i class="comments">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M26.25 14.375C26.2543 16.0248 25.8688 17.6524 25.125 19.125C24.243 20.8897 22.8872 22.374 21.2093 23.4116C19.5314 24.4492 17.5978 24.9992 15.625 25C13.9752 25.0043 12.3476 24.6188 10.875 23.875L3.75 26.25L6.125 19.125C5.38116 17.6524 4.9957 16.0248 5 14.375C5.00076 12.4022 5.55076 10.4686 6.5884 8.79069C7.62603 7.11282 9.11032 5.75696 10.875 4.87501C12.3476 4.13117 13.9752 3.7457 15.625 3.75001H16.25C18.8554 3.89374 21.3163 4.99346 23.1614 6.83858C25.0065 8.6837 26.1063 11.1446 26.25 13.75V14.375Z" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>                        
                </i>Чат</li>
                <li id="teams-button"><i class="file-alt">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <g clip-path="url(#clip0_212_194)">
                        <path d="M21.25 26.25V23.75C21.25 22.4239 20.7232 21.1521 19.7855 20.2145C18.8479 19.2768 17.5761 18.75 16.25 18.75H6.25C4.92392 18.75 3.65215 19.2768 2.71447 20.2145C1.77678 21.1521 1.25 22.4239 1.25 23.75V26.25M28.75 26.25V23.75C28.7492 22.6422 28.3804 21.566 27.7017 20.6904C27.023 19.8148 26.0727 19.1895 25 18.9125M20 3.9125C21.0755 4.18788 22.0288 4.81338 22.7095 5.69039C23.3903 6.5674 23.7598 7.64604 23.7598 8.75625C23.7598 9.86646 23.3903 10.9451 22.7095 11.8221C22.0288 12.6991 21.0755 13.3246 20 13.6M16.25 8.75C16.25 11.5114 14.0114 13.75 11.25 13.75C8.48858 13.75 6.25 11.5114 6.25 8.75C6.25 5.98858 8.48858 3.75 11.25 3.75C14.0114 3.75 16.25 5.98858 16.25 8.75Z" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                        </g>
                        <defs>
                        <clipPath id="clip0_212_194">
                        <rect width="30" height="30" fill="white"/>
                        </clipPath>
                        </defs>
                    </svg>                                           
                </i>Команды</li>
                <li id="tasks-button"><i class="tasks">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M11.25 13.75L15 17.5L27.5 5M26.25 15V23.75C26.25 24.413 25.9866 25.0489 25.5178 25.5178C25.0489 25.9866 24.413 26.25 23.75 26.25H6.25C5.58696 26.25 4.95107 25.9866 4.48223 25.5178C4.01339 25.0489 3.75 24.413 3.75 23.75L3.75 6.25C3.75 5.58696 4.01339 4.95107 4.48223 4.48223C4.95107 4.01339 5.58696 3.75 6.25 3.75L20 3.75" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>                    
                </i>Задачи</li>
                <li id="files-button"><i class="file-alt">
                    <svg width="30" height="30" viewBox="0 0 30 30" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M5 24.375C5 23.5462 5.32924 22.7513 5.91529 22.1653C6.50134 21.5792 7.2962 21.25 8.125 21.25H25M5 24.375C5 25.2038 5.32924 25.9987 5.91529 26.5847C6.50134 27.1708 7.2962 27.5 8.125 27.5H25V2.5H8.125C7.2962 2.5 6.50134 2.82924 5.91529 3.41529C5.32924 4.00134 5 4.7962 5 5.625V24.375Z" stroke="#01A274" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>                    
                </i>Отчёты</li>
            </ul>
        </div>
        <div class="content" id="content">
        </div>
    </div>
    );
  }
}

export default Base;